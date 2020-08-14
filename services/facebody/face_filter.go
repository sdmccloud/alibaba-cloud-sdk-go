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

// FaceFilter invokes the facebody.FaceFilter API synchronously
// api document: https://help.aliyun.com/api/facebody/facefilter.html
func (client *Client) FaceFilter(request *FaceFilterRequest) (response *FaceFilterResponse, err error) {
	response = CreateFaceFilterResponse()
	err = client.DoAction(request, response)
	return
}

// FaceFilterWithChan invokes the facebody.FaceFilter API asynchronously
// api document: https://help.aliyun.com/api/facebody/facefilter.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) FaceFilterWithChan(request *FaceFilterRequest) (<-chan *FaceFilterResponse, <-chan error) {
	responseChan := make(chan *FaceFilterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.FaceFilter(request)
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

// FaceFilterWithCallback invokes the facebody.FaceFilter API asynchronously
// api document: https://help.aliyun.com/api/facebody/facefilter.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) FaceFilterWithCallback(request *FaceFilterRequest, callback func(response *FaceFilterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *FaceFilterResponse
		var err error
		defer close(result)
		response, err = client.FaceFilter(request)
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

// FaceFilterRequest is the request struct for api FaceFilter
type FaceFilterRequest struct {
	*requests.RpcRequest
	Strength     requests.Float `position:"Body" name:"Strength"`
	ResourceType string         `position:"Body" name:"ResourceType"`
	ImageURL     string         `position:"Body" name:"ImageURL"`
}

// FaceFilterResponse is the response struct for api FaceFilter
type FaceFilterResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateFaceFilterRequest creates a request to invoke FaceFilter API
func CreateFaceFilterRequest() (request *FaceFilterRequest) {
	request = &FaceFilterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("facebody", "2019-12-30", "FaceFilter", "", "")
	request.Method = requests.POST
	return
}

// CreateFaceFilterResponse creates a response to parse from FaceFilter response
func CreateFaceFilterResponse() (response *FaceFilterResponse) {
	response = &FaceFilterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
