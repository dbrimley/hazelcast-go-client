// Copyright (c) 2008-2018, Hazelcast, Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License")
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package protocol

import (
	. "github.com/hazelcast/hazelcast-go-client/internal/common"
	. "github.com/hazelcast/hazelcast-go-client/internal/serialization"
)

type MapExecuteOnKeyResponseParameters struct {
	Response *Data
}

func MapExecuteOnKeyCalculateSize(name *string, entryProcessor *Data, key *Data, threadId int64) int {
	// Calculates the request payload size
	dataSize := 0
	dataSize += StringCalculateSize(name)
	dataSize += DataCalculateSize(entryProcessor)
	dataSize += DataCalculateSize(key)
	dataSize += INT64_SIZE_IN_BYTES
	return dataSize
}

func MapExecuteOnKeyEncodeRequest(name *string, entryProcessor *Data, key *Data, threadId int64) *ClientMessage {
	// Encode request into clientMessage
	clientMessage := NewClientMessage(nil, MapExecuteOnKeyCalculateSize(name, entryProcessor, key, threadId))
	clientMessage.SetMessageType(MAP_EXECUTEONKEY)
	clientMessage.IsRetryable = false
	clientMessage.AppendString(name)
	clientMessage.AppendData(entryProcessor)
	clientMessage.AppendData(key)
	clientMessage.AppendInt64(threadId)
	clientMessage.UpdateFrameLength()
	return clientMessage
}

func MapExecuteOnKeyDecodeResponse(clientMessage *ClientMessage) *MapExecuteOnKeyResponseParameters {
	// Decode response from client message
	parameters := new(MapExecuteOnKeyResponseParameters)

	if !clientMessage.ReadBool() {
		parameters.Response = clientMessage.ReadData()
	}
	return parameters
}
