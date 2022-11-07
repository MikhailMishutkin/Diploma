package app

import (
	"fmt"
	"graduatework/internal/handler"
	dcollect "graduatework/internal/infrastructure/microservices"
	"graduatework/internal/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type server struct {
	Server http.Server
	Router *mux.Router
	micro  *dcollect.MicroServiceStr
}

func NewServer() *server {
	return &server{
		Server: http.Server{},
		Router: mux.NewRouter(),
	}

}

//...
func RunServer() {
	s := NewServer()
	microServ := s.micro.MicroService()
	service := service.NewServiceManage(microServ)
	handle := handler.NewHandler(service)
	handle.RegisterR(s.Router)
	//	s.Router.HandleFunc("/", handler.HandleConnection)
	fmt.Print("server starts at port 8282 \n")
	err := http.ListenAndServe("localhost:8282", s)
	if err != nil {
		log.Fatal("server didn't start: ", err)
	}

}

//...
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}
