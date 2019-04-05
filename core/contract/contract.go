package contract

import (
	"fmt"

	pb "../../protoc/contractcode"

	"github.com/gogo/protobuf/proto"

	"github.com/leesper/couchdb-golang"
)

//var db *couchdb.Database

const (
	// DefaultBaseURL is the default address of CouchDB server.
	DefaultBaseURL = "http://10.0.2.15:5984/"
)

// ContractDetail for details regarding contract
type ContractDetail struct {
	contractID string
	db         *couchdb.Database
	channel    chan []map[string]interface{}
}

// TransactionDetail for details regarding transaction
type TransactionDetail struct {
	initStream   pb.Contract_InitContractServer
	invokeStream pb.Contract_InvokeServer
}

// RequiredCode interface for stating the code to behave
type RequiredCode interface {
	Init(detail *ContractDetail, txDetail *TransactionDetail) (code int)
	Invoke(detail *ContractDetail, tx string, args []string, txDetail *TransactionDetail) (code int)
}

// SupportCodeInterface interface for stating the code to behave
type SupportCodeInterface interface {
	ConnectDB()
	GetState(id string, txDetail *TransactionDetail) (docsQuery []map[string]interface{}, err error)
	GetCollectionState(query string) (docsQuery []map[string]interface{}, err error)
	PutState(newState []map[string]interface{}, txDetail *TransactionDetail)
	GenerateUUID(uuid string)
}

func (detail *ContractDetail) InitHandler(tDetail *TransactionDetail) {
	rq.Init(detail, tDetail)
}

func (detail *ContractDetail) InvokeHandler(tx string, args []string, txDetail *TransactionDetail) {

	//detail.channel = make(chan []map[string]interface{}, 1)
	rq.Invoke(detail, tx, args, txDetail)
	//close(detail.channel)
}

// ConnectDB used for connection to database
func (detail *ContractDetail) ConnectDB() {

	s, err := couchdb.NewServer(DefaultBaseURL)
	if err != nil {
		fmt.Println("hey")
		fmt.Println(err)
	}
	s.Login("admin", "admin")
	// var db *couchdb.Database
	detail.db, err = s.Get("new")

	if err != nil {
		fmt.Println("heyiiiiiii")
		fmt.Println(err)
	}
	err = detail.db.Available()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("ConnectDB called")
}

func (detail *ContractDetail) GetState(id string, txDetail *TransactionDetail) (docsQuery []map[string]interface{}, err error) {

	docsQuery, err = detail.db.Query(nil, `_id=="`+id+`"`, nil, nil, nil, nil)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	if txDetail.initStream != nil {
		out := &pb.GetState{Key: id, Version: fmt.Sprint(docsQuery[0]["_rev"])}
		payload, err := proto.Marshal(out)
		if err != nil {
			fmt.Println(err)
		}
		txDetail.initStream.Send(&pb.TransactionResponse{Query: pb.TransactionResponse_GETSTATE, Payload: payload})
	} else {
		out := &pb.GetState{Key: id, Version: fmt.Sprint(docsQuery[0]["_rev"])}
		payload, err := proto.Marshal(out)
		if err != nil {
			fmt.Println(err)
		}
		txDetail.invokeStream.Send(&pb.TransactionResponse{Query: pb.TransactionResponse_GETSTATE, Payload: payload})
	}
	return docsQuery, err
}

func (detail *ContractDetail) GetCollectionState(query string) (docsQuery []map[string]interface{}, err error) {

	docsQuery, err = detail.db.Query(nil, query, nil, nil, nil, nil)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	return docsQuery, err
}

// PutState used for Get data from database
// It takes  newState as parameter to run on database
func (detail *ContractDetail) PutState(newState []map[string]interface{}, txDetail *TransactionDetail) {

	if txDetail.initStream != nil {
		// out := &pb.GetState{Key: id, Version: fmt.Sprintln(docsQuery[0]["_rev"])}
		for _, state := range newState {
			id := fmt.Sprintln(state["_id"])
			arr := make([]*pb.ValueInfo, 0)
			for k := range state {
				val := fmt.Sprintln(state[k])
				arr = append(arr, &pb.ValueInfo{Key: k, Value: val})
				//out := &pb.PutState{Id: state["_id"], Value: state[k]}

			}
			out := &pb.PutState{Id: id, Val: arr}
			payload, err := proto.Marshal(out)
			if err != nil {
				fmt.Println(err)
			}
			txDetail.initStream.Send(&pb.TransactionResponse{Query: pb.TransactionResponse_PUTSTATE, Payload: payload})
		}
	} else {
		for _, state := range newState {
			id := fmt.Sprintln(state["_id"])
			arr := make([]*pb.ValueInfo, 0)
			for k := range state {
				val := fmt.Sprintln(state[k])
				arr = append(arr, &pb.ValueInfo{Key: k, Value: val})
				//out := &pb.PutState{Id: state["_id"], Value: state[k]}

			}
			out := &pb.PutState{Id: id, Val: arr}
			payload, err := proto.Marshal(out)
			if err != nil {
				fmt.Println(err)
			}
			txDetail.invokeStream.Send(&pb.TransactionResponse{Query: pb.TransactionResponse_PUTSTATE, Payload: payload})
		}
		//txDetail.invokeStream.Send(&pb.TransactionResponse{Query: pb.TransactionResponse_PUTSTATE, Payload: payload})
	}

	// NI
	if newState[0]["_id"] == nil {
		err := detail.db.Available()
		fmt.Println(err)
		if err == nil {
			detail.ConnectDB()
		}
		fmt.Println("here np")
		uuid := couchdb.GenerateUUID()
		for {
			if detail.db.Contains(uuid) != nil {
				break
			} else {
				uuid = couchdb.GenerateUUID()
			}
		}

		newState[0]["_id"] = uuid
		_, _, err = detail.db.Save(newState[0], nil)
		if err != nil {
			fmt.Print("could not save")
			fmt.Println(err)
		}
	} else {
		detail.db.Update(newState, nil)
	}
}

func (detail *ContractDetail) TxOut() (out []map[string]interface{}) {

	return out
}

//Start for chaincode listener
func Start(rq RequiredCode) {

	detail := new(ContractDetail)
	detail.ConnectDB()
	fmt.Println("reaches here")
	RegisterListener(rq, detail)
	// lis, err := net.Listen("tcp", contractPort)
	// if err != nil {
	// 	log.Fatalf("failed to listen: %v", err)
	// }
	// s := grpc.NewServer()
	// pb.RegisterContractServer(s, &server{})
	// if err := s.Serve(lis); err != nil {
	// 	log.Fatalf("failed to serve: %v", err)
	// }
}
