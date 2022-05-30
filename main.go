package main

import (
	"fmt"

	filterint64 "github.com/rgalicia0729/filter-int64-go"
)

func main() {
	nums := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	result := filterint64.Where(nums, func(value int64) bool {
		return value <= 5
	})

	fmt.Println(result)
}
