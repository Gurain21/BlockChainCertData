"# BlockChianCertData"
存用户数据到区块链上的时候
        
        certRecord := models.CertRecord{
		CertId:   fileMd5HashBytes,//这里要放hash后的切片,而不是哈希后的16进制的string强转[]byte
		CertHash: fileSha256HashBytes,//这里要放hash后的切片,而不是哈希后的16进制的string强转[]byte
		CertName: user.Name,
		Phone:    user.Phone,
		CertCard: user.Card,
		FileName: fileHeader.Filename,
		FileSize: fileHeader.Size,
		CertTime: time.Now().Unix(),
	}

从区块链上查询数据的时候
       
        if hex.EncodeToString(certRecord.CertId) == cert_id { //if成立，找到区块了
				block = eachBlock
				break
			}
			eachBig.SetBytes(eachBlock.PrevHash)
			if eachBig.Cmp(zeroBig) == 0 { //到创世区块了，停止遍历
				break
			}
			eachHash = eachBlock.PrevHash
		}
		


    ...
	type CertRecord struct {
    	CertId         []byte //文件的md5 哈希值
    	CertHash       []byte //文件的sha256哈希值
    	CertName       string //保存人姓名
    	Phone          string //保存人电话
    	CertCard       string //保存人身份证
    	FileName       string //文件名
    	FileSize       int64  //文件大小
    	CertTime       int64  //文件保存时间
    	CertIdFormat   string //仅用于md5 []byte格式化输出
    	CertHashFormat string //仅用于sha256 []byte格式化输出
    	CertTimeFormat string //仅用于时间格式化输出
    }
然后返回数据给证书展示页面时做一个转化
    
        certRecord.CertIdFormat = hex.EncodeToString(certRecord.CertId)
	certRecord.CertHashFormat = hex.EncodeToString(certRecord.CertHash)
	certRecord.CertTimeFormat = utils_BCCDP.TimeFormat(certRecord.CertTime,utils_BCCDP.TIME_FORMAT_ONE)	
	
一些小细节问题
模块语法 未遍历时结构体时先访问类,再访问属性	

	    {{.CertRecord.CertHashFormat}}
		
遍历时	

    {{range .CertRecord}}
    {{.CertRecord的属性名}}
    {{.CertRecord的属性名}}
    {{.CertRecord的属性名}}
    {{end}}