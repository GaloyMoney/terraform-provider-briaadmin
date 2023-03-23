package grpcclient

import (
	"context"
	"fmt"

	"github.com/GaloyMoney/terraform-provider-bria/services/bria_admin/v1"
	"google.golang.org/grpc"
)

type Client struct {
	conn    *grpc.ClientConn
	service v1.AdminServiceClient
}

func NewClient(address, apiKey string) (*Client, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithPerRPCCredentials(newCustomCreds(apiKey)))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the gRPC server: %w", err)
	}

	return &Client{
		conn:    conn,
		service: v1.NewAdminServiceClient(conn),
	}, nil
}

func (c *Client) Close() {
	c.conn.Close()
}

func (c *Client) CreateAccount(ctx context.Context, name string) (*v1.AccountCreateResponse, error) {
	request := &v1.AccountCreateRequest{Name: name}
	response, err := c.service.AccountCreate(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("failed to create account: %w", err)
	}
	return response, nil
}

type customCreds struct {
	apiKey string
}

func newCustomCreds(apiKey string) *customCreds {
	return &customCreds{apiKey: apiKey}
}

func (c *customCreds) GetRequestMetadata(ctx context.Context, uri ...

