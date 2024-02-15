package oidcmiddleware

import (
	"crypto/tls"
	"encoding/json"
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
	m.getAccessKey(c)
	userInfo := m.getUserInfo(c)

	c.Set("UserInfo", userInfo)
	c.Next()
}

func (m *middleware) getAccessKey(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
		c.Abort()
		return
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or missing Bearer token"})
		c.Abort()
		return
	}

	m.accessToken = headerParts[1]
}

func (m *middleware) getUserInfo(c *gin.Context) UserInfo {
	if m.accessToken == "" {
		c.JSON(http.StatusBadRequest, "no access token, access token is required in header")
		c.Abort()
	}

	req, err := http.NewRequest("GET", m.userInfoEndpoint, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "error creating request for OIDC validation")
		c.Abort()
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
		c.JSON(http.StatusInternalServerError, "error sending request to oidc")
		c.Abort()
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "error reading body from OIDC")
		c.Abort()
	}

	var result UserInfo
	err = json.Unmarshal(body, &result)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "binding responses")
		c.Abort()
	}

	return result
}
