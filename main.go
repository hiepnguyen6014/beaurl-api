package main

import (
	"fmt"
	"os"

	"beaurl.vn/api/src/configs"
	"beaurl.vn/api/src/routes"
)

func main() {
	configs.LoadEnv()

	router := routes.Routers()

	configs.ConnectDB()
	configs.CommonConfigs(router)

	fmt.Println("http://127.0.0.1:" + os.Getenv("PORT"))
	router.Run(":" + os.Getenv("PORT"))

}
