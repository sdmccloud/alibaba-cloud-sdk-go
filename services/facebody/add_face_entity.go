package facebody

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

// AddFaceEntity invokes the facebody.AddFaceEntity API synchronously
// api document: https://help.aliyun.com/api/facebody/addfaceentity.html
func (client *Client) AddFaceEntity(request *AddFaceEntityRequest) (response *AddFaceEntityResponse, err error) {
	response = CreateAddFaceEntityResponse()
	err = client.DoAction(request, response)
	return
}

// AddFaceEntityWithChan invokes the facebody.AddFaceEntity API asynchronously
// api document: https://help.aliyun.com/api/facebody/addfaceentity.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddFaceEntityWithChan(request *AddFaceEntityRequest) (<-chan *AddFaceEntityResponse, <-chan error) {
	responseChan := make(chan *AddFaceEntityResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddFaceEntity(request)
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

// AddFaceEntityWithCallback invokes the facebody.AddFaceEntity API asynchronously
// api document: https://help.aliyun.com/api/facebody/addfaceentity.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddFaceEntityWithCallback(request *AddFaceEntityRequest, callback func(response *AddFaceEntityResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddFaceEntityResponse
		var err error
		defer close(result)
		response, err = client.AddFaceEntity(request)
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

// AddFaceEntityRequest is the request struct for api AddFaceEntity
type AddFaceEntityRequest struct {
	*requests.RpcRequest
	EntityId string `position:"Body" name:"EntityId"`
	Labels   string `position:"Body" name:"Labels"`
	DbName   string `position:"Body" name:"DbName"`
}

// AddFaceEntityResponse is the response struct for api AddFaceEntity
type AddFaceEntityResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAddFaceEntityRequest creates a request to invoke AddFaceEntity API
func CreateAddFaceEntityRequest() (request *AddFaceEntityRequest) {
	request = &AddFaceEntityRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("facebody", "2019-12-30", "AddFaceEntity", "", "")
	request.Method = requests.POST
	return
}

// CreateAddFaceEntityResponse creates a response to parse from AddFaceEntity response
func CreateAddFaceEntityResponse() (response *AddFaceEntityResponse) {
	response = &AddFaceEntityResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
