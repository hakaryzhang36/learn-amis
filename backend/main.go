package main

import "hakaryzhang.com/amis-test/router"

func main() {
	r := router.Router()
	r.Run("localhost:8080")
}
