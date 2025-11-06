// Package prosearch 增强搜索
package prosearch

import (
	"context"
	"log"
	"time"

	"github.com/bytedance/sonic"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"

	"prosearchmcp/constants"
)

type searchImpl struct {
	toolName constants.ToolName
}

func newSearchImpl(toolName constants.ToolName) *searchImpl {
	return &searchImpl{toolName: toolName}
}

// CloudApi 云 api 联网搜索
// wsa.tencentcloudapi.com
func (s *searchImpl) CloudApi(ctx context.Context, req *CloudAPIRequest) (rsp string, e error) {
	start := time.Now()
	credential := common.NewCredential(req.SecretID, req.SecretKey)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = constants.ProSearchProxyURL
	cpf.HttpProfile.ReqTimeout = constants.Timeout
	cpf.HttpProfile.ReqMethod = "POST"
	client := common.NewCommonClient(credential, "", cpf).WithLogger(log.Default())
	request := tchttp.NewCommonRequest("wsa", "2025-05-08", "SearchPro")
	params, _ := sonic.MarshalString(req)
	log.Printf("prosearch query=%s, params=%v", req.Query, params)
	err := request.SetActionParameters(params)
	if err != nil {
		return
	}
	response := tchttp.NewCommonResponse()
	err = client.Send(request, response)
	if err != nil {
		return "", err
	}
	res := string(response.GetBody())
	log.Printf("prosearch query=%s, cost=%v, result=%v", req.Query, time.Since(start).Milliseconds(), res)
	return res, nil
}
