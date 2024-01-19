package upload

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"go.uber.org/zap"
)

type Qiniu struct{}

//@author: [piexlmax](https://github.com/piexlmax)
//@author: [ccfish86](https://github.com/ccfish86)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@object: *Qiniu
//@function: UploadFile
//@description: Upload file
//@param: file *multipart.FileHeader
//@return: string, string, error

func (*Qiniu) UploadFile(file *multipart.FileHeader) (string, string, error) {
	putPolicy := storage.PutPolicy{Scope: global.GVA_CONFIG.Qiniu.Bucket}
	mac := qbox.NewMac(global.GVA_CONFIG.Qiniu.AccessKey, global.GVA_CONFIG.Qiniu.SecretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := qiniuConfig()
	formUploader := storage.NewFormUploader(cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{Params: map[string]string{"x:name": "github logo"}}

	f, openError := file.Open()
	if openError != nil {
		global.GVA_LOG.Error("function file.Open() failed", zap.Any("err", openError.Error()))

		return "", "", errors.New("function file.Open() failed, err:" + openError.Error())
	}
	defer f.Close()                                                  // Create file defer close
	fileKey := fmt.Sprintf("%d%s", time.Now().Unix(), file.Filename) // The file name format can be changed by yourself. It is recommended to ensure uniqueness.
	putErr := formUploader.Put(context.Background(), &ret, upToken, fileKey, f, file.Size, &putExtra)
	if putErr != nil {
		global.GVA_LOG.Error("function formUploader.Put() failed", zap.Any("err", putErr.Error()))
		return "", "", errors.New("function formUploader.Put() failed, err:" + putErr.Error())
	}
	return global.GVA_CONFIG.Qiniu.ImgPath + "/" + ret.Key, ret.Key, nil
}

//@author: [piexlmax](https://github.com/piexlmax)
//@author: [ccfish86](https://github.com/ccfish86)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@object: *Qiniu
//@function: DeleteFile
//@description: delete file
//@param: key string
//@return: error

func (*Qiniu) DeleteFile(key string) error {
	mac := qbox.NewMac(global.GVA_CONFIG.Qiniu.AccessKey, global.GVA_CONFIG.Qiniu.SecretKey)
	cfg := qiniuConfig()
	bucketManager := storage.NewBucketManager(mac, cfg)
	if err := bucketManager.Delete(global.GVA_CONFIG.Qiniu.Bucket, key); err != nil {
		global.GVA_LOG.Error("function bucketManager.Delete() failed", zap.Any("err", err.Error()))
		return errors.New("function bucketManager.Delete() failed, err:" + err.Error())
	}
	return nil
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@object: *Qiniu
//@function: qiniuConfig
//@description: Return the configuration of Qiniu Cloud according to the configuration file
//@return: *storage.Config

func qiniuConfig() *storage.Config {
	cfg := storage.Config{
		UseHTTPS:      global.GVA_CONFIG.Qiniu.UseHTTPS,
		UseCdnDomains: global.GVA_CONFIG.Qiniu.UseCdnDomains,
	}
	switch global.GVA_CONFIG.Qiniu.Zone { // Initialize the computer room corresponding to the space according to the configuration file
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
