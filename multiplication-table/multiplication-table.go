package multiplication_table

import (
	"fmt"
	"strconv"
)

func MultiplicationTable(count int) {
	var l = strconv.Itoa(len(strconv.Itoa(count*count)) + 1)
	var format = "%" + l + "s"

	var res = [][]string{{" "}}
	for i := 1; i <= count; i++ {
		res[0] = append(res[0], strconv.Itoa(i))
	}

	for i := 1; i <= count; i++ {
		res = append(res, []string{strconv.Itoa(i)})
		for k := 1; k <= count; k++ {
			res[i] = append(res[i], strconv.Itoa(k*(i)))
		}
	}

	for _, r1 := range res {
		for _, r2 := range r1 {
			fmt.Printf(format, r2)
		}
		fmt.Print("\n")
	}
}
