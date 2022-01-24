/*
Copyright 2019 The Knative Authors
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

// This Adapter generates events at a regular interval.
package main

import (
	myadapter "github.com/baum/noobaa-source/pkg/adapter"
	"knative.dev/eventing/pkg/adapter/v2"
)

func main() {
	adapter.Main("noobaa-source", myadapter.NewEnv, myadapter.NewAdapter)
}
