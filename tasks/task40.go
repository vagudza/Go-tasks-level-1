package tasks

import (
	"fmt"
	"strconv"
)

/*
Дано два числа без ведущих нулей
Необходимо проверить, можно ли получить первое из второго перестановкой цифр
*/
func Task40() {
	X := 2021539
	Y := 2012395

	fmt.Println("isEqual=", isEqual(X, Y))
	fmt.Println("isEqual2=", isEqual2(X, Y))
}

func isEqual(X, Y int) bool {
	digestX := numToDigest(X)
	digestY := numToDigest(Y)

	if len(digestX) != len(digestY) {
		return false
	}

	for kx, vx := range digestX {
		if vy, ok := digestY[kx]; !ok || (ok && vx != vy) {
			return false
		}
	}

	return true
}

func numToDigest(X int) map[uint8]int {
	Xstr := strconv.Itoa(X)
	res := make(map[uint8]int, len(Xstr))

	for i := range Xstr {
		res[Xstr[i]]++
	}

	return res
}

func isEqual2(X, Y int) bool {
	digestX := numToDigestV2(X)
	digestY := numToDigestV2(Y)

	for i := 0; i < 10; i++ {
		if digestX[i] != digestY[i] {
			return false
		}
	}

	return true
}

func numToDigestV2(x int) [10]int {
	var res [10]int
	for x != 0 {
		res[x%10]++
		x /= 10
	}

	return res
}
