package main

import (
	"fmt"
	"net/http"

	"github.com/TheMysteriousVincent/liveinlife-restapi/pkg/config"
	"github.com/bukalapak/snowboard/adapter/drafter"
	snowboard "github.com/bukalapak/snowboard/parser"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"github.com/urfave/negroni"
)

var (
	engine        snowboard.Parser
	configuration *viper.Viper
)

func main() {
	engine = drafter.Engine{}
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
		HandlerFunc(RenderIt)
	return router
}
