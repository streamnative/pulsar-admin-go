// Copyright 2023 StreamNative, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package utils

type ConsumerConfig struct {
	SchemaType         string            `json:"schemaType,omitempty"`
	SerdeClassName     string            `json:"serdeClassName,omitempty"`
	IsRegexPattern     bool              `json:"isRegexPattern,omitempty"`
	ReceiverQueueSize  int               `json:"receiverQueueSize,omitempty"`
	SchemaProperties   map[string]string `json:"schemaProperties,omitempty"`
	ConsumerProperties map[string]string `json:"consumerProperties,omitempty"`
	CryptoConfig       CryptoConfig      `json:"cryptoConfig,omitempty"`
	PoolMessages       bool              `json:"poolMessages,omitempty"`
}
