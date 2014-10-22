package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func readFile(filename string) ([]int, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(content), "\r\n")
	data := make([]int, len(lines))

	for i, v := range lines {
		val, err := strconv.Atoi(v)
		if err == nil {
			data[i] = val
		}
	}

	return data, nil
}

func MergeAndCountSplitInv(left []int, right []int) ([]int, int64) {
	leftLength := len(left)
	length := leftLength + len(right)
	half := length / 2

	result := make([]int, length)

	var invCount int64 = 0
	i := 0
	j := 0

	for k := 0; k < length; k++ {
		if i >= half {
			result[k] = right[j]
			j++
		} else if j >= length-half {
			result[k] = left[i]
			i++
		} else {
			if left[i] <= right[j] {
				result[k] = left[i]
				i++
			} else {
				invCount += int64(leftLength - i)
				result[k] = right[j]
				j++
			}
		}
	}

	return result, invCount
}

func SortAndCount(data []int) ([]int, int64) {
	length := len(data)

	if length == 1 {
		return data, 0
	}

	half := length / 2

	left, leftCount := SortAndCount(data[:half])
	right, rightCount := SortAndCount(data[half:])
	result, splitCount := MergeAndCountSplitInv(left, right)

	return result, leftCount + rightCount + splitCount
}

func GetInversionsCount(data []int) int64 {
	_, count := SortAndCount(data)
	return count
}

func main() {
	data, err := readFile("./IntegerArray.txt")

	if err != nil {
		log.Fatal(err)
	}

	invCount := GetInversionsCount(data)
	fmt.Println("Inversions count: ", invCount)
}
