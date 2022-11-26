package admincontroller

import (
	"encoding/json"
	"log"
	"net/http"

	initializers "github.com/nidhey27/ticket-booking-app/initilizers"
	"github.com/nidhey27/ticket-booking-app/models"
	"github.com/nidhey27/ticket-booking-app/utils"
	"github.com/nidhey27/ticket-booking-app/validators"
)

func Register(w http.ResponseWriter, r *http.Request) {
	admin := &models.Admin{}

	utils.Parsebody(r, admin)

	ok := validators.AdminRegister(admin)

	log.Print(ok)
	json.NewEncoder(w).Encode(admin)
	return
	w.Header().Set("Content-Type", "application/json")

	// user := &models.Admin{Name: "Jinzhu", Email: "Jinzhu@gmail.com", Password: "Hola"}

	// result := initializers.DB.Create(&user)
	// log.Printf(result)

	var users []models.Admin
	initializers.DB.Find(&users)

	json.NewEncoder(w).Encode(users)
}
