package handlers

import (
	"github.com/noodles623/csp/errors"
	"github.com/noodles623/csp/objects"
	"io/ioutil"
	"net/http"

	"github.com/noodles623/csp/store"
)

type AssetHandler interface {
	Get(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	store store.AssetStore
}

func NewAssetHandler(store store.AssetStore) AssetHandler {
	return &handler{store: store}
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		WriteError(w, errors.ErrValidAssetIdIsRequired)
	}
	asst, err := h.store.Get(r.Context(), &objects.GetRequest{Id: id})
	if err != nil {
		WriteError(w, err)
		return
	}
	WriteResponse(w, &objects.AssetResponseWrapper{Asset: asst})
}

func (h *handler) List(w http.ResponseWriter, r *http.Request) {
	limit, err := IntFromString(w, r.URL.Query().Get("limit"))
	if err != nil {
		return
	}
	list, err := h.store.List(r.Context(), &objects.ListRequest{Limit: limit})
	if err != nil {
		WriteError(w, err)
	}
	WriteResponse(w, &objects.AssetResponseWrapper{Assets: list})
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		WriteError(w, errors.ErrUnprocessableEntity)
		return
	}
	asst := &objects.Asset{}
	if Unmarshal(w, data, asst) != nil {
		return
	}
	if err := h.store.Create(r.Context(), &objects.CreateRequest{Asset: asst}); err != nil {
		WriteError(w, err)
		return
	}
	WriteResponse(w, &objects.AssetResponseWrapper{Asset: asst})
}

func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		WriteError(w, errors.ErrValidAssetIdIsRequired)
		return
	}
	if _, err := h.store.Get(r.Context(), &objects.GetRequest{Id: id}); err != nil {
		WriteError(w, err)
		return
	}
	if err := h.store.Delete(r.Context(), &objects.DeleteRequest{Id: id}); err != nil {
		WriteError(w, err)
		return
	}
	WriteResponse(w, &objects.AssetResponseWrapper{})
}
