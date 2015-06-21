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

//MyInfo provides information about your public ip
func MyInfo() *IPInfo {
	return ForeignInfo("")
}

//ForeignInfo provides information about the passed ip
func ForeignInfo(ip string) *IPInfo {
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
