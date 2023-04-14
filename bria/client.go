package bria

import (
	"context"
	"fmt"

	adminv1 "github.com/GaloyMoney/terraform-provider-briaadmin/bria/proto/admin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type AdminClient struct {
	conn    *grpc.ClientConn
	service adminv1.AdminServiceClient
}

func NewAdminClient(endpoint string, apiKey string) (*AdminClient, error) {
	opts := []grpc.DialOption{grpc.WithInsecure()}

	// Inject API key as an HTTP header
	interceptor := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		newCtx := metadata.AppendToOutgoingContext(ctx, "x-bria-admin-api-key", apiKey)
		return invoker(newCtx, method, req, reply, cc, opts...)
	}
	opts = append(opts, grpc.WithUnaryInterceptor(interceptor))

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

func (c *AdminClient) ReadAccount(accountID string) (*adminv1.Account, error) {
	ctx := context.Background()

	listAccountsRequest := &adminv1.ListAccountsRequest{}
	listAccountsResponse, err := c.service.ListAccounts(ctx, listAccountsRequest)
	if err != nil {
		return nil, fmt.Errorf("error fetching accounts: %w", err)
	}

	var foundAccount *adminv1.Account
	for _, account := range listAccountsResponse.Accounts {
		if account.Id == accountID {
			foundAccount = account
			break
		}
	}

	return foundAccount, nil
}
