package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	initializers "github.com/nidhey27/ticket-booking-app/initilizers"
	"github.com/nidhey27/ticket-booking-app/routes/admin"
)

type ResponseWriter struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error"`
	Data    []any  `json:"data"`
}

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		res := &ResponseWriter{}

		res.Status = true
		res.Message = "Ticket Booking App by Nidhey Indurkar"
		res.Data = make([]any, 0)

		json.NewEncoder(w).Encode(res)
	})

	admin.AdminRoutes(r)

	http.ListenAndServe(":8080", r)
}
