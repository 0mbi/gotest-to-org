package main

import (
	"bufio"
	"fmt"
	"os"
)

func mkOrg(lines []string) (string, error) {
	if len(lines) <= 1 {
		return "", fmt.Errorf("Input is too small: %+v", lines)
	}
	mainPara := ""
	para := ""
	subPara := make([]string, 0, len(lines))
	blocks := make([]string, 0, len(lines))
	inBlock := false
	for _, l := range lines {
		if !inBlock && len(l) >= 1 && l[:1] != "=" && l[:1] != "-" && l[:1] != " " {
			if len(l) >= 7 {
				mainPara = fmt.Sprintf("* %s\n", l)
				fmt.Println(mainPara)
				for _, b := range blocks {
					fmt.Println(b)
				}
				blocks = make([]string, 0, len(lines))
			}
		}
		if len(l) >= 7 && l[:7] == "=== RUN" {
			inBlock = true
			continue
		}
		if inBlock {
			if len(l) >= 4 && l[:3] == "---" {
				stripDashes := l[4:]
				para = fmt.Sprintf("** %s\n", stripDashes)
				blocks = append(blocks, para)
				blocks = append(blocks, subPara...)
				subPara = make([]string, 0, len(lines))
				inBlock = false
				continue
			}
			subPara = append(subPara, l)
		}
	}
	return "", nil
}

func main() {
	lines := make([]string, 0)
	if len(os.Args) > 1 {
		f, err := os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
		defer f.Close()
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
	}
	out, err := mkOrg(lines)
	if err != nil {
		panic(err)
	}
	println(out)
}
