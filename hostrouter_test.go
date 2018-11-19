package hostrouter

import (
	"net/http"
	"fmt"
	"github.com/go-chi/chi"
	//"github.com/go-chi/hostrouter"   // Uncomment this line to reproduce
	"testing"
)

func TestHostRouterPortSupport(t *testing.T) {
	ehr := chi.NewRouter()

	ehr.Get("/hello", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, fmt.Sprintf("hello"))
	})

	ihr := chi.NewRouter()
	ihr.Get("/world", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, fmt.Sprintf("world"))
	})

	hr := New()  // comment out this line with the line below to reproduce
	//hr := hostrouter.New()
	hr.Map("api.hello.com", ehr)
	hr.Map("api.world.com", ihr)

	r := chi.NewRouter()
	r.Mount("/", hr)
	http.ListenAndServe(":3333", r)
}