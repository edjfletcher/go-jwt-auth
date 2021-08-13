package go_jwt_auth

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
)

type FirebaseIsValid func(ctx context.Context, idToken string) (error, *auth.Token)

func FirebaseIsValidFactory(app *firebase.App) FirebaseIsValid {
	return func(ctx context.Context, idToken string) (error, *auth.Token) {
		client, err := app.Auth(ctx)
		if err != nil {
			return err, nil
		}

		token, err := client.VerifyIDToken(ctx, idToken)
		if err != nil {
			return err, nil
		}

		return nil, token
	}
}
