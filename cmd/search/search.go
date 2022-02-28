package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

type Commands struct {
	Behavior   string
	TargetWord string
	Filepaths  []string
}

func NewCommands() *Commands {
	args := os.Args[1:]

	return &Commands{
		Behavior:   args[0],
		TargetWord: args[1],
		Filepaths:  args[2:],
	}
}

type LineInfo struct {
	lineNo int
	line   string
}

func NewLineInfo(lineNo int, line string) LineInfo {
	return LineInfo{lineNo: lineNo, line: line}
}

type FileInfo struct {
	filename string
	lines    []LineInfo
}

func NewFileInfo(filename string, lines []LineInfo) FileInfo {
	return FileInfo{filename: filename, lines: lines}
}

const (
	FIND = "find"
)

var BEHAVIORS = []string{FIND}

func contains(arr []string, word string) bool {
	for _, v := range arr {
		if strings.Contains(v, word) {
			return true
		}
	}
	return false
}

func validate() error {
	args := os.Args[1:]
	if len(args) != 3 {
		fmt.Errorf("arguments must be 3 commands")
	}

	behavior := args[0]
	if !contains(BEHAVIORS, behavior) {
		fmt.Errorf("it is not supported behavior")
	}

	return nil
}

func main() {
	err := validate()
	if err != nil {
		panic("occurred an error when validate commands")
	}

	commands := NewCommands()

	if err != nil {
		fmt.Errorf("failed to read files by path(%s)\n", commands.Filepaths)
	}

	findFilesContainedWord(commands.Filepaths, commands)
}

func findFilesContainedWord(files []string, commands *Commands) {
	size := len(files)
	filenameCh := make(chan string, size)
	ch := make(chan *FileInfo, size)
	var wg sync.WaitGroup
	wg.Add(2)

	go findFileContainedWord(commands.TargetWord, filenameCh, ch, &wg)
	go printFoundLineInfo(ch, &wg)

	for _, filename := range files {
		filenameCh <- filename
	}

	close(filenameCh)

	wg.Wait()
}

func printFoundLineInfo(ch chan *FileInfo, wg *sync.WaitGroup) {
	for fileInfo := range ch {
		if fileInfo != nil && len(fileInfo.lines) > 0 {
			fmt.Println("----- Start to print -----")
			fmt.Printf("filename: %s\n", fileInfo.filename)
			for _, lineInfo := range fileInfo.lines {
				fmt.Printf("no: %d, line: %s\n", lineInfo.lineNo, lineInfo.line)
			}
			fmt.Println("----- The end the line -----")
		}
	}
	wg.Done()
}

func findFileContainedWord(targetWord string, filenameCh chan string, ch chan *FileInfo, wg *sync.WaitGroup) {
	for filename := range filenameCh {
		f, err := os.Open(filename)
		if err != nil {
			panic("failed to read a file")
		}
		fmt.Printf("*** info: start to find word in the file(%s)\n", filename)

		var lines []LineInfo
		fileInfo := NewFileInfo(filename, lines)
		s := bufio.NewScanner(f)
		lineNo := 1
		for s.Scan() {
			line := s.Text()
			if strings.Contains(line, targetWord) {
				fileInfo.lines = append(fileInfo.lines, NewLineInfo(lineNo, line))
			}

			lineNo++
		}

		if len(fileInfo.lines) > 0 {
			ch <- &fileInfo
		}
		ch <- nil
	}
	close(ch)
	wg.Done()
}
