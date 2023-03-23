package account

import (
	"context"
	"encoding/base64"
	"fmt"

	briav1 "github.com/GaloyMoney/terraform-provider-bria/bria/proto/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type AccountClient struct {
	conn    *grpc.ClientConn
	service briav1.BriaServiceClient
}

func NewAccountClient(endpoint string, apiKey string) (*AccountClient, error) {
	opts := []grpc.DialOption{grpc.WithInsecure()}

	// Inject API key as an HTTP header
	interceptor := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		newCtx := metadata.AppendToOutgoingContext(ctx, "x-bria-api-key", apiKey)
		return invoker(newCtx, method, req, reply, cc, opts...)
	}
	opts = append(opts, grpc.WithUnaryInterceptor(interceptor))

	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return nil, err
	}

	client := briav1.NewBriaServiceClient(conn)

	return &AccountClient{
		conn:    conn,
		service: client,
	}, nil
}

func (c *AccountClient) Close() {
	c.conn.Close()
}

func (c *AccountClient) ImportXpub(name, xpub, derivation string) (*briav1.ImportXpubResponse, error) {
	req := &briav1.ImportXpubRequest{
		Name:       name,
		Xpub:       xpub,
		Derivation: derivation,
	}
	ctx := context.Background()
	res, err := c.service.ImportXpub(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AccountClient) CreateLndSignerConfig(xpub string, lndConfig []interface{}) error {
	lnd := lndConfig[0].(map[string]interface{})

	certBase64 := base64.StdEncoding.EncodeToString([]byte(lnd["cert"].(string)))

	lndSignerConfig := &briav1.LndSignerConfig{
		Endpoint:       lnd["endpoint"].(string),
		CertBase64:     certBase64,
		MacaroonBase64: lnd["macaroon_base64"].(string),
	}

	req := &briav1.SetSignerConfigRequest{
		XpubRef: xpub,
		Config: &briav1.SetSignerConfigRequest_Lnd{
			Lnd: lndSignerConfig,
		},
	}

	_, err := c.service.SetSignerConfig(context.Background(), req)
	if err != nil {
		return fmt.Errorf("error setting LND signer config: %w", err)
	}

	return nil
}

func (c *AccountClient) CreateWallet(name string, xpubRefs []string) (*briav1.CreateWalletResponse, error) {
	req := &briav1.CreateWalletRequest{
		Name:     name,
		XpubRefs: xpubRefs,
	}
	ctx := context.Background()
	res, err := c.service.CreateWallet(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AccountClient) CreateBatchGroup(name string) (*briav1.CreateBatchGroupResponse, error) {
	req := &briav1.CreateBatchGroupRequest{
		Name: name,
	}
	ctx := context.Background()
	res, err := c.service.CreateBatchGroup(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
