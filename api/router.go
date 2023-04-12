package api

import (
	"NewProUser/service"

	"go.uber.org/zap"

	"github.com/gorilla/mux"
)

type routes struct {
	root    *mux.Router
	apiRoot *mux.Router
}
type api struct {
	routes  *routes
	userser service.UserService
	logger  *zap.Logger
}

func Init(
	root *mux.Router,
	userser service.UserService,
	logger *zap.Logger) {
	r := routes{
		root:    root,
		apiRoot: root.PathPrefix("/api").Subrouter(),
	}
	api := api{
		routes:  &r,
		userser: userser,
	}
	api.initUser()
}
func (api *api) initUser() {
	api.routes.apiRoot.HandleFunc("/v1/sign-up-user", api.SignUpUser).Methods("POST")
}
