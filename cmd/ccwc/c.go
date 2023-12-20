package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/fs"
	"os"
	"strconv"
)

func init() {
	rootCmd.AddCommand(cCmd)
}

var cCmd = &cobra.Command{
	Use:   "c",
	Short: "c is outputs the number of bytes in a file",
	Long: `c is outputs the number of bytes in a files.
It is a clone of the wc command.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		if len(args) != 1 {
			_ = fmt.Errorf("please input right file path")
		} else {
			fmt.Println("c called, arg:", args)
			fileInfos, err := getFileInfo(args[0])
			if err != nil {
				_ = fmt.Errorf("please input right file path")
			}
			fmt.Println("size:" + strconv.FormatInt(fileInfos.Size(), 10))
		}

	},
}

func getFileInfo(path string) (fs.FileInfo, error) {
	info, err := os.Stat(path)

	if err != nil {
		return nil, err
	}
	return info, nil

}
