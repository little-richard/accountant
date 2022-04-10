package cmd

import (
	. "github.com/ahmetalpbalkan/go-linq"
	"github.com/jedib0t/go-pretty/table"
	"github.com/little-richard/accountant/database"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var transactionsCmd = &cobra.Command{
	Use:   "transactions",
	Short: "All transactions for a particular user",
	Long: `This command return all transactions for a particular user.
Usage: accountant transactions <username> `,
	Run: func(cmd *cobra.Command, args []string) {

		debitStatus, _ := cmd.Flags().GetBool("debit")
		creditStatus, _ := cmd.Flags().GetBool("credit")
		allStatusF, _ := cmd.Flags().GetBool("all")
		allStatus := (debitStatus == false && creditStatus == false) || (debitStatus == true && creditStatus == true) || allStatusF

		if len(args) < 1 {
			log.Fatal("Username not specified")
		}

		username := args[0]
		user, err := database.FindOrCreateUser(username)

		if err != nil {
			log.Fatal(err)
		}

		if user == nil {
			log.Fatal("Username is not exist")
		}

		if allStatus {
			printTable(user.Transactions, user.Balance)
		} else if debitStatus {
			debitTransactions := filterTransaction(user.Transactions, "debit")
			printTable(debitTransactions, user.Balance)
		} else if creditStatus {
			creditTransactions := filterTransaction(user.Transactions, "credit")
			printTable(creditTransactions, user.Balance)
		}
	},
}

func init() {
	rootCmd.AddCommand(transactionsCmd)
	transactionsCmd.Flags().BoolP("debit", "d", false, "Debit transactions")
	transactionsCmd.Flags().BoolP("credit", "c", false, "Credit transactions")
	transactionsCmd.Flags().BoolP("all", "a", false, "All transactions")
}

func printTable(transactions []database.Transaction, balance int64) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Type", "Narration", "Amount"})

	for index, transaction := range transactions {
		t.AppendRow([]interface{}{index, transaction.Type, transaction.Narration, transaction.Amount})
	}
	t.AppendFooter(table.Row{"", "", "TOTAL", balance})
	t.Render()
}

func filterTransaction(transactions []database.Transaction, typeTransaction string) []database.Transaction {
	var transactionsFilter []database.Transaction

	From(transactions).Where(func(i interface{}) bool {
		return i.(database.Transaction).Type == typeTransaction
	}).Select(func(i interface{}) interface{} {
		return i.(database.Transaction)
	}).ToSlice(&transactionsFilter)

	return transactionsFilter
}
