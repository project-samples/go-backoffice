package term

import (
	"context"
	sv "github.com/core-go/core"
	"github.com/core-go/core/builder"
	"github.com/core-go/search"
	"net/http"
	"reflect"
)

type TermTransport interface {
	Search(w http.ResponseWriter, r *http.Request)
	Load(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Patch(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewTermHandler(find func(context.Context, interface{}, interface{}, int64, ...int64) (int64, string, error), service TermService,
	generateId func(context.Context) (string, error), status sv.StatusConfig, logError func(context.Context, string, ...map[string]interface{}),
	validate func(ctx context.Context, model interface{}) ([]sv.ErrorMessage, error),
	tracking builder.TrackingConfig, action *sv.ActionConfig, writeLog func(context.Context, string, string, bool, string) error) TermTransport {
	searchModelType := reflect.TypeOf(TermFilter{})
	modelType := reflect.TypeOf(Term{})
	builder := builder.NewBuilderWithIdAndConfig(generateId, modelType, tracking)
	patchHandler, params := sv.CreatePatchAndParams(modelType, &status, logError, service.Patch, validate, builder.Patch, action, writeLog)
	searchHandler := search.NewSearchHandler(find, modelType, searchModelType, logError, params.Log)
	return &TermHandler{service: service, builder: builder, PatchHandler: patchHandler, SearchHandler: searchHandler, Params: params}
}

type TermHandler struct {
	service TermService
	builder sv.Builder
	*sv.PatchHandler
	*search.SearchHandler
	*sv.Params
}

func (h *TermHandler) Load(w http.ResponseWriter, r *http.Request) {
	id := sv.GetRequiredParam(w, r)
	if len(id) > 0 {
		result, err := h.service.Load(r.Context(), id)
		sv.RespondModel(w, r, result, err, h.Error, nil)
	}
}
func (h *TermHandler) Create(w http.ResponseWriter, r *http.Request) {
	var term Term
	er1 := sv.Decode(w, r, &term, h.builder.Create)
	if er1 == nil {
		errors, er2 := h.Validate(r.Context(), &term)
		if !sv.HasError(w, r, errors, er2, *h.Status.ValidationError, h.Error, h.Log, h.Resource, h.Action.Create) {
			result, er3 := h.service.Create(r.Context(), &term)
			sv.AfterCreated(w, r, &term, result, er3, h.Status, h.Error, h.Log, h.Resource, h.Action.Create)
		}
	}
}
func (h *TermHandler) Update(w http.ResponseWriter, r *http.Request) {
	var term Term
	er1 := sv.DecodeAndCheckId(w, r, &term, h.Keys, h.Indexes, h.builder.Update)
	if er1 == nil {
		errors, er2 := h.Validate(r.Context(), &term)
		if !sv.HasError(w, r, errors, er2, *h.Status.ValidationError, h.Error, h.Log, h.Resource, h.Action.Update) {
			result, er3 := h.service.Update(r.Context(), &term)
			sv.HandleResult(w, r, &term, result, er3, h.Status, h.Error, h.Log, h.Resource, h.Action.Update)
		}
	}
}
func (h *TermHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := sv.GetRequiredParam(w, r)
	if len(id) > 0 {
		result, err := h.service.Delete(r.Context(), id)
		sv.HandleDelete(w, r, result, err, h.Error, h.Log, h.Resource, h.Action.Delete)
	}
}
