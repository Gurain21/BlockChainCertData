package block_chain

import (
	"BlockChainCertDataPorject/utils_BCCDP"
	"bytes"
	"math/big"
	"time"
)

type ProofOfWork struct {
	Target *big.Int
	Block  Block
}

var Difficulty uint = 10

func NewPoW(block Block) ProofOfWork {
	t := big.NewInt(1)
	t.Lsh(t, 255-Difficulty)
	pow := ProofOfWork{
		Target: t,
		Block:  block,
	}
	return pow
}

func (p ProofOfWork) Run() (int64, []byte, int64) {
	var nonce int64 = 0
	blockBytes, _ := Block2Bytes(p.Block)
	var blockHash []byte
	var timeFree int64
	startTime := time.Now().Unix()
	for {
		nonceBytes, _ := utils_BCCDP.Int64ToByte(nonce)
		blockBytes := bytes.Join(
			[][]byte{
				blockBytes, nonceBytes,
			}, []byte{})
		blockHash = utils_BCCDP.SHA256HashByte(blockBytes)
		blockHashBig := new(big.Int)
		blockHashBig.SetBytes(blockHash)
		if blockHashBig.Cmp(p.Target) == -1 {
			endTime := time.Now().Unix()
			timeFree = endTime - startTime
			break
		}
		nonce++
	}
	return nonce, blockHash, timeFree

}
