/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// fbiCmd represents the fbi command
var fbiCmd = &cobra.Command{
	Use:   "fbi",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("fbi called")
	},
}

func init() {
	rootCmd.AddCommand(fbiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fbiCmd.PersistentFlags().String("foo", "", "A help for foo")
	fbiCmd.PersistentFlags().String("url", "", "Url")
	fbiCmd.PersistentFlags().String("host", "", "IP do dispositivo")
	fbiCmd.PersistentFlags().IntP("port", "p", 80, "Porta do dispositivo")
	fbiCmd.PersistentFlags().String("user", "admin", "Usuário")
	fbiCmd.PersistentFlags().String("password", "dankia77", "Senha")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fbiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
