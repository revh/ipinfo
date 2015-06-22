package ipinfo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type IPInfo struct {
	IP       string
	Hostname string
	City     string
	Country  string
	Loc      string
	Org      string
}

//MyIP provides information about your public ip
func MyIP() (*IPInfo, error) {
	return ForeignIP("")
}

//ForeignIP provides information about the passed ip
func ForeignIP(ip string) (*IPInfo, error) {
	if ip != "" {
		ip += "/"
	}

	response, err := http.Get(fmt.Sprintf("http://ipinfo.io/%sjson", ip))

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	var ipinfo IPInfo
	err = json.Unmarshal(contents, &ipinfo)

	if err != nil {
		return nil, err
	}

	return &ipinfo, nil
}
