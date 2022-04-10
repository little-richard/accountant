package cmd

import (
	"fmt"
	"github.com/little-richard/accountant/database"
	"github.com/spf13/cobra"
	"log"
)

var creditCmd = &cobra.Command{
	Use:   "credit",
	Short: "Create a credit transaction",
	Long: `This command creates a credit transaction for a particular user.
Usage: accountant credit <username> --amount=<amount> --narration=<narration>.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("Username not specified")
		}
		username := args[0]
		user, err := database.FindOrCreateUser(username)
		if err != nil {
			log.Fatal(err)
		}
		user.Balance = user.Balance + creditAmount
		creditTransaction := database.Transaction{Amount: creditAmount, Type: "credit", Narration: creditNarration}
		user.Transactions = append(user.Transactions, creditTransaction)
		database.UpdateUser(user)
		fmt.Println("Transaction created successfully. Balance:", user.Balance)
	},
}

var creditNarration string
var creditAmount int64

func init() {
	rootCmd.AddCommand(creditCmd)
	creditCmd.Flags().StringVarP(&creditNarration, "narration", "n", "", "Narration for this credit transaction")
	creditCmd.Flags().Int64VarP(&creditAmount, "amount", "a", 0, "Amount to be credited")
	creditCmd.MarkFlagRequired("narration")
	creditCmd.MarkFlagRequired("amount")
}
