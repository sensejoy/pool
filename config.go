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

type Config struct {
	maxConnections int //max connections in pool.
	minConnections int
	waitTimeout    int //max wait time in seconds for acquiring a connection
	product        func() (Connection, error)
}

func (cfg *Config) SetMaxConnections(max int) {
	cfg.maxConnections = max
}

func (cfg *Config) SetMinConnections(min int) {
	cfg.minConnections = min
}

func (cfg *Config) SetWaitTimeout(wait int) {
	cfg.waitTimeout = wait
}

func DefaultConfig() *Config {
	cfg := new(Config)
	cfg.maxConnections = defaultMaxConnections
	cfg.minConnections = defaultMinConnections
	cfg.waitTimeout = defaultWaitTimeout
	cfg.product = nil
	return cfg
}
