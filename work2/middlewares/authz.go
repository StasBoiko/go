// middlewares/authz.go

package middlewares

import (
	"context"

	"go.uber.org/zap"

	"log"
	"net/http"
	"strings"
	"work2/auth"
	"work2/models"
)

type Values struct {
	m map[string]string
}

func (v Values) Get(key string) string {
	return v.m[key]
}

type Service interface {
	GetFirstUser(ctx context.Context, email string, token string, user models.User) (models.User, error)
}

type Middleware struct {
	s   Service
	log *zap.Logger
}

func NewMiddleware(s Service, log *zap.Logger) *Middleware {
	return &Middleware{s: s, log: log}
}

func (m *Middleware) Authz(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var clientToken string
		extractedToken := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(extractedToken) == 2 {
			clientToken = strings.TrimSpace(extractedToken[1])
		} else {
			m.log.Error("Malformed token")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Malformed Token"))
			return
		}
		jwtWrapper := auth.JwtWrapper{
			SecretKey: "verysecretkey",
			Issuer:    "AuthService",
		}
		claims, err := jwtWrapper.ValidateToken(clientToken)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}

		ctx := context.WithValue(r.Context(), "email", claims.Email)
		ctx = context.WithValue(ctx, "token", clientToken)

		var user models.User
		user, err = m.s.GetFirstUser(ctx, claims.Email, clientToken, user)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Unauthorized"))
			return
		}
		models.HttpReqs.WithLabelValues(user.Name).Add(1)
		ctx = context.WithValue(ctx, "user_id", user.ID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
