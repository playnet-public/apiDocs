package main

import (
	"fmt"
	"net/http"

	"github.com/bukalapak/snowboard/adapter/drafter"
	snowboard "github.com/bukalapak/snowboard/parser"
	"github.com/gorilla/mux"
	"github.com/playnet-public/apiDocs/pkg/config"
	"github.com/playnet-public/apiDocs/pkg/render"
	"github.com/spf13/viper"
	"github.com/urfave/negroni"
)

var (
	endpoints     *render.Endpoints
	engine        snowboard.Parser
	configuration *viper.Viper
)

func main() {
	engine = &drafter.Engine{}
	endpoints = render.NewEndpoints(engine)
	configuration = config.NewConfig("config.yml")
	handleArgs()
	handleEnvVars()

	n := negroni.Classic()
	n.UseHandler(getRouter())

	http.ListenAndServe(
		fmt.Sprintf("%s:%d", configuration.GetString("http.host"), configuration.GetInt("http.port")),
		n,
	)
}

func handleArgs() {
}

func handleEnvVars() {
}

func getRouter() *mux.Router {
	router := mux.NewRouter()
	router.
		Methods("GET").
		Path("/render").
		Name("RenderIt").
		HandlerFunc(endpoints.renderIt)
	return router
}
