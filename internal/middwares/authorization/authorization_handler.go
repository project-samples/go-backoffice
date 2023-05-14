package authorization

import (
	"context"
	log "github.com/core-go/log/zap"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type AuthorizationChecker struct {
	GetAndVerifyToken func(authorization string, secret string) (bool, string, map[string]interface{}, int64, int64, error)
	Secret            string
	Ip                string
	Authorization     string
	Author            string
	UserId            string
	CheckBlacklist    func(id string, token string, createAt time.Time) string
	Key               string
	CheckWhitelist    func(id string, token string) bool
}

func NewDefaultAuthorizationChecker(verifyToken func(string, string) (bool, string, map[string]interface{}, int64, int64, error), secret string, key string, author string, userId string, options ...string) *AuthorizationChecker {
	var authorization string
	if len(options) >= 1 {
		authorization = options[0]
	}
	return &AuthorizationChecker{Authorization: authorization, GetAndVerifyToken: verifyToken, Secret: secret, Key: key, Author: author, UserId: userId}
}

func (c *AuthorizationChecker) AuthorizationReactionChecker(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var ctx = r.Context()
		log.LogInfo(ctx, "Doing Authorization Comment Checker")
		userID := getStringValue(ctx, "userId")
		if len(userID) == 0 {
			http.Error(w, "No permission", http.StatusForbidden)
			return
		}
		pAuthor := mux.Vars(r)[c.Author]
		pUserId := mux.Vars(r)[c.UserId]
		if pAuthor == "" {
			http.Error(w, "missing authorId", http.StatusBadRequest)
			return
		}
		if pUserId == "" {
			pUserId = pAuthor
		}
		if pUserId != userID {
			http.Error(w, "no permission reaction", http.StatusForbidden)
			return
		}
		ctx = context.WithValue(ctx, "userid", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getStringValue(ctx context.Context, key string) string {
	d := ctx.Value(key)
	if d != nil {
		v, ok := d.(string)
		if ok {
			return v
		}
	}
	return ""
}
