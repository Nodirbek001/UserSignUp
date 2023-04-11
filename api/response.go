package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type response struct {
	Error bool        `json: "error"`
	Data  interface{} `json:"data"`
}

type errorInfo struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func BodyParser(r *http.Request, body interface{}) error {
	return json.NewDecoder(r.Body).Decode(&body)
}
func writeJSON(w http.ResponseWriter, data interface{}) {
	bytes, _ := json.MarshalIndent(data, "", " ")
	w.Header().Set("Content-Type", "Application/json")
	w.Write(bytes)
}

func WriteJSONWithSuccess(w http.ResponseWriter, data interface{}) {
	data = response{
		Error: false,
		Data:  data,
	}
	bytes, _ := json.MarshalIndent(data, "", " ")
	w.Header().Set("Content-Type", "Aplication/json")
	w.Write(bytes)
}

func HandleInternalWithMessage(w http.ResponseWriter, err error, message string) error {
	if err == nil {
		return nil
	}
	log.Println(message+" ", err)
	w.WriteHeader(http.StatusInternalServerError)
	writeJSON(w, response{
		Error: true,
		Data: errorInfo{
			Status:  http.StatusInternalServerError,
			Message: message,
		},
	})
	return err
}
func HandleBadRequestErrWithMessage(w http.ResponseWriter, err error, message string) error {
	if err == nil {
		return nil
	}
	log.Println(message+" ", err)
	w.WriteHeader(http.StatusBadRequest)
	writeJSON(w, response{Error: true,
		Data: errorInfo{
			Status:  http.StatusBadRequest,
			Message: message + ": " + err.Error(),
		},
	})
	return err

}
