package ynabhelper

import (
	"fmt"
	"os"
	"time"

	"github.com/jsuar/ynab-bitcoin-balance-tracker/pkg/envhelper"
	"github.com/ryanuber/columnize"
	"go.bmvs.io/ynab"
	"go.bmvs.io/ynab/api"
	"go.bmvs.io/ynab/api/account"
	"go.bmvs.io/ynab/api/budget"
	"go.bmvs.io/ynab/api/transaction"
)

const baseURL string = "https://api.youneedabudget.com/v1"

// YnabHelper provides helper functions
type YnabHelper struct {
	BearerToken string
	verbose     bool
}

// Init will initialize the YnabHelper object
func (yh *YnabHelper) Init(verbose bool) {
	var err error
	yh.verbose = verbose
	yh.BearerToken, err = envhelper.GetRequiredEnv("YNAB_BEARER_TOKEN")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// ListBudgets lists all the budgets associated with the bearer token
func (yh *YnabHelper) ListBudgets() {
	var output []string

	c := ynab.NewClient(yh.BearerToken)
	budgets, err := c.Budget().GetBudgets()
	if err != nil {
		panic(err)
	}

	output = append(output, "Name|ID")
	for _, budget := range budgets {
		output = append(output, fmt.Sprintf("%s|%s", budget.Name, budget.ID))
	}

	result := columnize.SimpleFormat(output)
	fmt.Printf("%s\n\n", result)
}

func (yh *YnabHelper) getBudgetByName(budgetName string) (*budget.Summary, error) {
	c := ynab.NewClient(yh.BearerToken)
	budgets, err := c.Budget().GetBudgets()
	if err != nil {
		panic(err)
	}

	for _, budget := range budgets {
		if budget.Name == budgetName {
			return budget, nil
		}
	}

	return nil, fmt.Errorf("No budget found with name %s", budgetName)
}

// ListAccounts list all the accounts under a budget
func (yh *YnabHelper) ListAccounts(budgetName string) {
	var output []string

	budget, err := yh.getBudgetByName(budgetName)
	if err != nil {
		panic(err)
	}

	c := ynab.NewClient(yh.BearerToken)
	accounts, err := c.Account().GetAccounts(budget.ID)
	if err != nil {
		panic(err)
	}

	output = append(output, "Name|ID")
	for _, account := range accounts {
		output = append(output, fmt.Sprintf("%s|%s", account.Name, account.ID))
	}

	result := columnize.SimpleFormat(output)
	fmt.Printf("%s\n\n", result)
}

func (yh *YnabHelper) getAccountByName(budgetName string, accountName string) (*account.Account, error) {
	budget, err := yh.getBudgetByName(budgetName)
	if err != nil {
		panic(err)
	}

	c := ynab.NewClient(yh.BearerToken)
	accounts, err := c.Account().GetAccounts(budget.ID)
	if err != nil {
		panic(err)
	}

	for _, account := range accounts {
		if account.Name == accountName {
			return account, nil
		}
	}

	return nil, fmt.Errorf("No account found with name %s under budget %s", accountName, budgetName)
}

// GetAccountBalance returns the balance for a specificed budget and account
func (yh *YnabHelper) GetAccountBalance() int64 {
	budgetID, err := envhelper.GetRequiredEnv("YNAB_BUDGET_ID")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	accountID, err := envhelper.GetRequiredEnv("YNAB_ACCOUNT_ID")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	c := ynab.NewClient(yh.BearerToken)
	account, err := c.Account().GetAccount(budgetID, accountID)
	if err != nil {
		panic(err)
	}

	return account.Balance
}

// CreateTransaction creates a YNAB transaction
func (yh *YnabHelper) CreateTransaction(amount int64) error {
	c := ynab.NewClient(yh.BearerToken)

	budgetID, err := envhelper.GetRequiredEnv("YNAB_BUDGET_ID")
	if err != nil {
		return err
	}

	accountID, err := envhelper.GetRequiredEnv("YNAB_ACCOUNT_ID")
	if err != nil {
		return err
	}

	payloadMemo := "Auto filled with YNAB Bitcoin Balance Tracker"
	p := transaction.PayloadTransaction{
		AccountID:  accountID,
		Date:       api.Date{time.Now()},
		Amount:     amount,
		Cleared:    transaction.ClearingStatusCleared,
		Approved:   true,
		PayeeID:    nil,
		CategoryID: nil,
		Memo:       &payloadMemo,
		FlagColor:  nil,
	}
	transaction := c.Transaction()
	createdTransactions, err := transaction.CreateTransaction(budgetID, p)
	if err != nil {
		return err
	}

	if yh.verbose {
		for _, t := range createdTransactions.Transactions {
			fmt.Printf("Transaction %s with amount %.2f was created on %s\n", t.ID, float64(t.Amount)/1000.0, t.Date.Format("Mon Jan _2 15:04:05 2006"))
		}
	}
	return nil
}
