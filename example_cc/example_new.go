package main

import (
	"fmt"
	"strconv"
	cont "study/GitHub/consensus/core/contract"

	pb "../protoc/contractcode"
)

type SimpleCode struct{}

func (t *SimpleCode) Init(sup *cont.SupportCodeInterface, tDetail transactionDetail) (code int) {

	code = 200
	return code
}

func (t *SimpleCode) Invoke(sup *cont.SupportCodeInterface, tDetail transactionDetail) (code int) {
	channel = ch
	if tx == "AddUser" {
		if len(args) == 2 {
			return addUser(sup, args, tDetail)
		}

	} else if in.Transaction == "TransferAmount" {
		return transferAmount(sup, args, tDetail)
	}
	return &pb.Response{Result: nil}, nil
	code = 200
	return code
}

func addUser(sup *cont.SupportCodeInterface, args []string, tDetail transactionDetail) {

	// uuid := sup.GenerateUUID()
	doc := make([]map[string]interface{})

	doc[0]["_id"] = nil
	doc[0]["name"] = args[0]
	doc[0]["balance"] = strconv.Itoa(int(args[1]))

	id, _, err = sup.PutState(doc, tDetail)
	return 200

}

// TransferAmount for transfer amount from one user to other
func transferAmount(sup *cont.SupportCodeInterface, args []string, tDetail transactionDetail) {

	// args[0] = UserId
	// args[1] = FromAccID
	// args[2] = ToAccID
	// args[3] = Amount
	if args[0] == args[1] {
		docsQuery, err := sup.GetState(args[1], tDetail)

		if err != nil {
			fmt.Println(err)
			return &pb.StateInfo{Msg: pb.StateInfo_ERROR}, nil
		}

		getState := make([]map[string]interface{}, 1)
		getState[0] = docsQuery[0]

		if d.Contains(args[2]) == nil {
			fromBalance, err := strconv.Atoi(docsQuery[0]["balance"].(string))
			amount, err := strconv.Atoi(args[3])

			if err == nil && fromBalance >= amount {
				docsQuery2, err := sup.GetState(args[2], tDetail)
				getState = append(getState, docsQuery2[0])
				toBalance, err := strconv.Atoi(docsQuery2[0]["balance"].(string))

				if err == nil {

					docsQuery[0]["balance"] = strconv.Itoa(fromBalance - amount)
					docsQuery2[0]["balance"] = strconv.Itoa(toBalance + amount)

					sup.PutState(docsQuery[0], tDetail)
					sup.PutState(docsQuery2[0], tDetail)

				}
			}
		}
	}
}

func main() {

	cont.Start(t)
}
