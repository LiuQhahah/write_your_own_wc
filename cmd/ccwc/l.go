package main

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

func init() {
	rootCmd.AddCommand(lCmd)
}

var lCmd = &cobra.Command{
	Use:   "l",
	Short: "option -l that outputs the number of lines in a file",
	Long:  `option -l that outputs the number of lines in a file.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			scanner := bufio.NewScanner(os.Stdin)
			var count int

			for scanner.Scan() {
				count += 1
			}
			if scanner.Err() != nil {
				fmt.Printf("Error reading input: %s\n", scanner.Err())
				return
			}
			fmt.Println("size:" + strconv.FormatInt(int64(count), 10))

		} else {
			fmt.Println("l called, arg:", args)
			lines, err := getFileLine(args[0])
			if err != nil {
				_ = fmt.Errorf("please input right file path")
			}
			fmt.Println("size:" + strconv.FormatInt(lines, 10))

		}

	},
}

func getFileLine(path string) (int64, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			_ = fmt.Errorf("please input right file path")
		}
	}(file)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var count int
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return int64(count), nil
}
