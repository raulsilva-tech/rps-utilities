/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/raulsilva-tech/RPSUtilities/internal/usecase"
	"github.com/spf13/cobra"
)

// clickCmd represents the click command
var clickCmd = &cobra.Command{
	Use:   "click",
	Short: "Envia um click na coordenada recebida",
	Long:  ``,
	RunE:  clickFunc(),
}

func clickFunc() RunEFun {
	return func(cmd *cobra.Command, args []string) error {

		x, _ := cmd.Flags().GetInt("x")
		y, _ := cmd.Flags().GetInt("y")

		fmt.Println(x, y)
		ucClick := usecase.NewClickUseCase()
		ucClick.Execute(x, y)

		return nil
	}
}

func init() {
	rootCmd.AddCommand(clickCmd)

	// Here you will define your flags and configuration settings.

	clickCmd.PersistentFlags().IntP("x", "x", 0, "Coordenada X")
	clickCmd.PersistentFlags().IntP("y", "y", 0, "Coordenada X")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clickCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clickCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
