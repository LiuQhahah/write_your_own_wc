package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(dCmd)
}

var dCmd = &cobra.Command{
	Use:   "a",
	Short: "show default info",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("a called, arg:", args)
		infos, err := getDefaultOutput(args[0])
		if err != nil {
			_ = fmt.Errorf("please input right file path")
		}
		fmt.Printf("word: %d, size:%d, infos.word:%d ", +infos.word, infos.size, infos.line)

	},
}

type info struct {
	word int64
	size int64
	line int64
}

func getDefaultOutput(path string) (*info, error) {
	words, err := getCharacterCount(path)
	if err != nil {
		_ = fmt.Errorf("please input right file path")
	}
	fileInfos, err := getFileInfo(path)
	size := fileInfos.Size()
	if err != nil {
		_ = fmt.Errorf("please input right file path")

	}
	lines, err := getFileLine(path)
	if err != nil {
		_ = fmt.Errorf("please input right file path")

	}

	return &info{
		words,
		size,
		lines,
	}, nil

}
