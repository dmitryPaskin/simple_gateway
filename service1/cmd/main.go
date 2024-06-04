package main

import "API/service1/internal/server"

func main() {
	s := server.New()
	s.Start()
}
