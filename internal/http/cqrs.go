package http

import (
	"fmt"
	"net/http"

	"github.com/dddsphere/quanta/internal/core/errors"
	"github.com/dddsphere/quanta/internal/system"
)

type (
	CQRSHandler struct {
		system.Worker
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

func (h *CQRSHandler) Error(err error, w http.ResponseWriter) {
	h.Log().Error(err.Error())
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
