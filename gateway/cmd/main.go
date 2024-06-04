package main

import "API/gateway/internal/router"

func main() {
	r := router.New()
	r.Start()
}
