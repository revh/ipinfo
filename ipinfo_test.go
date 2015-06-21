package ipinfo

import "testing"

func TestGoogleIP(t *testing.T) {
	ipinfo := ForeignInfo("8.8.8.8")

	if ipinfo.IP != "8.8.8.8" {
		t.Error(ipinfo)
	}
}

func TestMyRemoteIP(t *testing.T) {
	ipinfo := MyInfo()

	if ipinfo.IP == "8.8.8.8" {
		t.Error(ipinfo)
	}
}
