// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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
	"github.com/xwy27/Go-Agenda/global"
	"github.com/xwy27/Go-Agenda/operation"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// queryMeetingsCmd represents the queryMeetings command
var queryMeetingsCmd = &cobra.Command{
	Use:   "queryMeetings",
	Short: "Query attended meetings",
	Long: `Query user attended meetings by a given time interval which returns all the matched meetings
with their title, sponsor, participator, startTime and endTime`,

	Run: func(cmd *cobra.Command, args []string) {
		startTime, err := cmd.Flags().GetString("startTime")
		global.PrintError(err, "")
		endTime, err := cmd.Flags().GetString("endTime")
		global.PrintError(err, "")
		// Error handle
		fmt.Println("queryMeetings called by " + startTime)
		fmt.Println("queryMeetings called by " + endTime)
		// Query
		meetings, err := operation.QueryMeetings(startTime+":00", endTime+":00")
		global.PrintError(err, "Query results:")
		if len(meetings) != 0 {
			for index, m := range meetings {
				fmt.Println("==========================================")
				fmt.Printf("Meeting %d:\n", index+1)
				fmt.Println("-Title: " + m.Title)
				fmt.Println("-Sponsor: " + m.Sponsor)
				start := time.Unix(m.StartTime, 0)
				fmt.Println("-StartTime: " + start.Format(time.RFC1123))
				end := time.Unix(m.EndTime, 0)
				fmt.Println("-EndTime: " + end.Format(time.RFC1123))
				fmt.Println("-Participators: ")
				for _, p := range m.Participators {
					fmt.Println("  -" + p.Username)
				}
			}
		} else {
			fmt.Println("Empty.")
		}
		fmt.Println("===============End of List================")
	},
}

func init() {
	rootCmd.AddCommand(queryMeetingsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// queryMeetingsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	queryMeetingsCmd.Flags().StringP("startTime", "s", "", "start time for query")
	queryMeetingsCmd.Flags().StringP("endTime", "e", "", "end time for query")
}
