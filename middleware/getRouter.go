package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
)

func GetAllRouter(app *fiber.App) {
	fmt.Println("=======路由列表=======")
	for _, route := range app.GetRoutes() {
		if len(route.Path) > 1 {
			fmt.Printf("[%s] %-25s -> %s\n", route.Method, route.Path, route.Name)
		}
	}
}
