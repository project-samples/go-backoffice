package company

import (
	"context"
	"net/http"
	"reflect"

	sv "github.com/core-go/core"
	"github.com/core-go/core/builder"
	"github.com/core-go/search"
)

type CompanyTransport interface {
	Search(w http.ResponseWriter, r *http.Request)
	Load(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Patch(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewCompanyHandler(
	find func(context.Context, interface{}, interface{}, int64, ...int64) (int64, string, error),
	companieservice CompanyRepository,
	conf sv.WriterConfig,
	logError func(context.Context, string, ...map[string]interface{}),
	generateId func(context.Context) (string, error),
	validate func(context.Context, interface{}) ([]sv.ErrorMessage, error),
	tracking builder.TrackingConfig,
	writeLog func(context.Context, string, string, bool, string) error,
) CompanyTransport {
	modelType := reflect.TypeOf(Company{})
	searchModelType := reflect.TypeOf(CompanyFilter{})
	builder := builder.NewBuilderWithIdAndConfig(generateId, modelType, tracking)
	patchHandler, params := sv.CreatePatchAndParams(modelType, conf.Status, logError, companieservice.Patch, validate, builder.Patch, conf.Action, writeLog)
	searchHandler := search.NewSearchHandler(find, modelType, searchModelType, logError, writeLog)
	return &CompanyHandler{service: companieservice, builder: builder, PatchHandler: patchHandler, SearchHandler: searchHandler, Params: params}
}

type CompanyHandler struct {
	service CompanyRepository
	builder sv.Builder
	*sv.PatchHandler
	*search.SearchHandler
	*sv.Params
}

func (h *CompanyHandler) Load(w http.ResponseWriter, r *http.Request) {
	id := sv.GetRequiredParam(w, r)
	if len(id) > 0 {
		result, err := h.service.Load(r.Context(), id)
		sv.RespondModel(w, r, result, err, h.Error, nil)
	}
}
func (h *CompanyHandler) Create(w http.ResponseWriter, r *http.Request) {
	var company Company
	er1 := sv.Decode(w, r, &company, h.builder.Create)
	if er1 == nil {
		errors, er2 := h.Validate(r.Context(), &company)
		if !sv.HasError(w, r, errors, er2, *h.Status.ValidationError, h.Error, h.Log, h.Resource, h.Action.Create) {
			result, er3 := h.service.Create(r.Context(), &company)
			sv.AfterCreated(w, r, &company, result, er3, h.Status, h.Error, h.Log, h.Resource, h.Action.Create)
		}
	}
}
func (h *CompanyHandler) Update(w http.ResponseWriter, r *http.Request) {
	var company Company
	er1 := sv.DecodeAndCheckId(w, r, &company, h.Keys, h.Indexes, h.builder.Update)
	if er1 == nil {
		errors, er2 := h.Validate(r.Context(), &company)
		if !sv.HasError(w, r, errors, er2, *h.Status.ValidationError, h.Error, h.Log, h.Resource, h.Action.Update) {
			result, er3 := h.service.Update(r.Context(), &company)
			sv.HandleResult(w, r, &company, result, er3, h.Status, h.Error, h.Log, h.Resource, h.Action.Update)
		}
	}
}
func (h *CompanyHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := sv.GetRequiredParam(w, r)
	if len(id) > 0 {
		result, err := h.service.Delete(r.Context(), id)
		sv.HandleDelete(w, r, result, err, h.Error, h.Log, h.Resource, h.Action.Delete)
	}
}
