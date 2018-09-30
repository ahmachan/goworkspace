package storage

import (
	//"fmt"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"dm01/secure"
)

// PutPolicy 表示文件策略
type PutPolicy struct {
	Scope               string `json:"scope"`
	Expires             uint32 `json:"deadline"` // 截止时间（以秒为单位）
	FsizeLimit          int64  `json:"fsizeLimit,omitempty"`
	MimeLimit           string `json:"mimeLimit,omitempty"`
	DeleteAfterDays     int    `json:"deleteAfterDays,omitempty"`
}

// UploadToken 方法用来进行上传凭证的生成
func (p *PutPolicy) UploadToken(mac *secure.Mac) (token string) {
	if p.Expires == 0 {
		p.Expires = 3600 // 1 hour
	}
	p.Expires += uint32(time.Now().Unix())

	putPolicyJSON, _ := json.Marshal(p)
	
	token = mac.SignWithData(putPolicyJSON)
	return
}

func (p *PutPolicy) DecodeWithToken(token string) (ak, bucket string, err error){
	items := strings.Split(token, ":")
	if len(items) != 3 {
		err = errors.New("invalid upload token, format error")
		return
	}

	ak = items[0]
	policyBytes, dErr := base64.URLEncoding.DecodeString(items[2])
	if dErr != nil {
		err = errors.New("invalid upload token, invalid put policy")
		return
	}

	putPolicy := PutPolicy{}
	uErr := json.Unmarshal(policyBytes, &putPolicy)
	//fmt.Println(uErr)
	if uErr != nil {
		err = errors.New("invalid upload token, invalid put policy")
		return
	}

	bucket = strings.Split(putPolicy.Scope, ":")[0]
	return
}