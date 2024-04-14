package http

import "github.com/gorilla/mux"

func MapWSRoutes(r *mux.Router, w *websocket.websocket) {
	r.HandleFunc("/lp", w.changeLPEvent())
	r.HandleFunc("/order", w.newOrderEvent())
}
