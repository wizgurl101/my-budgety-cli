# My Budgety CLI

CLI program for my budgety project.

### How to use CLI

Put the two csv files to be merged and have duplicate entries to be removed in files folder. ENsure only two csv files
are in this folder at a time.

```Golang
    go run . merge-csv
```

## Set the budget amount for a given year and starting month

This set each month of a given year budget to a default of 2000 to My Budgety Database

```Golang
    go run . set-year-budget --year 2025 --start-month 3
```
