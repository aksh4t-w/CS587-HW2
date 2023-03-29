package main

import (
	"ethos/altEthos"
	"ethos/syscall"
	"ethos/myRpc"
	"log"
)

// var name = "me"

func init() {

	myRpc.SetupMyRpcTransferReply(transferReply)
	myRpc.SetupMyRpcDepositReply(depositReply)
	myRpc.SetupMyRpcGetBalanceReply(getBalanceReply)
	myRpc.SetupMyRpcWithdrawReply(withdrawReply)
	
}

func getBalanceReply(balance uint64, err string) (myRpc.MyRpcProcedure) {
	if (err != "ErrorNone") {
		log.Printf("Error: %v", err)
		return nil
	}
	log.Printf("Available Balance: %v\n", balance)
	return nil
}

func depositReply(reply string, err string) (myRpc.MyRpcProcedure) {
	if (err != "ErrorNone") {
		log.Printf("Error: %v", err)
		return nil
	}
	log.Printf(reply)
	return nil
}

func withdrawReply(reply string, err string) (myRpc.MyRpcProcedure) {
	if (err != "ErrorNone") {
		log.Printf("Error: %v", err)
		return nil
	}
	log.Printf(reply)
	return nil
}

func transferReply(reply string, err string) (myRpc.MyRpcProcedure) {
	if (err != "ErrorNone") {
		log.Printf("Error: %v", err)
		return nil
	}
	log.Printf(reply)
	return nil
}


func main () {

	altEthos.LogToDirectory("test/bankingClient")
	
	log.Println("Performing calls now:\n")

	name := altEthos.GetUser()
	log.Println("Client User: ", name)

	fd, status := altEthos.IpcRepeat("myRpc", "", nil)
	if status != syscall.StatusOk {
		log.Printf("Ipc failed: %v\n", status)
		altEthos.Exit(status)
	}

	call1 := myRpc.MyRpcGetBalance{name}
	
	status = altEthos.ClientCall(fd, &call1)
	if status != syscall.StatusOk {
		log.Printf("clientCall failed: %v\n", status)
		altEthos.Exit(status)
	}

	fd, status = altEthos.IpcRepeat("myRpc", "", nil)
	if status != syscall.StatusOk {
		log.Printf("Ipc failed: %v\n", status)
		altEthos.Exit(status)
	}

	call2 := myRpc.MyRpcDeposit{name, 2000}
	
	status = altEthos.ClientCall(fd, &call2)
	if status != syscall.StatusOk {
		log.Printf("clientCall failed: %v\n", status)
		altEthos.Exit(status)
	}

	fd, status = altEthos.IpcRepeat("myRpc", "", nil)
	if status != syscall.StatusOk {
		log.Printf("Ipc failed: %v\n", status)
		altEthos.Exit(status)
	}

	call3 := myRpc.MyRpcWithdraw{name, 1000}
	
	status = altEthos.ClientCall(fd, &call3)
	if status != syscall.StatusOk {
		log.Printf("clientCall failed: %v\n", status)
		altEthos.Exit(status)
	}

	fd, status = altEthos.IpcRepeat("myRpc", "", nil)
	if status != syscall.StatusOk {
		log.Printf("Ipc failed: %v\n", status)
		altEthos.Exit(status)
	}

	call4 := myRpc.MyRpcTransfer{name, "gabriel", 2000}
	
	status = altEthos.ClientCall(fd, &call4)
	if status != syscall.StatusOk {
		log.Printf("clientCall failed: %v\n", status)
		altEthos.Exit(status)
	}

	log.Println("bankingClient: done")
}