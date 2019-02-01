package debug

import (
	"github.com/wdnb/gene/blockchain"
	"testing"
)

func TestSend(t *testing.T)  {
	blockchain.Send("1QCvT5FFo3G3nEkKWRzDZebMcXVucLUFsD","1KC1fEfnooCuVSFVXMR6yntHRcEEpjPTJY",1,"3000",true);
}
