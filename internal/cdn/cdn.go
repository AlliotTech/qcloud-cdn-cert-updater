package cdn

import (
	"fmt"

	cdn "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdn/v20180606"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

type Client struct {
	client *cdn.Client
}

func NewClient(secretId, secretKey string) *Client {
	credential := common.NewCredential(secretId, secretKey)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "cdn.tencentcloudapi.com"
	client, _ := cdn.NewClient(credential, "", cpf)
	return &Client{client: client}
}

func (c *Client) UpdateDomainConfig(domain, certId string) error {
	request := cdn.NewModifyDomainConfigRequest()
	request.Domain = common.StringPtr(domain)
	request.Route = common.StringPtr("Https.CertInfo.CertId")
	request.Value = common.StringPtr(fmt.Sprintf("{\"update\":\"%s\"}", certId))
	response, err := c.client.ModifyDomainConfig(request)
	//fmt.Println("req: ", request.ToJsonString())
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		return fmt.Errorf("an API error has returned: %s", err)
	}
	if err != nil {
		return err
	}
	response.ToJsonString()
	return nil
}
