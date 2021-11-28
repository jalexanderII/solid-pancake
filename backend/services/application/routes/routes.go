// package routes
//
// import (
// 	"github.com/gofiber/fiber/v2"
// 	ApplicationH "github.com/jalexanderII/solid-pancake/services/application/handlers"
// )
//
// func SetupApplicationRoutes(v1 fiber.Router) {
// 	// Application endpoints
// 	application := v1.Group("/application")
// 	application.Post("/apply", ApplicationH.Apply)
// 	application.Post("/:id/upload", ApplicationH.Upload)
// 	// Application Request endpoints
// 	appReq := application.Group("/request")
// 	appReq.Get("/", ApplicationH.GetApplications)
// 	appReq.Get("/:id", ApplicationH.GetApplication)
// 	appReq.Delete("/:id", ApplicationH.DeleteApplication)
// 	// Application Response endpoints
// 	appResp := application.Group("/response")
// 	appResp.Get("/:id", ApplicationH.GetApplicationResponse)
// 	appResp.Delete("/:id", ApplicationH.DeleteApplicationResponse)
// }
