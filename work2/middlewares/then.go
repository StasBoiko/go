// middlewares/authz.go

package middlewares

//import (
//	"fmt"
//	"log"
//	"net/http"
//
//	"github.com/DataDog/datadog-go/statsd"
//)
//
//// StatsdClient is an HTTP middleware that sends counters and timers for each
//// API endpoint to the local statsd daemon.
//type StatsdClient struct{}
//
//// Then middleware
//func (s StatsdClient) Then(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		c, err := statsd.New("127.0.0.1:8125")
//		if err != nil {
//			log.Fatal(err)
//		}
//		defer c.Close()
//
//		c.Increment(fmt.Sprintf("%s.%s", r.Method, r.URL.Path[1:]))
//		next.ServeHTTP(w, r)
//	})
//}
