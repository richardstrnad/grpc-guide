package main

import (
	"context"
	"crypto/x509"
	"fmt"
	"log"

	pb "grpc-gcloud/ping"

	"google.golang.org/api/idtoken"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
)

var (
	server     = "ping-leipieppba-uc.a.run.app"
	serverAuth = "ping-auth-leipieppba-uc.a.run.app"
)

func auth(ctx context.Context) oauth.TokenSource {
	audience := fmt.Sprintf("https://%s", serverAuth)
	// We load the service account from the json file
	idTokenSource, err := idtoken.NewTokenSource(ctx, audience, idtoken.WithCredentialsFile("service-account.json"))
	if err != nil {
		log.Fatalf("unable to create TokenSource: %v", err)
	}
	return oauth.TokenSource{idTokenSource}
}

func main() {
	ctx := context.Background()
	perRpc := auth(ctx)

	pool, _ := x509.SystemCertPool()
	creds := credentials.NewClientTLSFromCert(pool, "")

	// Set up a connection to the server. Without auth.
	conn, err := grpc.Dial(
		server+":443",
		grpc.WithTransportCredentials(creds),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	// Set up a connection to the server. With auth.
	connAuth, err := grpc.Dial(
		serverAuth+":443",
		grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(perRpc),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	defer connAuth.Close()

	c := pb.NewPingerClient(conn)
	cAuth := pb.NewPingerClient(connAuth)

	r, err := c.Ping(ctx, &pb.PingRequest{Message: "Go Client"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())

	rAuth, err := cAuth.Ping(ctx, &pb.PingRequest{Message: "Go Client"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", rAuth.GetMessage())
}
