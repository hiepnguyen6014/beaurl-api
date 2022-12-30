package main

import (
	"os"

	"beaurl.vn/api/src/configs"
	"beaurl.vn/api/src/routers"
)

func main() {
	configs.LoadEnv()

	router := routers.Routers()

	router.Run(":" + os.Getenv("PORT"))
}
