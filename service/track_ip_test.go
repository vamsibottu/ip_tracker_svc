package service

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"

	"github.com/ip-tracker-svc/config"
	"github.com/spf13/viper"

	"github.com/ip-tracker-svc/models"
)

func Test_TrackIP(t *testing.T) {

	expectedipdetails := &models.IPAddressDetails{
		IP:             "104.211.153.237",
		Type:           "ipv4",
		ContinentCode:  "AS",
		ContinentName:  "Asia",
		CountryCode:    "IN",
		CountryName:    "India",
		Latitude:       19.076000213623047,
		Longitude:      72.87770080566406,
		IsAcceptableIP: true,
	}

	successJSON, _ := json.Marshal(expectedipdetails)

	successReq := genServ(func() string { return string(successJSON) })
	defer successReq.Close()

	srvErr := genServError(func() int { return http.StatusNotFound })
	defer srvErr.Close()

	type args struct {
		ipaddress string
		countries []string
		client    *http.Client
		URL       string
	}
	tests := []struct {
		name string
		args args
		want *models.IPAddressDetails
	}{
		{
			name: "Success - Tracked IP Address",
			args: args{ipaddress: "tata.com", countries: []string{"india", "china"}, client: successReq.Client(), URL: successReq.URL},
			want: expectedipdetails,
		},
		{
			name: "fail - Invalid Client configurations",
			args: args{ipaddress: "tata.com", countries: []string{"india", "china"}, client: srvErr.Client(), URL: srvErr.URL},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			viper.Set(config.IPStackURL, tt.args.URL)
			s := Service{Client: tt.args.client}
			if got := s.TrackIP(tt.args.ipaddress, tt.args.countries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TrackIP() = %v, want %v", got, tt.want)
			}
		})
	}
}
