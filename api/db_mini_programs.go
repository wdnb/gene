package api

import (
	"errors"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"os"
)
const dbFile = "./db/users.db"
const miniProgramsBucket = "mini_programs_users"

type UserDb struct {
	db  *bolt.DB
}

//获取用户钱包地址
func(udb *UserDb) GetAddress(uname string)  ([]byte,error){
	var address []byte
	err := udb.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(miniProgramsBucket))
		address = b.Get([]byte(uname))
		if nil==address {
			err := errors.New("用户不存在")
			return err
		}
		return nil
	})
	if err != nil {
		return nil,err
	}
	return  address,nil
}

func(udb *UserDb) Insert(user UserName,address string)  error{
	err := udb.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(miniProgramsBucket))
		unm := b.Put([]byte(user.UserName),[]byte(address))//key 用户名 value 钱包地址
		fmt.Println(unm)
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	return  nil
}

func(udb *UserDb) IsExist(user UserName)  error{
	err := udb.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(miniProgramsBucket))
		unm := b.Get([]byte(user.UserName))
		if nil==unm {
			err := errors.New("用户不存在")
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}

	return  nil
}

//获取用户数据库对象
func NewUserDB() *UserDb{
	dbFile := fmt.Sprintf(dbFile)
	if dbExists(dbFile) == false {
		fmt.Println("No existing "+dbFile+" found. Create one first.")
		os.Exit(1)
	}
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		tx.Bucket([]byte(miniProgramsBucket))
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	bc := UserDb{ db}
	return &bc
}

//创建数据库
func CreateUserDB() *UserDb {
	dbFile := fmt.Sprintf(dbFile)
	if dbExists(dbFile) {
		fmt.Println(dbFile+" already exists.")
		os.Exit(1)
	}
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}
	bc := UserDb{ db}
	return &bc
}

//创建小程序用户桶
func (udb *UserDb)CreateMiniProgramUserBucket() error{
	err := udb.db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte(miniProgramsBucket))
		if err != nil {
			log.Panic(err)
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	return nil
}

func dbExists(dbFile string) bool {
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		return false
	}
	return true
}
