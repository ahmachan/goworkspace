package storage

import (
	"math/rand"
	"fmt"
	"errors"	
)

// PutExtra 为表单上传的额外可选项
type PutExtra struct {
	// 可选，用户自定义参数，必须以 "x:" 开头。若不以x:开头，则忽略。
	Params map[string]string
	// 可选，当为 "" 时候，服务端自动判断。
	MimeType string
}

// FormDmeroader 表示一个表单DM的对象
type FormDmer struct {
	FormId  string
	Cfg      *Config
}

//首字母大写对外，参数首字母小写不对外private
type FormValues struct {
	FormId       string
	Token        string
	Key          string
	Size         int64	
}

type formData struct {
	name      string
	hash     string
}
func (p *formData) renderForm(name string ) *formData {
	hash:=fmt.Sprintf("f_%s_%d", name,rand.Int())
	return &formData{
		name:name,
		hash:hash,
	}
}

// PutRet 为七牛标准的上传回复内容。
// 如果使用了上传回调或者自定义了returnBody，那么需要根据实际情况，自己自定义一个返回值结构体
type PutRet struct {
	Hash         string `json:"hash"`
	Key          string `json:"key"`
}


// NewFormUploader 用来构建一个表单DM的对象
func NewFormDmer(cfg *Config,form_id string) *FormDmer {
	if cfg == nil {
		cfg = &Config{}
	}

	//fmt.Printf("\n Config数据:")
	//fmt.Println(cfg)

	return &FormDmer{
		FormId : form_id,
		Cfg:     cfg,
	}
}


func fetchFormValues(form_id string,key string,token string,size int64) *FormValues {
	data:=formData{}
	rf:=data.renderForm(form_id)
	//return FormValues{}是不行的，必须前加&取“指针”
	return &FormValues{
		FormId       :rf.hash,
		Token        :token,
		Key          :key,
		Size         :size,
	}
}


func (p *FormDmer) PutFetch(uptoken string,key string, size int64, extra *PutExtra) (fetchRes *FormValues,err error) {
	if extra == nil {
		extra = &PutExtra{}
	}

	fetchRes=fetchFormValues("D1888",key,uptoken,size)
	if fetchRes == nil {
		err = errors.New("invalid fetch value")
		return
	}

	fmt.Printf("\n form value:\n")
	fmt.Println(fetchRes)
	return
}


func (p *FormDmer) UpHost(ak, bucket string) (upHost string, err error) {
	var zone *Zone
	if p.Cfg.Zone != nil {
		zone = p.Cfg.Zone
	} else {
		if v, zoneErr := GetZone(ak, bucket); zoneErr != nil {
			err = zoneErr
			return
		} else {
			zone = v
		}
	}

	scheme := "http://"
	if p.Cfg.UseHTTPS {
		scheme = "https://"
	}

	host := zone.SrcUpHosts[0]
	if p.Cfg.UseCdnDomains {
		host = zone.CdnUpHosts[0]
	}

	upHost = fmt.Sprintf("%s%s", scheme, host)
	return
}