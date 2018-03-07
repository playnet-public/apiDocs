package main

import (
	"fmt"
	"net/http"

	"github.com/bukalapak/snowboard/adapter/drafter"
	snowboard "github.com/bukalapak/snowboard/parser"
	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"github.com/playnet-public/apiDocs/pkg/render"
	"github.com/playnet-public/flagenv"
	"github.com/urfave/negroni"
)

var (
	endpoints                 *render.Endpoints
	engine                    snowboard.Parser
	httpHost, defaultTemplate string
	httpPort                  int
)

func main() {
	flagenv.EnvPrefix = "APIDOCSD_"
	handleArgs()
	glog.Info("server is setting up")
	engine = drafter.Engine{}
	endpoints = render.NewEndpoints(engine, defaultTemplate)

	glog.Infof("server is listening on %s:%d", httpHost, httpPort)
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
