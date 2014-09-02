package main

import (
	"fmt"
	"strings"
	"strconv"
	"os"
	"bufio"
	"log"
	"sort"
)

func GrabLines(args string) []string {
	var lines []string
	file, err := os.Open(args)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func Reversed(blocks []float64) []float64 {
	var reversedblocks []float64
	for _, block := range blocks {
		reversedblocks = append(reversedblocks, block)
	}
	return reversedblocks
}

func IntBlocks(blocks []string) []float64{
	var intblocks []float64
	for _, block := range blocks{
		s,_ := strconv.ParseFloat(block, 64)
		intblocks = append(intblocks, s)
		sort.Float64s(intblocks)
	}
	return intblocks 
}
func HonestNaomi(naomisblocks []float64, kensblocks []float64) int {
	points := 0
	var kenchoice float64
	var block float64
	returnslice := naomisblocks
	kensnewblocks := kensblocks
	for i := 1; i <= len(naomisblocks); i++ {
		block, returnslice = Shift(returnslice)
		kenchoice, kensnewblocks = KenTurn(block, kensnewblocks)
		if block > kenchoice {
			points += 1
		}
	}
	return points
}

func Delete(blocks []float64, index int) []float64 {
	var returnslice []float64
	for i, block := range blocks {
		if i != index {
			returnslice = append(returnslice, block)
		}
	}
	return returnslice
}

func KenTurn(naomiblock float64, kensblocks []float64) (float64, []float64) {
	var returnblock float64
	var returnslice []float64
	for i, kenblock := range kensblocks{
		if kenblock > naomiblock {
			returnblock = kenblock
			returnslice = Delete(kensblocks, i)
			return returnblock, returnslice
		} else {
			returnblock, returnslice = Shift(kensblocks)
		}
	}
	return returnblock, returnslice
}

func DeceitfulNaomi(naomisblocks []float64, kensblocks []float64) int {
	naomischangedblocks := naomisblocks
	points := 0
	var naomisturn float64
	var kenblock float64
	kenschangedblocks := kensblocks
	for i, naomisblock := range naomischangedblocks {
		if naomisblock < kenschangedblocks[0] {
			naomisturn = kenschangedblocks[len(kenschangedblocks)-1] - .0001
			naomischangedblocks = Delete(naomischangedblocks, i)
		} else {
			naomisturn = naomischangedblocks[len(naomischangedblocks)-1]
			_, naomischangedblocks = Shift(naomischangedblocks)
		}
		kenblock, kenschangedblocks = KenTurn(naomisturn, kenschangedblocks)
		if naomisturn > kenblock {
			points += 1
		}
	}
	return points
}

func Shift(inslice []float64) (float64, []float64) {
	firstitem := inslice[0]
	returnslice := inslice[1:]
	return firstitem, returnslice
}
func main() {
	argsWithoutProgram := os.Args[1]
	lines := GrabLines(argsWithoutProgram)
	cases,_ := strconv.Atoi(lines[0])
	index := 1
	for i := 1; i<= cases; i++ {
		naomisblocks := IntBlocks(strings.Split(lines[index+1], " "))
		kensblocks := IntBlocks(strings.Split(lines[index+2], " "))
		index += 3
		deceitfulturn := DeceitfulNaomi(naomisblocks, kensblocks)
		honestturn := HonestNaomi(naomisblocks, kensblocks)
		fmt.Printf("Case #%d: ", i)
		fmt.Println(deceitfulturn, honestturn)
	}
}
