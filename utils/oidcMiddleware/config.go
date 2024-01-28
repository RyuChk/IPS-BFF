package oidcmiddleware

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog/log"
)

type Config struct {
	ConfigURL    string `envconfig:"OIDC_MIDDLEWARE_CONFIG_URL"`
	ClientID     string `envconfig:"OIDC_MIDDLEWARRE_CLIENT_ID" required:"true"`
	ClientSecret string `envconfig:"OIDC_MIDDLEWARE_CLIENT_SECRET" required:"true"`

	//Can be loaded from config URL
	Issuer                string `envconfig:"OIDC_MIDDLEWARE_ISSUER" json:"issuer"`
	AuthorizationEndpoint string `envconfig:"OIDC_MIDDLEWARE_AUTH_ENDPOINT" json:"authorization_endpoint"`
	TokenEndpoint         string `envconfig:"OIDC_MIDDLEWARE_TOKEN_ENDPOINT" json:"token_endpoint"`
	UserinfoEndpoint      string `envconfig:"OIDC_MIDDLEWARE_USERINFO_ENDPOINT" json:"userinfo_endpoint"`
	EndSessionEndpoint    string `envconfig:"OIDC_MIDDLEWARE_END_SESSION_ENDPOINT" json:"end_session_endpoint"`
}

func LoadConfig() (cfg *Config) {
	config := &Config{}
	prefix := "OIDC_MIDDLEWARRE"
	if val, ok := os.LookupEnv("OIDC_MIDDLEWARRE_PREFIX"); ok {
		prefix = val
	}
	envconfig.MustProcess(prefix, config)
	config.validateConfig()
	config.updateConfig()
	return config
}

func (c *Config) updateConfig() {
	if c.ConfigURL == "" {
		return
	}

	req, _ := http.NewRequest("GET", c.ConfigURL, nil)
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Error().Msg("error loading config from config URL fallback to load config")
		c.ConfigURL = ""
		c.validateConfig()
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error().Msg("error loading config from config URL fallback to load config")
		c.ConfigURL = ""
		c.validateConfig()
		return
	}

	var temp Config
	err = json.Unmarshal(body, &temp)
	if err != nil {
		log.Error().Msg("error loading config from config URL fallback to load config")
		c.ConfigURL = ""
		c.validateConfig()
		return
	}

	c.Issuer = temp.Issuer
	c.TokenEndpoint = temp.TokenEndpoint
	c.AuthorizationEndpoint = temp.AuthorizationEndpoint
	c.UserinfoEndpoint = temp.UserinfoEndpoint
	c.EndSessionEndpoint = temp.EndSessionEndpoint
}

type condition struct {
	key        string
	constraint bool
}

func (c *Config) validateConfig() {
	conditions := []condition{
		{
			"OIDC_MIDDLEWARRE_CLIENT_ID",
			c.UserinfoEndpoint == "",
		},
		{
			"OIDC_MIDDLEWARE_CLIENT_SECRET",
			c.TokenEndpoint == "",
		},
	}

	if c.ConfigURL == "" {
		conditions = append(conditions,
			condition{
				"OIDC_MIDDLEWARE_ISSUER",
				c.Issuer == "",
			},
			condition{
				"OIDC_MIDDLEWARE_USERINFO_ENDPOINT",
				c.UserinfoEndpoint == "",
			},
			condition{
				"OIDC_MIDDLEWARE_TOKEN_ENDPOINT",
				c.TokenEndpoint == "",
			})
	}

	for _, v := range conditions {
		if v.constraint {
			panic(fmt.Sprintf("error missing config : %s", v.key))
		}
	}
}
