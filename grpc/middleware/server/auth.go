package server

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	ClientHeaderKey = "client-id"
	ClientSecretKey = "client-secret"
)

func NewClientCredential(ak, sk string) metadata.MD {
	return metadata.MD{
		ClientHeaderKey: []string{ak},
		ClientSecretKey: []string{sk},
	}
}

func NewAuthUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return (&GrpcAuther{}).UnaryServerInterceptor
}

type GrpcAuther struct {
}

func (a *GrpcAuther) UnaryServerInterceptor(
	ctx context.Context, req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (
	resp interface{}, err error) {
	// 1.读取凭证,凭证放在meta信息[http2 header]
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("ctx is not an grpc incoming context")
	}

	// 从meta data中获取客户端传递过来的凭证
	clientId, clientSecret := a.getClientCredentialsFromMeta(md)
	if err := a.validateServiceCredential(clientId, clientSecret); err != nil {
		return nil, err
	}
	handler(ctx, req)
	return nil, nil

}

// stream rpc interceptor
func (a *GrpcAuther) StreamServerInterceptor(
	srv interface{},
	ss grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler) error {

	return nil
}

func (a *GrpcAuther) getClientCredentialsFromMeta(md metadata.MD) (
	clientId, clientSecret string) {
	cakList := md[ClientHeaderKey]
	if len(cakList) > 0 {
		clientId = cakList[0]
	}
	cskList := md[ClientSecretKey]
	if len(cskList) > 0 {
		clientSecret = cskList[0]
	}
	return
}

func (a *GrpcAuther) validateServiceCredential(
	clientId, clientSecret string) error {
	if !(clientId == "admin" && clientSecret == "123456") {
		// 返回一个认证错误，并结束RPC调用
		return status.Errorf(codes.Unauthenticated, "client_id or client_secret not conrect")
	}
	return nil
}
