package cloudflare

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSpectrumApp(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "GET", "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"result": {
				"id": "f68579455bd947efb65ffa1bcf33b52c",
				"protocol": "tcp/22",
				"ipv4": true,
				"dns": {
					"type": "CNAME",
					"name": "spectrum.example.com"
				},
				"origin_direct": [
					"tcp://192.0.2.1:22"
				],
				"ip_firewall": true,
				"proxy_protocol": false,
				"tls": "off",
				"created_on": "2018-03-28T21:25:55.643771Z",
				"modified_on": "2018-03-28T21:25:55.643771Z"
			},
			"success": true,
			"errors": [],
			"messages": []
		}`)
	}

	mux.HandleFunc("/zones/01a7362d577a6c3019a474fd6f485823/spectrum/apps/f68579455bd947efb65ffa1bcf33b52c", handler)
	createdOn, _ := time.Parse(time.RFC3339, "2018-03-28T21:25:55.643771Z")
	modifiedOn, _ := time.Parse(time.RFC3339, "2018-03-28T21:25:55.643771Z")
	want := SpectrumApp{
		ID:         "f68579455bd947efb65ffa1bcf33b52c",
		CreatedOn:  &createdOn,
		ModifiedOn: &modifiedOn,
		Protocol:   "tcp/22",
		IPv4:       true,
		DNS: SpectrumAppDNS{
			Name: "spectrum.example.com",
			Type: "CNAME",
		},
		OriginDirect:  []string{"tcp://192.0.2.1:22"},
		IPFirewall:    true,
		ProxyProtocol: false,
		TLS:           "off",
	}

	actual, err := client.SpectrumApp("01a7362d577a6c3019a474fd6f485823", "f68579455bd947efb65ffa1bcf33b52c")
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}
