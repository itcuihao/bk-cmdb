/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.
 * Copyright (C) 2017-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package client

import (
	"configcenter/src/source_controller/client/v3"
	"net/http"
)

type Interface interface {
	V3() v3.Interface
}

var _ Interface = &Client{}

var baseclient = &Client{}

// SetBaseClient set the based client
func SetBaseClient(cli *Client) {
	baseclient = cli
}

// New returns a new clien
func New() *Client {
	return baseclient.clone()
}

// NewWithFoward returns a client that will tramsmit ForwardParam when calling request
func NewWithFoward(forward *ForwardParam) *Client {
	clone := baseclient.clone()
	clone.forward = forward
	return clone
}

// NewWithReq returns a client that will tramsmit ForwardParam when calling request
// @deprecated please use NewWithFoward instead
func NewWithReq(forward *http.Request) *Client {
	panic("please use NewWithFoward instead")
}

func (cli *Client) clone() *Client {
	clone := *cli
	return &clone
}

// V3 returns a v3 client
func (cli *Client) V3() v3.Interface {
	return cli.ccv3
}

// NewTransaction set this client in transaction
func (cli *Client) NewTransaction() *Client {
	clone := baseclient.clone()
	clone.isTransaction = true
	return clone
}
