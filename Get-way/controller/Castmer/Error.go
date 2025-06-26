package castmer

import (
	"encoding/json"
	"net/http"


)


func (app *Application) HandleError(w http.ResponseWriter, code int, msg string, err error) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "invalid",
		"message": msg,
		"problem": err.Error(),
	})
}
