package clients

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"

	"github.com/ip-tracker-svc/config"
	"github.com/ip-tracker-svc/models"
)

// IPDetails is used to capture the details of IP Address
func IPDetails(client *http.Client, ipaddress string) (*models.IPAddressDetails, error) {

	logg().Infof("Tracking IP details for address %s", ipaddress)

	// build a request
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?access_key=%s", viper.GetString(config.IPStackURL), ipaddress, viper.GetString(config.IPStackAccessKey)), nil)
	if err != nil {
		return nil, err
	}

	// add header to accept the json response
	req.Header.Set("Accept", "application/json")

	// make request and get the response back
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// read and close the body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	ipdetails := models.IPAddressDetails{}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Failed to retrieve IP details")
	}

	// unmarshal response body
	err = json.Unmarshal(body, &ipdetails)
	if err != nil {
		return nil, err
	}

	return &ipdetails, nil
}
