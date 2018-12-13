package config

import (
	"golang.org/x/oauth2"
	"os"
)

var wserverHost = os.Getenv("WSERVER_HOST")

func initGRPC() {
	/*perRPC := oauth.NewOauthAccess(oauth.TokenSource{})

	opts := []grpc.DialOption{
		// In addition to the following grpc.DialOption, callers may also use
		// the grpc.CallOption grpc.PerRPCCredentials with the RPC invocation
		// itself.
		// See: https://godoc.org/google.golang.org/grpc#PerRPCCredentials
		grpc.WithPerRPCCredentials(perRPC),
		// oauth.NewOauthAccess requires the configuration of transport
		// credentials.
		grpc.WithTransportCredentials(
			credentials.NewTLS(&tls.Config{InsecureSkipVerify: true}),
		),
	}
	conn, err := grpc.Dial(":8080", opts...)*/
	//TODO implement service to Deserialize player FameStats from Characters
}

type OauthImpl struct {
	token        string
	refreshToken string
}

func (tk *OauthImpl) Token() (*oauth2.Token, error) {
	return &oauth2.Token{AccessToken: tk.token, RefreshToken: tk.refreshToken, TokenType: "bearer"}, nil
}
