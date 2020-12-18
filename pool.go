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

type Connection interface {
	Ping() error  //test alive
	Close() error //in case resource lease
}

type Pool interface {
	Acquire() (Connection, error) //get an exist connection
	Release(Connection)           //release the connection to pool
	Shutdown() error              //close the pool
}

type Product func(args ...interface{}) (Connection, error)
