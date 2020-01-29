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

package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"zblog2md/pkg/config"
	"zblog2md/pkg/model"
	"zblog2md/pkg/utils"
)

type options struct {
	DBname string
	OutPutDir string
	PageSize uint
}

// var posts []Post // := make([]Post, 0)

func zblog2mdCmd() *cobra.Command {
	var (
		optionitem options
	)
	z2md := &cobra.Command{
		Use:   "2md",
		Short: "read posts from mysql and translate to markdown file ",
		PreRun: func(cmd *cobra.Command, args []string) {
			var cf config.Config
			config.Initialize(cf)
			utils.InitFilterMap()
			fmt.Printf("Inside rootCmd PreRun with args: %v\n", args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Fprintf(cmd.OutOrStdout(), "OK")
			fmt.Println("Echo: " + strings.Join(args, " "))
			fmt.Println("DBname: " + optionitem.DBname + " OutPutDir: " + optionitem.OutPutDir + " PageSize: " + string(optionitem.PageSize))

			total, _, err:= model.ZbpPostPagedQuery("log_ID > 0", 1, uint(1))
			if err != nil {
				fmt.Printf("err: %s", err.Error())
			} else {
				fmt.Printf("total: %d", total)
				page := uint(0)
				for ( (uint(page) * optionitem.PageSize) < (total + 20)) {
					readAndWritePosts(optionitem.PageSize, page, optionitem.OutPutDir)
					page = page + 1
				} 
			}
			return nil
		},
	}
	z2md.PersistentFlags().StringVar(&optionitem.DBname, "DBname", "zblog",
		"the DBname to read posts.")
	z2md.PersistentFlags().StringVar(&optionitem.OutPutDir, "output", "./output/",
		"the dir to write posts to.")
	z2md.PersistentFlags().UintVar(&optionitem.PageSize, "pagesize", 20, "pagesize read from db one time.")

	return z2md
}

func readAndWritePosts(pagesize, page uint, outputdir string) error {
	_, rows, err:= model.ZbpPostPagedQuery("log_ID > 0", pagesize, uint(page))
	if err != nil {
		fmt.Printf("err: %s", err.Error())
		return err
	} 
	// default output dir, default to ./output/
	dealPosts(rows, outputdir)

	return nil
}

func dealPosts(posts []*model.ZbpPost, outputdir string)  {
	i := len(posts)
	for i > 0 {
		i = i - 1
		fmt.Printf("LogCateID: %d ", posts[i].LogCateID)
		category := "life"
		cate, _ := model.GetZbpCategory("cate_ID = ?", posts[i].LogCateID)
		if (cate != nil) {
			category = cate.CateName
		}
		var tags []string
		// {46}{72}{91}{108}{110}
		retags, _ := utils.Tags2ID(posts[i].LogTag)
		fmt.Printf("logtags1: %v ", retags)
		v := len(retags)
		for v > 0 {
			v = v - 1
			tag, _ := model.GetZbpTag("tag_ID = ?", retags[v])
			if (tag != nil) {
				fmt.Printf("tag.TagName: %s ", tag.TagName)
				tags = append(tags, tag.TagName)
			}
		}
		fmt.Printf("logtags2: %v ", tags)
		utils.Write2hugomd(posts[i], tags, category, outputdir)
	}
}