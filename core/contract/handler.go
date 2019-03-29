package contract

import (
	"log"
	"net"

	pb "study/GitHub/consensus/protoc/contractcode"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

const (
	contractPort = ":50053"
)

var rq RequiredCode
var detail *contractDetail

func (s *server) InitContract(in *empty.Empty, stream pb.Contract_InitContractServer) error {

	tDetail := new(transactionDetail)
	tDetail.initStream = stream
	tDetail.invokeStream = nil
	detail.InitHandler(tDetail)

	return nil
}

func (s *server) Invoke(in *pb.ContractInfo, stream pb.Contract_InvokeServer) error {

	tDetail := new(transactionDetail)
	tDetail.invokeStream = stream
	tDetail.initStream = nil
	detail.InvokeHandler(in.Transaction, in.Args, tDetail)
	return nil
}

// RegisterListener registers a server created for contract execution
func RegisterListener(rc RequiredCode, dtl *contractDetail) {

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
