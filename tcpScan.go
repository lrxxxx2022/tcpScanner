package main

import (
	"bufio"
	"fmt"
	"github.com/gookit/color"
	"net"
	"os"
	"sync"
	"time"
)

func main() {
	var wg = sync.WaitGroup{}
	start := time.Now()
	//CheckPortStatus(&wg)
	lines, _, _ := ReadLines("ip.txt")
	for line := range lines {
		wg.Add(1)
		go func(line int) {
			_, err := net.DialTimeout("tcp", lines[line], 4*time.Second)
			time.Sleep(time.Second * 4)
			if err != nil {
				//fmt.Printf("%40v   \033[1;31;40m%s\033[0m\n", lines[line], "close")
				color.Red.Printf("%40v    close\n", lines[line])
				//ColorPrint(lines[line], "close", FontColor.red)
			} else {
				//fmt.Printf("%50v   \033[1;32;40m%s\033[0m\n", lines[line], "open")
				//ColorPrint(lines[line], "open", FontColor.green)
				color.Green.Printf("%40v    open\n", lines[line])
			}
			wg.Done()
		}(line)
	}
	wg.Wait()
	fmt.Println("--------------------------------------------------------------------------")
	fmt.Println("扫描用时:", time.Since(start))
}

func ReadLines(path string) ([]string, int, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("\"ip.txt\" open failed!")
		return nil, 0, err
	}
	defer file.Close()
	var lines []string
	lineCount := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		lineCount++
	}
	return lines, lineCount, scanner.Err()
}
