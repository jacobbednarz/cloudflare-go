package cloudflare

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

// AccessIdentityProvider is the structure of the provider object.
type AccessIdentityProvider struct {
	ID     string      `json:"id,omitemtpy"`
	Name   string      `json:"name"`
	Type   string      `json:"type"`
	Config interface{} `json:"config"`
}

// AzureADConfiguration is the representation of the Azure AD identity
// provider.
//
// API reference: https://developers.cloudflare.com/access/configuring-identity-providers/azuread/
type AzureADConfiguration struct {
	ClientID      string `json:"client_id"`
	ClientSecret  string `json:"client_secret"`
	DirectoryID   string `json:"directory_id"`
	SupportGroups bool   `json:"support_groups"`
}

// CentrifyConfiguration is the representation of the Centrify identity
// provider.
//
// API reference: https://developers.cloudflare.com/access/configuring-identity-providers/centrify/
type CentrifyConfiguration struct {
	ClientID        string `json:"client_id"`
	ClientSecret    string `json:"client_secret"`
	CentrifyAccount string `json:"centrify_account"`
	CentrifyAppID   string `json:"centrify_app_id"`
}

// CentrifySAMLConfiguration is the representation of the Centrify
// identity provider using SAML.
//
// API reference: https://developers.cloudflare.com/access/configuring-identity-providers/saml-centrify/
type CentrifySAMLConfiguration struct {
	IssuerURL          string   `json:"issuer_url"`
	SsoTargetURL       string   `json:"sso_target_url"`
	Attributes         []string `json:"attributes"`
	EmailAttributeName string   `json:"email_attribute_name"`
	SignRequest        bool     `json:"sign_request"`
	IdpPublicCert      string   `json:"idp_public_cert"`
}

// FacebookConfiguration is the representation of the Facebook identity
// provider.
//
// API reference: https://developers.cloudflare.com/access/configuring-identity-providers/facebook-login/
type FacebookConfiguration struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

// GSuiteConfiguration is the representation of the GSuite identity
// provider.
//
// API reference: https://developers.cloudflare.com/access/configuring-identity-providers/gsuite/
type GSuiteConfiguration struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	AppsDomain   string `json:"apps_domain"`
}

// GenericOIDCConfiguration is the representation of the generic OpenID
// Connect (OIDC) connector.
//
// API reference: https://developers.cloudflare.com/access/configuring-identity-providers/generic-oidc/
type GenericOIDCConfiguration struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	AuthURL      string `json:"auth_url"`
	TokenURL     string `json:"token_url"`
	CertsURL     string `json:"certs_url"`
}

// GitHubConfiguration is the representation of the GitHub identity
// provider.
//
// API reference: https://developers.cloudflare.com/access/configuring-identity-providers/github/
type GitHubConfiguration struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

// GoogleConfiguration is the representation of the Google identity
// provider.
//
// API reference: https://developers.cloudflare.com/access/configuring-identity-providers/google/
type GoogleConfiguration struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

// JumpCloudSAMLConfiguration is the representation of the Jump Cloud
// identity provider using SAML.
//
// API reference: https://developers.cloudflare.com/access/configuring-identity-providers/jumpcloud-saml/
type JumpCloudSAMLConfiguration struct {
	IssuerURL          string   `json:"issuer_url"`
	SsoTargetURL       string   `json:"sso_target_url"`
	Attributes         []string `json:"attributes"`
	EmailAttributeName string   `json:"email_attribute_name"`
	SignRequest        bool     `json:"sign_request"`
	IdpPublicCert      string   `json:"idp_public_cert"`
}

// OktaConfiguration is the representation of the Okta identity provider.
//
// API reference: https://developers.cloudflare.com/access/configuring-identity-providers/okta/
type OktaConfiguration struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	OktaAccount  string `json:"okta_account"`
}

// OktaSAMLConfiguration is the representation of the Okta identity
// provider using SAML.
//
// API reference: https://developers.cloudflare.com/access/configuring-identity-providers/saml-okta/
type OktaSAMLConfiguration struct {
	IssuerURL          string   `json:"issuer_url"`
	SsoTargetURL       string   `json:"sso_target_url"`
	Attributes         []string `json:"attributes"`
	EmailAttributeName string   `json:"email_attribute_name"`
	SignRequest        bool     `json:"sign_request"`
	IdpPublicCert      string   `json:"idp_public_cert"`
}

// OneTimePinCnfiguration is the representation of the default One Time
// Pin identity provider.
//
// API reference: https://developers.cloudflare.com/access/configuring-identity-providers/one-time-pin/
type OneTimePinCnfiguration struct{}

// OneLoginOIDCConfiguration is the representation of the OneLogin
// OpenID connector as an identity provider.
//
// API reference: https://developers.cloudflare.com/access/configuring-identity-providers/onelogin-oidc/
type OneLoginOIDCConfiguration struct {
	ClientID        string `json:"client_id"`
	ClientSecret    string `json:"client_secret"`
	OneloginAccount string `json:"onelogin_account"`
}

// OneLoginSAMLConfiguration is the representation of the OneLogin
// identity provider using SAML.
//
// API reference: https://developers.cloudflare.com/access/configuring-identity-providers/onelogin-saml/
type OneLoginSAMLConfiguration struct {
	IssuerURL          string   `json:"issuer_url"`
	SsoTargetURL       string   `json:"sso_target_url"`
	Attributes         []string `json:"attributes"`
	EmailAttributeName string   `json:"email_attribute_name"`
	SignRequest        bool     `json:"sign_request"`
	IdpPublicCert      string   `json:"idp_public_cert"`
}

// PingSAMLConfiguration is the representation of the Ping identity
// provider using SAML.
//
// API reference: https://developers.cloudflare.com/access/configuring-identity-providers/ping-saml/
type PingSAMLConfiguration struct {
	IssuerURL          string   `json:"issuer_url"`
	SsoTargetURL       string   `json:"sso_target_url"`
	Attributes         []string `json:"attributes"`
	EmailAttributeName string   `json:"email_attribute_name"`
	SignRequest        bool     `json:"sign_request"`
	IdpPublicCert      string   `json:"idp_public_cert"`
}

// YandexConfiguration is the representation of the Yandex identity provider.
//
// API reference: https://developers.cloudflare.com/access/configuring-identity-providers/yandex/
type YandexConfiguration struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

// ADSAMLConfiguration is the representation of the Active Directory
// identity provider using SAML.
//
// API reference: https://developers.cloudflare.com/access/configuring-identity-providers/adfs/
type ADSAMLConfiguration struct {
	IssuerURL          string   `json:"issuer_url"`
	SsoTargetURL       string   `json:"sso_target_url"`
	Attributes         []string `json:"attributes"`
	EmailAttributeName string   `json:"email_attribute_name"`
	SignRequest        bool     `json:"sign_request"`
	IdpPublicCert      string   `json:"idp_public_cert"`
}

// AccessIdentityProvidersListResponse is the API response for multiple
// Access Identity Providers.
type AccessIdentityProvidersListResponse struct {
	Success  bool                     `json:"success"`
	Errors   []string                 `json:"errors"`
	Messages []string                 `json:"messages"`
	Result   []AccessIdentityProvider `json:"result"`
}

// AccessIdentityProviderListResponse is the API response for a single
// Access Identity Provider.
type AccessIdentityProviderListResponse struct {
	Success  bool                   `json:"success"`
	Errors   []string               `json:"errors"`
	Messages []string               `json:"messages"`
	Result   AccessIdentityProvider `json:"result"`
}

// AccessIdentityProviders returns all Access Identity Providers for an
// account.
//
// API reference: https://api.cloudflare.com/#access-identity-providers-list-access-identity-providers
func (api *API) AccessIdentityProviders(accountID string) ([]AccessIdentityProvider, error) {
	uri := "/accounts/" + accountID + "/access/identity_providers"

	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return []AccessIdentityProvider{}, errors.Wrap(err, errMakeRequestError)
	}

	var accessIdentityProviderResponse AccessIdentityProvidersListResponse
	err = json.Unmarshal(res, &accessIdentityProviderResponse)
	if err != nil {
		return []AccessIdentityProvider{}, errors.Wrap(err, errUnmarshalError)
	}

	return accessIdentityProviderResponse.Result, nil
}

// AccessIdentityProviderDetails returns a single Access Identity
// Provider for an account.
//
// API reference: https://api.cloudflare.com/#access-identity-providers-access-identity-providers-details
func (api *API) AccessIdentityProviderDetails(accountID, identityProviderID string) (AccessIdentityProvider, error) {
	uri := fmt.Sprintf(
		"/accounts/%s/access/identity_providers/%s",
		accountID,
		identityProviderID,
	)

	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return AccessIdentityProvider{}, errors.Wrap(err, errMakeRequestError)
	}

	var accessIdentityProviderResponse AccessIdentityProviderListResponse
	err = json.Unmarshal(res, &accessIdentityProviderResponse)
	if err != nil {
		return AccessIdentityProvider{}, errors.Wrap(err, errUnmarshalError)
	}

	return accessIdentityProviderResponse.Result, nil
}

// CreateAccessIdentityProvider creates a new Access Identity Provider.
//
// API reference: https://api.cloudflare.com/#access-identity-providers-create-access-identity-provider
func (api *API) CreateAccessIdentityProvider(accountID string, identityProviderConfiguration AccessIdentityProvider) (AccessIdentityProvider, error) {
	uri := "/accounts/" + accountID + "/access/identity_providers"

	res, err := api.makeRequest("POST", uri, identityProviderConfiguration)
	if err != nil {
		return AccessIdentityProvider{}, errors.Wrap(err, errMakeRequestError)
	}

	var accessIdentityProviderResponse AccessIdentityProviderListResponse
	err = json.Unmarshal(res, &accessIdentityProviderResponse)
	if err != nil {
		return AccessIdentityProvider{}, errors.Wrap(err, errUnmarshalError)
	}

	return accessIdentityProviderResponse.Result, nil
}

// UpdateAccessIdentityProvider updates an existing Access Identity
// Provider.
//
// API reference: https://api.cloudflare.com/#access-identity-providers-create-access-identity-provider
func (api *API) UpdateAccessIdentityProvider(accountID, identityProviderUUID string, identityProviderConfiguration AccessIdentityProvider) (AccessIdentityProvider, error) {
	uri := fmt.Sprintf(
		"/accounts/%s/access/identity_providers/%s",
		accountID,
		identityProviderUUID,
	)

	res, err := api.makeRequest("PUT", uri, identityProviderConfiguration)
	if err != nil {
		return AccessIdentityProvider{}, errors.Wrap(err, errMakeRequestError)
	}

	var accessIdentityProviderResponse AccessIdentityProviderListResponse
	err = json.Unmarshal(res, &accessIdentityProviderResponse)
	if err != nil {
		return AccessIdentityProvider{}, errors.Wrap(err, errUnmarshalError)
	}

	return accessIdentityProviderResponse.Result, nil
}

// DeleteAccessIdentityProvider deletes an Access Identity Provider.
//
// API reference: https://api.cloudflare.com/#access-identity-providers-create-access-identity-provider
func (api *API) DeleteAccessIdentityProvider(accountID, identityProviderUUID string) (AccessIdentityProvider, error) {
	uri := fmt.Sprintf(
		"/accounts/%s/access/identity_providers/%s",
		accountID,
		identityProviderUUID,
	)

	res, err := api.makeRequest("DELETE", uri, nil)
	if err != nil {
		return AccessIdentityProvider{}, errors.Wrap(err, errMakeRequestError)
	}

	var accessIdentityProviderResponse AccessIdentityProviderListResponse
	err = json.Unmarshal(res, &accessIdentityProviderResponse)
	if err != nil {
		return AccessIdentityProvider{}, errors.Wrap(err, errUnmarshalError)
	}

	return accessIdentityProviderResponse.Result, nil
}
