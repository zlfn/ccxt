// -------------------------------------------------------------------------------

// PLEASE DO NOT EDIT THIS FILE, IT IS GENERATED AND WILL BE OVERWRITTEN:
// https://github.com/ccxt/ccxt/blob/master/CONTRIBUTING.md#how-to-contribute-code

// -------------------------------------------------------------------------------

package ccxt

func (this *coinspot) PublicGetLatest (args ...interface{}) <-chan interface{} {
   return this.callEndpointAsync("publicGetLatest", args...)
}

func (this *coinspot) PrivatePostOrders (args ...interface{}) <-chan interface{} {
   return this.callEndpointAsync("privatePostOrders", args...)
}

func (this *coinspot) PrivatePostOrdersHistory (args ...interface{}) <-chan interface{} {
   return this.callEndpointAsync("privatePostOrdersHistory", args...)
}

func (this *coinspot) PrivatePostMyCoinDeposit (args ...interface{}) <-chan interface{} {
   return this.callEndpointAsync("privatePostMyCoinDeposit", args...)
}

func (this *coinspot) PrivatePostMyCoinSend (args ...interface{}) <-chan interface{} {
   return this.callEndpointAsync("privatePostMyCoinSend", args...)
}

func (this *coinspot) PrivatePostQuoteBuy (args ...interface{}) <-chan interface{} {
   return this.callEndpointAsync("privatePostQuoteBuy", args...)
}

func (this *coinspot) PrivatePostQuoteSell (args ...interface{}) <-chan interface{} {
   return this.callEndpointAsync("privatePostQuoteSell", args...)
}

func (this *coinspot) PrivatePostMyBalances (args ...interface{}) <-chan interface{} {
   return this.callEndpointAsync("privatePostMyBalances", args...)
}

func (this *coinspot) PrivatePostMyOrders (args ...interface{}) <-chan interface{} {
   return this.callEndpointAsync("privatePostMyOrders", args...)
}

func (this *coinspot) PrivatePostMyBuy (args ...interface{}) <-chan interface{} {
   return this.callEndpointAsync("privatePostMyBuy", args...)
}

func (this *coinspot) PrivatePostMySell (args ...interface{}) <-chan interface{} {
   return this.callEndpointAsync("privatePostMySell", args...)
}

func (this *coinspot) PrivatePostMyBuyCancel (args ...interface{}) <-chan interface{} {
   return this.callEndpointAsync("privatePostMyBuyCancel", args...)
}

func (this *coinspot) PrivatePostMySellCancel (args ...interface{}) <-chan interface{} {
   return this.callEndpointAsync("privatePostMySellCancel", args...)
}

func (this *coinspot) PrivatePostRoMyBalances (args ...interface{}) <-chan interface{} {
   return this.callEndpointAsync("privatePostRoMyBalances", args...)
}

func (this *coinspot) PrivatePostRoMyBalancesCointype (args ...interface{}) <-chan interface{} {
   return this.callEndpointAsync("privatePostRoMyBalancesCointype", args...)
}

func (this *coinspot) PrivatePostRoMyDeposits (args ...interface{}) <-chan interface{} {
   return this.callEndpointAsync("privatePostRoMyDeposits", args...)
}

func (this *coinspot) PrivatePostRoMyWithdrawals (args ...interface{}) <-chan interface{} {
   return this.callEndpointAsync("privatePostRoMyWithdrawals", args...)
}

func (this *coinspot) PrivatePostRoMyTransactions (args ...interface{}) <-chan interface{} {
   return this.callEndpointAsync("privatePostRoMyTransactions", args...)
}

func (this *coinspot) PrivatePostRoMyTransactionsCointype (args ...interface{}) <-chan interface{} {
   return this.callEndpointAsync("privatePostRoMyTransactionsCointype", args...)
}

func (this *coinspot) PrivatePostRoMyTransactionsOpen (args ...interface{}) <-chan interface{} {
   return this.callEndpointAsync("privatePostRoMyTransactionsOpen", args...)
}

func (this *coinspot) PrivatePostRoMyTransactionsCointypeOpen (args ...interface{}) <-chan interface{} {
   return this.callEndpointAsync("privatePostRoMyTransactionsCointypeOpen", args...)
}

func (this *coinspot) PrivatePostRoMySendreceive (args ...interface{}) <-chan interface{} {
   return this.callEndpointAsync("privatePostRoMySendreceive", args...)
}

func (this *coinspot) PrivatePostRoMyAffiliatepayments (args ...interface{}) <-chan interface{} {
   return this.callEndpointAsync("privatePostRoMyAffiliatepayments", args...)
}

func (this *coinspot) PrivatePostRoMyReferralpayments (args ...interface{}) <-chan interface{} {
   return this.callEndpointAsync("privatePostRoMyReferralpayments", args...)
}
