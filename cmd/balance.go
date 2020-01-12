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
	"os"

	"github.com/jsuar/ynab-bitcoin-balance-tracker/pkg/bitcoinhelper"
	"github.com/jsuar/ynab-bitcoin-balance-tracker/pkg/ynabhelper"
	"github.com/ryanuber/columnize"
	"github.com/spf13/cobra"
)

func handleError(err error, exit bool) {
	if err != nil {
		fmt.Println(err)
		if exit {
			os.Exit(1)
		}
	}
}

// balanceCmd represents the balance command
var balanceCmd = &cobra.Command{
	Use:   "balance",
	Short: "Display balances and sync to YNAB",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		exitOnError := true
		// Handle flags
		verbose, err := cmd.Flags().GetBool("verbose")
		handleError(err, exitOnError)
		currency, err := cmd.Flags().GetString("currency")
		handleError(err, exitOnError)
		sync, err := cmd.Flags().GetBool("sync")
		handleError(err, exitOnError)

		// Handle bitcoin info
		btcHelper := new(bitcoinhelper.BitcoinHelper)
		btcHelper.Init(verbose)
		exchangeRate, err := btcHelper.GetMarketPrice(currency, "Last")
		handleError(err, exitOnError)

		btcBalance, err := btcHelper.GetAddressBalance()
		handleError(err, exitOnError)

		var output []string
		output = append(output, "Current Value|$")

		convertedBalance := float64(btcBalance) / 100000000 * exchangeRate
		output = append(output, fmt.Sprintf("Bitcoin balance (%s)|%.2f", currency, convertedBalance))

		ynabhelper := new(ynabhelper.YnabHelper)
		ynabhelper.Init(verbose)
		accountBalance := ynabhelper.GetAccountBalance()
		output = append(output, fmt.Sprintf("YNAB Account|%.2f", float64(accountBalance)/1000.0))

		delta := int64(convertedBalance*1000) - accountBalance
		output = append(output, fmt.Sprintf("Delta|%.2f", float64(delta)/1000.0))

		result := columnize.SimpleFormat(output)
		fmt.Printf("%s\n\n", result)

		if sync {
			err = ynabhelper.CreateTransaction(delta)
			handleError(err, exitOnError)
		}
	},
}

func init() {
	rootCmd.AddCommand(balanceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	balanceCmd.PersistentFlags().BoolP("verbose", "v", false, "Verbose output")
	balanceCmd.PersistentFlags().String("currency", "USD", "Currency to retrieve (USD, CAD, HKD, etc.)")

	balanceCmd.PersistentFlags().BoolP("sync", "s", false, "Sync balance delta with YNAB")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// balanceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
