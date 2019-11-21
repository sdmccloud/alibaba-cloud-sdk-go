package cr_20181201

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

// UpdateInstanceEndpointStatus invokes the cr.UpdateInstanceEndpointStatus API synchronously
// api document: https://help.aliyun.com/api/cr/updateinstanceendpointstatus.html
func (client *Client) UpdateInstanceEndpointStatus(request *UpdateInstanceEndpointStatusRequest) (response *UpdateInstanceEndpointStatusResponse, err error) {
	response = CreateUpdateInstanceEndpointStatusResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateInstanceEndpointStatusWithChan invokes the cr.UpdateInstanceEndpointStatus API asynchronously
// api document: https://help.aliyun.com/api/cr/updateinstanceendpointstatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateInstanceEndpointStatusWithChan(request *UpdateInstanceEndpointStatusRequest) (<-chan *UpdateInstanceEndpointStatusResponse, <-chan error) {
	responseChan := make(chan *UpdateInstanceEndpointStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateInstanceEndpointStatus(request)
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

// UpdateInstanceEndpointStatusWithCallback invokes the cr.UpdateInstanceEndpointStatus API asynchronously
// api document: https://help.aliyun.com/api/cr/updateinstanceendpointstatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateInstanceEndpointStatusWithCallback(request *UpdateInstanceEndpointStatusRequest, callback func(response *UpdateInstanceEndpointStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateInstanceEndpointStatusResponse
		var err error
		defer close(result)
		response, err = client.UpdateInstanceEndpointStatus(request)
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

// UpdateInstanceEndpointStatusRequest is the request struct for api UpdateInstanceEndpointStatus
type UpdateInstanceEndpointStatusRequest struct {
	*requests.RpcRequest
	InstanceId   string           `position:"Query" name:"InstanceId"`
	EndpointType string           `position:"Query" name:"EndpointType"`
	Enable       requests.Boolean `position:"Query" name:"Enable"`
	ModuleName   string           `position:"Query" name:"ModuleName"`
}

// UpdateInstanceEndpointStatusResponse is the response struct for api UpdateInstanceEndpointStatus
type UpdateInstanceEndpointStatusResponse struct {
	*responses.BaseResponse
	UpdateInstanceEndpointStatusIsSuccess bool   `json:"IsSuccess" xml:"IsSuccess"`
	Code                                  string `json:"Code" xml:"Code"`
	RequestId                             string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateInstanceEndpointStatusRequest creates a request to invoke UpdateInstanceEndpointStatus API
func CreateUpdateInstanceEndpointStatusRequest() (request *UpdateInstanceEndpointStatusRequest) {
	request = &UpdateInstanceEndpointStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cr", "2018-12-01", "UpdateInstanceEndpointStatus", "cr", "openAPI")
	return
}

// CreateUpdateInstanceEndpointStatusResponse creates a response to parse from UpdateInstanceEndpointStatus response
func CreateUpdateInstanceEndpointStatusResponse() (response *UpdateInstanceEndpointStatusResponse) {
	response = &UpdateInstanceEndpointStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
