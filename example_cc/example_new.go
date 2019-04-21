package main

import (
	"fmt"
	"strconv"

	cont "../core/contract"
)

// SimpleCode struct for contractCode
type SimpleCode struct{}

// Init for initialization
func (t *SimpleCode) Init(sup *cont.ContractDetail, tDetail *cont.TransactionDetail) (code int) {

	code = 200
	return code
}

// Invoke for tx invokation
func (t *SimpleCode) Invoke(sup *cont.ContractDetail, tx string, args []string, tDetail *cont.TransactionDetail) (code int) {
	// channel = ch
	if tx == "AddUser" {
		if len(args) == 2 {
			i := addUser(sup, args, tDetail)
			return i
		}

	} else if tx == "TransferAmount" {
		i := transferAmount(sup, args, tDetail)
		return i
	}
	// return &pb.Response{Result: nil}, nil
	code = 200
	return code
}

func addUser(sup *cont.ContractDetail, args []string, tDetail *cont.TransactionDetail) int {

	doc := make(map[string]interface{})

	doc["_id"] = nil
	doc["name"] = args[0]
	doc["balance"], _ = strconv.Atoi((args[1]))
	// uuid := sup.GenerateUUID()
	docs := make([]map[string]interface{}, 0)

	// doc[0]["_id"] = nil
	// doc[0]["name"] = args[0]
	// doc[0]["balance"], _ = strconv.Atoi((args[1]))
	//sup.PutState()
	docs = append(docs, doc)
	sup.PutState(docs, tDetail)
	return 200
}

// TransferAmount for transfer amount from one user to other
func transferAmount(sup *cont.ContractDetail, args []string, tDetail *cont.TransactionDetail) int {

	// args[0] = UserId
	// args[1] = FromAccID
	// args[2] = ToAccID
	// args[3] = Amount
	if args[0] == args[1] {
		docsQuery, err := sup.GetState(args[1], tDetail)

		if err != nil {
			fmt.Println(err)
			return 500
			// return &pb.StateInfo{Msg: pb.StateInfo_ERROR}, nil
		}

		// getState := make([]map[string]interface{}, 1)
		// getState[0] = docsQuery[0]
		// fmt.Printf(docsQuery[0])
		balA := fmt.Sprint(docsQuery[0]["balance"])
		fromBalance, err := strconv.Atoi(balA)
		amount, err := strconv.Atoi(args[3])

		if err == nil && fromBalance >= amount {
			docsQuery2, err := sup.GetState(args[2], tDetail)
			// getState = append(getState, docsQuery2[0])
			balB := fmt.Sprint(docsQuery2[0]["balance"])
			toBalance, err := strconv.Atoi(balB)

			if err == nil {

				docsQuery[0]["balance"] = strconv.Itoa(fromBalance - amount)
				docsQuery2[0]["balance"] = strconv.Itoa(toBalance + amount)

				sup.PutState(docsQuery, tDetail)
				sup.PutState(docsQuery2, tDetail)

			} else {
				return 500
			}
		} else {
			return 500
		}
	} else {
		return 500
	}
	return 200
}

func main() {

	cont.Start(new(SimpleCode))
}
