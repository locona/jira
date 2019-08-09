package cmd

import (
	"bufio"
	"fmt"
	"os"
	"syscall"

	"github.com/3-shake/jira/pkg/auth"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

// authCmd represents the auth command
var authCmd = &cobra.Command{
	Use: "auth",
	RunE: func(cmd *cobra.Command, args []string) error {
		return Auth()
	},
}

func init() {
	rootCmd.AddCommand(authCmd)
}

func Auth() error {
	fmt.Print("Username: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	username := scanner.Text()

	fmt.Print("Password: ")
	password, _ := terminal.ReadPassword(int(syscall.Stdin))
	auth.Write(username, string(password))
	return nil
}
