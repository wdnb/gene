package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/wdnb/gene/api"
	"github.com/wdnb/gene/gecko"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	"io"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
)

var debug = true
//var path  = "D:\workspace\go\src\github.com\wdnb\blockchain-tutorial"

//type  Message struct {
//	gecko.Gecko
//}

var geckos []gecko.Gecko
// Block represents each 'item' in the blockchain
type Block struct {
	Index     int
	Timestamp string
	BPM       int
	Hash      string
	PrevHash  string
	Validator string
}



// Blockchain is a series of validated Blocks
var Blockchain []Block
var tempBlocks []Block

// candidateBlocks handles incoming blocks for validation
var candidateBlocks = make(chan Block)

// announcements broadcasts winning validator to all nodes
var announcements = make(chan string)

var mutex = &sync.Mutex{}

// validators keeps track of open validators and balances
var validators = make(map[string]int)



func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%p, %T\n", err, err)
	// create genesis blockchain
	t := time.Now()
	genesisBlock := Block{}
	genesisBlock = Block{0, t.String(), 0, calculateBlockHash(genesisBlock), "", ""}
	spew.Dump(genesisBlock)
	Blockchain = append(Blockchain, genesisBlock)

	tcpPort := os.Getenv("TCP_PORT")
	httpPort := os.Getenv("HTTP_PORT")
	//接收前端http请求
	err2 := handleHttp(httpPort)
	if err2 !=nil {
		log.Fatal(err2)
	}

	//fmt.Printf("%p, %T\n", err, err)

	// start TCP and serve TCP server
	server, err3 := net.Listen("tcp", ":"+tcpPort)
	if err3 != nil {
		log.Fatal(err3)
	}
	log.Println("TCP Server Listening on port :", tcpPort)
	defer server.Close()

	go func() {
		for candidate := range candidateBlocks {
			mutex.Lock()
			tempBlocks = append(tempBlocks, candidate)
			mutex.Unlock()
		}
	}()

	go func() {
		for {
			pickWinner()
		}
	}()

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)
	}
}
//输入用户名密码-》
//账户登陆-》服务器认证-》账户操作==私钥签名 签名通过服务器进行
func handleHttp(httpPort string) error{
	//mux := makeMuxRouter()
	r := gin.Default()
	v1 := r.Group("/api/v1")
	v1.POST("/user/register", api.Register)
	v1.POST("/user/login", api.Login)
	v1.Use(api.TokenMiddleware())
	{
		v1.POST("/user/inquire", api.Inquire)
		l := v1.Group("leopard")
		l.POST("/create", api.LeopardCreate)
		l.POST("/retrieve", api.LeopardRetrieve)
		l.POST("/basegene", api.BaseGeneLoad)

		v1.POST("/BMI/c", api.BMICreate)
		v1.POST("/BMI/r", api.BMIRetrieve)
		v1.POST("/BMI/u", api.BMIUpdate)
		v1.POST("/BMI/d", api.BMIDelete)
	}

	log.Println("HTTP Server Listening on port :", httpPort)
	//router.Handle()
	//r.Run(httpPort)
	//if e:=r.Run(httpPort);e!=nil{
	//	log.Fatal(e)
	//}

	s := &http.Server{
		Addr:           ":" + httpPort,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

//func  ss(c *gin.Context )  {
//		c.JSON(200, gin.H{
//			"status":  "posted",
//			//"message": message,
//			//"nick":    nick,
//		})
//	return
//}

// create handlers
func makeMuxRouter() http.Handler {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", handleGetBlockchain).Methods("GET")

	muxRouter.HandleFunc("/api/v1/BMI", handleWriteBlock).Methods("POST")
	muxRouter.HandleFunc("/api/v1/egg", handleWriteBlock).Methods("POST")
	muxRouter.HandleFunc("/api/v1/reproduction", handleWriteBlock).Methods("POST")
	muxRouter.HandleFunc("/api/v1/care", handleWriteBlock).Methods("POST")
	//muxRouter.HandleFunc("/api/v1/test", api.Test()).Methods("POST")
	return muxRouter
}

// write blockchain when we receive an http request
func handleGetBlockchain(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.MarshalIndent(Blockchain, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}
// takes JSON payload as an input for heart rate (gecko)
func handleWriteBlock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var msg gecko.Gecko

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&msg); err != nil {
		api.RespondWithJSON(w, r, http.StatusBadRequest, r.Body)
		return
	}
	if err1 := api.DataVerification(msg);err1!=nil{//数据合法性检查
		api.RespondWithJSON(w, r, http.StatusOK, err1)
		return
	}

	defer r.Body.Close()
	//fmt.Println(msg.Hash)
	mutex.Lock()
	//prevBlock := Blockchain[len(Blockchain)-1]
	//newBlock := generateBlock(prevBlock, msg.BPM)
	//if isBlockValid(newBlock, prevBlock) {
	//	Blockchain = append(Blockchain, newBlock)
	//	spew.Dump(Blockchain)
	//}
	mutex.Unlock()
	//fmt.Println(msg)
	api.RespondWithJSON(w, r, http.StatusCreated, msg)
}

//type MyData struct {
//	Code   int    `json:"code"`
//	Msg   string    `json:"msg"`
//}

func dataVerification(msg gecko.Gecko)  interface{}{
	//if msg.Hash=="" {
	//	//err := fmt.Errorf("%s", "")
	//	//var err error = errors.New("哈希呢")
	//	return api.ErrorRepsonse(12,"哈希呢s")
	//}
	return nil
}

func respondWithJSON1(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
	response, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		return
	}
	w.WriteHeader(code)
	w.Write(response)
}

// pickWinner creates a lottery pool of validators and chooses the validator who gets to forge a blockchain to the blockchain
// by random selecting from the pool, weighted by amount of tokens staked
func pickWinner() {
	time.Sleep(30 * time.Second)
	mutex.Lock()
	temp := tempBlocks
	mutex.Unlock()

	lotteryPool := []string{}
	if len(temp) > 0 {

		// slightly modified traditional proof of stake algorithm
		// from all validators who submitted a blockchain, weight them by the number of staked tokens
		// in traditional proof of stake, validators can participate without submitting a blockchain to be forged
	OUTER:
		for _, block := range temp {
			// if already in lottery pool, skip
			for _, node := range lotteryPool {
				if block.Validator == node {
					continue OUTER
				}
			}

			// lock list of validators to prevent data race
			mutex.Lock()
			setValidators := validators
			mutex.Unlock()

			k, ok := setValidators[block.Validator]
			if ok {
				for i := 0; i < k; i++ {
					lotteryPool = append(lotteryPool, block.Validator)
				}
			}
		}

		// randomly pick winner from lottery pool
		s := rand.NewSource(time.Now().Unix())
		r := rand.New(s)
		lotteryWinner := lotteryPool[r.Intn(len(lotteryPool))]

		// add blockchain of winner to blockchain and let all the other nodes know
		for _, block := range temp {
			if block.Validator == lotteryWinner {
				mutex.Lock()
				Blockchain = append(Blockchain, block)
				mutex.Unlock()
				for _ = range validators {
					announcements <- "\nwinning validator: " + lotteryWinner + "\n"
				}
				break
			}
		}
	}

	mutex.Lock()
	tempBlocks = []Block{}
	mutex.Unlock()
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	go func() {
		for {
			msg := <-announcements
			io.WriteString(conn, msg)
		}
	}()
	// validator address
	var address string

	// allow user to allocate number of tokens to stake
	// the greater the number of tokens, the greater chance to forging a new blockchain
	io.WriteString(conn, "Enter token balance:")
	scanBalance := bufio.NewScanner(conn)
	for scanBalance.Scan() {
		balance, err := strconv.Atoi(scanBalance.Text())
		if err != nil {
			log.Printf("%v not a number: %v", scanBalance.Text(), err)
			return
		}
		t := time.Now()
		address = calculateHash(t.String())
		validators[address] = balance
		fmt.Println(validators)
		break
	}

	io.WriteString(conn, "\nEnter a new BPM:")

	scanBPM := bufio.NewScanner(conn)

	go func() {
		for {
			// take in BPM from stdin and add it to blockchain after conducting necessary validation
			for scanBPM.Scan() {
				bpm, err := strconv.Atoi(scanBPM.Text())
				// if malicious party tries to mutate the chain with a bad input, delete them as a validator and they lose their staked tokens
				if err != nil {
					log.Printf("%v not a number: %v", scanBPM.Text(), err)
					delete(validators, address)
					conn.Close()
				}

				mutex.Lock()
				oldLastIndex := Blockchain[len(Blockchain)-1]
				mutex.Unlock()

				// create newBlock for consideration to be forged
				newBlock, err := generateBlock(oldLastIndex, bpm, address)
				if err != nil {
					log.Println(err)
					continue
				}
				if isBlockValid(newBlock, oldLastIndex) {
					candidateBlocks <- newBlock
				}
				io.WriteString(conn, "\nEnter a new BPM:")
			}
		}
	}()

	// simulate receiving broadcast
	for {
		time.Sleep(time.Minute)
		mutex.Lock()
		output, err := json.Marshal(Blockchain)
		mutex.Unlock()
		if err != nil {
			log.Fatal(err)
		}
		io.WriteString(conn, string(output)+"\n")
	}

}

// isBlockValid makes sure blockchain is valid by checking index
// and comparing the hash of the previous blockchain
func isBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if calculateBlockHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}

// SHA256 hasing
// calculateHash is a simple SHA256 hashing function
func calculateHash(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

//calculateBlockHash returns the hash of all blockchain information
func calculateBlockHash(block Block) string {
	record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PrevHash
	return calculateHash(record)
}

// generateBlock creates a new blockchain using previous blockchain's hash
func generateBlock(oldBlock Block, BPM int, address string) (Block, error) {

	var newBlock Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.BPM = BPM
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateBlockHash(newBlock)
	newBlock.Validator = address

	return newBlock, nil
}
