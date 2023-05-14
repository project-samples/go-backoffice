package app

import (
	"context"
	"go-service/internal/middwares/chain"
	"net/http"

	. "github.com/core-go/security"
	"github.com/gorilla/mux"
)

const (
	role      = "role"
	user      = "user"
	entity    = "entity"
	company   = "company"
	product   = "product"
	article   = "article"
	term      = "term"
	question  = "question"
	test      = "test"
	ticket    = "ticket"
	audit_log = "audit-log"
)

func Route(r *mux.Router, ctx context.Context, conf Config) error {
	app, err := NewApp(ctx, conf)
	if err != nil {
		return err
	}
	r.Use(app.Authorization.HandleAuthorization)
	sec := &SecurityConfig{SecuritySkip: conf.SecuritySkip, Check: app.AuthorizationChecker.Check, Authorize: app.Authorizer.Authorize}
	rcMiddleware := chain.New(app.AuthorizationChecker.Check, app.AuthorizationCommentChecker.AuthorizationReactionChecker)

	Handle(r, "/health", app.Health.Check, GET)
	Handle(r, "/authenticate", app.Authentication.Authenticate, POST)

	r.Handle("/code/{code}", app.AuthorizationChecker.Check(http.HandlerFunc(app.Code.Load))).Methods(GET)

	HandleWithSecurity(sec, r, "/privileges", app.Privileges.All, role, ActionRead, GET)
	roles := r.PathPrefix("/roles").Subrouter()
	HandleWithSecurity(sec, roles, "/search", app.Role.Search, role, ActionRead, POST, GET)
	HandleWithSecurity(sec, roles, "/{roleId}", app.Role.Load, role, ActionRead, GET)
	HandleWithSecurity(sec, roles, "", app.Role.Create, role, ActionWrite, POST)
	HandleWithSecurity(sec, roles, "/{roleId}", app.Role.Update, role, ActionWrite, PUT)
	HandleWithSecurity(sec, roles, "/{roleId}/assign", app.Role.AssignRole, role, ActionWrite, PUT)
	HandleWithSecurity(sec, roles, "/{userId}", app.Role.Patch, user, ActionWrite, PATCH)
	HandleWithSecurity(sec, roles, "/{roleId}", app.Role.Delete, role, ActionWrite, DELETE)

	HandleWithSecurity(sec, r, "/roles", app.Roles.Load, user, ActionRead, GET)
	users := r.PathPrefix("/users").Subrouter()
	HandleWithSecurity(sec, users, "", app.User.GetUserByRole, user, ActionRead, GET)
	HandleWithSecurity(sec, users, "/search", app.User.Search, user, ActionRead, GET, POST)
	HandleWithSecurity(sec, users, "/{userId}", app.User.Load, user, ActionRead, GET)
	HandleWithSecurity(sec, users, "", app.User.Create, user, ActionWrite, POST)
	HandleWithSecurity(sec, users, "/{userId}", app.User.Update, user, ActionWrite, PUT)
	HandleWithSecurity(sec, users, "/{userId}", app.User.Patch, user, ActionWrite, PATCH)
	HandleWithSecurity(sec, users, "/{userId}", app.User.Delete, user, ActionWrite, DELETE)

	entities := r.PathPrefix("/entities").Subrouter()
	HandleWithSecurity(sec, entities, "", app.Entity.Search, entity, ActionRead, GET)
	HandleWithSecurity(sec, entities, "/search", app.Entity.Search, entity, ActionRead, GET, POST)
	HandleWithSecurity(sec, entities, "/{entityId}", app.Entity.Load, entity, ActionRead, GET)
	HandleWithSecurity(sec, entities, "", app.Entity.Create, entity, ActionWrite, POST)
	HandleWithSecurity(sec, entities, "/{entityId}", app.Entity.Update, entity, ActionWrite, PUT)
	HandleWithSecurity(sec, entities, "/{entityId}", app.Entity.Patch, entity, ActionWrite, PATCH)
	HandleWithSecurity(sec, entities, "/{entityId}", app.Entity.Delete, entity, ActionWrite, DELETE)

	companies := r.PathPrefix("/companies").Subrouter()
	HandleWithSecurity(sec, companies, "", app.Company.Search, company, ActionRead, GET)
	HandleWithSecurity(sec, companies, "/search", app.Company.Search, company, ActionRead, GET, POST)
	HandleWithSecurity(sec, companies, "/{companyId}", app.Company.Load, company, ActionRead, GET)
	HandleWithSecurity(sec, companies, "", app.Company.Create, company, ActionWrite, POST)
	HandleWithSecurity(sec, companies, "/{companyId}", app.Company.Update, company, ActionWrite, PUT)
	HandleWithSecurity(sec, companies, "/{companyId}", app.Company.Patch, company, ActionWrite, PATCH)
	HandleWithSecurity(sec, companies, "/{companyId}", app.Company.Delete, company, ActionWrite, DELETE)

	products := r.PathPrefix("/products").Subrouter()
	HandleWithSecurity(sec, products, "", app.Product.Search, product, ActionRead, GET)
	HandleWithSecurity(sec, products, "/search", app.Product.Search, product, ActionRead, GET, POST)
	HandleWithSecurity(sec, products, "/{id}", app.Product.Load, product, ActionRead, GET)
	HandleWithSecurity(sec, products, "", app.Product.Create, product, ActionWrite, POST)
	HandleWithSecurity(sec, products, "/{id}", app.Product.Update, product, ActionWrite, PUT)
	HandleWithSecurity(sec, products, "/{id}", app.Product.Patch, product, ActionWrite, PATCH)
	HandleWithSecurity(sec, products, "/{id}", app.Product.Delete, product, ActionWrite, DELETE)

	articles := r.PathPrefix("/articles").Subrouter()
	HandleWithSecurity(sec, articles, "", app.Article.Search, article, ActionRead, GET)
	HandleWithSecurity(sec, articles, "/search", app.Article.Search, article, ActionRead, GET, POST)
	HandleWithSecurity(sec, articles, "/{id}", app.Article.Load, article, ActionRead, GET)
	HandleWithSecurity(sec, articles, "", app.Article.Create, article, ActionWrite, POST)
	HandleWithSecurity(sec, articles, "/{id}", app.Article.Update, article, ActionWrite, PUT)
	HandleWithSecurity(sec, articles, "/{id}", app.Article.Patch, article, ActionWrite, PATCH)
	HandleWithSecurity(sec, articles, "/{id}", app.Article.Delete, article, ActionWrite, DELETE)

	terms := r.PathPrefix("/terms").Subrouter()
	HandleWithSecurity(sec, terms, "", app.Term.Search, term, ActionRead, GET)
	HandleWithSecurity(sec, terms, "/search", app.Term.Search, term, ActionRead, GET, POST)
	HandleWithSecurity(sec, terms, "/{id}", app.Term.Load, term, ActionRead, GET)
	HandleWithSecurity(sec, terms, "", app.Term.Create, term, ActionWrite, POST)
	HandleWithSecurity(sec, terms, "/{id}", app.Term.Update, term, ActionWrite, PUT)
	HandleWithSecurity(sec, terms, "/{id}", app.Term.Patch, term, ActionWrite, PATCH)
	HandleWithSecurity(sec, terms, "/{id}", app.Term.Delete, term, ActionWrite, DELETE)

	questions := r.PathPrefix("/questions").Subrouter()
	HandleWithSecurity(sec, questions, "", app.Question.Search, question, ActionRead, GET)
	HandleWithSecurity(sec, questions, "/list", app.Question.GetByIds, question, ActionRead, POST)
	HandleWithSecurity(sec, questions, "/search", app.Question.Search, question, ActionRead, GET, POST)
	HandleWithSecurity(sec, questions, "/{id}", app.Question.Load, question, ActionRead, GET)
	HandleWithSecurity(sec, questions, "", app.Question.Create, question, ActionWrite, POST)
	HandleWithSecurity(sec, questions, "/{id}", app.Question.Update, question, ActionWrite, PUT)
	HandleWithSecurity(sec, questions, "/{id}", app.Question.Patch, question, ActionWrite, PATCH)
	HandleWithSecurity(sec, questions, "/{id}", app.Question.Delete, question, ActionWrite, DELETE)

	tests := r.PathPrefix("/tests").Subrouter()
	HandleWithSecurity(sec, tests, "", app.Test.Search, test, ActionRead, GET)
	HandleWithSecurity(sec, tests, "/search", app.Test.Search, test, ActionRead, GET, POST)
	HandleWithSecurity(sec, tests, "/{id}", app.Test.Load, test, ActionRead, GET)
	HandleWithSecurity(sec, tests, "", app.Test.Create, test, ActionWrite, POST)
	HandleWithSecurity(sec, tests, "/{id}", app.Test.Update, test, ActionWrite, PUT)
	HandleWithSecurity(sec, tests, "/{id}", app.Test.Patch, test, ActionWrite, PATCH)
	HandleWithSecurity(sec, tests, "/{id}", app.Test.Delete, test, ActionWrite, DELETE)

	tickets := r.PathPrefix("/tickets").Subrouter()
	HandleWithSecurity(sec, tickets, "", app.Ticket.Search, ticket, ActionRead, GET)
	HandleWithSecurity(sec, tickets, "/search", app.Ticket.Search, ticket, ActionRead, GET, POST)
	HandleWithSecurity(sec, tickets, "/{id}", app.Ticket.Load, ticket, ActionRead, GET)
	HandleWithSecurity(sec, tickets, "", app.Ticket.Create, ticket, ActionWrite, POST)
	HandleWithSecurity(sec, tickets, "/{id}", app.Ticket.Update, ticket, ActionWrite, PUT)
	HandleWithSecurity(sec, tickets, "/{id}", app.Ticket.Patch, ticket, ActionWrite, PATCH)
	HandleWithSecurity(sec, tickets, "/{id}", app.Ticket.Delete, ticket, ActionWrite, DELETE)

	HandleWithSecurity(sec, r, "/audit-logs", app.AuditLog.Search, audit_log, ActionRead, GET, POST)
	HandleWithSecurity(sec, r, "/audit-logs/search", app.AuditLog.Search, audit_log, ActionRead, GET, POST)

	ticketComment := "/tickets/comment"
	r.HandleFunc(ticketComment+"/search", app.SearchTicketCommentThread.Search).Methods(GET, POST)
	r.HandleFunc(ticketComment+"/{commentThreadId}/reply", app.TicketCommentReply.GetReplyComments).Methods(POST)
	r.HandleFunc(ticketComment+"/{id}/{author}/reply/{commentThreadId}", rcMiddleware.ThenFn(app.TicketCommentReply.Reply)).Methods(POST)
	r.HandleFunc(ticketComment+"/reply/{commentId}/{author}", rcMiddleware.ThenFn(app.TicketCommentReply.UpdateReply)).Methods(PATCH)
	r.HandleFunc(ticketComment+"/{commentThreadId}/reply/{commentId}/{author}", rcMiddleware.ThenFn(app.TicketCommentReply.Delete)).Methods(DELETE)
	r.HandleFunc(ticketComment+"/{id}/{author}", rcMiddleware.ThenFn(app.TicketComment.Comment)).Methods(POST)
	r.HandleFunc(ticketComment+"/{commentId}/{author}", rcMiddleware.ThenFn(app.TicketComment.Update)).Methods(PATCH)
	r.HandleFunc(ticketComment+"/{commentId}/{author}", rcMiddleware.ThenFn(app.TicketComment.Delete)).Methods(DELETE)

	r.HandleFunc("/tickets/{id}/upload", app.TicketUploadHandler.UploadFile).Methods(POST)
	r.HandleFunc("/tickets/{id}/upload", app.TicketUploadHandler.DeleteFile).Methods(DELETE)
	return nil
}

func Handle(r *mux.Router, path string, f func(http.ResponseWriter, *http.Request), methods ...string) *mux.Route {
	return r.HandleFunc(path, f).Methods(methods...)
}
func HandleWithSecurity(authorizer *SecurityConfig, r *mux.Router, path string, f func(http.ResponseWriter, *http.Request), menuId string, action int32, methods ...string) *mux.Route {
	finalHandler := http.HandlerFunc(f)
	if authorizer.SecuritySkip {
		return r.HandleFunc(path, finalHandler).Methods(methods...)
	}
	authorize := func(next http.Handler) http.Handler {
		return authorizer.Authorize(next, menuId, action)
	}
	return r.Handle(path, authorizer.Check(authorize(finalHandler))).Methods(methods...)
}
