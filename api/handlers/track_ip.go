package handlers

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/ip-tracker-svc/service"
)

// TrackIPHandler is used to receive a request and send response back
func TrackIPHandler(w http.ResponseWriter, r *http.Request) {
	// receive params from URL
	params := mux.Vars(r)
	ipaddress := params["address"]

	// validate white listed countries
	countries := r.URL.Query()["countries"]

	logg().Infof("Received request and sent to service layer")

	// comnfigure the service layer
	s := service.Service{Client: newHTTPClient()}

	// send request to service layer, get the output
	ipdetails := s.TrackIP(ipaddress, countries)

	// encode the response
	rawJSON, err := json.Marshal(ipdetails)
	if err != nil {
		logg().Error(err)
	}

	fmt.Fprintln(w, string(rawJSON))
}

// newHTTPClient returns new httpClient
// This is configurable client, we can add more timeout if needed
func newHTTPClient() *http.Client {

	netTransport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		DialContext: (&net.Dialer{
			Timeout: time.Duration(10 * time.Second),
		}).DialContext,
	}

	client := &http.Client{
		Timeout:   time.Duration(30 * time.Second),
		Transport: netTransport,
	}
	return client
}
