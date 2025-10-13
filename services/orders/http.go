package main

import (
	"log"
	"net/http"

	handler "github.com/uchidas-rogue/kitchen/services/orders/handler/orders"
	"github.com/uchidas-rogue/kitchen/services/orders/service"
)

type httpService struct {
	addr string
}

func NewHTTPServer(addr string) *httpService {
	return &httpService{addr: addr}
}

func (s *httpService) Run() error {
	// Implement HTTP server logic here
	router := http.NewServeMux()

	orderService := service.NewOrderService()
	orderHandler := handler.NewHttpOrdersHandler(orderService)
	orderHandler.RegisterRouter(router)

	log.Println("Starting server on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
