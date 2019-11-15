package cdn

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ModifyCdnDomain invokes the cdn.ModifyCdnDomain API synchronously
// api document: https://help.aliyun.com/api/cdn/modifycdndomain.html
func (client *Client) ModifyCdnDomain(request *ModifyCdnDomainRequest) (response *ModifyCdnDomainResponse, err error) {
	response = CreateModifyCdnDomainResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyCdnDomainWithChan invokes the cdn.ModifyCdnDomain API asynchronously
// api document: https://help.aliyun.com/api/cdn/modifycdndomain.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyCdnDomainWithChan(request *ModifyCdnDomainRequest) (<-chan *ModifyCdnDomainResponse, <-chan error) {
	responseChan := make(chan *ModifyCdnDomainResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyCdnDomain(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// ModifyCdnDomainWithCallback invokes the cdn.ModifyCdnDomain API asynchronously
// api document: https://help.aliyun.com/api/cdn/modifycdndomain.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyCdnDomainWithCallback(request *ModifyCdnDomainRequest, callback func(response *ModifyCdnDomainResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyCdnDomainResponse
		var err error
		defer close(result)
		response, err = client.ModifyCdnDomain(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// ModifyCdnDomainRequest is the request struct for api ModifyCdnDomain
type ModifyCdnDomainRequest struct {
	*requests.RpcRequest
	TopLevelDomain  string           `position:"Query" name:"TopLevelDomain"`
	Sources         string           `position:"Query" name:"Sources"`
	DomainName      string           `position:"Query" name:"DomainName"`
	OwnerId         requests.Integer `position:"Query" name:"OwnerId"`
	ResourceGroupId string           `position:"Query" name:"ResourceGroupId"`
	SecurityToken   string           `position:"Query" name:"SecurityToken"`
}

// ModifyCdnDomainResponse is the response struct for api ModifyCdnDomain
type ModifyCdnDomainResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyCdnDomainRequest creates a request to invoke ModifyCdnDomain API
func CreateModifyCdnDomainRequest() (request *ModifyCdnDomainRequest) {
	request = &ModifyCdnDomainRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "ModifyCdnDomain", "", "")
	return
}

// CreateModifyCdnDomainResponse creates a response to parse from ModifyCdnDomain response
func CreateModifyCdnDomainResponse() (response *ModifyCdnDomainResponse) {
	response = &ModifyCdnDomainResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
