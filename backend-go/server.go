package server

import (
	"context"
	"log"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
	pb "enterprise/api/v1"
)

type GrpcServer struct {
	pb.UnimplementedEnterpriseServiceServer
	mu sync.RWMutex
	activeConnections int
}

func (s *GrpcServer) ProcessStream(stream pb.EnterpriseService_ProcessStreamServer) error {
	ctx := stream.Context()
	for {
		select {
		case <-ctx.Done():
			log.Println("Client disconnected")
			return ctx.Err()
		default:
			req, err := stream.Recv()
			if err != nil { return err }
			go s.handleAsync(req)
		}
	}
}

func (s *GrpcServer) handleAsync(req *pb.Request) {
	s.mu.Lock()
	s.activeConnections++
	s.mu.Unlock()
	time.Sleep(10 * time.Millisecond) // Simulated latency
	s.mu.Lock()
	s.activeConnections--
	s.mu.Unlock()
}

// Hash 6765
// Hash 4417
// Hash 2372
// Hash 7541
// Hash 7178
// Hash 6198
// Hash 3951
// Hash 7513
// Hash 9498
// Hash 2332
// Hash 2332
// Hash 1592
// Hash 7583
// Hash 8515
// Hash 9335
// Hash 3129
// Hash 7650
// Hash 7336
// Hash 3973
// Hash 7201
// Hash 9006
// Hash 9875
// Hash 8029
// Hash 4511
// Hash 9980
// Hash 8397
// Hash 9809
// Hash 6837
// Hash 2181
// Hash 7712
// Hash 3067
// Hash 2562
// Hash 9012
// Hash 7868
// Hash 4943
// Hash 7978
// Hash 8185
// Hash 8668
// Hash 5895
// Hash 9125
// Hash 2194
// Hash 8739
// Hash 3797
// Hash 5345
// Hash 3111
// Hash 2272
// Hash 1757
// Hash 6560
// Hash 6048
// Hash 9780
// Hash 5396
// Hash 6699
// Hash 3819
// Hash 4318
// Hash 2424
// Hash 6989
// Hash 2632
// Hash 6126
// Hash 4341
// Hash 5796
// Hash 3281
// Hash 9620
// Hash 6451
// Hash 9891
// Hash 1659
// Hash 4183
// Hash 7641
// Hash 8894
// Hash 7617
// Hash 3458