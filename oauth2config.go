package oauth2config

import (
	"context"
	"net/http"

	"golang.org/x/oauth2"
)

type Config interface {
	AuthCodeURL(string, ...oauth2.AuthCodeOption) string
	Client(context.Context, *oauth2.Token) *http.Client
	Exchange(context.Context, string, ...oauth2.AuthCodeOption) (*oauth2.Token, error)
	PasswordCredentialsToken(context.Context, string, string) (*oauth2.Token, error)
	TokenSource(context.Context, *oauth2.Token) oauth2.TokenSource
}
