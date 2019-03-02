package oauth2config

import (
	"context"
	"net/http"

	"golang.org/x/oauth2"
)

type FakeConfig struct {
	Config *oauth2.Config
}

func (f FakeConfig) AuthCodeURL(state string, opts ...oauth2.AuthCodeOption) string {
	return f.Config.AuthCodeURL(state, opts...)
}

func (f FakeConfig) Client(ctx context.Context, t *oauth2.Token) *http.Client {
	return f.Config.Client(ctx, t)
}

func (f FakeConfig) Exchange(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	return f.Config.Exchange(ctx, code, opts...)
}

func (f FakeConfig) PasswordCredentialsToken(ctx context.Context, username, password string) (*oauth2.Token, error) {
	return f.Config.PasswordCredentialsToken(ctx, username, password)
}

func (f FakeConfig) TokenSource(ctx context.Context, t *oauth2.Token) oauth2.TokenSource {
	return f.Config.TokenSource(ctx, t)
}
