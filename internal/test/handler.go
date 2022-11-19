package test

import (
	"context"
	sv "github.com/core-go/core"
	"github.com/core-go/search"
	"net/http"
	"reflect"
)

type TestTransport interface {
	Search(w http.ResponseWriter, r *http.Request)
	Load(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Patch(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewTestHandler(find func(context.Context, interface{}, interface{}, int64, ...int64) (int64, string, error), service TestService, status sv.StatusConfig, logError func(context.Context, string, ...map[string]interface{}), validate func(context.Context, interface{}) ([]sv.ErrorMessage, error), action *sv.ActionConfig) *TestHandler {
	filterType := reflect.TypeOf(TestFilter{})
	modelType := reflect.TypeOf(Test{})
	params := sv.CreateParams(modelType, &status, logError, validate, action)
	searchHandler := search.NewSearchHandler(find, modelType, filterType, logError, params.Log)
	return &TestHandler{service: service, SearchHandler: searchHandler, Params: params}
}

type TestHandler struct {
	service TestService
	*search.SearchHandler
	*sv.Params
}

func (h *TestHandler) Load(w http.ResponseWriter, r *http.Request) {
	id := sv.GetRequiredParam(w, r)
	if len(id) > 0 {
		res, err := h.service.Load(r.Context(), id)
		sv.RespondModel(w, r, res, err, h.Error, nil)
	}
}
func (h *TestHandler) Create(w http.ResponseWriter, r *http.Request) {
	var test Test
	er1 := sv.Decode(w, r, &test)
	if er1 == nil {
		errors, er2 := h.Validate(r.Context(), &test)
		if !sv.HasError(w, r, errors, er2, *h.Status.ValidationError, h.Error, h.Log, h.Resource, h.Action.Create) {
			res, er3 := h.service.Create(r.Context(), &test)
			sv.AfterCreated(w, r, &test, res, er3, h.Status, h.Error, h.Log, h.Resource, h.Action.Create)
		}
	}
}
func (h *TestHandler) Update(w http.ResponseWriter, r *http.Request) {
	var test Test
	er1 := sv.DecodeAndCheckId(w, r, &test, h.Keys, h.Indexes)
	if er1 == nil {
		errors, er2 := h.Validate(r.Context(), &test)
		if !sv.HasError(w, r, errors, er2, *h.Status.ValidationError, h.Error, h.Log, h.Resource, h.Action.Update) {
			res, er3 := h.service.Update(r.Context(), &test)
			sv.HandleResult(w, r, &test, res, er3, h.Status, h.Error, h.Log, h.Resource, h.Action.Update)
		}
	}
}
func (h *TestHandler) Patch(w http.ResponseWriter, r *http.Request) {
	var test Test
	r, json, er1 := sv.BuildMapAndCheckId(w, r, &test, h.Keys, h.Indexes)
	if er1 == nil {
		errors, er2 := h.Validate(r.Context(), &test)
		if !sv.HasError(w, r, errors, er2, *h.Status.ValidationError, h.Error, h.Log, h.Resource, h.Action.Patch) {
			res, er3 := h.service.Patch(r.Context(), json)
			sv.HandleResult(w, r, json, res, er3, h.Status, h.Error, h.Log, h.Resource, h.Action.Patch)
		}
	}
}
func (h *TestHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := sv.GetRequiredParam(w, r)
    if len(id) > 0 {
		res, err := h.service.Delete(r.Context(), id)
		sv.HandleDelete(w, r, res, err, h.Error, h.Log, h.Resource, h.Action.Delete)
	}
}
