package main

import (
	"errors"
	"flag"
	"fmt"
	"keto/kfx/pb"
	"keto/kfx/route"
	"log"
	"net"

	"github.com/gogo/protobuf/proto"

	"google.golang.org/grpc"

	"google.golang.org/grpc/credentials"
)

var (
	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile   = flag.String("cert_file", "", "The TLS cert file")
	keyFile    = flag.String("key_file", "", "The TLS key file")
	jsonDBFile = flag.String("json_db_file", "testdata/route_guide_db.json", "A json file containing a list of features")
	port       = flag.Int("port", 10000, "The server port")
)

type Router interface {
	Route(stream pb.KfxService_DoWorksServer) error
}

type kfxServer struct {
	routers map[string]Router
}

type AccountRouter struct {
	authed bool
}

func (server *kfxServer) DoWorks(stream pb.KfxService_DoWorksServer) error {
	return route.HandleStream(stream)
}

func (server *kfxServer) Route(stream pb.KfxService_DoWorksServer) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}

		switch pb.Codes(req.Code) {
		case pb.Codes_ConfigService:
			body := pb.ConfigServiceRequest{}
			if err := proto.Unmarshal(req.Payload, &body); err != nil {
				return errors.New("bad payload")
			}
			return server.configService(stream, &body)
		}
		break
	}

	return nil
}

func (server *kfxServer) configService(stream pb.KfxService_DoWorksServer, cfg *pb.ConfigServiceRequest) error {
	m := route.Mt4Route{}
	return m.Route(stream)
}

func newKfxServer() *kfxServer {
	s := new(kfxServer)
	return s
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	if *tls {
		if *certFile == "" {
			*certFile = "server1.pem"
		}
		if *keyFile == "" {
			*keyFile = "server1.key"
		}
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			log.Fatalf("Failed to generate credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterKfxServiceServer(grpcServer, newKfxServer())
	grpcServer.Serve(lis)
}
