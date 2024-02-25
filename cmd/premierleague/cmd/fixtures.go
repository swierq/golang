/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/swierq/golang/internal/premierleague"
)

// fixturesCmd represents the fixtures command
var fixturesCmd = &cobra.Command{
	Use:   "fixtures",
	Short: "Show upcomming fixtures",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		team, _ := cmd.Flags().GetString("team")
		dback, _ := cmd.Flags().GetInt("dbck")
		dfwd, _ := cmd.Flags().GetInt("dfwd")
		allFixtures(team, dback, dfwd)
	},
}

func init() {
	rootCmd.AddCommand(fixturesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fixturesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fixturesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	fixturesCmd.Flags().StringP("team", "t", "", "Team, eg MUN.")
	fixturesCmd.Flags().Int("dbck", 10, "Number of days to look back.")
	fixturesCmd.Flags().Int("dfwd", 10, "Number of days to look forward.")
}

func allFixtures(team string, dbck, dfwd int) {
	client := http.Client{}
	err := premierleague.AllFixtures(team, &client, dbck, dfwd)
	if err != nil {
		fmt.Printf("Something is wrong: %v", err)
	}
}
