package main

import "API/service2/internal/server"

func main() {
	s := server.New()
	s.Start()
}
