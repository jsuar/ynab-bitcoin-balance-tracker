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
	"fmt"

	"github.com/jsuar/ynab-bitcoin-balance-tracker/pkg/bitcoinhelper"
	"github.com/jsuar/ynab-bitcoin-balance-tracker/pkg/ynabhelper"
	"github.com/spf13/cobra"
)

// balanceCmd represents the balance command
var balanceCmd = &cobra.Command{
	Use:   "balance",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetBool("verbose")
		currency, err := cmd.Flags().GetString("currency")
		if err != nil {
			panic(err)
		}
		btcHelper := new(bitcoinhelper.BitcoinHelper)
		btcHelper.Init(verbose)
		conversionPrice := btcHelper.GetMarketPrice(currency, "Last")
		btcBalance := btcHelper.GetAddressBalance()
		convertedBalance := float64(btcBalance) / 100000000 * conversionPrice
		fmt.Printf("Current balance (%s): %.2f\n", currency, convertedBalance)

		ynabhelper := new(ynabhelper.YnabHelper)
		ynabhelper.Init(verbose)
		accountBalance := ynabhelper.GetAccountBalance()
		fmt.Printf("Current account balance: %d\n", accountBalance)

		ynabhelper.CreateTransaction()
	},
}

func init() {
	rootCmd.AddCommand(balanceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	balanceCmd.PersistentFlags().BoolP("verbose", "v", false, "Verbose output")
	balanceCmd.PersistentFlags().String("currency", "USD", "Currency to retrieve (USD, CAD, HKD, etc.)")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// balanceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
