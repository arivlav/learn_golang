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
	"strings"
	"time"
)

const (
	LAYUOT_TIME = "15:04:05"
	YES         = "YES"
	NO          = "NO"
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
		result := YES
		timeIntervals := make([][]int, n, n)
		for i := 0; i < n; i++ {
			var str string
			fmt.Fscan(in, &str)
			strSplit := strings.Split(str, "-")
			start, errStart := time.Parse(LAYUOT_TIME, strSplit[0])
			end, errEnd := time.Parse(LAYUOT_TIME, strSplit[1])
			if errStart == nil && errEnd == nil && (end.Equal(start) || end.After(start)) {
				timeIntervals[i] = []int{int(start.Unix()), int(end.Unix())}
				if i > 0 {
					for j := i - 1; j >= 0; j-- {
						if len(timeIntervals[j]) > 0 {
							fmt.Fprintln(out, timeIntervals[i][0], timeIntervals[i][1], timeIntervals[j][0], timeIntervals[j][1])
							if timeIntervals[i][0] < timeIntervals[j][1] {
								if timeIntervals[i][1] >= timeIntervals[j][1] {
									result = NO
								}

							} else if timeIntervals[i][0] == timeIntervals[j][1] {
								result = NO
							}
						}
					}
				}
			} else {
				result = NO
			}
		}
		fmt.Fprintln(out, result)
	}
	fmt.Fprintln(out)
}
