package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/camtrik/gRPC-blog-tag-management/proto"
)

func main() {
	ctx := context.Background()
	clientConn, _ := GetClientConn("localhost:8080", nil)
	defer clientConn.Close()

	targetServiceClient := pb.NewTagServiceClient(clientConn)
	resp, _ := targetServiceClient.GetTagList(ctx, &pb.GetTagListRequest{Name: "go"})

	log.Printf("resp: %v", resp)
}

func GetClientConn(target string, opts []grpc.DialOption) (*grpc.ClientConn, error) {
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	return grpc.NewClient(target, opts...)
}
