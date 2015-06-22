package ipinfo

import "testing"

func TestForeignIP(t *testing.T) {

	if _, err := ForeignIP("8.8.8.8"); err != nil {
		t.Error(err)
	}
}

func TestMyIP(t *testing.T) {

	if _, err := MyIP(); err != nil {
		t.Error(err)
	}
}
