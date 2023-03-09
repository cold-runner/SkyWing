package oss

import (
	"context"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"go.uber.org/zap"
	"testing"
)

var (
	ak   = "pEBtMkmdrODi59i306RyfjACD1YhdzoybxYnTar0"
	sk   = "Kk6-bK3eacxUO9IR89KDE0T4Na_KgZjmHFCe0xnL"
	iUrl = "rjl8hsvex.hb-bkt.clouddn.com"
	buck = "skylab-org-cn"
	m    = qbox.NewMac(ak, sk)
)

func TestUploadBs64File(t *testing.T) {
	putPolicy := storage.PutPolicy{
		Scope: buck,
	}
	upToken := putPolicy.UploadToken(m)
	cfg := storage.Config{
		UseHTTPS:      false,
		UseCdnDomains: false,
		Zone:          &storage.ZoneHuabei,
	}
	putExtra := storage.PutExtra{}
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)
	if err != nil {
		zap.L().Error("上传文件失败", zap.Error(err))
	}
	url := imgUrl + ret.Key
}
