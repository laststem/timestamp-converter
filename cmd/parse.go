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

// parseCmd represents the parse command
var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "parse time to unix timestamp",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		format := cmd.Flag("format").Value.String()
		t, err := time.Parse(format, args[0])
		if err != nil {
			cmd.PrintErrln(err)
			return
		}

		cmd.Println(t.Unix())
	},
}

func init() {
	parseCmd.Flags().StringP("format", "f", time.RFC3339, `Print by format. example) --format="2006-01-02T15:04:05Z07:00", default) YYYY-MM-DDTHH:mm:ssZ`)
	rootCmd.AddCommand(parseCmd)
}
