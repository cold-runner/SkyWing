package oss

import (
	"Skywing/settings"
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"go.uber.org/zap"
	"mime/multipart"
)

func UploadToQiNiu(file *multipart.FileHeader, keyWord string) (string, error) {
	// 配置属性
	putPolicy := storage.PutPolicy{Scope: settings.Conf.QiniuConf.Bucket}
	mac := qbox.NewMac(settings.Conf.QiniuConf.Ak, settings.Conf.QiniuConf.Sk)
	upToken := putPolicy.UploadToken(mac)
	cfg := qiniuConfig()
	putExtra := storage.PutExtra{}
	formUploader := storage.NewFormUploader(cfg)
	ret := storage.PutRet{}

	// 打开文件
	f, openError := file.Open()
	if openError != nil {
		zap.L().Error("七牛云上传文件打开失败", zap.Error(openError))
	}
	defer f.Close()
	// 文件路径
	fileKey := fmt.Sprintf("userPhoto/%s.jpg", keyWord)
	// 上传文件
	err := formUploader.Put(context.Background(), &ret, upToken, fileKey, f, file.Size, &putExtra)
	if err != nil {
		zap.L().Error("上传文件失败", zap.Error(err))
		return "", err
	}
	url := settings.Conf.QiniuConf.ImgPath + "/" + ret.Key
	return url, nil
}

func qiniuConfig() *storage.Config {
	cfg := storage.Config{
		UseHTTPS:      settings.Conf.QiniuConf.UseHttps,
		UseCdnDomains: settings.Conf.QiniuConf.UseCdnDomains,
	}
	switch settings.Conf.QiniuConf.Zone { // 根据配置文件进行初始化空间对应的机房
	case "ZoneHuadong":
		cfg.Zone = &storage.ZoneHuadong
	case "ZoneHuabei":
		cfg.Zone = &storage.ZoneHuabei
	case "ZoneHuanan":
		cfg.Zone = &storage.ZoneHuanan
	case "ZoneBeimei":
		cfg.Zone = &storage.ZoneBeimei
	case "ZoneXinjiapo":
		cfg.Zone = &storage.ZoneXinjiapo
	}
	return &cfg
}
func DeleteFileFromQiniu(key string) error {
	bucketManager := storage.NewBucketManager(qbox.NewMac(settings.Conf.QiniuConf.Ak, settings.Conf.QiniuConf.Sk), qiniuConfig())
	err := bucketManager.Delete(settings.Conf.QiniuConf.Bucket, key)
	if err != nil {
		zap.L().Error("删除文件失败", zap.Error(err))
		return err
	}
	return nil
}
