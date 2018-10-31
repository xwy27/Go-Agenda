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
	"Go-Agenda/global"
	"Go-Agenda/model"
	"Go-Agenda/operation"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// createMeetingCmd represents the createMeeting command
var createMeetingCmd = &cobra.Command{
	Use:   "createMeeting",
	Short: "Create a meeting",
	Long: `With login user, he/she can create a meeting by providing the title, participators, startTime and endTime.

Validation:
Title, participators, startTime and endTime are all required. EndTime must be after startTime.
Input requirement is shown in below example.

Example:
Agenda createMeeting -t=Title -p=user1,user2,user3 -s=2006-01-02 15:04 -e=2006-01-03 15:04`,

	Run: func(cmd *cobra.Command, args []string) {
		title, err := cmd.Flags().GetString("title")
		global.PrintError(err, "")
		participators, err := cmd.Flags().GetString("participators")
		global.PrintError(err, "")
		startTime, err := cmd.Flags().GetString("startTime")
		global.PrintError(err, "")
		endTime, err := cmd.Flags().GetString("endTime")
		global.PrintError(err, "")
		fmt.Println("createMeeting called by " + title)
		fmt.Println("createMeeting called by " + participators)
		fmt.Println("createMeeting called by " + startTime)
		fmt.Println("createMeeting called by " + endTime)
		temp := strings.Split(participators, ",")
		var p []model.Participator
		for _, t := range temp {
			p = append(p, model.Participator{Username: t})
		}
		global.PrintError(operation.AddMeeting(title, p, startTime+":00", endTime+":00"),
			"Meeting: "+title+" successfully created")
	},
}

func init() {
	rootCmd.AddCommand(createMeetingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createMeetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	createMeetingCmd.Flags().StringP("title", "t", "", "Title for meeting")
	createMeetingCmd.Flags().StringP("participators", "p", "", "participators for meeting")
	createMeetingCmd.Flags().StringP("startTime", "s", "", "start time for meeting")
	createMeetingCmd.Flags().StringP("endTime", "e", "", "end time for meeting")
}
