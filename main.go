package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'queensAttack' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER k
 *  3. INTEGER r_q
 *  4. INTEGER c_q
 *  5. 2D_INTEGER_ARRAY obstacles
 */

func queensAttack(n int32, k int32, r_q int32, c_q int32, obstacles [][]int32) int32 {
	// Write your code here
	obstMap := obstaclesToMap(k, obstacles)
	directions := []string{"N", "NE", "E", "SE", "S", "SW", "W", "NW"}
	allMoves := 0
	orig_r, orig_c := r_q, c_q
	for _, dir := range directions {
		dirMoves := 0
		r_q, c_q := orig_r, orig_c
		for {
			r_q, c_q = moveQueen(r_q, c_q, dir)

			if isObstacleThere(r_q, c_q, obstMap) {
				break
			}

			if !((r_q >= 1 && r_q <= n) && (c_q >= 1 && c_q <= n)) {
				break
			}
			dirMoves += 1
		}
		fmt.Printf("Moves for %s: %d\n", dir, dirMoves)
		allMoves += dirMoves
	}
	fmt.Println(allMoves)
	return int32(allMoves)
}

func obstaclesToMap(k int32, obstacles [][]int32) map[int32][]int32 {
	o := make(map[int32][]int32, k)
	for _, rc := range obstacles {
		row, column := rc[0], rc[1]
		// fmt.Println(row, column)
		if _, ok := o[row]; ok {
			o[row] = append(o[row], column)
		} else {
			o[row] = []int32{column}
		}
	}
	fmt.Println(o)
	return o
}

func moveQueen(r_q, c_q int32, direction string) (int32, int32) {
	switch direction {
	case "N":
		r_q += 1
	case "NE":
		r_q += 1
		c_q += 1
	case "E":
		c_q += 1
	case "SE":
		r_q -= 1
		c_q += 1
	case "S":
		r_q -= 1
	case "SW":
		r_q -= 1
		c_q -= 1
	case "W":
		c_q -= 1
	case "NW":
		r_q += 1
		c_q -= 1
	}
	return r_q, c_q
}

func isObstacleThere(r_q, c_q int32, obstacles map[int32][]int32) bool {
	if cols, ok := obstacles[r_q]; ok {
		for _, c := range cols {
			if c_q == c {
				return true
			}
		}
	}
	return false
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	secondMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	r_qTemp, err := strconv.ParseInt(secondMultipleInput[0], 10, 64)
	checkError(err)
	r_q := int32(r_qTemp)

	c_qTemp, err := strconv.ParseInt(secondMultipleInput[1], 10, 64)
	checkError(err)
	c_q := int32(c_qTemp)

	var obstacles [][]int32
	for i := 0; i < int(k); i++ {
		obstaclesRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var obstaclesRow []int32
		for _, obstaclesRowItem := range obstaclesRowTemp {
			obstaclesItemTemp, err := strconv.ParseInt(obstaclesRowItem, 10, 64)
			checkError(err)
			obstaclesItem := int32(obstaclesItemTemp)
			obstaclesRow = append(obstaclesRow, obstaclesItem)
		}

		if len(obstaclesRow) != 2 {
			panic("Bad input")
		}

		obstacles = append(obstacles, obstaclesRow)
	}

	result := queensAttack(n, k, r_q, c_q, obstacles)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
