package client

import (
	"crypto/tls"
	"crypto/x509"
	"os"
	"strconv"
	"time"

	"ddd/internal/common/genproto/auth"
	"ddd/internal/common/genproto/twofa"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

func NewAuthClient() (client auth.AuthServiceClient, close func() error, err error) {
	grpcAddr := os.Getenv("AUTH_GRPC_ADDR")
	if grpcAddr == "" {
		return nil, func() error { return nil }, errors.New("empty env AUTH_GRPC_ADDR")
	}

	opts, err := grpcDialOpts( /* grpcAddr */ )
	if err != nil {
		return nil, func() error { return nil }, err
	}

	conn, err := grpc.NewClient(grpcAddr, opts...)
	if err != nil {
		return nil, func() error { return nil }, err
	}

	return auth.NewAuthServiceClient(conn), conn.Close, nil
}

func WaitForAuthService(timeout time.Duration) bool {
	return waitForPort(os.Getenv("AUTH_GRPC_ADDR"), timeout)
}

func NewTwofaClient() (client twofa.TwoFAServiceClient, close func() error, err error) {
}

func grpcDialOpts( /*grpcAddr string*/ ) ([]grpc.DialOption, error) {
	if noTLS, _ := strconv.ParseBool(os.Getenv("GRPC_NO_TLS")); noTLS {
		return []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}, nil
	}

	systemRoots, err := x509.SystemCertPool()
	if err != nil {
		return nil, errors.Wrap(err, "cannot load root CA cert")
	}

	creds := credentials.NewTLS(&tls.Config{
		RootCAs:    systemRoots,
		MinVersion: tls.VersionTLS12,
	})

	return []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
		// grpc.WithPerRPCCredentials(newMetadataServerToken(grpcAddr))
	}, nil
}
