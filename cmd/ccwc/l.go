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
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			_ = fmt.Errorf("please input right file path")
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
	//info, err := os.Stat(path)
	//if err != nil {
	//	return 0, err
	//}
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()
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
