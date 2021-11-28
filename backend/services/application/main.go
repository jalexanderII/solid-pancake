package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	"google.golang.org/grpc"
	applicationpb "github.com/jalexanderII/solid-pancake/gen/applicaiton"
)


type userID int32

var (
	applicationAddr              = "localhost:9093"
)

func main() {
	grpcServer, lis := setupApplicationServer()
	grpcServer.Serve(lis)
}

func setupApplicationServer() (*grpc.Server, net.Listener) {
	lis, err := net.Listen("tcp", applicationAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// setup and register currency service
	// create a new gRPC server, use WithInsecure to allow http connections
	grpcServer := grpc.NewServer()
	applicationpb.RegisterApplicationServer(grpcServer, newApplicationServer())

	return grpcServer, lis
}

func newApplicationServer() *casinoServer {
	return &casinoServer{
		stockPrice:     100,
		userToTokens:   map[userID]int32{},
		userToPayments: map[userID][]int32{},
		userToStocks:   map[userID]int32{},
	}
}

type applicaitonServer struct {
	stockPrice int32

	userToTokens   map[userID]int32
	userToPayments map[userID][]int32
	userToStocks   map[userID]int32
}

func (c *casinoServer) BuyTokens(ctx context.Context, payment *commonpb.Payment) (*casinopb.Tokens, error) {
	log.Printf("BuyTokens invoked with payment %v\n", payment)

	usrID := userID(payment.User.GetId())
	tokens := payment.GetAmount() * tokensPerDollar

	c.userToTokens[usrID] += tokens
	c.userToPayments[usrID] = append(c.userToPayments[usrID], -payment.Amount)

	return &casinopb.Tokens{Count: tokens}, nil
}

func (c *casinoServer) Withdraw(ctx context.Context, withdrawReq *casinopb.WithdrawRequest) (*commonpb.Payment, error) {
	toWithdraw := withdrawReq.GetTokensCnt()
	log.Printf("Withdraw invoked with tokens %v\n", toWithdraw)

	usrID := userID(withdrawReq.User.GetId())
	log.Println(c.userToTokens[usrID])
	if !c.hasEnoughTokens(usrID, toWithdraw) {
		return nil, fmt.Errorf("not enough tokens to withdraw")
	}

	amount := toWithdraw / tokensPerDollar
	c.userToTokens[usrID] -= toWithdraw
	c.userToPayments[usrID] = append(c.userToPayments[usrID], amount)

	return &commonpb.Payment{User: withdrawReq.User, Amount: amount}, nil
}

func (c *casinoServer) GetTokenBalance(_ context.Context, user *commonpb.User) (*casinopb.Tokens, error) {
	log.Printf("GetTokenBalance invoked with user %v\n", user)

	usrID := userID(user.GetId())
	return &casinopb.Tokens{Count: c.userToTokens[usrID]}, nil
}

func (c *casinoServer) GetPayments(user *commonpb.User, stream casinopb.Casino_GetPaymentsServer) error {
	log.Printf("GetPayments invoked with user %v", user)

	usrID := userID(user.GetId())
	payments := c.userToPayments[usrID]
	for _, payment := range payments {
		err := stream.Send(&commonpb.Payment{
			User:   user,
			Amount: payment,
		})
		if err != nil {
			return fmt.Errorf("failed sending payment through stream: %w", err)
		}
	}

	return nil
}

func (c *casinoServer) GetPaymentStatement(ctx context.Context, user *commonpb.User) (*commonpb.PaymentStatement, error) {
	log.Printf("GetPaymentStatement invoked with user %v\n", user)

	stream, err := paymentStatementsClient.CreateStatement(ctx)
	if err != nil {
		return nil, fmt.Errorf("couldn't create payment statements stream: %w", err)
	}

	usrID := userID(user.GetId())
	payments := c.userToPayments[usrID]
	for _, payment := range payments {
		err := stream.Send(&commonpb.Payment{User: user, Amount: payment})
		if err != nil {
			return nil, fmt.Errorf("failed sending payment to payments_statements: %w", err)
		}
	}

	return stream.CloseAndRecv()
}

func (c *casinoServer) Gamble(stream casinopb.Casino_GambleServer) error {
	log.Println("Gamble invoked...")

	errc := make(chan error, 2)
	go iterateStreamWithHandler(errc, stream, c.handleUserGamblingAction)
	go iterateStreamWithHandler(errc, stream, c.incrementAndSendStockPrice)

	err := <-errc
	log.Println("Gambling ending with err " + err.Error())

	return err
}

func iterateStreamWithHandler(errc chan error, stream casinopb.Casino_GambleServer, handler streamHandler) {
	for {
		select {
		case <-errc:
			return
		default:
		}

		err := handler(stream)
		if err != nil {
			errc <- err
			break
		}
	}
}

func (c *casinoServer) handleUserGamblingAction(stream casinopb.Casino_GambleServer) error {
	action, err := stream.Recv()
	if err != nil {
		return err
	}

	usrID := userID(action.User.GetId())
	targetTokens := action.StocksCount * c.stockPrice
	switch action.Type {
	case casinopb.ActionType_BUY:
		if !c.hasEnoughTokens(usrID, targetTokens) {
			return stream.Send(&casinopb.GambleInfo{
				Type:   casinopb.GambleType_ACTION_RESULT,
				Result: &casinopb.ActionResult{Msg: "you don't have enough tokens"},
			})
		}

		c.userToTokens[usrID] -= targetTokens
		c.userToStocks[usrID] += action.StocksCount
	case casinopb.ActionType_SELL:
		if !c.hasEnoughStocks(usrID, action.StocksCount) {
			return stream.Send(&casinopb.GambleInfo{
				Type:   casinopb.GambleType_ACTION_RESULT,
				Result: &casinopb.ActionResult{Msg: "you don't have enough stocks to sell"},
			})
		}

		c.userToTokens[usrID] += targetTokens
		c.userToStocks[usrID] -= action.StocksCount
	default:
		return errors.New("unknown operation")
	}

	return stream.Send(&casinopb.GambleInfo{
		Type:   casinopb.GambleType_ACTION_RESULT,
		Result: &casinopb.ActionResult{Msg: "operation executed successfully"},
	})

}

func (c *casinoServer) incrementAndSendStockPrice(stream casinopb.Casino_GambleServer) error {
	time.Sleep(10 * time.Second)
	c.stockPrice += int32(rand.Intn(14) + 1)
	c.stockPrice -= int32(rand.Intn(14) + 1)

	log.Println("sending stock price", c.stockPrice)
	return stream.Send(&casinopb.GambleInfo{
		Type: casinopb.GambleType_STOCK_INFO,
		Info: &casinopb.StockInfo{
			Name:  "AwesomeStock",
			Price: c.stockPrice,
		},
	})
}
