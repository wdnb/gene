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

func(udb *UserDb) Insert(user UserName)  error{

	//udb := NewUserDB()
	err := udb.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(miniProgramsBucket))
		unm := b.Put([]byte(user.UserName),[]byte(user.UserName))
		fmt.Println(unm)
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	//defer udb.db.Close()
	//fmt.Println("注册 successed!")
	return  nil
}

func(udb *UserDb) IsExist(user UserName)  error{
	err := udb.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(miniProgramsBucket))
		unm := b.Get([]byte(user.UserName))
		if nil==unm {
			err := errors.New("用户不存在")
			//fmt.Println(err)
			return err
		}
		//fmt.Println(unm)
		return nil
	})

	if err != nil {
		return err
		//log.Panic(err)
	}

	return  nil
	//fmt.Println(unm)
	//fmt.Println("认证成功")
}

//获取数据库对象
func NewUserDB() *UserDb{
	dbFile := fmt.Sprintf(dbFile)
	if dbExists(dbFile) == false {
		fmt.Println("No existing "+dbFile+" found. Create one first.")
		os.Exit(1)
	}

	//var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		tx.Bucket([]byte(miniProgramsBucket))
		//tip = b.Get([]byte("l"))

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
