package models

// IPAddressDetails holds the details of an IP address
type IPAddressDetails struct {
	IP             string  `json:"ip,omitempty"`
	Type           string  `json:"type,omitempty"`
	ContinentCode  string  `json:"continent_code,omitempty"`
	ContinentName  string  `json:"continent_name,omitempty"`
	CountryCode    string  `json:"country_code,omitempty"`
	CountryName    string  `json:"country_name,omitempty"`
	Latitude       float64 `json:"latitude,omitempty"`
	Longitude      float64 `json:"longitude,omitempty"`
	IsAcceptableIP bool    `json:"is_acceptable_ip"`
	Error          string  `json:"error,omitempty"`
}
