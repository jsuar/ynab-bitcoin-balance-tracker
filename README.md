# YNAB Bitcoin Balance Tracker

This simple CLI helps sync bitcoin price changes to your YNAB account(s). It also has some functions to help find the necessary IDs for automating the syncing process since YNAB doesn't provide end users with easy access to budget IDs, account IDs, etc.

Note: I built this for my minor understanding of Bitcoin. 

# How to use

## Prerequisites

Required environment variables:
```
YNAB_BEARER_TOKEN
YNAB_BUDGET_ID
YNAB_ACCOUNT_ID
BITCOIN_ADDR
```

A [bearer token](https://api.youneedabudget.com/#authentication-overview) can be created from [your My Account page](https://app.youneedabudget.com/settings) in YNAB web app.

Use the [resource functions](#resource-functions) to obtain the `YNAB_BUDGET_ID` and `YNAB_ACCOUNT_ID` values.

And set `BITCOIN_ADDR` for the transaction balance you would like to track.

## Commands

The `balance` command will display your YNAB balance and the current value of a Bitcoin transaction.

```
$ ynab-bitcoin-balance-tracker balance
Current Value          $
Bitcoin balance (USD)  5222.84
YNAB Account           5270.11
Delta                  -52.72
```

Adding the `--sync` flag will create a transction in YNAB based on the account and budget IDs set as environment variables.

```
$ ynab-bitcoin-balance-tracker balance --sync
Current Value          $
Bitcoin balance (USD)  5280.51
YNAB Account           5279.99
Delta                  0.52

Transaction of amount 0.52 successfully created.
```

## Resource Functions

The `budgets` command will display all budgets under your account. I think "My Budget" is the default name given for your first account.

```
$ ynab-bitcoin-balance-tracker budgets
Name         ID
My Budget    DF5F6915-EB1F-442B-A6D9-40C9EBA36D82
Work Budget  EDB6DF43-2638-4C10-A9F8-04A938C6BCF1
```

The accounts command will default to "My Budget" but any budget can be listed by using the `--budgetName` flag.

```
$ ynab-bitcoin-balance-tracker accounts --budgetName "Work Budget"
Name      ID
Checking  9CF5EC96-9334-41EA-AA59-F17C46DBAF63
Credit    026D03F1-41DF-4FFF-9ED0-B95FE3F3D382
```