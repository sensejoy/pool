/*
Copyright [2020] [sensejoy@github.com]

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package pool

import (
	"errors"
)

const (
	defaultMaxConnections = 10
	defaultMinConnections = 0
	defaultWaitTimeout    = 5
)

var (
	ErrTimeout    = errors.New("wait time out")
	ErrConfig     = errors.New("invalid config")
	ErrConnection = errors.New("no valid connection")
	ErrProduct    = errors.New("invalid product function")
)
