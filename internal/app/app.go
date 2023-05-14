package app

import (
	"context"
	sv "github.com/core-go/core"
	"github.com/core-go/reaction/commentthread"
	"github.com/core-go/search"
	"github.com/core-go/storage"
	"github.com/lib/pq"
	"go-service/internal/uploads"
	"reflect"

	"github.com/core-go/auth"
	ah "github.com/core-go/auth/handler"
	. "github.com/core-go/auth/mock"
	as "github.com/core-go/auth/sql"
	"github.com/core-go/code"
	"github.com/core-go/core/authorization"
	"github.com/core-go/core/shortid"
	"github.com/core-go/core/unique"
	v10 "github.com/core-go/core/v10"
	. "github.com/core-go/health"
	log "github.com/core-go/log/zap"
	"github.com/core-go/search/convert"
	"github.com/core-go/search/query"
	"github.com/core-go/search/template"
	. "github.com/core-go/security"
	. "github.com/core-go/security/jwt"
	. "github.com/core-go/security/sql"
	q "github.com/core-go/sql"

	a "go-service/internal/article"
	"go-service/internal/audit-log"
	c "go-service/internal/company"
	e "go-service/internal/entity"
	p "go-service/internal/product"
	qu "go-service/internal/question"
	r "go-service/internal/role"
	tm "go-service/internal/term"
	t "go-service/internal/test"
	tk "go-service/internal/ticket"
	u "go-service/internal/user"

	commentthreadreply "github.com/core-go/reaction/commentthread/comment"
	muxcomment "github.com/core-go/reaction/commentthread/comment/mux"
	muxcommentthread "github.com/core-go/reaction/commentthread/mux"
	commentthreadsearch "github.com/core-go/reaction/commentthread/search"

	"github.com/core-go/storage/google"
	authComment "go-service/internal/middwares/authorization"
)

type ApplicationContext struct {
	SkipSecurity                bool
	Health                      *Handler
	Authorization               *authorization.Handler
	AuthorizationChecker        *AuthorizationChecker
	AuthorizationCommentChecker *authComment.AuthorizationChecker
	Authorizer                  *Authorizer
	Authentication              *ah.AuthenticationHandler
	Privileges                  *ah.PrivilegesHandler
	Code                        *code.Handler
	Roles                       *code.Handler
	Role                        r.RoleTransport
	User                        u.UserTransport
	Entity                      e.EntityTransport
	Company                     c.CompanyTransport
	Product                     p.ProductTransport
	Article                     a.ArticleTransport
	Term                        tm.TermTransport
	Question                    qu.QuestionTransport
	Test                        t.TestTransport
	Ticket                      tk.TicketTransport
	SearchTicketCommentThread   *commentthreadsearch.CommentThreadSearchHandler
	TicketCommentReply          muxcomment.CommentHandler
	TicketComment               muxcommentthread.CommentThreadHandler
	TicketUploadHandler         *uploads.Handler
	AuditLog                    *audit.AuditLogHandler
}

func NewApp(ctx context.Context, conf Config) (*ApplicationContext, error) {
	db, er0 := q.Open(conf.DB)
	if er0 != nil {
		return nil, er0
	}
	sqlHealthChecker := q.NewHealthChecker(db)
	var healthHandler *Handler

	logError := log.LogError
	generateId := shortid.Generate
	cloudService, _ := CreateCloudService(ctx, conf)
	var writeLog func(ctx context.Context, resource string, action string, success bool, desc string) error

	if conf.AuditLog.Log {
		auditLogDB, er1 := q.Open(conf.AuditLog.DB)
		if er1 != nil {
			return nil, er1
		}
		logWriter := q.NewActionLogWriter(auditLogDB, "auditLog", conf.AuditLog.Config, conf.AuditLog.Schema, generateId)
		writeLog = logWriter.Write
		auditLogHealthChecker := q.NewSqlHealthChecker(auditLogDB, "audit_log")
		healthHandler = NewHandler(sqlHealthChecker, auditLogHealthChecker)
	} else {
		healthHandler = NewHandler(sqlHealthChecker)
	}
	modelStatus := sv.InitializeStatus(conf.ModelStatus)
	action := sv.InitializeAction(conf.Action)
	buildParam := q.GetBuild(db)
	validator := v10.NewValidator()
	sqlPrivilegeLoader := NewPrivilegeLoader(db, conf.Sql.PermissionsByUser)

	userId := conf.Tracking.User
	tokenService := NewTokenService()
	authorizationHandler := authorization.NewHandler(tokenService.GetAndVerifyToken, conf.Auth.Token.Secret)
	authorizationChecker := NewDefaultAuthorizationChecker(tokenService.GetAndVerifyToken, conf.Auth.Token.Secret, userId)
	authorizer := NewAuthorizer(sqlPrivilegeLoader.Privilege, true, userId)

	authStatus := auth.InitStatus(conf.Auth.Status)
	ldapAuthenticator, er2 := NewDAPAuthenticatorByConfig(conf.Ldap, authStatus)
	if er2 != nil {
		return nil, er2
	}
	userInfoService, er3 := as.NewUserRepository(db, conf.Auth.Query, conf.Auth.DB, conf.Auth.UserStatus)
	if er3 != nil {
		return nil, er3
	}
	privilegeLoader, er4 := as.NewSqlPrivilegesLoader(db, conf.Sql.PrivilegesByUser, 1, true)
	if er4 != nil {
		return nil, er4
	}
	authenticator := auth.NewBasicAuthenticator(authStatus, ldapAuthenticator.Authenticate, userInfoService, tokenService.GenerateToken, conf.Auth.Token, conf.Auth.Payload, privilegeLoader.Load)
	authenticationHandler := ah.NewAuthenticationHandler(authenticator.Authenticate, authStatus.Error, authStatus.Timeout, logError, writeLog)
	authorizationCommentHandler := authComment.NewDefaultAuthorizationChecker(tokenService.GetAndVerifyToken, conf.Auth.Token.Secret, conf.Auth.Payload.Id, "author", "userId", "Authorization")

	privilegeReader, er5 := as.NewPrivilegesReader(db, conf.Sql.Privileges)
	if er5 != nil {
		return nil, er5
	}
	privilegeHandler := ah.NewPrivilegesHandler(privilegeReader.Privileges)

	// codeLoader := code.NewDynamicSqlCodeLoader(db, "select code, name, status as text from codeMaster where master = ? and status = 'A'", 1)
	codeLoader := code.NewSqlCodeLoader(db, "codeMaster", conf.Code.Loader)
	codeHandler := code.NewCodeHandlerByConfig(codeLoader.Load, conf.Code.Handler, logError)

	templates, err := template.LoadTemplates(template.Trim, "configs/query.xml")
	if err != nil {
		return nil, err
	}
	// rolesLoader := code.NewDynamicSqlCodeLoader(db, "select roleName as name, roleId as id, status as code from roles where status = 'A'", 0)
	rolesLoader := code.NewSqlCodeLoader(db, "roles", conf.Role.Loader)
	rolesHandler := code.NewCodeHandlerByConfig(rolesLoader.Load, conf.Role.Handler, logError)

	roleType := reflect.TypeOf(r.Role{})
	queryRole, err := template.UseQuery(conf.Template, query.UseQuery(db, "roles", roleType, buildParam), "role", templates, &roleType, convert.ToMap, buildParam)
	roleSearchBuilder, err := q.NewSearchBuilder(db, roleType, queryRole)
	// roleValidator := user.NewRoleValidator(db, conf.Sql.Role.Duplicate, validator.validateFileName)
	roleValidator := unique.NewUniqueFieldValidator(db, "roles", "rolename", reflect.TypeOf(r.Role{}), validator.Validate)
	roleRepository, er6 := r.NewRoleAdapter(db, conf.Sql.Role.Check)
	if er6 != nil {
		return nil, er6
	}
	roleService := r.NewRoleService(roleRepository)
	generateRoleId := shortid.Func(conf.AutoRoleId)
	roleHandler := r.NewRoleHandler(roleSearchBuilder.Search, roleService, conf.Writer, logError, generateRoleId, roleValidator.Validate, conf.Tracking, writeLog)

	userType := reflect.TypeOf(u.User{})
	queryUser, err := template.UseQuery(conf.Template, query.UseQuery(db, "users", userType, buildParam), "user", templates, &userType, convert.ToMap, buildParam)
	if err != nil {
		return nil, err
	}
	userSearchBuilder, err := q.NewSearchBuilder(db, userType, queryUser)
	if err != nil {
		return nil, err
	}
	// userValidator := user.NewUserValidator(db, conf.Sql.User, validator.validateFileName)
	userValidator := unique.NewUniqueFieldValidator(db, "users", "username", reflect.TypeOf(u.User{}), validator.Validate)
	userRepository, er7 := u.NewUserRepository(db)
	if er7 != nil {
		return nil, er7
	}
	userService := u.NewUserService(userRepository)
	generateUserId := shortid.Func(conf.AutoUserId)
	userHandler := u.NewUserHandler(userSearchBuilder.Search, userService, conf.Writer, logError, generateUserId, userValidator.Validate, conf.Tracking, writeLog)

	entityType := reflect.TypeOf(e.Entity{})
	queryEntity, err := template.UseQuery(conf.Template, query.UseQuery(db, "entities", entityType, buildParam), "entity", templates, &entityType, convert.ToMap, buildParam)
	if err != nil {
		return nil, err
	}
	entitySearchBuilder, err := q.NewSearchBuilder(db, entityType, queryEntity)
	if err != nil {
		return nil, err
	}
	entityValidator := unique.NewUniqueFieldValidator(db, "entities", "entityname", reflect.TypeOf(e.Entity{}), validator.Validate)
	entityRepository, er7 := e.NewEntityRepository(db)
	if er7 != nil {
		return nil, er7
	}
	entityService := e.NewentitieService(entityRepository)
	generateEntityId := shortid.Func(conf.AutoEntityId)
	entityHandler := e.NewEntityHandler(entitySearchBuilder.Search, entityService, conf.Writer, logError, generateEntityId, entityValidator.Validate, conf.Tracking, writeLog)

	companyType := reflect.TypeOf(c.Company{})
	queryCompany, err := template.UseQuery(conf.Template, query.UseQuery(db, "companies", companyType, buildParam), "company", templates, &companyType, convert.ToMap, buildParam)
	if err != nil {
		return nil, err
	}
	companySearchBuilder, err := q.NewSearchBuilder(db, companyType, queryCompany)
	if err != nil {
		return nil, err
	}
	companyValidator := unique.NewUniqueFieldValidator(db, "companies", "companyname", reflect.TypeOf(c.Company{}), validator.Validate)
	companyRepository, er7 := c.NewCompanyRepository(db)
	if er7 != nil {
		return nil, er7
	}
	companyService := c.NewCompanyService(companyRepository)
	generateCompanyId := shortid.Func(conf.AutoCompanyId)
	companyHandler := c.NewCompanyHandler(companySearchBuilder.Search, companyService, conf.Writer, logError, generateCompanyId, companyValidator.Validate, conf.Tracking, writeLog)

	articleType := reflect.TypeOf(a.Article{})
	queryArticle, err := template.UseQuery(conf.Template, query.UseQuery(db, "article", companyType, buildParam), "article", templates, &articleType, convert.ToMap, buildParam)
	if err != nil {
		return nil, err
	}
	articleSearchBuilder, err := q.NewSearchBuilderWithArray(db, articleType, queryArticle, pq.Array)
	if err != nil {
		return nil, err
	}
	articlesRepository := a.NewArticleRepository(db, pq.Array)
	articleService := a.NewArticleService(articlesRepository, generateId)
	generateArticleId := shortid.Func(conf.AutoArticleId)
	articleHandler := a.NewArticleHandler(articleSearchBuilder.Search, articleService, generateArticleId, modelStatus, logError, validator.Validate, conf.Tracking, &action, writeLog)

	termType := reflect.TypeOf(tm.Term{})
	queryTerm, err := template.UseQuery(conf.Template, query.UseQuery(db, "term", companyType, buildParam), "term", templates, &termType, convert.ToMap, buildParam)
	if err != nil {
		return nil, err
	}
	termSearchBuilder, err := q.NewSearchBuilderWithArray(db, termType, queryTerm, pq.Array)
	if err != nil {
		return nil, err
	}
	termsRepository, err := tm.NewTermRepository(db, pq.Array)
	if err != nil {
		return nil, err
	}
	termService := tm.NewTermService(termsRepository, generateId)
	termHandler := tm.NewTermHandler(termSearchBuilder.Search, termService, nil, modelStatus, logError, validator.Validate, conf.Tracking, &action, writeLog)

	productType := reflect.TypeOf(p.Product{})
	queryProduct, err := template.UseQuery(conf.Template, query.UseQuery(db, "product", companyType, buildParam), "product", templates, &productType, convert.ToMap, buildParam)
	if err != nil {
		return nil, err
	}
	productSearchBuilder, err := q.NewSearchBuilderWithArray(db, productType, queryProduct, pq.Array)
	if err != nil {
		return nil, err
	}
	productsRepository, err := p.NewProductRepository(db, pq.Array)
	if err != nil {
		return nil, err
	}
	productService := p.NewProductService(productsRepository, generateId)
	productHandler := p.NewProductHandler(productSearchBuilder.Search, productService, nil, modelStatus, logError, validator.Validate, conf.Tracking, &action, writeLog)

	questionType := reflect.TypeOf(qu.Question{})
	questionQuery := query.UseQuery(db, "questions", questionType)
	questionSearchBuilder, err := q.NewSearchBuilderWithArray(db, questionType, questionQuery, pq.Array)
	if err != nil {
		return nil, err
	}
	questionRepository, err := qu.NewQuestionRepository(db, pq.Array)
	if err != nil {
		return nil, err
	}
	questionService := qu.NewQuestionService(questionRepository, generateId)
	questionHandler := qu.NewQuestionHandler(questionSearchBuilder.Search, questionService, modelStatus, logError, validator.Validate, &action)

	testType := reflect.TypeOf(t.Test{})
	testQuery := query.UseQuery(db, "tests", testType)
	testSearchBuilder, err := q.NewSearchBuilderWithArray(db, testType, testQuery, pq.Array)
	if err != nil {
		return nil, err
	}
	testRepository, err := t.NewTestRepository(db, pq.Array)
	if err != nil {
		return nil, err
	}
	testService := t.NewTestService(testRepository, generateId)
	testHandler := t.NewTestHandler(testSearchBuilder.Search, testService, modelStatus, logError, validator.Validate, &action)

	ticketType := reflect.TypeOf(tk.Ticket{})
	ticketQuery := query.UseQuery(db, "tickets", ticketType)
	ticketSearchBuilder, err := q.NewSearchBuilderWithArray(db, ticketType, ticketQuery, pq.Array)
	if err != nil {
		return nil, err
	}
	ticketRepository, err := tk.NewTicketRepository(db, pq.Array)
	if err != nil {
		return nil, err
	}
	ticketService := tk.NewTicketService(ticketRepository, generateId)
	ticketHandler := tk.NewTicketHandler(ticketSearchBuilder.Search, ticketService, modelStatus, logError, validator.Validate, &action)

	ticketCommentThreadType := reflect.TypeOf(commentthread.CommentThread{})
	ticketCommentQuery, err := template.UseQueryWithArray(conf.Template, nil, "ticket_comment", templates, &ticketCommentThreadType, convert.ToMap, buildParam)
	if err != nil {
		return nil, err
	}

	queryUserInfo := commentthreadsearch.NewQueryInfo(db, "users", "imageURL", "userid", "username", "displayname", pq.Array)
	ticketCommentBuilder, err := commentthreadsearch.NewCommentThreadSearchService(db, ticketCommentQuery, pq.Array, queryUserInfo.Load, q.BuildFromQuery, search.GetOffset)
	if err != nil {
		return nil, err
	}

	ticketCommentSearchHandler := commentthreadsearch.NewSearchCommentThreadHandler(ticketCommentBuilder)
	commentThreadService := commentthread.NewCommentThreadService(db, pq.Array,
		"ticketcommentthread", "commentId", "id", "author",
		"histories", "comment", "time", "updatedat",
		"ticketreplycomment", "commentid", "commentthreadid",
		"ticketcommentthreadinfo", "commentid",
		"ticketreplycommentinfo", "commentid",
		"ticketcommentthreadreaction", "commentid",
		"ticketreplycommentreaction", "commentId")
	ticketCommentHandler := muxcommentthread.NewCommentThreadHandler(commentThreadService, shortid.Generate, "commentId", "author", "id")

	tkUserInforeply := commentthreadreply.NewQueryInfo(db, "users", "imageURL", "userid", "username", "displayname", pq.Array)
	locationCommentThreadReplyService := commentthreadreply.NewCommentService(db, "ticketreplycomment", "commentId", "author", "id", "updatedat", "comment", "userId", "time", "histories", "commentthreadId", "reaction",
		"ticketreplycommentreaction", "commentId", "users", "id", "username", "imageurl", "ticketcommentthreadinfo", "usefulcount",
		"commentId", "ticketcommentthreadinfo", "commentId",
		"replycount", "usefulcount", tkUserInforeply.Load, pq.Array)
	ticketCommentReplyHandler := muxcomment.NewCommentHandler(locationCommentThreadReplyService, "commentThreadId", "userId", "author", "id", "commentId", generateId)

	uploadTicketRepository := uploads.NewRepository(db, "tickets", uploads.FieldColumn{Id: "id", File: "attachments"}, pq.Array)

	ticketUploadService := uploads.NewUploadService(uploadTicketRepository, cloudService, conf.Provider, conf.GeneralDirectory, conf.KeyFile, conf.Storage.Directory)

	ticketUploadHandler := uploads.NewHandler(ticketUploadService, logError, conf.KeyFile, generateId, conf.AllowedExtensions)

	reportDB, er8 := q.Open(conf.AuditLog.DB)
	if er8 != nil {
		return nil, er8
	}
	auditLogQuery, er9 := audit.NewAuditLogQuery(reportDB)
	if er9 != nil {
		return nil, er9
	}
	auditLogHandler := audit.NewAuditLogHandler(auditLogQuery, logError, writeLog)

	app := &ApplicationContext{
		Health:                      healthHandler,
		SkipSecurity:                conf.SecuritySkip,
		Authorization:               authorizationHandler,
		AuthorizationChecker:        authorizationChecker,
		AuthorizationCommentChecker: authorizationCommentHandler,
		Authorizer:                  authorizer,
		Authentication:              authenticationHandler,
		Privileges:                  privilegeHandler,
		Code:                        codeHandler,
		Roles:                       rolesHandler,
		Role:                        roleHandler,
		User:                        userHandler,
		Entity:                      entityHandler,
		Company:                     companyHandler,
		Product:                     productHandler,
		Article:                     articleHandler,
		Term:                        termHandler,
		Question:                    questionHandler,
		Test:                        testHandler,
		Ticket:                      ticketHandler,
		SearchTicketCommentThread:   ticketCommentSearchHandler,
		TicketComment:               ticketCommentHandler,
		TicketCommentReply:          ticketCommentReplyHandler,
		TicketUploadHandler:         ticketUploadHandler,
		AuditLog:                    auditLogHandler,
	}
	return app, nil
}

func CreateCloudService(ctx context.Context, config Config) (storage.StorageService, error) {
	return google.NewGoogleStorageServiceWithCredentials(ctx, []byte(config.GoogleCredentials), config.Storage)
}
