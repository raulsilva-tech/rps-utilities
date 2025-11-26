/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"os"

	"github.com/raulsilva-tech/RPSUtilities/internal/usecase"
	"github.com/spf13/cobra"
)

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Recebe url http para envio de requisições",
	Long:  ``,
	RunE:  httpFunc(),
}

func httpFunc() RunEFun {
	return func(cmd *cobra.Command, args []string) error {

		user, _ := cmd.Flags().GetString("user")
		password, _ := cmd.Flags().GetString("password")
		url, _ := cmd.Flags().GetString("url")
		auth, _ := cmd.Flags().GetString("auth")
		method, _ := cmd.Flags().GetString("method")
		stream, _ := cmd.Flags().GetBool("stream")

		if method == "GET" {

			if stream {
				ucGetStreamRequest := usecase.NewSendGETStreamRequestUseCase()
				output, err := ucGetStreamRequest.Execute(url, user, password, auth)
				json.NewEncoder(os.Stdout).Encode(output)
				return err
			} else {
				ucGetRequest := usecase.NewSendGETRequestUseCase()
				output, err := ucGetRequest.Execute(url, user, password, auth)
				json.NewEncoder(os.Stdout).Encode(output)
				return err
			}

		}

		return nil
	}
}

func init() {
	rootCmd.AddCommand(httpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	httpCmd.PersistentFlags().String("url", "", "Url")
	httpCmd.PersistentFlags().String("user", "", "Usuário")
	httpCmd.PersistentFlags().String("password", "", "Senha")
	httpCmd.PersistentFlags().StringP("method", "m", "", "Método")
	httpCmd.PersistentFlags().StringP("auth", "a", "", "Modo de autenticação")
	httpCmd.PersistentFlags().BoolP("stream", "s", false, "Modo de streaming de dados")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// httpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
