package push

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

type PushStat struct {
	MessageId              string `json:"MessageId" xml:"MessageId"`
	AcceptCount            int    `json:"AcceptCount" xml:"AcceptCount"`
	SentCount              int    `json:"SentCount" xml:"SentCount"`
	ReceivedCount          int    `json:"ReceivedCount" xml:"ReceivedCount"`
	OpenedCount            int    `json:"OpenedCount" xml:"OpenedCount"`
	DeletedCount           int    `json:"DeletedCount" xml:"DeletedCount"`
	SmsSentCount           int    `json:"SmsSentCount" xml:"SmsSentCount"`
	SmsSkipCount           int    `json:"SmsSkipCount" xml:"SmsSkipCount"`
	SmsFailedCount         int    `json:"SmsFailedCount" xml:"SmsFailedCount"`
	SmsReceiveSuccessCount int    `json:"SmsReceiveSuccessCount" xml:"SmsReceiveSuccessCount"`
	SmsReceiveFailedCount  int    `json:"SmsReceiveFailedCount" xml:"SmsReceiveFailedCount"`
}