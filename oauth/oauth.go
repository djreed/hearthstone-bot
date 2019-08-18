package oauth

import (
	"context"

	"golang.org/x/oauth2"
	clientAuth "golang.org/x/oauth2/clientcredentials"

	"github.com/djreed/hearthstone-bot/battlenet"
	"github.com/djreed/hearthstone-bot/certs"
)

func BlizzardOAuthClient(id, secret string) *battlenet.Client {
	endpoint := Endpoint("us")

	oauthCfg := &clientAuth.Config{
		ClientID:     id,
		ClientSecret: secret,
		TokenURL:     endpoint.TokenURL,
	}

	ctx := context.TODO()
	httpsClient := certs.HTTPSClient()
	ctx = context.WithValue(ctx, oauth2.HTTPClient, httpsClient)
	authClient := oauthCfg.Client(ctx)
	client := battlenet.NewClient("us", authClient)

	return client
}
