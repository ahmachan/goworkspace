package storage

// 资源管理相关的默认域名
const (
	DefaultRsHost  = "rs.qiniu.com"
	DefaultRsfHost = "rsf.qiniu.com"
	DefaultAPIHost = "api.qiniu.com"
	DefaultPubHost = "pu.qbox.me:10200"
)


// Config 为文件上传，资源管理等配置
type Config struct {
	Zone          *Zone  //空间所在的机房
	UseHTTPS      bool   //是否使用https域名
	UseCdnDomains bool   //是否使用cdn加速域名
	CentralRsHost string //中心机房的RsHost，用于list bucket
}