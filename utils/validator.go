package utils

import (
	"bytes"
	"github.com/davdwhyte87/travelfy/models"
	"github.com/thedevsaddam/govalidator"
	"io/ioutil"
	"net/http"
)


// CreateUserValidator ... 
func CreateUserValidator (w http.ResponseWriter, r *http.Request) bool {
	b, bodyReaderErr := ioutil.ReadAll(r.Body)
	if bodyReaderErr != nil {
		RespondWithError(w, http.StatusInternalServerError, "Error reading request body")
	}

	r.Body = ioutil.NopCloser(bytes.NewBuffer(b))

	rules := govalidator.MapData{
		"name": []string{"required",  "min:4", "max:20"},
		"email":    []string{"required", "email"},
		"password":      []string{"required", "min:4", "max:20"},
	}
	//
	//messages := govalidator.MapData{
	//	"name": []string{"Your name is required", "between"},
	//	"email":    []string{"An email is required", "min:4", "Max: 20 ", "A valid email is required"},
	//	"password": []string{"A password is required" },
	//}
	var user models.User
	opts := govalidator.Options{
		Request:         r,        // request object
		Rules:           rules,    // rules map
		// custom message map (Optional)
		Data:&user,
		RequiredDefault: false,     // all the field to be pass the rules
	}
	v := govalidator.New(opts)
	e := v.ValidateJSON()

	r.Body = ioutil.NopCloser(bytes.NewBuffer(b))
	if len(e) != 0  {
		err := map[string]interface{}{"validationError": e}
		w.Header().Set("Content-type", "application/json")
		var returnData = ReturnData{ Status:http.StatusBadRequest, Error:err }
		RespondWithJson(w,http.StatusBadRequest,returnData)
		return false
	}
	return true
}
