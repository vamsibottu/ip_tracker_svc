package service

import (
	"net/http"
	"strings"

	"github.com/ip-tracker-svc/clients"
	"github.com/ip-tracker-svc/models"
)

// Service holds the default configurations of the service layer
type Service struct {
	Client *http.Client
}

// TrackIP will process the request and return back isacceptable IP or not
func (s *Service) TrackIP(ipaddress string, countries []string) *models.IPAddressDetails {

	logg().Infof("started tracking IPaddress %s", ipaddress)

	if len(countries) > 0 {
		countries = strings.Split(countries[0], ",")
	}

	// pass the ipaddress to client and get the details of ipaddress
	ipdetails, err := clients.IPDetails(s.Client, ipaddress)
	if err != nil {
		logg().Error(err)
		return ipdetails
	}

	// check if the IPaddress belongs to particular country or not
	for i := range countries {
		if strings.EqualFold(strings.ToLower(ipdetails.CountryName), strings.ToLower(countries[i])) {
			ipdetails.IsAcceptableIP = true
		}
	}
	return ipdetails
}
