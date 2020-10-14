package api

import (
	"clean-api/api/action"
	"github.com/gorilla/mux"
	"github.com/labstack/gommon/log"
	"net/http"
)

type HTTPServer struct {
}

func (s HTTPServer) Listen() {
	router := mux.NewRouter()
	router.HandleFunc("/healthcheck", action.HealthCheck)
	log.Printf("initiated connection on port 3001")
	if err := http.ListenAndServe(":3001", router); err != nil {
		log.Fatal(err)
	}
}
