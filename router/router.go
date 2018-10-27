package router

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
	"github.com/sysu-go-online/auth-service/controller"
	"github.com/sysu-go-online/public-service/types"
	"github.com/urfave/negroni"
)

var upgrader = websocket.Upgrader{}

// GetServer return web server
func GetServer() *negroni.Negroni {
	r := mux.NewRouter()

	r.Handle("", types.ErrorHandler(controller.LogInHandler)).Methods("POST")
	r.Handle("/", types.ErrorHandler(controller.LogInHandler)).Methods("POST")
	r.Handle("", types.ErrorHandler(controller.LogOutHandler)).Methods("DELETE")
	r.Handle("/", types.ErrorHandler(controller.LogOutHandler)).Methods("DELETE")

	// project collection

	// Use classic server and return it
	handler := cors.Default().Handler(r)
	s := negroni.Classic()
	s.UseHandler(handler)
	return s
}
