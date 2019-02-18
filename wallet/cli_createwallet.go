package wallet

func CreateWallet(nodeID string) string {
	wallets, _ := NewWallets(nodeID)
	address := wallets.CreateWallet()
	wallets.SaveToFile(nodeID)
	return address
}
