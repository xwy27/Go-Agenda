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
	"Go-Agenda/operation"

	"github.com/spf13/cobra"
)

// quitMeetingCmd represents the quitMeeting command
var quitMeetingCmd = &cobra.Command{
	Use:   "quitMeeting",
	Short: "Quit a meeting",
	Long:  `Quit a specified meeting with given title.`,

	Run: func(cmd *cobra.Command, args []string) {
		title, err := cmd.Flags().GetString("title")
		global.PrintError(err, "")
		// fmt.Println("quitMeeting called by " + title)
		// Quit meeting
		err = operation.QuitMeeting(title)
		global.PrintError(err, "Quit "+title+" successfully")
	},
}

func init() {
	rootCmd.AddCommand(quitMeetingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// quitMeetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	quitMeetingCmd.Flags().StringP("title", "t", "", "title for meeting")
}
