package question

import (
	"context"
	sv "github.com/core-go/core"
	"github.com/core-go/search"
	"net/http"
	"reflect"
)

type QuestionTransport interface {
	Search(w http.ResponseWriter, r *http.Request)
	Load(w http.ResponseWriter, r *http.Request)
	GetByIds(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Patch(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewQuestionHandler(find func(context.Context, interface{}, interface{}, int64, ...int64) (int64, string, error), service QuestionService, status sv.StatusConfig, logError func(context.Context, string, ...map[string]interface{}), validate func(context.Context, interface{}) ([]sv.ErrorMessage, error), action *sv.ActionConfig) *QuestionHandler {
	filterType := reflect.TypeOf(QuestionFilter{})
	modelType := reflect.TypeOf(Question{})
	params := sv.CreateParams(modelType, &status, logError, validate, action)
	searchHandler := search.NewSearchHandler(find, modelType, filterType, logError, params.Log)
	return &QuestionHandler{service: service, SearchHandler: searchHandler, Params: params}
}

type QuestionHandler struct {
	service QuestionService
	*search.SearchHandler
	*sv.Params
}

func (h *QuestionHandler) Load(w http.ResponseWriter, r *http.Request) {
	id := sv.GetRequiredParam(w, r)
	if len(id) > 0 {
		res, err := h.service.Load(r.Context(), id)
		sv.RespondModel(w, r, res, err, h.Error, nil)
	}
}
func (h *QuestionHandler) GetByIds(w http.ResponseWriter, r *http.Request) {
	var req QuestionListRequest
	er1 := sv.Decode(w, r, &req)
	if er1 != nil {
		return
	}
	res, err := h.service.GetQuestions(r.Context(), req.IDs)
	sv.RespondModel(w, r, res, err, h.Error, nil)

}
func (h *QuestionHandler) Create(w http.ResponseWriter, r *http.Request) {
	var question Question
	er1 := sv.Decode(w, r, &question)
	if er1 == nil {
		errors, er2 := h.Validate(r.Context(), &question)
		if !sv.HasError(w, r, errors, er2, *h.Status.ValidationError, h.Error, h.Log, h.Resource, h.Action.Create) {
			res, er3 := h.service.Create(r.Context(), &question)
			sv.AfterCreated(w, r, &question, res, er3, h.Status, h.Error, h.Log, h.Resource, h.Action.Create)
		}
	}
}
func (h *QuestionHandler) Update(w http.ResponseWriter, r *http.Request) {
	var question Question
	er1 := sv.DecodeAndCheckId(w, r, &question, h.Keys, h.Indexes)
	if er1 == nil {
		errors, er2 := h.Validate(r.Context(), &question)
		if !sv.HasError(w, r, errors, er2, *h.Status.ValidationError, h.Error, h.Log, h.Resource, h.Action.Update) {
			res, er3 := h.service.Update(r.Context(), &question)
			sv.HandleResult(w, r, &question, res, er3, h.Status, h.Error, h.Log, h.Resource, h.Action.Update)
		}
	}
}
func (h *QuestionHandler) Patch(w http.ResponseWriter, r *http.Request) {
	var question Question
	r, json, er1 := sv.BuildMapAndCheckId(w, r, &question, h.Keys, h.Indexes)
	if er1 == nil {
		errors, er2 := h.Validate(r.Context(), &question)
		if !sv.HasError(w, r, errors, er2, *h.Status.ValidationError, h.Error, h.Log, h.Resource, h.Action.Patch) {
			res, er3 := h.service.Patch(r.Context(), json)
			sv.HandleResult(w, r, json, res, er3, h.Status, h.Error, h.Log, h.Resource, h.Action.Patch)
		}
	}
}
func (h *QuestionHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := sv.GetRequiredParam(w, r)
	if len(id) > 0 {
		res, err := h.service.Delete(r.Context(), id)
		sv.HandleDelete(w, r, res, err, h.Error, h.Log, h.Resource, h.Action.Delete)
	}
}
