package user

import (
	"math/rand"
	"fmt"
	"dm01/storage"
	"dm01/secure"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	accessKey = "QDM_ACCESS_KEY"
	secretKey = "QDM_SECRET_KEY"
	bucket    = "QDM_TEST_BUCKET"
	testBucket = "TEST_BUCKET_1001"
)

// HealthCheck shows `OK` as the ping-pong result.
func HealthCheck(c *gin.Context) {
	message := "OK"
	c.String(http.StatusOK, "\n"+message)
}

// get user detail info
func GetUserInfo(c *gin.Context) {
	//message := "User Info"
	//c.String(http.StatusOK, "\n"+message)
	c.JSON(http.StatusOK, gin.H{"user_id": 10821038, "nickname": "leely"})
}

func DisMac(c *gin.Context) {
	// 简单上传凭证
	putPolicy := storage.PutPolicy{
		Scope: bucket,
		FsizeLimit:1024,//int64
		MimeLimit:"PNG",
		DeleteAfterDays:7,
	}

	macObj := secure.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(macObj)
	fmt.Println(putPolicy)
	c.JSON(http.StatusOK, gin.H{"user_id": 10821038, "token": upToken})
}

func GetAkWithToken(c *gin.Context) {	
	putPolicy := storage.PutPolicy{}		
	upToken :="QDM_ACCESS_KEY:xOtx4yibGoENAY3Wvgj0I7h9gOU=:eyJzY29wZSI6IlFETV9URVNUX0JVQ0tFVCIsImRlYWRsaW5lIjoxNTM4MDI2OTUwLCJmc2l6ZUxpbWl0IjoxMDI0LCJtaW1lTGltaXQiOiJQTkciLCJkZWxldGVBZnRlckRheXMiOjd9"
	//upToken := c.Param("token")
	//upToken :=c.PostForm("token")
	ak, _, gErr := putPolicy.DecodeWithToken(upToken)
	if gErr != nil {
		fmt.Println(gErr)
		//return
	}
	fmt.Println(putPolicy)
	c.JSON(http.StatusOK, gin.H{"ak": ak, "token": upToken})
}

func GetUpHostByAkBucket(c *gin.Context) {
	//var putRet storage.PutRet
	
	putPolicy := storage.PutPolicy{
		Scope:           testBucket,
		DeleteAfterDays: 7,
	}
	macInstance := secure.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(macInstance)
	testKey := fmt.Sprintf("testPutFileKey_%d", rand.Int())
	

    ak:=fmt.Sprintf("testAk_%d", rand.Int())
	bucket:=fmt.Sprintf("testCDN_%d", rand.Int())
	akcfg:=storage.Config{
		Zone:&storage.ZoneHuadong,
		UseHTTPS:true,
		UseCdnDomains:true,
		CentralRsHost:"",
	}

	/**/
	formDmer:=storage.FormDmer{
		FormId:"D10888",
		Cfg:&akcfg,
	}		
	//formDmer:=storage.NewFormDmer(&akcfg,"D10888")

	// 可选配置
	/*
	putExtra:=storage.PutExtra{}
	putExtra.MimeType="png";
	putExtra.Params=map[string]string{
		"user_type":"2",
		"sign_flag":"extra_auth",
	}
	*/	
	putExtra := &storage.PutExtra{
		MimeType:"png",
		Params: map[string]string{
			"user_type":"2",
		    "sign_flag":"extra_auth",
		},
	}
	
	putRes,putErr:=formDmer.PutFetch(upToken,testKey,188,putExtra)
	if putErr != nil {		
		fmt.Println(putErr)
		return
	}
	
	upHost,err := formDmer.UpHost(ak,bucket)
	if err != nil {		
		fmt.Println(err)
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"ak": ak, "bucket": bucket,"host":upHost})
}
