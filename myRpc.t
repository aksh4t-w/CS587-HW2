MyRpc interface {
  Transfer(from string, to string, amount uint64) (reply string, err string)
  Deposit(user string, amount uint64) (reply string, err string)
  Withdraw(user string, amount uint64) (reply string, err string)
  GetBalance(user string) (reply uint64, err string)
}
