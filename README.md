# :rocket: My Account CLI Application in Golang

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

This is the basic account cli Application built in golang using cobra library.

## :computer: Tech used

1. golang
2. cobra

A few examples.

Credit:

    accountant credit <username> --amount=<amount> --narration=<narration>

Debit:

    accountant debit <username> --amount=<amount> --narration=<narration>

Balance:

    accountant balance <username>

Transactions:

    accountant transactions <username> <--all, --credit, --debit> 

This project is open source and available under the [MIT License](LICENSE).