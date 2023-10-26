package main

import (
	"chuanshu/KERNEL"
	"chuanshu/NET"
	"context"
	"fmt"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

type object_storage struct {
	accessKey string
	secretKey string
	//mac *auth.Credentials
	bucket    string
	putPolicy storage.PutPolicy
	upToken   string
}

func (this *object_storage) gettok() string {
	if this.upToken != "" {
		return this.upToken
	}
	mac := qbox.NewMac(this.accessKey, this.secretKey)
	bucket := "visual-field-of-creativity"
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	//ac := qbox.NewMac(accessKey, secretKey)
	this.upToken = putPolicy.UploadToken(mac)

	return this.upToken
}

// 上传文件，根据文件地址上传
func (this *object_storage) Uploadfile(filepath string, filename string) {

	this.gettok()

	localFile := filepath //本地文件位置
	key := filename       //上传成功后存储的对象主键

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Region = &storage.ZoneHuanan
	// 是否使用https域名
	cfg.UseHTTPS = true
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false

	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	// 可选配置
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}
	err := formUploader.PutFile(context.Background(), &ret, this.upToken, key, localFile, &putExtra)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(ret.Key, ret.Hash) //文件主键以及md5值
}

func main() {
	gt := object_storage{
		accessKey: "ZUNUMAntP3XgFkwbp7opEBKWZUiuZHomq2klZxCq",
		secretKey: "pEHQ4FVyabT4Dmipe3fRLkcnuuVQeDQCF4bmUt12",
	}

	gt.Uploadfile("C:/Users/16435/Videos/dabcbb7b6194371d167abe4933dd2750.mp4", "xiaoshi.mp4")

	domain := "http://s31bw3fji.hn-bkt.clouddn.com/"
	key := "xiaoshi.mp4"
	publicAccessURL := storage.MakePublicURL(domain, key)
	fmt.Println(publicAccessURL)

	ke := KERNEL.Kernel{}
	ke.DosomSething("adsafs")
	ne := NET.MyNet{}
	ne.Netinit()

	select {}
}
