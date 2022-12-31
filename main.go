package main

import (
	"os"

	"beaurl.vn/api/src/configs"
)

func main() {

	router := configs.Router()

	router.Run(":" + os.Getenv("PORT"))

}
