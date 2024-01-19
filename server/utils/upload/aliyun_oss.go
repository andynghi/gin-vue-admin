package upload

import (
	"errors"
	"mime/multipart"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"go.uber.org/zap"
)

type AliyunOSS struct{}

func (*AliyunOSS) UploadFile(file *multipart.FileHeader) (string, string, error) {
	bucket, err := NewBucket()
	if err != nil {
		global.GVA_LOG.Error("function AliyunOSS.NewBucket() Failed", zap.Any("err", err.Error()))
		return "", "", errors.New("function AliyunOSS.NewBucket() Failed, err:" + err.Error())
	}

	//Read local file.
	f, openError := file.Open()
	if openError != nil {
		global.GVA_LOG.Error("function file.Open() Failed", zap.Any("err", openError.Error()))
		return "", "", errors.New("function file.Open() Failed, err:" + openError.Error())
	}
	defer f.Close() // Create file defer close
	// Upload Alibaba Cloud path file name format. You can change it yourself. It is recommended to ensure uniqueness.
	// yunFileTmpPath := filepath.Join("uploads", time.Now().Format("2006-01-02")) + "/" + file.Filename
	yunFileTmpPath := global.GVA_CONFIG.AliyunOSS.BasePath + "/" + "uploads" + "/" + time.Now().Format("2006-01-02") + "/" + file.Filename

	//Upload file stream.
	err = bucket.PutObject(yunFileTmpPath, f)
	if err != nil {
		global.GVA_LOG.Error("function formUploader.Put() Failed", zap.Any("err", err.Error()))
		return "", "", errors.New("function formUploader.Put() Failed, err:" + err.Error())
	}

	return global.GVA_CONFIG.AliyunOSS.BucketUrl + "/" + yunFileTmpPath, yunFileTmpPath, nil
}

func (*AliyunOSS) DeleteFile(key string) error {
	bucket, err := NewBucket()
	if err != nil {
		global.GVA_LOG.Error("function AliyunOSS.NewBucket() Failed", zap.Any("err", err.Error()))
		return errors.New("function AliyunOSS.NewBucket() Failed, err:" + err.Error())
	}

	//Delete a single file. objectName indicates that when deleting an OSS file, you need to specify the full path including the file suffix, such as abc/efg/123.jpg.
	// If you want to delete a folder, set objectName to the corresponding folder name. If the folder is not empty, you need to delete all objects under the folder before you can delete the folder.
	err = bucket.DeleteObject(key)
	if err != nil {
		global.GVA_LOG.Error("function bucketManager.Delete() failed", zap.Any("err", err.Error()))
		return errors.New("function bucketManager.Delete() failed, err:" + err.Error())
	}

	return nil
}

func NewBucket() (*oss.Bucket, error) {
	//Create an OSSClient instance.
	client, err := oss.New(global.GVA_CONFIG.AliyunOSS.Endpoint, global.GVA_CONFIG.AliyunOSS.AccessKeyId, global.GVA_CONFIG.AliyunOSS.AccessKeySecret)
	if err != nil {
		return nil, err
	}

	// Get storage space.
	bucket, err := client.Bucket(global.GVA_CONFIG.AliyunOSS.BucketName)
	if err != nil {
		return nil, err
	}

	return bucket, nil
}
