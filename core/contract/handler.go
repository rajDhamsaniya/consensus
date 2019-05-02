package contract

import (
	"errors"
	"log"
	"net"

	pb "protoc/contractcode"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

const (
	contractPort = ":50053"
)

var rq RequiredCode
var detail *ContractDetail

type server struct{}

func (s *server) InitContract(in *empty.Empty, stream pb.Contract_InitContractServer) error {

	tDetail := new(TransactionDetail)
	tDetail.initStream = stream
	tDetail.invokeStream = nil
	code := detail.InitHandler(tDetail)
	if code == 200 {
		return nil
	}
	return errors.New("Can't Init Contract")

}

func (s *server) Invoke(in *pb.ContractInfo, stream pb.Contract_InvokeServer) error {

	tDetail := new(TransactionDetail)
	tDetail.invokeStream = stream
	tDetail.initStream = nil
	code := detail.InvokeHandler(in.Transaction, in.Args, tDetail)
	if code == 200 {
		return nil
	}
	return errors.New("Can't Invoke Contract")
}

// RegisterListener registers a server created for contract execution
func RegisterListener(rc RequiredCode, dtl *ContractDetail) {

	rq = rc
	detail = dtl
	lis, err := net.Listen("tcp", contractPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterContractServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
