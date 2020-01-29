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
	"fmt"
	"time"
	"strings"

	"zblog2md/pkg/model"
)

/*
+++
title = "我的二十年感言"
date = "2009-01-15T13:47:08+02:00"
tags = ["life"]
categories = ["life"]
banner = "img/banners/banner-2.jpg"
draft = false
author = "helight"
authorlink = "https://helight.cn"
summary = ""
keywords = ["life"]
+++
*/
// Write2md post to file
func Write2md(post *model.ZbpPost, tags []string, category string)  {
	posttime := time.Unix(int64(post.LogPostTime), 0)
	fmt.Printf("LogPostTime: %d \r\n", posttime.Year())
}

// Tags2ID {46}{72}{91}{108}{110}
func Tags2ID(tags string) ([]string, error) {
	results := make([]string, 0)

	tags = strings.TrimSpace(strings.Replace(tags, "{", "", -1))
	fmt.Printf("tags: %s \r\n", tags)
	t := strings.Split(tags, "}")
	v := len(t)
	fmt.Printf("len tags: %d \r\n", v)
	for v > 0 {
		v = v - 1
		
		if(len(t[v]) > 0) {
			fmt.Printf("len tags xxx: %s \r\n", t[v])
			results = append(results, t[v])
		}
	}
	return results, nil
}