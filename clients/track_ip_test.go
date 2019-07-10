package clients

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"

	"github.com/spf13/viper"

	"github.com/ip-tracker-svc/config"
	"github.com/ip-tracker-svc/models"
)

func Test_IPDetails(t *testing.T) {

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

	marshalErrReq := genServ(func() string { return "this is an invalid data" })
	defer marshalErrReq.Close()

	type args struct {
		client    *http.Client
		ipaddress string
		URL       string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.IPAddressDetails
		wantErr bool
	}{
		{
			name:    "Success - got the ipaddress details",
			args:    args{client: successReq.Client(), ipaddress: "tata.com", URL: successReq.URL},
			want:    expectedipdetails,
			wantErr: false,
		},
		{
			name:    "fail - failed to retrieve IP details for given ip address",
			args:    args{client: srvErr.Client(), ipaddress: "invalidip.com", URL: srvErr.URL},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "fail - invalid IP details",
			args:    args{client: marshalErrReq.Client(), ipaddress: "invaliddata.com", URL: marshalErrReq.URL},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			viper.Set(config.IPStackURL, tt.args.URL)
			got, err := IPDetails(tt.args.client, tt.args.ipaddress)
			if (err != nil) != tt.wantErr {
				t.Errorf("IPDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IPDetails() = %v, want %v", got, tt.want)
			}
		})
	}
}
