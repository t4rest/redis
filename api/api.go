package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	"github.com/t4rest/redis/conf"
	"github.com/t4rest/redis/service"
)

// API .
type API struct {
	srv    service.Worker
	cfg    conf.HttpConf
	Router *httprouter.Router
}

// New .
func New(cfg conf.AppConf, srv service.Worker) *API {
	return &API{cfg: cfg.Http, srv}
}

// StartServe .
func (api *API) StartServe() {

	api.Router = httprouter.New()
	api.Router.GET("/", api.Index)

	logrus.Infof("Listening on port %s", api.cfg.ListenOnPort)
	err := http.ListenAndServe(api.cfg.ListenOnPort, api.Router)
	if err == nil {
		logrus.Fatal("Error can't launch the server on port: " + api.cfg.ListenOnPort)
	}
}

// Index .
func (api *API) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")

	work := api.srv.Work()

	w.Write([]byte(work))
}
