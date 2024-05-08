package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter text: ")
		raw, _ := reader.ReadString('\n')
		input := raw[:len(raw)-2]
		if input == "q" {
			break
		}
		decoded, err := decode(input)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		fmt.Println("Decoded text:", decoded)
	}

}

func decode(input string) (string, error) {
	decodeArray := make([]int, len(input)+1)
	min := 0
	for i, char := range input {
		switch char {
		case 'L':
			decodeArray[i+1] = decodeArray[i] - 1
		case 'R':
			decodeArray[i+1] = decodeArray[i] + 1
		case '=':
			decodeArray[i+1] = decodeArray[i]
		default:
			return "", errors.New("invalid character")
		}
		if decodeArray[i+1] < min {
			min = decodeArray[i+1]
		}
	}
	if min < 0 {
		for i := range decodeArray {
			decodeArray[i] -= min
		}
	}
	minimizeDecodeArray := minimize(decodeArray)

	decoded := ""
	for _, number := range minimizeDecodeArray {
		decoded += fmt.Sprintf("%v", number)
	}
	return decoded, nil
}

func minimize(decodeArray []int) []int {
	arraySize := len(decodeArray)
	result := make([]int, arraySize)
	copy(result, decodeArray)

	sectionMin := result[0]
	sectionStartIndex := 0
	// skip the first element
	for i := 1; i < arraySize-1; i++ {
		if sectionMin == 0 {
			break
		}
		if result[i-1] >= result[i] {
			sectionMin = result[i]
		} else {
			if sectionMin > 0 {
				for j := sectionStartIndex; j <= i-1; j++ {
					result[j] -= sectionMin
				}
			}
			sectionStartIndex = i
		}
	}

	sectionMin = result[arraySize-1]
	sectionStartIndex = arraySize - 1
	// skip the last element
	for i := arraySize - 2; i >= 0; i-- {
		if sectionMin == 0 {
			break
		}
		if result[i+1] >= result[i] {
			sectionMin = result[i]
		} else {
			if sectionMin > 0 {
				for j := i + 1; j <= sectionStartIndex; j++ {
					result[j] -= sectionMin
				}
			}
			sectionStartIndex = i
		}
	}
	return result
}
