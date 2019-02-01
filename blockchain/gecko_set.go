package blockchain

import (
	"encoding/hex"
	"github.com/boltdb/bolt"
	"log"
)

const geckoBucket = "chaingecko"

//面向守宫编程
type GeckoSet struct {
	Blockchain *Blockchain
}

func (g GeckoSet) FindGecko(pubKeyHash []byte) []TXOutput {
	return nil
	//var Geckos []TXOutput
	//db := g.Blockchain.db
	//
	//err := db.View(func(tx *bolt.Tx) error {
	//	b := tx.Bucket([]byte(geckoBucket))
	//	c := b.Cursor()
	//
	//	for k, v := c.First(); k != nil; k, v = c.Next() {
	//		outs := DeserializeOutputs(v)
	//
	//		for _, out := range outs.Outputs {
	//			if out.IsLockedWithKey(pubKeyHash) {
	//				Geckos = append(Geckos, out)
	//			}
	//		}
	//	}
	//
	//	return nil
	//})
	//if err != nil {
	//	log.Panic(err)
	//}
	//
	//return Geckos
}

func (u GeckoSet) Reindex() {
	// Reindex rebuilds the UTXO set
	db := u.Blockchain.db
	bucketName := []byte(geckoBucket)

	err := db.Update(func(tx *bolt.Tx) error {
		err := tx.DeleteBucket(bucketName)
		if err != nil && err != bolt.ErrBucketNotFound {
			log.Panic(err)
		}

		_, err = tx.CreateBucket(bucketName)
		if err != nil {
			log.Panic(err)
		}

		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	Gecko := u.Blockchain.FindUTXO()

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)

		for txID, outs := range Gecko {
			key, err := hex.DecodeString(txID)
			if err != nil {
				log.Panic(err)
			}

			err = b.Put(key, outs.Serialize())
			if err != nil {
				log.Panic(err)
			}
		}

		return nil
	})
}
