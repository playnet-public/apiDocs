package main

import (
	"fmt"
	"net/http"

	"github.com/bukalapak/snowboard/adapter/drafter"
	snowboard "github.com/bukalapak/snowboard/parser"
	"github.com/gorilla/mux"
	"github.com/playnet-public/flagenv"
	"github.com/urfave/negroni"
	"gitlab.allgameplay.de/Vincent/apiDocs/pkg/render"
)

var (
	endpoints                 *render.Endpoints
	engine                    snowboard.Parser
	httpHost, defaultTemplate string
	httpPort                  int
)

func main() {
	engine = drafter.Engine{}
	flagenv.EnvPrefix = "APIDOCSD_"
	handleArgs()

	endpoints = render.NewEndpoints(engine, defaultTemplate)

	http.ListenAndServe(
		fmt.Sprintf("%s:%d", httpHost, httpPort),
		getHandler(),
	)
}

func handleArgs() {
	flagenv.StringVar(&httpHost, "address", "localhost", "The host of the server.")
	flagenv.IntVar(&httpPort, "port", 8088, "The listening port of the server.")
	flagenv.StringVar(&defaultTemplate, "defaultTemplate", "", "Path to default template.")
	flagenv.Parse()
}

func getRouter() *mux.Router {
	router := mux.NewRouter()

	router.
		Methods("POST").
		Path("/render").
		Name("RenderIt").
		HandlerFunc(endpoints.RenderIt)

	return router
}

func getHandler() http.Handler {
	n := negroni.Classic()
	n.UseHandler(getRouter())
	return n
}
