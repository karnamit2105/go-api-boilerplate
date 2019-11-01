package oauth2

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/vardius/go-api-boilerplate/cmd/user/internal/application/config"
	"golang.org/x/oauth2"
)

// NewConfig provides oauth2 config
func NewConfig() oauth2.Config {
	return oauth2.Config{
		ClientID:     config.Env.Auth.ClientID,
		ClientSecret: config.Env.Auth.ClientSecret,
		Scopes:       []string{"all"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  fmt.Sprintf("http://%s:%d/authorize", config.Env.Auth.Host, config.Env.HTTP.Port),
			TokenURL: fmt.Sprintf("http://%s:%d/token", config.Env.Auth.Host, config.Env.HTTP.Port),
		},
	}
}
