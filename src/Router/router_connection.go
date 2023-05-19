package Router

import (
	"CEO_ASSIST/Controllers"
	"CEO_ASSIST/Router/Newly_Added_Routes"
	"CEO_ASSIST/db"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var Cfg = &Controllers.Config{
	Collector: db.DATABASE_CONNECTION(),
	Router:    Router_Engine(),
}

func Router_Engine() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	fs2 := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fs2))

	RoutersGrouping()

	err := http.ListenAndServe(":8080", router)

	if err != nil {
		log.Println("Failed to start the server", err.Error())
	} else if err == nil {
		log.Println("Engine sucessfully starts at port", 8080)
	}

	return router
}

func RoutersGrouping() {
	Newly_Added_Routes.NewlyAdded()
}
