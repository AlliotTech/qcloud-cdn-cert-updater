package ssl

import (
	"fmt"
	"os"
	"time"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	ssl "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ssl/v20191205"
)

type Client struct {
	client *ssl.Client
}

func NewClient(secretId, secretKey string) *Client {
	credential := common.NewCredential(secretId, secretKey)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "ssl.tencentcloudapi.com"
	client, _ := ssl.NewClient(credential, "", cpf)
	return &Client{client: client}
}

func (c *Client) UploadCertificate(certPath, keyPath, alias string) (string, error) {
	certData, err := os.ReadFile(certPath)
	if err != nil {
		return "", fmt.Errorf("failed to read certificate file: %v", err)
	}

	keyData, err := os.ReadFile(keyPath)
	if err != nil {
		return "", fmt.Errorf("failed to read key file: %v", err)
	}

	request := ssl.NewUploadCertificateRequest()
	request.CertificatePublicKey = common.StringPtr(string(certData))
	request.CertificatePrivateKey = common.StringPtr(string(keyData))
	request.CertificateType = common.StringPtr("SVR")
	request.Alias = common.StringPtr(alias + time.Now().Format("20060102"))
	request.CertificateUse = common.StringPtr("CDN")
	//request.Tags = []*ssl.Tags{
	//
	//	&ssl.Tags{
	//		TagKey:   common.StringPtr("managed"),
	//		TagValue: common.StringPtr("qcloud-cdn-cert-updater"),
	//	},
	//	&ssl.Tags{
	//		TagKey:   common.StringPtr("update"),
	//		TagValue: common.StringPtr(time.Now().Format("20060102")),
	//	},
	//}

	response, err := c.client.UploadCertificate(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		return "", fmt.Errorf("an API error has returned: %s", err)
	}
	if err != nil {
		return "", err
	}

	return *response.Response.CertificateId, nil
}
