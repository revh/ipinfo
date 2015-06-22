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
func MyIP() *IPInfo {
	return ForeignIP("")
}

//ForeignIP provides information about the passed ip
func ForeignIP(ip string) *IPInfo {
	if ip != "" {
		ip += "/"
	}

	response, _ := http.Get(fmt.Sprintf("http://ipinfo.io/%sjson", ip))

	defer response.Body.Close()
	contents, _ := ioutil.ReadAll(response.Body)

	var ipinfo IPInfo
	json.Unmarshal(contents, &ipinfo)
	return &ipinfo
}
