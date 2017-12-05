package main

import (
	"context"
	"flag"
	"keto/kfx/pb"
	"log"
	"sync/atomic"

	"github.com/golang/protobuf/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containning the CA root cert file")
	serverAddr         = flag.String("server_addr", "127.0.0.1:10000", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.youtube.com", "The server name use to verify the hostname returned by TLS handshake")
)

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	if *tls {
		if *caFile == "" {
			*caFile = "ca.pem"
		}
		creds, err := credentials.NewClientTLSFromFile(*caFile, *serverHostOverride)
		if err != nil {
			log.Fatalf("Failed to create TLS credentials %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewKfxServiceClient(conn)
	stream, err := client.DoWorks(context.Background())
	if err != nil {
		panic(err)
	}

	var rid int64
	newRid := func() int64 {
		return atomic.AddInt64(&rid, 1)
	}

	newRequest := func(code pb.Codes, msg proto.Message) *pb.Request {
		b, _ := proto.Marshal(msg)
		return &pb.Request{
			Code:      int32(code),
			Rid:       newRid(),
			Payload:   b,
			Debug:     true,
			SplitSize: 30,
		}
	}

	for {
		stream.Send(newRequest(pb.Codes_GetPositions, &pb.GetPositionsRequest{}))
	}
}
