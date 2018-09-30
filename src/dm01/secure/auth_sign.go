package secure

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
)

// Mac 七牛AK/SK的对象，AK/SK可以从 https://portal.qiniu.com/user/key 获取。
type Mac struct {
	AccessKey string
	SecretKey []byte
}

// NewMac 构建一个新的拥有AK/SK的对象
func NewMac(accessKey, secretKey string) (mac *Mac) {
	return &Mac{accessKey, []byte(secretKey)}
}

// Sign 对数据进行签名，一般用于私有空间下载用途
func (mac *Mac) Sign(data []byte) (token string) {
	h := hmac.New(sha1.New, mac.SecretKey)
	h.Write(data)

	sign := base64.URLEncoding.EncodeToString(h.Sum(nil))
	return fmt.Sprintf("%s:%s", mac.AccessKey, sign)
}

// SignWithData 对数据进行签名，一般用于上传凭证的生成用途
func (mac *Mac) SignWithData(b []byte) (token string) {
	encodedData := base64.URLEncoding.EncodeToString(b)
	h := hmac.New(sha1.New, mac.SecretKey)
	h.Write([]byte(encodedData))
	digest := h.Sum(nil)
	sign := base64.URLEncoding.EncodeToString(digest)
	return fmt.Sprintf("%s:%s:%s", mac.AccessKey, sign, encodedData)
}