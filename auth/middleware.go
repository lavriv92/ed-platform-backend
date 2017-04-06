package auth

import (
	"log"
	"strings"
	"net/http"
)

func AuthenticateMiddleware(inner http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
		s := strings.Split(r.Header.Get("Authorization"), " ")
		if len(s) != 2 {
			http.Error(w, "Not authorized", 401)
			return
		}
		userId, err := ValidateToken(s[1])
		if err != nil {
			http.Error(w, "Not authorized", 401)
			return
		}
		log.Printf("userid %s", userId)
		inner.ServeHTTP(w, r)
	}
}