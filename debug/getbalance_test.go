package debug

import (
	"github.com/wdnb/gene/blockchain"
	"testing"
)

func TestGetBalance(t *testing.T)  {
	blockchain.GetBalance("1QCvT5FFo3G3nEkKWRzDZebMcXVucLUFsD","3000")
}
