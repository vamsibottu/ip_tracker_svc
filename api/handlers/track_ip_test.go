package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func Test_TrackIPHandler(t *testing.T) {

	// build the request
	req, _ := http.NewRequest(http.MethodGet, "localhost:8080/v1/trackip", nil)

	req = mux.SetURLVars(req, map[string]string{
		"countries": "india,china",
		"address":   "tata.com",
	})

	// build response handler
	resp := httptest.NewRecorder()

	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success - triggered valid request",
			args: args{w: resp, r: req},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TrackIPHandler(tt.args.w, tt.args.r)
		})
	}
}
