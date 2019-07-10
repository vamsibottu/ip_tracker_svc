package clients

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func genServ(f func() string) *httptest.Server {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		content := f()

		if r.Method == http.MethodGet {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(content))
		}

		_, _ = fmt.Fprintln(w, "")
	}))

	return srv
}

func genServError(f func() int) *httptest.Server {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		status := f()
		w.Header().Set("Content-Type", "application/vnd.cia.v1+json")

		w.WriteHeader(status)
		_, _ = fmt.Fprintln(w, "")
	}))

	return srv
}
