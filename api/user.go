package api

import (
	"NewProUser/entity"
	"NewProUser/pkg/utils"
	"net/http"

	"github.com/google/uuid"
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
	password,err:=utils.HashPassword(body.Password)
	if err!=nil {
		HandleInternalWithMessage(w, err, "error with hashing with password")
		return
	}
	body.Password=password
	body.ID=uuid.NewString()

	id, err:=api.userser.SignUpUser(r.Context(), body)
	 if err!=nil {
		HandleInternalWithMessage(w, err,"error in SignUpUserApi")
		return
	 }

	 tokenCredentials:=map[string]string{
		"role":body.Role,
	 }

	 tokens, err:=utils.GenerateNewTokenForUser(id, tokenCredentials)

	 if err!=nil {
		HandleInternalWithMessage(w, err, "error in generate tokens")
		return
	 }
	 WriteJSONWithSuccess(w, entity.SignUpResModel{
		ID: body.ID,
		Refresh: tokens.Refresh,
		Access: tokens.Access,
	 })
	 
}
