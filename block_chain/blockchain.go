package block_chain

import (
	"errors"
	"github.com/boltdb/bolt"
	"math/big"
)

const BLOCKCHAINDB = "blockChain.db"
const LASTHASH = "lastHash"
const BUCKET_NAMEM_BLOCKS = "blocks"

type BlockChian struct {
	LastHash []byte //最新区块哈希
	BoltDB   *bolt.DB
}

var CHAIN *BlockChian

//该函数 返回一条区块链，如果没有，则创建一条区块链再返回
func NewBlockChain() *BlockChian {
	var bc *BlockChian
	//1、打开数据库
	db, err := bolt.Open(BLOCKCHAINDB, 0600, nil)
	db.Update(func(tx *bolt.Tx) error {
		//查看blockchain文件
		bucket := tx.Bucket([]byte(BUCKET_NAMEM_BLOCKS)) //先拿桶
		if bucket == nil {                               //如果没有这个桶，则新建一个桶
			bucket, err = tx.CreateBucket([]byte(BUCKET_NAMEM_BLOCKS))
			if err != nil {
				panic(err.Error())
			}
		}
		//查看桶中是否存在区块
		lastHash := bucket.Get([]byte(LASTHASH))
		if len(lastHash) == 0 {
			//桶中没有lastHash记录，需要新建创世区块
			genesis := NewBlock(0, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, nil)

			//把创世区块放到桶里去
			bucket.Put(genesis.Hash, genesis.Serialize())

			bucket.Put([]byte(LASTHASH), genesis.Hash)
			bc = &BlockChian{
				LastHash: genesis.Hash,
				BoltDB:   db,
			}
		} else {
			bc = &BlockChian{
				LastHash: lastHash,
				BoltDB:   db,
			}
		}
		return err
	})
	CHAIN = bc
	return bc
}
func (bc BlockChian) QueryAllBlocks() ([]*Block, error) {
	var err error
	db := bc.BoltDB
	blocks := make([]*Block, 0)
	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAMEM_BLOCKS))
		if bucket == nil {
			err = errors.New("查询区块链数据失败，请重试！")
			return err
		}
		eachHash := bc.LastHash
		eachBig := new(big.Int)
		zeroBig := big.NewInt(0)
		for {
			eachBlockBytes := bucket.Get(eachHash)
			eachBlock, err := DeSerialize(eachBlockBytes)
			if err != nil {
				return err
			}
			blocks = append(blocks, eachBlock)
			eachBig.SetBytes(eachBlock.PreHash)
			if eachBig.Cmp(zeroBig) == 0 {
				break
			}
			eachHash = eachBlock.PreHash
		}

		return err
	})
	return blocks, err
}
func (bc BlockChian) QueryBlockByHEight(height int64) (*Block, error) {
	var err error
	if height < 0 {
		err = errors.New("读取区块数据失败,请重新确认区块高度")
		return nil, err
	}
	db := bc.BoltDB
	var targetBlock *Block
	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAMEM_BLOCKS))
		if bucket == nil {
			err = errors.New("读取区块数据失败")
			return err
		}
		lastBlockBytes := bucket.Get(bc.LastHash)
		lastBlock, err := DeSerialize(lastBlockBytes)
		if lastBlock.Height < height {
			err = errors.New("读取区块链数据失败,您请确认输入的高度")
			return err
		}
		eachHash := bc.LastHash
		for {
			eachBlockBytes := bucket.Get(eachHash)
			targetBlock, err = DeSerialize(eachBlockBytes)
			if err != nil {
				return err
			}
			if targetBlock.Height == height {
				break
			} else {
				eachHash = targetBlock.PreHash
			}

		}

		return err
	})
	if err != nil {
		return nil,err
	}
	return targetBlock, err
}

//将用户的数据保存到新创建的一个区块中，然后把这个新的区块添加到区块链上去
func (bc *BlockChian) SaveData(data []byte) (Block, error) {
	//第一步：拿到区块链上最新的那个区块
	var err error
	var lastBlock *Block //新建区块需要最新区块的高度和它的hash
	db := bc.BoltDB
	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAMEM_BLOCKS))
		if bucket == nil {
			err = errors.New("读取区块链数据失败。")
			return err
		}
		lastBlockBytes := bucket.Get(bc.LastHash)
		lastBlock, err = DeSerialize(lastBlockBytes)
		if err != nil {
			return err
		}
		return err
	})
	//第二步：通过最新区块的hash值和高度，以及用户要保存的数据新建一个区块
	newBlock := NewBlock(lastBlock.Height+1, lastBlock.Hash, data)

	//第三步：把这个新建的区块保存到区块链上，并更新区块链的最新区块（lastHash）
	err = db.Update(func(tx *bolt.Tx) error {

		bucket := tx.Bucket([]byte(BUCKET_NAMEM_BLOCKS))
		bucket.Put(newBlock.Hash, newBlock.Serialize())
		bucket.Put([]byte(LASTHASH), newBlock.Hash)
		bc.LastHash = newBlock.Hash

		return err
	})
	return newBlock, err
}
