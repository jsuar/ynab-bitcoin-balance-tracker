# YNAB Bitcoin Balance Tracker

This simple CLI helps sync bitcoin price changes to your YNAB account(s). It also has some functions to help find the necessary IDs for automating the syncing process since YNAB doesn't provide end users with easy access to budget IDs, account IDs, etc.

Note: I built this for my minor understanding of Bitcoin. 

# How to use

Required environment variables:
```
YNAB_BEARER_TOKEN
YNAB_BUDGET_ID
YNAB_ACCOUNT_ID
BITCOIN_ADDR
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