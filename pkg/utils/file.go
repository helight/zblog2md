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
	"text/template"
	"os"
	"unicode/utf8"

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

type hugoPostTpl struct {
	LogTitle 		string
	LogDate   		string
	LogTags    		string
	LogCategories   string
	LogSummary    	string
	LogKeywords    	string
	LogContent    	string
 }

// CheckErr check
 func CheckErr(err error) {
	if err != nil {
	   panic(err)
	}
 }

// Write2md post to file
func Write2md(post *model.ZbpPost, tags []string, category string)  {
	posttime := time.Unix(int64(post.LogPostTime), 0)
	fmt.Printf("LogPostTime: %d \r\n", posttime.Year())
}

// Write2hugomd post to hugo markdown file
func Write2hugomd(post *model.ZbpPost, tags []string, category string, outputdir string)  { 
	// file data
	postData := hugoPostTpl{"wool", "wool", "wool", "wool", "wool", "wool", "wool"}

	posttime := time.Unix(int64(post.LogPostTime), 0)
	fmt.Printf("LogPostTime: %d \r\n", posttime.Year())

	postData.LogTitle = fmt.Sprintf("\"%s\"", post.LogTitle)
	postData.LogDate = fmt.Sprintf("%s", posttime.Format(time.RFC3339))
	postData.LogCategories = fmt.Sprintf("[\"%s\"]", category)
	postData.LogSummary = ""
	postData.LogContent = post.LogContent
	logtags := "["
	i := len(tags)
	for i > 0 {
		i = i - 1
		logtags = logtags + "\"" + tags[i] + "\""
		if (i > 0) {
			logtags = logtags + ","
		}
	}
	logtags = logtags + "]"
	if (utf8.RuneCountInString(logtags) < 3) {
		logtags = postData.LogCategories 
	}	
	postData.LogTags = logtags
	postData.LogKeywords = logtags

	// create output dir ./output/2009/
	filepath := fmt.Sprintf("%s/%d/", outputdir, posttime.Year())
	CreateDir(filepath)
	filename := fmt.Sprintf("%s/%d.md", filepath, post.LogID)

	// fill the template and write to file
	tmpl, err := template.ParseFiles("data/hugo-md.tpl")
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0755)
	CheckErr(err)
	err = tmpl.Execute(file, postData)
	CheckErr(err)
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

// CreateDir CreateDir
func CreateDir(dirpath string) {
	if _, err := os.Stat(dirpath); os.IsNotExist(err) {
		os.MkdirAll(dirpath, os.ModePerm) //os.ModePerm
		os.Chmod(dirpath, 0755)
	}
}

// PathExists check
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}