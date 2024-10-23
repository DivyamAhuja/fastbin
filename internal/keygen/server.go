package keygen

import (
	context "context"
	"math/rand"

	pb "fastbin/api/keygen"

	"google.golang.org/grpc"
)

type keygenServer struct {
	pb.UnimplementedKeygenServer
}

func (k *keygenServer) GenerateKey(ctx context.Context, req *pb.Empty) (*pb.Key, error) {
	b := make([]byte, 6)
	for i := range b {
		b[i] = 'a' + byte(rand.Intn(26))
	}

	key := pb.Key{Value: string(b)}
	return &key, nil
}

func NewKeygenServer() *grpc.Server {
	grpcSever := grpc.NewServer()
	pb.RegisterKeygenServer(grpcSever, &keygenServer{})
	return grpcSever
}
