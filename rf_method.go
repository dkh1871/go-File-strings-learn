package main

import (
	"bufio"
	"index/suffixarray"
	"log"
	"os"
	"strings"
)

func getonefile(FilePath string) []byte {
	dat, err := os.ReadFile(FilePath)
	if err != nil {
		log.Fatal()
	}
	return dat
}

func Ts_op_suffixarray(filepath string, search_txt []byte) []int {
	dat := getonefile(filepath)
	idx := suffixarray.New(dat)

	offsets := idx.Lookup(search_txt, -1)
	//fmt.Println(idx.Lookup(search_txt, -1))

	// for _, off := range offsets {
	// 	fmt.Println(off)
	// 	fmt.Println(string(dat[off : off+10]))
	// }

	return offsets
}

func readoneFile(FilePath string) int {
	f, err := os.Open(FilePath)
	if err != nil {
		log.Fatal()
	}
	defer f.Close()

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)
	var filelines []string

	//if I don't care about the arry do I add a counter object? most likely seem pretty strigit forword. I should test it.
	// I could also just check for the match and return the whole text line???? might be the easiest way since I should know the input file type.
	// might be a bit more tricky depeing on what I want to return, if I return the field only what if the user wants the table its attached too?
	// might not be though??? I'll have to think about the data structs and how they work
	for fileScanner.Scan() {
		val := fileScanner.Text()

		if strings.Contains(val, "Cool") == true {
			filelines = append(filelines, val)
		}

	}

	// with this method the line would be the idex of the arry
	return len(filelines)
}
