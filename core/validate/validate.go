package validate

import (
	"fmt"
	"sync"

	pb2 "protoc/contractcode"
	pb "protoc/gossip"

	"github.com/gogo/protobuf/proto"
	"github.com/leesper/couchdb-golang"
)

const (
	// DefaultBaseURL is the default address of CouchDB server.
	DefaultBaseURL = "http://10.0.2.15:5984/"
)

//BlockInfo regarding the block
type BlockInfo struct {
	endorsedTxs []*pb.EndorsedTx
	sign        string
	offset      string
	mask        []bool
}

var db *couchdb.Database

// ValidateBlock for validation
func (block *BlockInfo) ValidateBlock() {

	var wg sync.WaitGroup
	for i := 0; i < len(block.endorsedTxs); i++ {
		wg.Add(1)
		go block.CheckEndorcement(block.endorsedTxs[i], i, &wg)
	}
	wg.Wait()

	for i := 0; i < len(block.endorsedTxs); i++ {
		fmt.Println("Reaches here")
		if block.mask[i] {
			var txResponse pb2.Cargo
			err := proto.Unmarshal(block.endorsedTxs[i].Payload, &txResponse)
			if err != nil {
				fmt.Println(err)
				return
			}
			for j := 0; j < len(txResponse.Load); j++ {
				if txResponse.Load[j].Query == pb2.TransactionResponse_GETSTATE {
					var state pb2.GetState
					err = proto.Unmarshal(txResponse.Load[j].Payload, &state)
					if !checkState(&state) {
						block.mask[i] = false
					}
				}
			}
			for j := 0; j < len(txResponse.Load); j++ {
				if txResponse.Load[j].Query == pb2.TransactionResponse_PUTSTATE {
					var state pb2.PutState
					err = proto.Unmarshal(txResponse.Load[j].Payload, &state)
					putState(&state)
				} else if txResponse.Load[j].Query == pb2.TransactionResponse_DELSTATE {
					var state pb2.DeleteState
					err = proto.Unmarshal(txResponse.Load[j].Payload, &state)
					delState(&state)
				}
			}
		}
	}
}

func checkState(in *pb2.GetState) bool {
	docsQuery, err := db.Query(nil, `_id=="`+in.Key+`"`, nil, nil, nil, nil)
	if err != nil {
		fmt.Println(err)
		return false
	}
	if docsQuery[0]["_rev"] == in.Version {
		return true
	}
	return false
}

func putState(in *pb2.PutState) {
	connectDB()
	docsQuery := make([]map[string]interface{}, 1)
	doc := make(map[string]interface{})
	var err error
	fmt.Println(in)
	if in.Id == fmt.Sprint(nil) {
		fmt.Println("goes here")
		err := db.Available()
		fmt.Println(err)
		if err == nil {
			connectDB()
		}
		fmt.Println("here np")
		uuid := couchdb.GenerateUUID()
		for {
			if db.Contains(uuid) != nil {
				break
			} else {
				uuid = couchdb.GenerateUUID()
			}
		}

		in.Id = uuid
	} else {
		docsQuery, err = db.Query(nil, `_id=="`+in.Id+`"`, nil, nil, nil, nil)
		doc = docsQuery[0]
	}

	for i := 0; i < len(in.Val); i++ {
		key := in.Val[i].Key
		doc[key] = in.Val[i].Value
	}
	doc["_id"] = in.Id
	docsQuery[0] = doc
	if err != nil {
		return
	}
	_, _ = db.Update(docsQuery, nil)

}

func delState(in *pb2.DeleteState) {

}

func connectDB() {

	s, err := couchdb.NewServer(DefaultBaseURL)
	if err != nil {
		// fmt.Println("hey")
		fmt.Println(err)
	}
	s.Login("admin", "admin")
	// var db *couchdb.Database
	db, err = s.Get("new")

	if err != nil {
		// fmt.Println("heyiiiiiii")
		fmt.Println(err)
	}
	err = db.Available()
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println("ConnectDB called")
}

// CheckEndorcement for checking endoecement
func (block *BlockInfo) CheckEndorcement(in *pb.EndorsedTx, i int, wg *sync.WaitGroup) {
	block.mask[i] = true
	wg.Done()
}

// SetEndoresedTxs for set endoresedTxs value
func (block *BlockInfo) SetEndoresedTxs(in []*pb.EndorsedTx) {
	block.endorsedTxs = in
}

// SetSign for set sign
func (block *BlockInfo) SetSign(in string) {
	block.sign = in
}

// SetOffset for set offset
func (block *BlockInfo) SetOffset(in string) {
	block.offset = in
}

// SetMask for initializing the array of mask
func (block *BlockInfo) SetMask(n int) {
	block.mask = make([]bool, n)
}
