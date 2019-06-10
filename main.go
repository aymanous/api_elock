// @APIVersion 1.0.0
// @APITitle Teamwork Desk
// @APIDescription Bend Teamwork Desk to your will using these read and write endpoints
// @Contact support@teamwork.com
// @TermsOfServiceUrl https://www.teamwork.com/termsofservice
// @License BSD
// @LicenseUrl http://opensource.org/licenses/BSD-2-Clause

package main

import (
	"errors"
	"flag"
	"log"
	"net/http"
	"os"

	configuration "./Configuration"
	helperhttp "./Helper/Http"
	middlewares "./Middlewares"
	model "./Model"
	apiService "./Services/API"
	monitoringService "./Services/Monitoring"
)

var (
	gitHash   = ""
	buildDate = ""
)

func corsAllowed(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		next.ServeHTTP(w, r)
	})
}

func main() {
	// TODO en attendant de régler le problème des packages locaux non accessible avec -ldflags
	configuration.SetBuildInfo(gitHash, buildDate)

	// Création du logger de sortie
	logger := log.New(os.Stderr, "", log.LstdFlags)
	// Affichage des informations du logiciel
	logger.Print(configuration.String())

	// Parse application command line options
	configFile := flag.String("config", "", "Server configuration file path")
	flag.Parse()

	// Lecture de la configuration du service
	var err = errors.New("No configuration set")
	if *configFile != "" {
		err = configuration.ReadAndCreate(*configFile)
	} else {
		err = configuration.ReadSpringCloudConfig()
	}
	if err != nil {
		flag.Usage()
		logger.Fatal(err)
	}
	config := configuration.Get().Configuration.Data

	model, err := model.NewModel(config.DBConfig)
	if err != nil {
		logger.Fatal(err)
	}

	// Create services runner
	servicesRunner := helperhttp.Services{Logger: logger}

	// API endpoints
	apiRouter := apiService.NewRouter(model)
	apiRouter.Use(middlewares.NewPrometheus(configuration.Get().Name).Handler)
	apiRouter.Use(corsAllowed)
	servicesRunner.Add("API", config.API, apiRouter)

	// Monitoring endpoints
	servicesRunner.Add("Monitoring", config.Monitoring, monitoringService.NewRouter())

	// Run all services
	servicesRunner.Run()
}
