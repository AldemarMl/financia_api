package main

import (
		"fmt"
		"net/http"
		log "github.com/sirupsen/logrus"
		//"./controllers"
		"./api"
		"./config"
	)
const port string = "8000"

func main() {
	log.Info("App is running in port 8000")
	/*err := godotenv.Load("/config")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w,"hola mundo")
	})
	http.HandleFunc("/hello", controllers.Hello)*/
	log.SetLevel(log.DebugLevel)
	log.WithField("version",config.Version).Debug("starting server.")
	router, err := api.NewRouter()
	if err != nil {
		log.WithError(err).Fatal("Error building router")
	}
	server := http.ListenAndServe(fmt.Sprintf(":%s", port), router)
	log.Fatal(server)
}

