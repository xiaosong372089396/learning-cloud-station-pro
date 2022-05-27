package aliyun_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/xiaosong372089396/learning-cloud-station-pro/store/aliyun"
	"github.com/stretchr/testify/assert"
)

var (
	ep, ak, sk, bn string
)

// TDD: 测试驱动开发
func TestUpload(t *testing.T) {
	//fmt.Println("xxx", ep, ak, sk)
	// 断言对象
	should := assert.New(t)

	uploader, err := aliyun.NewAliYunOssUploader(ep, ak, sk)

	if should.NoError(err) {
		downloadURL, err := uploader.Upload(bn, "impl.go", "impl.go")

		wd, _ := os.Getwd()
		fmt.Println("work dir: ", wd)
		// 简化
		// if err != nil {
		// 	t.Fatal(err)
		// }

		// if downloadURL == "" {
		// 	t.Fatal("no down")
		// }

		// if err == nil

		if should.NoError(err) {
			should.NotEmpty(downloadURL)
		}
	}

}

// 测试时候 通过环境变量加载参数
func init() {
	ep = os.Getenv("ALI_OSS_ENDPOINT")
	ak = os.Getenv("ALI_AK")
	sk = os.Getenv("ALI_SK")
	bn = os.Getenv("ALI_BUCKET_NAME")
}
