package app

import (
	"context"
	"net/http"

	. "github.com/core-go/security"
	"github.com/gorilla/mux"
)

const (
	role      = "role"
	user      = "user"
	entity    = "entity"
	company   = "company"
	audit_log = "audit-log"
)

func Route(r *mux.Router, ctx context.Context, conf Config) error {
	app, err := NewApp(ctx, conf)
	if err != nil {
		return err
	}
	r.Use(app.Authorization.HandleAuthorization)
	sec := &SecurityConfig{SecuritySkip: conf.SecuritySkip, Check: app.AuthorizationChecker.Check, Authorize: app.Authorizer.Authorize}

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

	HandleWithSecurity(sec, r, "/audit-logs", app.AuditLog.Search, audit_log, ActionRead, GET, POST)
	HandleWithSecurity(sec, r, "/audit-logs/search", app.AuditLog.Search, audit_log, ActionRead, GET, POST)
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
