package admin

import (
	"github.com/gorilla/mux"
	admincontroller "github.com/nidhey27/ticket-booking-app/controllers/adminController"
)

var AdminRoutes = func(router *mux.Router) {
	router.HandleFunc("/admin/login", admincontroller.Register).Methods("POST")
}
