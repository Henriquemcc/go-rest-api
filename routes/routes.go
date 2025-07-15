package routes

import (
	"log"
	"net/http"

	"github.com/Henriquemcc/go-rest-api/controllers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades)
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade)
	log.Fatal(http.ListenAndServe(":8000", r))
}
