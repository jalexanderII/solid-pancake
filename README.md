# solid-pancake

RentEnd2End

## Quick start

1. Rename `.env.example` to `.env`
2. Install [Optional[Air]](https://github.com/cosmtrek/air) for server auto-reload.
3. Run project with this command:

```bash
start server by running 'air' in terminal 
If you did not download air you can run the server with 'go run main.go' from the root directory

# Routes will take the form: http://127.0.0.1:9092/api/v1/...
# For example http://127.0.0.1:9092/api/v1/apartments = GET method "get all apartments"
# Routes can be found in the following folders and can be testing with curl or Postman:
#   - solid-pancake/services/application/routes/routes.go
#   - solid-pancake/services/realestate/routes/routes.go
```