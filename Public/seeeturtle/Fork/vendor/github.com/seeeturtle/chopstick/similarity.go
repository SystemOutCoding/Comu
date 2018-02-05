package chopstick

// Similarity returns similarity of hangul string, a and b.
// It calculate with Jaccard Similarity.
func Similarity(a, b []rune) float64 {
	intersection := make([]int, 0)
	union := make([]int, 0)
	var longer, shorter []int
	sliceA := make([]int, 0)
	for _, r := range a {
		sliceA = append(sliceA, separate(r)...)
	}
	sliceA = cutByTwo(sliceA)
	sliceB := make([]int, 0)
	for _, r := range b {
		sliceB = append(sliceB, separate(r)...)
	}
	sliceB = cutByTwo(sliceB)
	if len(sliceA) >= len(sliceB) {
		longer = sliceA
		shorter = sliceB
	} else {
		longer = sliceB
		shorter = sliceA
	}
	for _, i := range shorter {
		if !inIntSlice(i, union) {
			union = append(union, i)
		}
		for _, j := range longer {
			if !inIntSlice(j, union) {
				union = append(union, j)
			}
			if i == j && !inIntSlice(i, intersection) {
				intersection = append(intersection, i)
			}
		}
	}
	intersectionLen := len(intersection)
	unionLen := len(union)
	return float64(intersectionLen) / float64(unionLen)
}

func inIntSlice(a int, b []int) bool {
	for _, i := range b {
		if a == i {
			return true
		}
	}
	return false
}

func separate(a rune) []int {
	var slice []int
	var hangulCodes = []string{"ㄱ", "ㄲ", "ㄴ", "ㄷ", "ㄸ", "ㄹ", "ㅁ", "ㅂ",
		"ㅃ", "ㅅ", "ㅆ", "ㅇ", "ㅈ", "ㅉ", "ㅊ", "ㅋ", "ㅌ", "ㅍ", "ㅎ"}
	if 12592 < int(a) && int(a) < 12687 {
		var a, b int = -1, -1
		for i, c := range hangulCodes {
			if string(a) == c {
				if inIntSlice(i, []int{1, 4, 8, 10, 13}) {
					a = i - 1
					b = i - 1
				} else {
					a = i
				}
				break
			}
		}
		if a != -1 {
			if b != -1 {
				return append(slice, a, b)
			}
			return append(slice, a)
		}
		return []int{}
	}
	code := int(a) - 44032
	jongSeong := code % 28
	jungSeong := ((code - jongSeong) / 28) % 21
	choSeong := ((code-jongSeong)/28 - jungSeong) / 21
	if inIntSlice(choSeong, []int{1, 4, 8, 10, 13}) {
		slice = append(slice, choSeong-1, choSeong-1)
	} else {
		slice = append(slice, choSeong)
	}
	slice = append(slice, jungSeong)
	if jongSeong != 0 {
		slice = append(slice, jongSeong)
	}

	return slice
}

func cutByTwo(a []int) []int {
	var result []int
	if len(a) <= 0 {
		return result
	}
	result = append(result, (a[0]+3)*30)
	for index, i := range a {
		if index == len(a)-1 {
			result = append(result, 100*(i+3))
			continue
		}
		result = append(result, (i+3)*a[index+1])
	}
	return result
}
