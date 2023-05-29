package http

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/dddsphere/quanta/internal/core/errors"
)

const (
	userID = "54c2b14a-c51d-427d-89c5-aed8b75e3b51"
)

func (h *CQRSHandler) CreateList(w http.ResponseWriter, r *http.Request) {
	// Auth service not implemented yet
	// Later will be acquired from there
	// Use default user ID for now

	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		err = errors.Wrap("not valid request body data", err)
		h.Error(err, w)
	}
	defer r.Body.Close()

	// Unmarshal JSON data into struct
	var data CreateListReq
	err = json.Unmarshal(body, &data)
	if err != nil {
		err = errors.Wrap("cannot unmarshall request data", err)
		h.Error(err, w)
	}
}
