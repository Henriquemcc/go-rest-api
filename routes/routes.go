package routes

import (
	"log"
	"net/http"

	"github.com/Henriquemcc/go-rest-api/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/personalidades", controllers.TodasPersonalidades)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
