// Package ipinfo provides info on IP address location
// using the ipinfo.io service.
package ipinfo

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var ipinfoURI = "http://ipinfo.io"

// IPInfo wraps ipinfo.io response
type IPInfo struct {
	IP       string `json:"ip"`
	Hostname string `json:"hostname"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
	Postal   string `json:"postal"`
}

// MyIP provides information about the public IP address of the client.
func MyIP() (*IPInfo, error) {
	return ForeignIP("")
}

// ForeignIP provides information about the given IP address (IPv4 or IPv6)
// if empty it behaves as MyIP()
func ForeignIP(ip string) (*IPInfo, error) {
	if ip != "" {
		ip += "/"
	}

	response, err := http.Get(fmt.Sprintf("%s/%sjson", ipinfoURI, ip))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var ipinfo IPInfo
	err = json.NewDecoder(response.Body).Decode(&ipinfo)
	if err != nil {
		return nil, err
	}

	return &ipinfo, nil
}
