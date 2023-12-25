package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

func init() {
	rootCmd.AddCommand(mCmd)
}

var mCmd = &cobra.Command{
	Use:   "m",
	Short: "m is a character count utility for file",
	Long:  `m is a character count utility for file.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			_ = fmt.Errorf("please input right file path")
		} else {
			fmt.Println("m called, arg:", args)
			words, err := getCharacterCount(args[0])
			if err != nil {
				_ = fmt.Errorf("please input right file path")
			}
			fmt.Println("size:" + strconv.FormatInt(words, 10))
		}
	},
}

func getCharacterCount(path string) (int64, error) {

	content, err := os.ReadFile(path)
	if err != nil {
		return 0, err
	}

	return int64(len(content)), nil
}
