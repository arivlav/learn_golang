/*
Директор IT-корпорации оценивает эффективность работы сотрудников по различным показателям и критериям. Один из этих критериев сформулирован следующим образом: приступив
к некоторому заданию, сотрудник должен завершить его, не переключаясь на другие задания.

Чтобы проверить сотрудников на соответствие этому критерию, директор потребовал от каждого сотрудника отчет о том, какие задания он выполнял в последние n дней.
Отчет — это последовательность из n целых чисел a1,a2,…,an, где ai — идентификатор задания, которое сотрудник выполнял в i-й день.

Вам необходимо написать программу, проверяющую, соответствует ли сотрудник критерию по его отчету. Сотрудник соответствует этому критерию, если не существует такого задания
x, которое выполнялось с перерывом (т. е. в некоторый день i сотрудник выполнял задание x, в дни с i+1 по j−1 он занимался другими заданиями, а в день j сотрудник продолжил
выполнение задания x, при этом j>i+1). Иными словами, каждое задание, которое выполнял сотрудник, должно занимать один непрерывный отрезок дней.

Неполные решения этой задачи (например, недостаточно эффективные) могут быть оценены частичным баллом.
Входные данные

В первой строке задано одно целое число t (1≤t≤10) — количество наборов входных данных.

Каждый набор входных данных состоит из двух строк. В первой строке задано одно целое число n
(3≤n≤50000). Во второй строке заданы n целых чисел a1,a2,…,an (1≤ai≤n) — отчет сотрудника.
Выходные данные

Для каждого набора входных данных выведите ответ на отдельной строке. Если отчет соответствует критерию, выведите YES, иначе выведите NO.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	MaxSize = 30
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var setCount int
	fmt.Fscanln(in, &setCount)

	for z := 0; z < setCount; z++ {
		var n int
		fmt.Fscan(in, &n)
		days := make(map[int]int)
		prev, current := 0, 0
		result := "YES"
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &current)
			if prev != current {
				_, ok := days[current]
				if !ok {
					days[current] = 1
				} else {
					days[current]++
				}
			}
			if days[current] > 1 {
				result = "NO"
			}
			prev = current
		}
		fmt.Fprintln(out, result)
		fmt.Fscanln(in)
	}
	fmt.Fprintln(out)
}