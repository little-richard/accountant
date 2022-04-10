package cmd

import (
	"fmt"
	"github.com/little-richard/accountant/database"
	"github.com/spf13/cobra"
	"log"
)

var debitCmd = &cobra.Command{
	Use:   "debit",
	Short: "Create a debit transaction",
	Long: `This command creates a debit transaction for a particular user.
Usage: accountant debit <username> --amount=<amount> --narration=<narration>.`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			log.Fatal("Username not specified")
		}

		username := args[0]
		user, err := database.FindOrCreateUser(username)
		if err != nil {
			log.Fatal(err)
		}

		if user.Balance > debitAmount {
			user.Balance -= debitAmount
			debitTransaction := database.Transaction{Amount: debitAmount, Type: "debit", Narration: debitNarration}
			user.Transactions = append(user.Transactions, debitTransaction)
			database.UpdateUser(user)
			fmt.Println("Transaction created successfully! Balance:", user.Balance)
		} else {
			fmt.Println("Insufficient funds!")
		}
	},
}

var debitNarration string
var debitAmount int64

func init() {
	rootCmd.AddCommand(debitCmd)
	debitCmd.Flags().StringVarP(&debitNarration, "narration", "n", "", "Narration for this debit transaction")
	debitCmd.Flags().Int64VarP(&debitAmount, "amount", "a", 0, "Amount to be debited")
	debitCmd.MarkFlagRequired("narration")
	debitCmd.MarkFlagRequired("amount")
}
