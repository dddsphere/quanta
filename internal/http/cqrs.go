package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"

	"github.com/dddsphere/quanta/internal/core/errors"
	"github.com/dddsphere/quanta/internal/event/handler"
	"github.com/dddsphere/quanta/internal/system"
)

type (
	CQRSHandler struct {
		system.Worker
		event handler.Handler
	}

	CommandResponse struct {
		RequestID string `json:"request_id"`
	}
)

const (
	name = "cqrs-handler"
)

func NewCQRSHadler(opts ...system.Option) *CQRSHandler {
	return &CQRSHandler{
		Worker: system.NewWorker(name, opts...),
	}
}

func (h *CQRSHandler) Dispatch(w http.ResponseWriter, r *http.Request, commandName string) {
	//path := strings.Split(r.URL.Path, "/")
	//name := path[len(path)-1]

	switch commandName {
	case "create-list":
		h.CreateList(w, r)

	//case "update-list":
	//	h.UpdateList(w, r)

	//case "delete-list":
	//	h.DeleteList(w, r)

	//case "add-item-to-list":
	//	h.AddItemToList(w, r)

	//case "add-item-to-list":
	//	h.AddItemToList(w, r)

	//case "remove-item-from-list":
	//	h.RemoveItemFromList(w, r)

	//case "add-tag-to-item":
	//	h.AddTagToItem(w, r)

	//case "add-place-to-item":
	//	h.AddPlaceToItem(w, r)

	//case "add-place-to-item":
	//	h.AddPlaceToItem(w, r)

	//case "set-category-to-item":
	//	h.SetCategoryToItem(w, r)

	//case "set-new-status-to-item":
	//	h.SetNewStatusToItem(w, r)

	default:
		msg := fmt.Sprintf("command '%s' not found", commandName)
		err := errors.Wrap(msg, CommandDispatchError)
		h.Error(err, w)
	}
}

func (h *CQRSHandler) genReqID(r *http.Request) string {
	reqID := r.Header.Get("X-Request-ID")
	if reqID == "" {
		reqID = uuid.New().String()
	}

	return reqID
}

func (h *CQRSHandler) CommandOK(w http.ResponseWriter, reqID string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	cr := CommandResponse{RequestID: reqID}
	resData, err := json.Marshal(cr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(resData)
	if err != nil {
		err = errors.Wrap("error writing command OK response", err)
		h.Log().Error(err.Error())
	}
}

func (h *CQRSHandler) Error(err error, w http.ResponseWriter) {
	h.Log().Error(err.Error())
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
