package oauth

import (
	"context"

	clientAuth "golang.org/x/oauth2/clientcredentials"

	"github.com/djreed/hearthstone-bot/battlenet"
)

func BlizzardOAuthClient(id, secret string) *battlenet.Client {
	endpoint := Endpoint("us")

	oauthCfg := &clientAuth.Config{
		ClientID:     id,
		ClientSecret: secret,
		TokenURL:     endpoint.TokenURL,
	}

	ctx := context.TODO()
	authClient := oauthCfg.Client(ctx)
	client := battlenet.NewClient("us", authClient)

	return client
}
