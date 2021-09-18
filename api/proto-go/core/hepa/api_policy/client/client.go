// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: api_policy.proto

package client

import (
	context "context"

	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/core/hepa/api_policy/pb"
	grpc1 "google.golang.org/grpc"
)

// Client provide all service clients.
type Client interface {
	// ApiPolicyService api_policy.proto
	ApiPolicyService() pb.ApiPolicyServiceClient
}

// New create client
func New(cc grpc.ClientConnInterface) Client {
	return &serviceClients{
		apiPolicyService: pb.NewApiPolicyServiceClient(cc),
	}
}

type serviceClients struct {
	apiPolicyService pb.ApiPolicyServiceClient
}

func (c *serviceClients) ApiPolicyService() pb.ApiPolicyServiceClient {
	return c.apiPolicyService
}

type apiPolicyServiceWrapper struct {
	client pb.ApiPolicyServiceClient
	opts   []grpc1.CallOption
}

func (s *apiPolicyServiceWrapper) GetPolicy(ctx context.Context, req *pb.GetPolicyRequest) (*pb.GetPolicyResponse, error) {
	return s.client.GetPolicy(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *apiPolicyServiceWrapper) SetPolicy(ctx context.Context, req *pb.SetPolicyRequest) (*pb.SetPolicyResponse, error) {
	return s.client.SetPolicy(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}