package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"

	pb "../protoc/contractcode"

	"github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/leesper/couchdb-golang"
	"google.golang.org/grpc"
)

type server struct{}

const (
	// DefaultBaseURL is the default address of CouchDB server.
	DefaultBaseURL = "http://10.0.2.15:5984/"
	contractPort   = ":50053"
)

var d *couchdb.Database

func connectDatabase() {
	s, err := couchdb.NewServer(DefaultBaseURL)
	if err != nil {
		fmt.Println("hey")
		fmt.Println(err)
	}
	s.Login("admin", "admin")
	// var d *couchdb.Database
	d, err = s.Get("new")

	if err != nil {
		fmt.Println("heyiiiiiii")
		fmt.Println(err)
	}
	err = d.Available()
	if err != nil {
		fmt.Println(err)
	}
}

// AddUser for adding the user into database
func AddUser(in *pb.UserInfo) (id string) {

	err := d.Available()
	fmt.Println(err)
	if err == nil {
		connectDatabase()
	}
	fmt.Println("here np")
	uuid := couchdb.GenerateUUID()
	for {
		if d.Contains(uuid) != nil {
			break
		} else {
			uuid = couchdb.GenerateUUID()
		}
	}
	doc := make(map[string]interface{})

	doc["_id"] = uuid
	doc["name"] = in.UserName
	doc["balance"] = strconv.Itoa(int(in.Balance))

	id, _, err = d.Save(doc, nil)
	if err != nil {
		fmt.Print("could not save")
		fmt.Println(err)
	}
	return id
}

// TransferAmount for transfer amount from one user to other
func TransferAmount(in *pb.TransferMessage) (*pb.StateInfo, error) {

	if in.UserId == in.FromAccId {
		//fmt.Println("here")
		docsQuery, err := d.Query(nil, `_id=="`+in.FromAccId+`"`, nil, nil, nil, nil)
		if err != nil {
			fmt.Println(err)
			return &pb.StateInfo{Msg: pb.StateInfo_ERROR}, nil
		}
		//fmt.Println("here2")
		getState := make([]map[string]interface{}, 1)
		getState[0] = docsQuery[0]
		if d.Contains(in.ToAccId) == nil {
			//fmt.Println("here3")
			fromBalance, err := strconv.Atoi(docsQuery[0]["balance"].(string))
			if err == nil && fromBalance >= int(in.Amount) {
				//fmt.Println("here4")
				docsQuery2, err := d.Query(nil, `_id=="`+in.ToAccId+`"`, nil, nil, nil, nil)
				getState = append(getState, docsQuery2[0])
				toBalance, err := strconv.Atoi(docsQuery2[0]["balance"].(string))
				if err == nil {
					docsQuery[0]["balance"] = strconv.Itoa(fromBalance - int(in.Amount))
					docsQuery2[0]["balance"] = strconv.Itoa(toBalance + int(in.Amount))
					updateInfo := append(docsQuery, docsQuery2[0])
					d.Update(updateInfo, nil)
					get := &pb.GetState{Key: "balance"}
					data, err := proto.Marshal(get)
					if err != nil {
						fmt.Print("Can't Marshal get")
					}
					return &pb.StateInfo{Msg: pb.StateInfo_SUCESS, Payload: data}, nil
				}
			}
		}
	}

	return &pb.StateInfo{Msg: pb.StateInfo_ERROR}, nil
	// fmt.Println(docsQuery)
}

func (s *server) InitContract(ctx context.Context, in *empty.Empty) (*pb.InfoMsg, error) {

	connectDatabase()
	return &pb.InfoMsg{Msg: "its a success", Type: pb.InfoMsg_SUCESS}, nil
}

func (s *server) Invoke(ctx context.Context, in *pb.ContractInfo) (*pb.Response, error) {

	if in.Transaction == "AddUser" {
		var userInfo pb.UserInfo
		proto.Unmarshal(in.Args, &userInfo)
		userInfo.UserId = AddUser(&userInfo)
		out, err := proto.Marshal(&userInfo)
		if err != nil {
			fmt.Println(err)
		}
		return &pb.Response{Result: out}, nil

	} else if in.Transaction == "TransferAmount" {
		var transferMessage pb.TransferMessage
		proto.Unmarshal(in.Args, &transferMessage)
		//fmt.Println(transferMessage)
		stateInfo, err := TransferAmount(&transferMessage)
		out, errr := proto.Marshal(stateInfo)
		if errr != nil {
			fmt.Println(errr)
		}
		return &pb.Response{Result: out}, err
	}
	return &pb.Response{Result: nil}, nil
}

func main() {

	connectDatabase()
	lis, err := net.Listen("tcp", contractPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterContractServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	fmt.Println("Hello")
}
