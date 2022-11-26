package validators

import (
	"github.com/nidhey27/ticket-booking-app/models"
	"github.com/softbrewery/gojoi/pkg/joi"
)

type Validators struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

func AdminRegister(body *models.Admin) Validators {
	schema := joi.Struct().Keys(joi.StructKeys{
		"Name":     joi.String().Required(),
		"Email":    joi.String().Email(nil).Required(),
		"Password": joi.String().Required(),
	})

	err := joi.Validate(body, schema)
	if err != nil {
		return Validators{Status: false, Message: err.Error()}
	}
	return Validators{Status: true, Message: "Valid"}
}
