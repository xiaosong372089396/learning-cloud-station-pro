package store

// oss 存储适配器, 通用解
// 阿里云OSS/Tencent OSS/IDC Minio 开源OSS/AWS S3
type OSSUploader interface {
	Upload(bucketName, objectKey, fileName string) (downloadUrl string, err error)
}
