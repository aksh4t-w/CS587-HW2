package main

import (
	"ethos/syscall"
	"ethos/altEthos"
	"ethos/myRpc"
	"log"
	"fmt"
)

var myRpc_increment_counter uint64 = 0

type Account struct {
	name string
	balance uint64
}

var balances map[string]uint64

func init() {
	myRpc.SetupMyRpcTransfer(transfer)
	myRpc.SetupMyRpcGetBalance(getBalance)
	myRpc.SetupMyRpcDeposit(deposit)
	myRpc.SetupMyRpcWithdraw(withdraw)

}

func getBalance(user string) (myRpc.MyRpcProcedure) {
	balance, ok := balances[user]

	if ok {
		log.Printf("Balance in the requested account for %s: %v", user, balance)
		// balance = Account1.balance
		return &myRpc.MyRpcGetBalanceReply{balance, "ErrorNone"}
	}

	return &myRpc.MyRpcGetBalanceReply{0, "No user registered by the given name."}
}

func transfer(from string, to string, amount uint64) (myRpc.MyRpcProcedure) {
	balance1, ok1 := balances[from]
	_, ok2 := balances[to]
	

	if ok1 && ok2 {
		if balance1 < amount {
			return &myRpc.MyRpcTransferReply{"Error", "Insufficient balance."}
		}
		balances[from] -= amount
		balances[to] += amount
		log.Printf("Amount withdrawn from the account of %s, Current balance: %v", from, balances[from])
		log.Printf("Amount deposited in the account of %s, Current balance: %v", to, balances[to])

		return &myRpc.MyRpcTransferReply{"Transfer successful!", "ErrorNone"}
	}
	
	return &myRpc.MyRpcTransferReply{"Error", "No user registered by the given name."}
}

func deposit(user string, amount uint64) (myRpc.MyRpcProcedure) {
	balance, ok := balances[user]

	if ok {
		balances[user] = balance + amount
		log.Printf("Amount deposited in the account for %s, Current balance: %v", user, balances[user])
		res := fmt.Sprintf("Amount deposited in the account for %s, Current balance: %v", user, balances[user])

		return &myRpc.MyRpcDepositReply{res, "ErrorNone"}
	}

	return &myRpc.MyRpcDepositReply{"Error", "No user registered by the given name."}
}

func withdraw(user string, amount uint64) (myRpc.MyRpcProcedure) {
	balance, ok := balances[user]

	if ok {
		if amount <= balance {
			balances[user] = balance - amount
			log.Printf("Amount withdrawn in the account for %s: Current balance: %v", user, balances[user])
			res := fmt.Sprintf("Amount withdrawn successfully from the account for %s\n: Current balance: %v", user, balances[user])
			return &myRpc.MyRpcWithdrawReply{res, "ErrorNone"}
		} else {
			return &myRpc.MyRpcWithdrawReply{"Error", "Insufficient balance."}
		}
	}
	return &myRpc.MyRpcWithdrawReply{"Error", "No user registered by the given name."}
}


func main () {
	balances = make(map[string]uint64)

	balances["me"] = 10000
	balances["pat"] = 20000
	balances["bennet"] = 50000
	balances["gabriel"] = 5000


	altEthos.LogToDirectory("test/bankingServer")

	listeningFd, status := altEthos.Advertise("myRpc")
	if status != syscall.StatusOk {
		log.Println("Advertising service failed: ", status)
		altEthos.Exit(status)
	}

	for {
		_, fd, status := altEthos.Import(listeningFd)
		if status != syscall.StatusOk {
			log.Printf("Error calling Import: %v\n", status)
			altEthos.Exit(status)
		}

		log.Println("new connection accepted")

		t := myRpc.MyRpc{}
		altEthos.Handle(fd, &t)
	}
}
