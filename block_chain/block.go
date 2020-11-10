package block_chain

import (
	"BlockChainCertDataPorject/utils_BCCDP"
	"bytes"
	"encoding/gob"
	"math/big"
	"time"
)

type Block struct {
	Height     int64  //区块的高度
	TimeStamp  int64  //区块创建的时间戳
	PrevHash    []byte //上一个区块的hash
	Data       []byte //本区块要存储的数据
	Hash       []byte //本区块的hash
	Nonce      int64  //随机数
	Version    string //版本号
	TimeFormat string //用于格式化输出时间
	TimeFree   int64  //创建区块所花费的时间
}

func NewBlock(height int64, prevHash, data []byte) Block {
	newblock := Block{
		Height:    height,
		TimeStamp: time.Now().Unix(),
		PrevHash:   prevHash,
		Data:      data,
		Hash:      nil,
		Nonce:     0,
		Version:   "0 x 01",
	}
	pow := NewPoW(newblock)
	newblock.Nonce, newblock.Hash, newblock.TimeFree = pow.Run()
	return newblock
}
func Block2Bytes(block Block) ([]byte, error) {
	heightBytes, err := utils_BCCDP.Int64ToByte(block.Height)
	if err != nil {
		return nil, err
	}
	timeBytes, err := utils_BCCDP.Int64ToByte(block.TimeStamp)
	if err != nil {
		return nil, err
	}
	versionBytes := []byte(block.Version)

	blockBytes := bytes.Join(
		[][]byte{
			heightBytes, timeBytes, versionBytes, block.PrevHash, block.Data,
		}, []byte{})
	return blockBytes, nil
}
func (b Block)Serialize()([]byte)  {
	buff := new(bytes.Buffer)
	gob.NewEncoder(buff).Encode(b)
	return buff.Bytes()
}
func DeSerialize(data []byte)(*Block,error)  {
	var block *Block
	//Decode（）要取变量的地址
	err := gob.NewDecoder(bytes.NewReader(data)).Decode(&block)
	if err != nil {
		return  nil,err
	}
	return block,nil
}