package main

import (
	"flag"
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/camtrik/gRPC-blog-tag-management/proto"
	"github.com/camtrik/gRPC-blog-tag-management/server"
	"github.com/soheilhy/cmux"
)

var port string
var grpcPort string
var httpPort string

func init() {
	flag.StringVar(&port, "port", "8080", "API server port")
	flag.StringVar(&grpcPort, "grpc_port", "7080", "gRPC port")
	flag.StringVar(&httpPort, "http_port", "9080", "HTTP port")
	flag.Parse()
}

func RunTCPServer(port string) (net.Listener, error) {
	return net.Listen("tcp", ":"+port)
}

func RunHttpServer(port string) *http.Server {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`pong`))
	})

	return &http.Server{
		Addr:    ":" + port,
		Handler: serveMux,
	}
}

func RunGrpcServer() *grpc.Server {
	s := grpc.NewServer()
	pb.RegisterTagServiceServer(s, server.NewTagServer())
	reflection.Register(s)

	return s
}

func main() {
	l, err := RunTCPServer(port)
	if err != nil {
		log.Fatalf("Run TCP server error: %v", err)
	}

	m := cmux.New(l)
	grpcL := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldPrefixSendSettings("content-type", "application/grpc"))
	httpL := m.Match(cmux.HTTP1Fast())

	grpcS := RunGrpcServer()
	httpS := RunHttpServer(port)
	go grpcS.Serve(grpcL)
	go httpS.Serve(httpL)

	err = m.Serve()
	if err != nil {
		log.Fatalf("Run server error: %v", err)

	}

	// errs := make(chan error)
	// go func() {
	// 	err := RunHttpServer(httpPort)
	// 	if err != nil {
	// 		errs <- err
	// 	}
	// }()

	// go func() {
	// 	err := RunGrpcServer(grpcPort)
	// 	if err != nil {
	// 		errs <- err
	// 	}
	// }()

	// select {
	// case err := <-errs:
	// 	log.Fatalf("error: %v", err)
	// }
}
