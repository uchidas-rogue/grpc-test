package main

func main() {

	httpService := NewHTTPServer(":8080")
	go httpService.Run()

	server := NewGRPCServer(":9000")
	server.Run()
}
