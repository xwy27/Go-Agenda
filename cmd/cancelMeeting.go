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
	"fmt"

	"github.com/spf13/cobra"
)

// cancelMeetingCmd represents the cancelMeeting command
var cancelMeetingCmd = &cobra.Command{
	Use:   "cancelMeeting",
	Short: "Cancel a meeting",
	Long:  `Cancel a sponsored meeting by given its title.`,

	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("title")
		// TODO:Error handle
		fmt.Println("cancelMeeting called by " + title)
		// TODO:cancel meeting
	},
}

func init() {
	rootCmd.AddCommand(cancelMeetingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cancelMeetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	cancelMeetingCmd.Flags().StringP("title", "t", "", "title for meeting")
}
