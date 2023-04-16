package main

import (
	"context"
	"github.com/kronosbro3886/go-grpc/invoicer"
	"google.golang.org/grpc"
	"log"
	"net"
)

type myInvoiceServer struct {
	invoicer.UnimplementedInvoicerServer
}

func (s myInvoiceServer) Create(ctx context.Context, req *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {
	return &invoicer.CreateResponse{
		Pdf:  []byte("test"),
		Docx: []byte("test"),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("cannot create listener : %s", err)
	}

	serverRegistrar := grpc.NewServer()
	service := &myInvoiceServer{}

	invoicer.RegisterInvoicerServer(serverRegistrar, service)
	serverRegistrar.Serve(lis)

	if err != nil {
		log.Fatalf("impossible to serve : %s", err)
	}
}
