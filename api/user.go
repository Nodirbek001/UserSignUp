package api

import (
	"NewProUser/entity"
	"net/http"
)

func (api *api) SignUpUser(w http.ResponseWriter, r *http.Request) {
	var body entity.SignUpModel
	if err := BodyParser(r, &body); err != nil {
		HandleBadRequestErrWithMessage(w, err, "error parsing json: ")
		return
	}
	if err:=entity.Validate(body.PhoneNumber,body.Password); err!=nil {
		HandleBadRequestErrWithMessage(w,err,"invalid json data:")
		return
	}
	password,err:=utils
}
