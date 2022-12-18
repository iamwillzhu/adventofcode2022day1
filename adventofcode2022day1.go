package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
)

func getCaloriesForAllElves(reader io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	caloriesForAllElves := make([]int, 0)

	totalCaloriesForOneElf := 0
	for scanner.Scan() {

		line := scanner.Text()

		if len(line) == 0 {
			caloriesForAllElves = append(caloriesForAllElves, totalCaloriesForOneElf)
			totalCaloriesForOneElf = 0
			continue
		}

		caloriesOfFood, err := strconv.Atoi(line)

		if err != nil {
			return nil, err
		}

		totalCaloriesForOneElf += caloriesOfFood
	}

	return caloriesForAllElves, nil
}

func main() {
	file, err := os.Open("/home/ec2-user/go/src/github.com/iamwillzhu/adventofcode2022day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	caloriesForAllElves, err := getCaloriesForAllElves(bufio.NewReader(file))

	if err != nil {
		log.Println(err)
		return
	}

	if len(caloriesForAllElves) == 0 {
		log.Println("The input list is empty")
		return
	}

	sort.Sort(sort.Reverse(sort.IntSlice(caloriesForAllElves)))
	maxCaloriesOfAllElves := caloriesForAllElves[0]

	sumOfTopThreeElvesCalories := caloriesForAllElves[0] + caloriesForAllElves[1] + caloriesForAllElves[2]

	fmt.Printf("The highest number of calories an elf has is %d\n", maxCaloriesOfAllElves)
	fmt.Printf("The sum of top three elves' calories is %d\n", sumOfTopThreeElvesCalories)

}
