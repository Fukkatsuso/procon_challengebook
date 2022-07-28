package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

const (
	initialBufSize = 100000
	maxBufSize     = 10000000
)

var (
	sc = bufio.NewScanner(os.Stdin)
	wt = bufio.NewWriter(os.Stdout)
)

func gets() string {
	sc.Scan()
	return sc.Text()
}

func getInt() int {
	i, _ := strconv.Atoi(gets())
	return i
}

func getInt2() (int, int) {
	return getInt(), getInt()
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func max(nums ...int) int {
	ret := nums[0]
	for _, v := range nums {
		if v > ret {
			ret = v
		}
	}
	return ret
}

// 区間[a,b)の素数判定
// iが素数 <=> isPrime[i-a] == true
func SegmentSieve(a, b int) []bool {
	sqrtB := math.Sqrt(float64(b))
	isPrimeSmall := make([]bool, int(math.Ceil(sqrtB)))
	isPrime := make([]bool, b-a)

	for i := 2; i*i < b; i++ {
		isPrimeSmall[i] = true
	}
	for i := 0; i < b-a; i++ {
		isPrime[i] = true
	}

	for i := 2; i*i < b; i++ {
		if isPrimeSmall[i] {
			for j := 2 * i; j*j < b; j += i {
				isPrimeSmall[j] = false
			}
			for j := max(2, (a+i-1)/i) * i; j < b; j += i {
				isPrime[j-a] = false
			}
		}
	}

	return isPrime
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	a, b := getInt2()

	isPrime := SegmentSieve(a, b)
	ans := 0
	for i := a; i < b; i++ {
		if isPrime[i-a] {
			ans++
		}
	}

	puts(ans)
}

// example1:
// 22 37
// 3

// example2:
// 22801763489 22801787297
// 1000
