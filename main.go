package main

import (
	"fiber/boot"
	_ "fiber/docs"
)

//	@title			FiberTemplate Services
//	@version		0.0.1
//	@description	RESTful Self-ie Academy API Services. Built to ensure Self-ie Services are good to be served!
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.email	fiber@swagger.io
//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		selfieapi.up.railway.app
//	@BasePath	/v1

// @accept		json
// @produce	json
func main() {
	boot.InitHttpServer()
}
