package getip_test

import (
	"testing"

	"github.com/hsmtkk/addhosts/getip"
	"github.com/stretchr/testify/assert"
)

func TestGetIPv4(t *testing.T) {
	want := "93.184.216.34"
	got, err := getip.New().GetIP("www.example.com", getip.IPv4)
	assert.Nil(t, err)
	assert.Equal(t, want, got)
}

func TestGetIPv6(t *testing.T) {
	want := "2606:2800:220:1:248:1893:25c8:1946"
	got, err := getip.New().GetIP("www.example.com", getip.IPv6)
	assert.Nil(t, err)
	assert.Equal(t, want, got)
}

func TestError(t *testing.T) {
	_, err := getip.New().GetIP("fuga.hoge.com", getip.IPv4)
	assert.Error(t, err)
}
