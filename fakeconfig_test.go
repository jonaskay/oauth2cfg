package oauth2config

import (
	"context"
	"testing"

	"golang.org/x/oauth2"
)

func newOAuth2Config() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     "CLIENT_ID",
		ClientSecret: "CLIENT_SECRET",
		RedirectURL:  "REDIRECT_URL",
		Scopes:       []string{"scope1", "scope2"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "http://www.example.com/auth",
			TokenURL: "http://www.example.com/token",
		},
	}
}

func newFakeConfig(c *oauth2.Config) *FakeConfig {
	return &FakeConfig{Config: c}
}

func TestAuthCodeURL(t *testing.T) {
	conf := newOAuth2Config()
	fakeConf := newFakeConfig(conf)

	want := conf.AuthCodeURL(
		"foo",
		oauth2.AccessTypeOffline,
		oauth2.ApprovalForce,
	)
	got := fakeConf.AuthCodeURL(
		"foo",
		oauth2.AccessTypeOffline,
		oauth2.ApprovalForce,
	)

	if got != want {
		t.Errorf("Auth code URL %v, want %v", got, want)
	}
}

func TestExchange(t *testing.T) {
	conf := newOAuth2Config()
	fakeConf := newFakeConfig(conf)

	ctx := context.TODO()

	token, err := fakeConf.Exchange(
		ctx,
		"code",
		oauth2.AccessTypeOffline,
		oauth2.ApprovalForce,
	)
	if err != nil {
		t.Fatal(err)
	}

	want := "ACCESS_TOKEN"
	got := token.AccessToken

	if got != want {
		t.Errorf("Access token %v, want %v", got, want)
	}
}

func TestPasswordCredentialsToken(t *testing.T) {
	conf := newOAuth2Config()
	fakeConf := newFakeConfig(conf)

	ctx := context.TODO()

	token, err := fakeConf.PasswordCredentialsToken(
		ctx,
		"username",
		"password",
	)
	if err != nil {
		t.Fatal(err)
	}

	want := "ACCESS_TOKEN"
	got := token.AccessToken

	if got != want {
		t.Errorf("Access token %v, want %v", got, want)
	}
}
