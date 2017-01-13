package controllers

import (
    "net/http"
    "github.com/julienschmidt/httprouter"
    "bitbucket.org/ebogun/music-storage/components"
    valid "github.com/asaskevich/govalidator"
    "gopkg.in/mgo.v2/bson"
    "bitbucket.org/ebogun/music-storage/models"
)

type RegistrationForm struct {
    Email    string `valid:"email,required"`
    Password string `valid:"alphanum,length(6|50),required"`
}

type UserResponse struct {
    User models.User `json:"user"`
}

func init() {
    valid.SetFieldsRequiredByDefault(true)
}

func HandleRegistration(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    registrationForm := &RegistrationForm{
        Email: r.FormValue("email"),
        Password: r.FormValue("password"),
    }

    _, validErr := valid.ValidateStruct(registrationForm)
    if validErr != nil {
        components.SendErrorResponse(w, validErr.Error())
        return
    }

    user := models.User{
        Id : bson.NewObjectId(),
        Email: registrationForm.Email,
        Password: registrationForm.Password,
    }
    insertErr := user.Insert()
    if insertErr != nil {
        components.SendErrorResponse(w, insertErr.Error())
        return
    }

    components.SendResponse(w, UserResponse{User:user}, http.StatusOK)
    return
}