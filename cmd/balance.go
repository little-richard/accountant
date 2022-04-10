package cmd

import (
	"fmt"
	"github.com/little-richard/accountant/database"
	"github.com/spf13/cobra"
	"log"
)

var balanceCmd = &cobra.Command{
	Use:   "balance",
	Short: "Balance for a particular user",
	Long: `This command return balance for a particular user.
Usage: accountant balance <username> `,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("Username not specified")
		}
		username := args[0]
		user, err := database.FindUser(username)

		if user == nil {
			log.Fatal("Username is not exist")
		}

		if err == nil {
			fmt.Println("The current balance for", username, "is: ", user.Balance)
		} else {

			log.Fatal("TEESTE", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(balanceCmd)
}
