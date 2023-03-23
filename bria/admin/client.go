package admin

import (
	"context"
	"crypto/tls"

	adminv1 "github.com/GaloyMoney/terraform-provider-bria/bria/proto/admin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type AdminClient struct {
	conn    *grpc.ClientConn
	service adminv1.AdminServiceClient
}

func NewAdminClient(endpoint string, apiKey string) (*AdminClient, error) {
	creds := credentials.NewTLS(&tls.Config{})
	opts := []grpc.DialOption{grpc.WithTransportCredentials(creds)}
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return nil, err
	}

	client := adminv1.NewAdminServiceClient(conn)

	return &AdminClient{
		conn:    conn,
		service: client,
	}, nil
}

func (c *AdminClient) Close() {
	c.conn.Close()
}

func (c *AdminClient) CreateAccount(name string) (*adminv1.AccountCreateResponse, error) {
	req := &adminv1.AccountCreateRequest{
		Name: name,
	}
	ctx := context.Background()
	res, err := c.service.AccountCreate(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
