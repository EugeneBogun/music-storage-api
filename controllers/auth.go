package controllers

import (
	valid "github.com/asaskevich/govalidator"
	"github.com/eugenebogun/music-storage/components"
	"github.com/eugenebogun/music-storage/models"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

// RegistrationForm is registration form struct with validation rules.
type RegistrationForm struct {
	Email    string `valid:"email,required"`
	Password string `valid:"alphanum,length(6|50),required"`
}

// UserResponse is response struct with user as response attribute.
type UserResponse struct {
	User models.User `json:"user"`
}

func init() {
	valid.SetFieldsRequiredByDefault(true)
}

// HandleRegistration handle registration request.
// Using registration form struct for validation registration form data.
// In case success validation create new user.
func HandleRegistration(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	registrationForm := &RegistrationForm{Email: r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	_, validErr := valid.ValidateStruct(registrationForm)
	if validErr != nil {
		components.SendErrorResponse(w, validErr.Error())
		return
	}

	user := models.User{
		Id:       bson.NewObjectId(),
		Email:    registrationForm.Email,
		Password: registrationForm.Password,
	}
	insertErr := user.Insert()
	if insertErr != nil {
		components.SendErrorResponse(w, insertErr.Error())
		return
	}

	components.SendResponse(w, UserResponse{User: user}, http.StatusOK)
	return
}
