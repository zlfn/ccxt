// -------------------------------------------------------------------------------

// PLEASE DO NOT EDIT THIS FILE, IT IS GENERATED AND WILL BE OVERWRITTEN:
// https://github.com/ccxt/ccxt/blob/master/CONTRIBUTING.md#how-to-contribute-code

// -------------------------------------------------------------------------------

package ccxt

func (this *hyperliquid) PublicPostInfo (args ...interface{}) <-chan interface{} {
   return this.callEndpointAsync("publicPostInfo", args...)
}

func (this *hyperliquid) PrivatePostExchange (args ...interface{}) <-chan interface{} {
   return this.callEndpointAsync("privatePostExchange", args...)
}
