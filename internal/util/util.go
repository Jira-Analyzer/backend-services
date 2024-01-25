package util

import (
	"encoding/json"
	"net/http"

	errorlib "github.com/Jira-Analyzer/backend-services/internal/error"
	log "github.com/sirupsen/logrus"
)

func WriteJSON(writer http.ResponseWriter, jsonStruct interface{}) {
	err := json.NewEncoder(writer).Encode(jsonStruct)
	if err != nil {
		jsonerr := errorlib.GetJSONError("json encoder error", errorlib.ErrHttpInternal)
		log.Error(err)

		if err = json.NewEncoder(writer).Encode(jsonerr); err != nil {
			log.Error(err)
		}
		writer.WriteHeader(jsonerr.Error.Code)
		return
	}
}
