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
	"sync"
)

/*
 * round robin pool.
 */
type DefaultPool struct {
	lock        sync.Mutex
	cfg         *Config
	connections []Connection
	use         []bool
	product     Product
}

func NewDefaultPool(config *Config, product Product) (Pool, error) {
	var cfg *Config
	if config == nil {
		cfg = DefaultConfig()
	} else {
		cfg = config
	}
	if cfg.maxConnections <= 0 || cfg.minConnections < 0 {
		return nil, ErrConfig
	}

	if product == nil {
		return nil, ErrProduct
	}

	dp := new(DefaultPool)
	dp.product = product
	dp.cfg = cfg

	connections := make([]Connection, cfg.maxConnections)
	use := make([]bool, cfg.maxConnections)
	for i := 0; i < cfg.minConnections; i++ {
		c, err := dp.product()
		if err != nil {
			return nil, ErrProduct
		}
		connections[i] = c
		use[i] = false
	}
	dp.connections = connections
	dp.use = use

	return dp, nil
}

func (dp *DefaultPool) Acquire() (Connection, error) {
	dp.lock.Lock()
	defer dp.lock.Unlock()
	idx := -1
	for i, c := range dp.connections {
		if c != nil && dp.use[i] == false {
			if nil == c.Ping() {
				dp.use[i] = true
				return c, nil
			}
		}
		if c == nil {
			idx = i
		}
	}
	if idx != -1 {
		c, err := dp.product()
		if err == nil {
			dp.connections[idx] = c
			dp.use[idx] = true
			return c, nil
		} else {
			return nil, err
		}
	}

	return nil, ErrConnection
}

func (dp *DefaultPool) Release(conn Connection) {
	dp.lock.Lock()
	defer dp.lock.Unlock()

	for idx, c := range dp.connections {
		if c == conn {
			dp.use[idx] = false
			break
		}
	}
}

func (dp *DefaultPool) Shutdown() error {
	for _, c := range dp.connections {
		if c != nil {
			c.Close()
		}
	}
	dp.connections = nil
	dp.use = nil
	dp = nil
	return nil
}
