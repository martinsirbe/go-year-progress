package main

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "yp",
	Short: "Year Progress Indicator",
	Long: `The Year Progress Indicator is a Go application that visually represents how much of 
the current year has elapsed. It calculates the percentage of the year that has passed based on 
today's date and displays this information through a progress bar composed of filled and unfilled 
segments.`,
	Run: func(cmd *cobra.Command, args []string) {
		runYearProgress()
	},
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringP("location", "l", "UTC", "Timezone location")
	rootCmd.PersistentFlags().IntP("total-blocks", "t", 50, "Total number of blocks in the progress bar")
	rootCmd.PersistentFlags().StringP("filled-char", "f", "█", "Character for filled sections of the progress bar")
	rootCmd.PersistentFlags().StringP("empty-char", "e", "▁", "Character for empty sections of the progress bar")

	if err := viper.BindPFlag("location", rootCmd.PersistentFlags().Lookup("location")); err != nil {
		fmt.Printf("failed to bind location flag: %s", err)
		os.Exit(1)
	}
	if err := viper.BindPFlag("total_blocks", rootCmd.PersistentFlags().Lookup("total-blocks")); err != nil {
		fmt.Printf("failed to bind total_blocks flag: %s", err)
		os.Exit(1)
	}
	if err := viper.BindPFlag("filled_char", rootCmd.PersistentFlags().Lookup("filled-char")); err != nil {
		fmt.Printf("failed to bind filled_char flag: %s", err)
		os.Exit(1)
	}
	if err := viper.BindPFlag("empty_char", rootCmd.PersistentFlags().Lookup("empty-char")); err != nil {
		fmt.Printf("failed to bind empty_char flag: %s", err)
		os.Exit(1)
	}
}

func initConfig() {
	viper.SetEnvPrefix("yp")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		_, notFound := err.(viper.ConfigFileNotFoundError)
		if !notFound {
			fmt.Printf("failed to read in config: %s\n", err)
			os.Exit(1)
		}
	}
}

func runYearProgress() {
	location := viper.GetString("location")
	totalBlocks := viper.GetInt("total_blocks")
	filledChar := viper.GetString("filled_char")
	emptyChar := viper.GetString("empty_char")

	loc, err := time.LoadLocation(location)
	if err != nil {
		fmt.Printf("Failed to load location: %s\n", err)
		return
	}

	now := time.Now().In(loc)
	startOfYear := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, loc)
	endOfYear := time.Date(now.Year(), 12, 31, 23, 59, 59, 0, loc)
	totalDays := endOfYear.Sub(startOfYear).Hours() / 24
	daysElapsed := now.Sub(startOfYear).Hours() / 24
	progress := (daysElapsed / totalDays) * 100

	progressBar := ""

	filledBlocks := int((progress / 100) * float64(totalBlocks))
	for i := 0; i < filledBlocks; i++ {
		progressBar += filledChar
	}

	emptyBlocks := totalBlocks - filledBlocks
	for i := 0; i < emptyBlocks; i++ {
		progressBar += emptyChar
	}

	fmt.Printf("%s %.2f%%\n", progressBar, progress)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
