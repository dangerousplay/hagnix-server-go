package config

import (
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"hagnix-server-go1/proto"
	"os"
)

var wserverHost = os.Getenv("WSERVER_HOST")
var gameClient hagnix.GameClient

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

	conn, err := grpc.Dial(wserverHost)

	if err != nil {
		panic(err)
	}

	gameClient = hagnix.NewGameClient(conn)
	//TODO implement service to Deserialize player FameStats from Characters
}

func GetGameClient() hagnix.GameClient {
	return gameClient
}

type OauthImpl struct {
	token        string
	refreshToken string
}

func (tk *OauthImpl) Token() (*oauth2.Token, error) {
	return &oauth2.Token{AccessToken: tk.token, RefreshToken: tk.refreshToken, TokenType: "bearer"}, nil
}
