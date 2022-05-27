package aliyun

import (
	"fmt"

	"github.com/xiaosong372089396/learning-cloud-station-pro/store"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"

	"github.com/go-playground/validator/v10"
)

var (
	// 校验器的实例对象
	validate = validator.New()
)

func NewAliYunOssUploader(endpoint, ak, sk string) (store.OSSUploader, error) {
	uploader := &impl{
		Endpoint: endpoint,
		AK:       ak,
		SK:       sk,
	}

	if err := uploader.Validate(); err != nil {
		return nil, fmt.Errorf("validate params error, %s", err)
	}

	uploader.listener = NewOssProgressListener()
	return uploader, nil
}

// 这个对象, 实现我们定义的接口
type impl struct {
	// oss 服务
	Endpoint string `validate:"required"`
	AK       string `validate:"required"`
	SK       string `validate:"required"`

	listener oss.ProgressListener // global one
}

func (i *impl) Validate() error {
	return validate.Struct(i)
}

func (i *impl) Upload(bucketName, objectKey, fileName string) (downloadUrl string, err error) {
	client, err := oss.New(i.Endpoint, i.AK, i.SK)
	if err != nil {
		err = fmt.Errorf("new client error, %s", err)
		return
	}

	bucket, err := client.Bucket(bucketName)
	if err != nil {
		err = fmt.Errorf("get bucket %s error, %s", bucketName, err)
		return
	}

	listener := oss.Progress(i.listener)
	err = bucket.PutObjectFromFile(objectKey, fileName, listener)
	if err != nil {
		err = fmt.Errorf("upload file %s error, %s", fileName, err)
		return
	}

	// 生成下载链接
	return bucket.SignURL(objectKey, oss.HTTPGet, 60*60*24*3)
}
