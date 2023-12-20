package main

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strconv"
	"strings"
)

func init() {
	rootCmd.AddCommand(wCmd)
}

var wCmd = &cobra.Command{
	Use:   "w",
	Short: "w is a word count utility for words",
	Long:  `w is a word count utility for words.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			_ = fmt.Errorf("please input right file path")
		} else {
			fmt.Println("l called, arg:", args)
			words, err := getWordCount(args[0])
			if err != nil {
				_ = fmt.Errorf("please input right file path")
			}
			fmt.Println("size:" + strconv.FormatInt(words, 10))
		}
	},
}

func getWordCount(path string) (int64, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	reader := bufio.NewReader(file)
	count := 0
	for {
		line, _ := reader.ReadString('\n')
		fields := strings.Fields(line)
		for _, _ = range fields {
			count++
		}
		if line == "" {
			break
		}
	}

	return int64(count), nil
}
