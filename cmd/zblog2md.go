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
//	"log"
	"strings"
// 	"time"

	"github.com/spf13/cobra"

	"zblog2md/pkg/config"
	"zblog2md/pkg/model"
)

type options struct {
	DBname string
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
			fmt.Printf("Inside rootCmd PreRun with args: %v\n", args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			total, rows, err:= model.ZbpPostPagedQuery("log_ID > 0", 20, uint(1))
			if err != nil {
				fmt.Printf("err: %s", err.Error())
			} else {
				fmt.Printf("total: %d", total)
				fmt.Printf("data: %v", rows)
	
				fmt.Fprintf(cmd.OutOrStdout(), "OK")
				fmt.Println("Echo: " + strings.Join(args, " "))
				fmt.Println("Echo: " + optionitem.DBname)
			}
			
			return nil
		},
	}
	z2md.PersistentFlags().StringVar(&optionitem.DBname, "DBname", "",
		"the DBname to read posts.")

	return z2md
}