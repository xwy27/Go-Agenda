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

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register an account",
	Long:  `Register an account in Agenda with username, password, email and telephone. Start with managing your own meetings and attending others' meetings!`,

	Run: func(cmd *cobra.Command, args []string) {
		username, err := cmd.Flags().GetString("username")
		global.PrintError(err, "")
		password, err := cmd.Flags().GetString("password")
		global.PrintError(err, "")
		email, err := cmd.Flags().GetString("email")
		global.PrintError(err, "")
		telephone, err := cmd.Flags().GetString("telephone")
		global.PrintError(err, "")

		err = operation.RegisterUser(username, password, email, telephone)
		global.PrintError(err, "Register successfully")
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	registerCmd.Flags().StringP("username", "u", "Anonymous", "Username of the account")
	registerCmd.Flags().StringP("password", "p", "123456", "Password of the account")
	registerCmd.Flags().StringP("email", "e", "Anonymous@Gmail.com", "Email of the account")
	registerCmd.Flags().StringP("telephone", "t", "137123456789", "Telephone of the account")
}
