package multiplication_table

import (
	"fmt"
	"strconv"
)

func MultiplicationTable(count int) {
	// Подготовка формата вывода
	var l = strconv.Itoa(len(strconv.Itoa(count*count)) + 1)
	var format = "%" + l + "s"

	// Подготовка таблицы
	var res = [][]string{{" "}}
	// Заполнение первой строки
	for i := 1; i <= count; i++ {
		res[0] = append(res[0], strconv.Itoa(i))
	}

	// Заполнение таблицы
	for i := 1; i <= count; i++ {
		// Добавление строки
		res = append(res, []string{strconv.Itoa(i)})
		// Заполнение строки
		for k := 1; k <= count; k++ {
			res[i] = append(res[i], strconv.Itoa(k*(i)))
		}
	}

	// Вывод таблицы
	for _, r1 := range res {
		for _, r2 := range r1 {
			fmt.Printf(format, r2)
		}
		fmt.Print("\n")
	}
}
