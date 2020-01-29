// Copyright 2020 helight Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package utils

import (
	"strings"
)

var filterMap map[string]string

// InitFilterMap init
func InitFilterMap() {
	filterMap = make(map[string]string)
	filterMap["<p>"] = ""
	filterMap["</p>"] = ""
	filterMap["<coolcode lang=\"java\">"] = "```c"
	filterMap["</coolcode>"] = "```"
}

// FilterHTML FilterHtml
func FilterHTML(content string) string {
	for k,v := range filterMap {
		content = strings.Replace(content, k, v, -1 )
	}
	return content
}

// CheckUpload check if there is upload in post
func CheckUpload(content string) bool {
	return false
}