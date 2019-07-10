package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"

	"github.com/ip-tracker-svc/config"
)

// Init is used to initialize the handlers
func Init() {

	logg().Infof("Started listening in port %s", viper.GetString(config.ListeningPort))

	// Initialize the new router
	router := mux.NewRouter()

	// Routes consist of a path and a handler function.
	router.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) { fmt.Fprintln(w, "Ip Tracker service is Up") }).Methods(http.MethodGet)
	router.HandleFunc("/v1/trackip/{address}", TrackIPHandler).Methods(http.MethodGet)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(viper.GetString(config.ListeningPort), router))
}
