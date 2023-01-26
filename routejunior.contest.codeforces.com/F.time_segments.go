/*
Вам задан набор отрезков времени. Каждый отрезок задан в формате HH:MM:SS-HH:MM:SS, то есть сначала заданы часы, минуты и секунды левой границы отрезка, а затем часы, минуты и
секунды правой границы.

Вам необходимо выполнить валидацию заданного набора отрезков времени. Иными словами, вам нужно проверить следующие условия:

	часы, минуты и секунды заданы корректно (то есть часы находятся в промежутке от 0 до 23, а минуты и секунды — в промежутке от 0 до 59);
	левая граница отрезка находится не позже его правой границы (но границы могут быть равны);
	никакая пара отрезков не пересекается (даже в граничных моментах времени).

Вам необходимо вывести YES, если заданный набор отрезков времени проходит валидацию, и NO в противном случае.

Вам необходимо ответить на t независимых наборов тестовых данных.

Неполные решения этой задачи (например, недостаточно эффективные) могут быть оценены частичным баллом.
Входные данные

Первая строка входных данных содержит одно целое число t (1≤t≤10) — количество наборов тестовых данных. Затем следуют t наборов.

Первая строка набора содержит одно целое число n (1≤n≤2⋅104) — количество отрезков времени. В следующих n строках следуют описания отрезков.

Описание отрезка времени задано в формате HH:MM:SS-HH:MM:SS, где HH, MM и SS — последовательности из двух цифр. Заметьте, что никаких пробелов в описании формата нет. Также
ни в одном описании нет пробелов в начале и конце строки.
Выходные данные

Для каждого набора тестовых данных выведите ответ — YES, если заданный набор отрезков времени проходит валидацию, и NO в противном случае.
Ответы выводите в порядке следования наборов во входных данных.
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
