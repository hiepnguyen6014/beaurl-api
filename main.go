package main

import (
	"fmt"
	"os"

	"beaurl.vn/api/src/configs"
	"beaurl.vn/api/src/routes"
)

func main() {
	configs.LoadEnv()

	fmt.Println("http://127.0.0.1:" + os.Getenv("PORT"))

	router := routes.Routers()

	router.Run(":" + os.Getenv("PORT"))

}
