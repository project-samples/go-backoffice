package ticket

import (
	"context"
	sv "github.com/core-go/core"
	"github.com/core-go/search"
	"net/http"
	"reflect"
)

type TicketTransport interface {
	Search(w http.ResponseWriter, r *http.Request)
	Load(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Patch(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewTicketHandler(find func(context.Context, interface{}, interface{}, int64, ...int64) (int64, string, error), service TicketService, status sv.StatusConfig, logError func(context.Context, string, ...map[string]interface{}), validate func(context.Context, interface{}) ([]sv.ErrorMessage, error), action *sv.ActionConfig) *TicketHandler {
	filterType := reflect.TypeOf(TicketFilter{})
	modelType := reflect.TypeOf(Ticket{})
	params := sv.CreateParams(modelType, &status, logError, validate, action)
	searchHandler := search.NewSearchHandler(find, modelType, filterType, logError, params.Log)
	return &TicketHandler{service: service, SearchHandler: searchHandler, Params: params}
}

type TicketHandler struct {
	service TicketService
	*search.SearchHandler
	*sv.Params
}

func (h *TicketHandler) Load(w http.ResponseWriter, r *http.Request) {
	id := sv.GetRequiredParam(w, r)
	if len(id) > 0 {
		res, err := h.service.Load(r.Context(), id)
		sv.RespondModel(w, r, res, err, h.Error, nil)
	}
}
func (h *TicketHandler) Create(w http.ResponseWriter, r *http.Request) {
	var ticket Ticket
	er1 := sv.Decode(w, r, &ticket)
	if er1 == nil {
		errors, er2 := h.Validate(r.Context(), &ticket)
		if !sv.HasError(w, r, errors, er2, *h.Status.ValidationError, h.Error, h.Log, h.Resource, h.Action.Create) {
			res, er3 := h.service.Create(r.Context(), &ticket)
			sv.AfterCreated(w, r, &ticket, res, er3, h.Status, h.Error, h.Log, h.Resource, h.Action.Create)
		}
	}
}
func (h *TicketHandler) Update(w http.ResponseWriter, r *http.Request) {
	var ticket Ticket
	er1 := sv.DecodeAndCheckId(w, r, &ticket, h.Keys, h.Indexes)
	if er1 == nil {
		errors, er2 := h.Validate(r.Context(), &ticket)
		if !sv.HasError(w, r, errors, er2, *h.Status.ValidationError, h.Error, h.Log, h.Resource, h.Action.Update) {
			res, er3 := h.service.Update(r.Context(), &ticket)
			sv.HandleResult(w, r, &ticket, res, er3, h.Status, h.Error, h.Log, h.Resource, h.Action.Update)
		}
	}
}
func (h *TicketHandler) Patch(w http.ResponseWriter, r *http.Request) {
	var ticket Ticket
	r, json, er1 := sv.BuildMapAndCheckId(w, r, &ticket, h.Keys, h.Indexes)
	if er1 == nil {
		errors, er2 := h.Validate(r.Context(), &ticket)
		if !sv.HasError(w, r, errors, er2, *h.Status.ValidationError, h.Error, h.Log, h.Resource, h.Action.Patch) {
			res, er3 := h.service.Patch(r.Context(), json)
			sv.HandleResult(w, r, json, res, er3, h.Status, h.Error, h.Log, h.Resource, h.Action.Patch)
		}
	}
}
func (h *TicketHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := sv.GetRequiredParam(w, r)
    if len(id) > 0 {
		res, err := h.service.Delete(r.Context(), id)
		sv.HandleDelete(w, r, res, err, h.Error, h.Log, h.Resource, h.Action.Delete)
	}
}
