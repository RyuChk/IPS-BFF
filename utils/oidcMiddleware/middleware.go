package oidcmiddleware

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type middleware struct {
	accessToken           string
	issuer                string
	clientID              string
	clientSecret          string
	userInfoEndpoint      string
	tokenEndpoint         string
	authorizationEndpoint string
}

func New(config *Config) gin.HandlerFunc {
	oidcMiddleware := newOIDCMiddleware(config)
	return func(c *gin.Context) {
		oidcMiddleware.applyOIDCMiddleWare(c)
	}
}

func newOIDCMiddleware(config *Config) *middleware {
	return &middleware{
		issuer:                config.Issuer,
		clientID:              config.ClientID,
		clientSecret:          config.ClientSecret,
		userInfoEndpoint:      config.UserinfoEndpoint,
		tokenEndpoint:         config.TokenEndpoint,
		authorizationEndpoint: config.AuthorizationEndpoint,
	}
}

func (m *middleware) applyOIDCMiddleWare(c *gin.Context) {
	if err := m.getAccessKey(c); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	userInfo, err := m.getUserInfo(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.Set("UserInfo", userInfo)
	c.Next()
}

func (m *middleware) getAccessKey(c *gin.Context) error {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		return errors.New("authorization header is missing")
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return errors.New("invalid or missing Bearer token")
	}

	m.accessToken = headerParts[1]
	return nil
}

func (m *middleware) getUserInfo(c *gin.Context) (UserInfo, error) {
	if m.accessToken == "" {
		return UserInfo{}, errors.New("no access token, access token is required in header")
	}

	req, err := http.NewRequest("GET", m.userInfoEndpoint, nil)
	if err != nil {
		return UserInfo{}, errors.New("error creating request for OIDC validation")
	}

	req.Header.Add("Authorization", "Bearer "+m.accessToken)

	// Create an HTTP client
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return UserInfo{}, errors.New("error sending request to oidc")
	}

	if resp.StatusCode != http.StatusOK {
		return UserInfo{}, errors.New("unauthorized")
	}

	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return UserInfo{}, errors.New("error reading body from OIDC")
	}

	var result UserInfo
	err = json.Unmarshal(body, &result)
	if err != nil {
		return UserInfo{}, errors.New("binding responses")
	}

	return result, nil
}
