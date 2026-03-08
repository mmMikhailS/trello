package client

//
// import (
// 	"context"
// 	"fmt"
// 	"strings"
//
// 	"cloud.google.com/go/compute/metadata"
// 	"github.com/pkg/errors"
// 	"google.golang.org/grpc/credentials"
// )
//
// type metadataServerToken struct {
// 	serviceURL string
// }
//
// func newMetadataServerToken(grpcAddr string) credentials.PerRPCCredentials {
// 	serviceURL := "htpp://" + strings.Split(grpcAddr, ":")[0]
//
// 	return metadataServerToken{serviceURL: serviceURL}
// }
//
// // GetRequestMetadata is called on every request, so we are sure that token is always not expired
// func (t metadataServerToken) GetRequestMetadata(ctx context.Context, in ...string) (map[string]string, error) {
// 	tokenURL := fmt.Sprintf("/instance/service-accounts/default/identity?audience=%s", t.serviceURL)
// 	idToken, err := metadata.GetWithContext(ctx, tokenURL)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "cannot query id token for gRPC")
// 	}
//
// 	return map[string]string{
// 		"authorization": "Bearer " + idToken,
// 	}, nil
// }
//
// func (metadataServerToken) RequireTransportSecurity() bool {
// 	return true
// }
