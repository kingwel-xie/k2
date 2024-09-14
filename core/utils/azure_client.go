package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type AzureUser struct {
	Id                string `json:"id,omitempty"`
	DisplayName       string `json:"displayName,omitempty" json:"displayName,omitempty"`
	GivenName         string `json:"givenName" json:"givenName,omitempty"`
	JobTitle          string `json:"jobTitle" json:"jobTitle,omitempty"`
	Mail              string `json:"mail,omitempty"`
	MobilePhone       string `json:"mobilePhone,omitempty"`
	OfficeLocation    string `json:"officeLocation,omitempty"`
	Surname           string `json:"surname,omitempty"`
	UserPrincipalName string `json:"userPrincipalName,omitempty"`
}

type AzureLightClient struct {
	tenantId   string
	clientId   string
	httpClient *http.Client
}

func NewAzureLightClientWithDefaultHttp(tenantId, clientId string) *AzureLightClient {
	return NewAzureLightClient(tenantId, clientId, http.DefaultClient)
}

func NewAzureLightClient(tenantId, clientId string, hc *http.Client) *AzureLightClient {
	return &AzureLightClient{
		tenantId:   tenantId,
		clientId:   clientId,
		httpClient: hc,
	}
}

func (c *AzureLightClient) Me(token string) (*AzureUser, error) {
	endpoint := "https://graph.microsoft.com/v1.0/me"
	headers := http.Header{}
	headers.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	req := &http.Request{
		Method: http.MethodGet,
		Header: headers,
	}

	a := AzureUser{}
	err := c.req(endpoint, req, &a)
	if err != nil {
		return nil, err
	}
	return &a, nil
}

func (c *AzureLightClient) EnumUsersDelta(secret string, fn func(*AzureUser) bool) error {
	token, err := c.AcquireTokenByClientSecret(secret, "https://graph.microsoft.com/.default")
	if err != nil {
		return err
	}

	endpoint := "https://graph.microsoft.com/v1.0/users/delta()"

	for {
		ud, err := c.usersDelta(token, endpoint)
		if err != nil {
			return err
		}

		for _, user := range ud.Users {
			fn(user)
		}
		if len(ud.NextLink) == 0 || len(ud.Users) == 0 {
			return nil
		}
		endpoint = ud.NextLink
	}
}

type usersDelta struct {
	NextLink string       `json:"@odata.nextLink,omitempty"`
	Users    []*AzureUser `json:"value"`
}

func (c *AzureLightClient) usersDelta(token string, endpoint string) (*usersDelta, error) {

	headers := http.Header{}
	headers.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	req := &http.Request{
		Method: http.MethodGet,
		Header: headers,
	}

	a := usersDelta{}
	err := c.req(endpoint, req, &a)
	if err != nil {
		return nil, err
	}
	return &a, nil
}

func (c *AzureLightClient) AcquireTokenByUsernamePassword(username, password string, scopes ...string) (string, error) {
	return c.acquireToken(func(qv *url.Values) {
		qv.Set("grant_type", "password")
		qv.Set("username", username)
		qv.Set("password", password)
	}, scopes)
}

func (c *AzureLightClient) AcquireTokenByClientSecret(secret string, scopes ...string) (string, error) {
	return c.acquireToken(func(qv *url.Values) {
		qv.Set("grant_type", "client_credentials")
		qv.Set("client_secret", secret)
	}, scopes)
}

func (c *AzureLightClient) acquireToken(build func(qv *url.Values), scopes []string) (string, error) {
	qv := url.Values{}
	qv.Set("client_id", c.clientId)
	qv.Set("client_info", "1")
	qv.Set("scope", strings.Join(append(scopes, "openid", "offline_access", "profile"), " "))

	build(&qv)

	endpoint := fmt.Sprintf("https://login.microsoftonline.com/%s/oauth2/v2.0/token", c.tenantId)

	headers := http.Header{}
	headers.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")

	enc := qv.Encode()

	req := &http.Request{
		Method:        http.MethodPost,
		Header:        headers,
		ContentLength: int64(len(enc)),
		Body:          io.NopCloser(strings.NewReader(enc)),
		GetBody: func() (io.ReadCloser, error) {
			return io.NopCloser(strings.NewReader(enc)), nil
		},
	}

	v := struct {
		AccessToken string `json:"access_token"`
	}{}

	err := c.req(endpoint, req, &v)
	if err != nil {
		return "", err
	}

	return v.AccessToken, nil
}

func (c *AzureLightClient) req(endpoint string, req *http.Request, v interface{}) error {
	u, err := url.Parse(endpoint)
	if err != nil {
		return err
	}
	req.URL = u

	res, err := c.httpClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = res.Body.Close()
	}()
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP Status %d", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}
