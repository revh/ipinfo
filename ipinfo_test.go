package ipinfo

import "testing"

func TestForeignIP(t *testing.T) {
	ipinfo := ForeignIP("8.8.8.8")

	if ipinfo.IP != "8.8.8.8" {
		t.Error(ipinfo)
	}
}

func TestMyIP(t *testing.T) {
	ipinfo := MyIP()

	if ipinfo.IP == "8.8.8.8" {
		t.Error(ipinfo)
	}
}
