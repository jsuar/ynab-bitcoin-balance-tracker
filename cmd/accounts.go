// Package cmd handles all CLI calls
/*
Copyright © 2020 John Suarez jsuar@users.noreply.github.com

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
	"github.com/jsuar/ynab-bitcoin-balance-tracker/pkg/ynabhelper"
	"github.com/spf13/cobra"
)

// accountsCmd represents the accounts command
var accountsCmd = &cobra.Command{
	Use:   "accounts",
	Short: "Lists the name and ID of all YNAB accounts",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		logger, err := InitLogger()
		if err != nil {
			panic(err)
		}
		budgetName, err := cmd.Flags().GetString("budgetName")
		HandleError(err, true, logger)

		yh := new(ynabhelper.YnabHelper)
		err = yh.Init(true, logger)
		HandleError(err, true, logger)
		yh.ListAccounts(budgetName)
	},
}

func init() {
	rootCmd.AddCommand(accountsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	accountsCmd.PersistentFlags().String("budgetName", "My Budget", "Name of the budget")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// accountsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
