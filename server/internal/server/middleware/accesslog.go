package middleware

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func AccessLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//fmt.Println("AccessLog->",r.RequestURI)
		log.Debugf("=>%s ,m=%s,q = %s, ip=%s ,ua = %s", r.Method,r.RequestURI, r.URL.RawQuery, r.RemoteAddr, r.Header.Get("User-Agent"))
		next.ServeHTTP(w, r)
	})
}
