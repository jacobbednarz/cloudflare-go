package cloudflare

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/pkg/errors"
)

// SpectrumApp defines a single Spectrum Application.
type SpectrumApp struct {
	ID            string         `json:"id"`
	Protocol      string         `json:"protocol"`
	IPv4          bool           `json:"ipv4"`
	DNS           SpectrumAppDNS `json:"dns"`
	OriginDirect  []string       `json:"origin_direct"`
	IPFirewall    bool           `json:"ip_firewall"`
	ProxyProtocol bool           `json:"proxy_protocol"`
	TLS           string         `json:"tls"`
	CreatedOn     *time.Time     `json:"created_on"`
	ModifiedOn    *time.Time     `json:"modified_on"`
}

// SpectrumAppDNS holds the origin DNS configuration for a Spectrum
// Application.
type SpectrumAppDNS struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

// SpectrumAppDetailResponse is the structure of the detailed response
// from the API.
type SpectrumAppDetailResponse struct {
	Success  bool        `json:"success"`
	Errors   []string    `json:"errors"`
	Messages []string    `json:"messages"`
	Result   SpectrumApp `json:"result"`
}

// SpectrumApp fetches a single Spectrum application based on the ID.
//
// API reference: https://developers.cloudflare.com/spectrum/api-reference/#list-spectrum-applications
func (api *API) SpectrumApp(zoneID string, applicationID string) (SpectrumApp, error) {
	uri := fmt.Sprintf(
		"/zones/%s/spectrum/apps/%s",
		zoneID,
		applicationID,
	)

	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return SpectrumApp{}, errors.Wrap(err, errMakeRequestError)
	}

	var spectrumApp SpectrumAppDetailResponse
	err = json.Unmarshal(res, &spectrumApp)
	if err != nil {
		return SpectrumApp{}, errors.Wrap(err, errUnmarshalError)
	}

	return spectrumApp.Result, nil
}
