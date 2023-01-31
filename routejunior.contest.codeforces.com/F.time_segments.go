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
	DAYSECONDS  = 86400
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
		var day [DAYSECONDS]byte
		for i := 0; i < n; i++ {
			var str string
			fmt.Fscan(in, &str)
			strSplit := strings.Split(str, "-")
			start, errStart := time.Parse(LAYUOT_TIME, strSplit[0])
			end, errEnd := time.Parse(LAYUOT_TIME, strSplit[1])
			if errStart == nil && errEnd == nil && (end.Equal(start) || end.After(start)) {
				hourStart, minStart, secStart := start.Clock()
				hourEnd, minEnd, secEnd := end.Clock()
				startTime := hourStart*3600 + minStart*60 + secStart
				endTime := hourEnd*3600 + minEnd*60 + secEnd
				for j := startTime; j <= endTime; j++ {
					if day[j] == 0 {
						day[j] = 1
					} else {
						result = NO
					}
					if result == NO {
						break
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
