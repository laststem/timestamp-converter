/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
package cmd

import (
	"time"

	"github.com/spf13/cobra"
)

var d = time.Second

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:     "gen",
	Short:   "generate current unix timestamp value",
	Example: "tc gen",
	Run: func(cmd *cobra.Command, args []string) {
		flags := cmd.Flags()
		if flags.NFlag() > 1 {
			_ = cmd.Usage()
			return
		}

		if cmd.Flag("ms").Value.String() == "true" {
			d = time.Millisecond
		}
		cmd.Println(time.Now().UnixNano() / int64(d))
	},
}

func init() {
	genCmd.Flags().BoolP("ms", "m", false, "Print Millisecond")
	rootCmd.AddCommand(genCmd)
}
