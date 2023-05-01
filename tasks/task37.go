package tasks

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
Task 37.
Двухмерный мир состоит из блоков размером 1x1 метр
Остров игрока представляет собой набор столбцов различной высоты,
состоящих из блоков камня и окруженных морем. Над островом прошел
сильный дождь, который заполнил все низины, а не поместившаяся в них
вода стекла в море, не увеличив его уровень. По ландшафту острова
определите, сколько блоков воды после дождя осталось в низинах на острове

Input:
[1 0 0 0 0 1 0 0 0 0]
[1 0 0 1 0 1 0 0 0 0]
[1 1 0 1 1 1 0 0 1 0]
[1 1 0 1 1 1 0 0 1 0]
[1 1 1 1 1 1 1 1 1 1]

Output:
[1 2 2 2 2 1 0 0 0 0]
[1 2 2 1 2 1 0 0 0 0]
[1 1 2 1 1 1 2 2 1 0]
[1 1 2 1 1 1 2 2 1 0]
[1 1 1 1 1 1 1 1 1 1]
*/
func Task37() {
	island := readFile37()
	printIsland(island)
	island, totalWaterBlocks := fillWater(island)
	fmt.Println("Answer: 2-is water, total water blocks is", totalWaterBlocks)
	printIsland(island)
	writeToFile37(island)
}

func fillWater(island [][]int) ([][]int, int) {
	X := len(island[0])
	Y := len(island)
	var leftBorder bool
	var totalWaterBlocks int

	levelDropsIdxs := make([]int, 0, X)
	for y := 0; y < Y; y++ {
		for x := 0; x < X; x++ {
			if island[y][x] == 1 {
				if !leftBorder {
					leftBorder = true
				} else {
					totalWaterBlocks += len(levelDropsIdxs)
					fillWaterLevel(levelDropsIdxs, island[y])
					levelDropsIdxs = levelDropsIdxs[:0]
				}
			} else {
				if leftBorder {
					levelDropsIdxs = append(levelDropsIdxs, x)
				}
			}
		}

		levelDropsIdxs = levelDropsIdxs[:0]
		leftBorder = false
	}

	return island, totalWaterBlocks
}

func fillWaterLevel(levelDropsIdxs []int, level []int) {
	for i := range levelDropsIdxs {
		level[levelDropsIdxs[i]] = 2
	}
}

func readFile37() [][]int {
	fileIn, err := os.Open("tasks/inputs/input37.txt")
	if err != nil {
		log.Fatal("open err=", err)
	}

	island := make([][]int, 0)
	level := 0

	scanner := bufio.NewScanner(fileIn)
	for scanner.Scan() {
		row := scanner.Text()
		blocks := strings.Split(row, " ")

		island = append(island, []int{})
		for i := range blocks {
			n, err := strconv.Atoi(blocks[i])
			if err != nil {
				log.Fatal("input file: can not convert line to int", err)
			}

			island[level] = append(island[level], n)
		}
		level++
	}

	return island
}

func printIsland(island [][]int) {
	for i := range island {
		fmt.Println(island[i])
	}
}

func writeToFile37(island [][]int) {
	fileOut, err := os.OpenFile("tasks/outputs/output37.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal("open file err=", err)
	}

	dw := bufio.NewWriter(fileOut)
	for _, level := range island {
		strLevel := ""
		for i := range level {
			strLevel += strconv.Itoa(level[i]) + " "
		}

		_, err = dw.WriteString(strLevel + "\n")
		if err != nil {
			log.Fatal("write string err=", err)
		}
	}

	dw.Flush()
}
