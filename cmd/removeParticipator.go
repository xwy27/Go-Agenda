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

	"github.com/spf13/cobra"
)

// removeParticipatorCmd represents the removeParticipator command
var removeParticipatorCmd = &cobra.Command{
	Use:   "removeParticipator",
	Short: "Remove a participator for a meeting",
	Long: `Remove a participator for a sponsored meeting.
Specify the meeting title and the new participator name.`,

	Run: func(cmd *cobra.Command, args []string) {
		title, err := cmd.Flags().GetString("title")
		global.PrintError(err, "")
		participator, err := cmd.Flags().GetString("participator")
		global.PrintError(err, "")

		err = operation.DeleteParticipator(title, participator)
		global.PrintError(err, "Delete "+participator+" successfully")
	},
}

func init() {
	rootCmd.AddCommand(removeParticipatorCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removeParticipatorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	removeParticipatorCmd.Flags().StringP("title", "t", "", "Meeting title")
	removeParticipatorCmd.Flags().StringP("participator", "p", "", "Meeting participator")
}
