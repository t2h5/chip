package cmd

import (
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "chip",
	Version: "0.0.1",
	Short: "CHeck IP address",
	Long:  `Check current IP address.`,
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := http.Get("https://checkip.amazonaws.com/")
		if err != nil {
			cmd.SetOutput(os.Stderr)
			cmd.Println(err.Error())
			os.Exit(1)
		}
		defer resp.Body.Close()

		ip, err := io.ReadAll(resp.Body)
		if err != nil {
			cmd.SetOutput(os.Stderr)
			cmd.Println(err.Error())
			os.Exit(1)
		}
		cmd.Printf("current ip: %s", string(ip))
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
