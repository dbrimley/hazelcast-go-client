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

package compatibility

import . "github.com/hazelcast/hazelcast-go-client/serialization"

type aDataSerializableFactory struct{}

func (*aDataSerializableFactory) Create(classId int32) IdentifiedDataSerializable {
	if classId == DATA_SERIALIZABLE_CLASS_ID {
		return &anIdentifiedDataSerializable{}
	}
	return nil
}
