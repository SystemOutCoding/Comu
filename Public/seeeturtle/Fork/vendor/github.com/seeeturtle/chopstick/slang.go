package chopstick

import "regexp"

var slangs = []string{
	"개",
	"걸레",
	"년",
	"놈",
	"느금마",
	"닥쳐",
	"등신",
	"또라이",
	"미친",
	"멍청",
	"병신",
	"새끼",
	"썅",
	"시발",
	"씨발",
	"씨팔",
	"씨발",
	"썖",
	"씹",
	"에미",
	"애미",
	"애비",
	"에비",
	"염병",
	"옘병",
	"좆",
	"좃",
	"좇",
	"지랄",
	"창",
	"호로",
	"후레",
	"호구",
	"후장",
}

// SlangSimilarity returns similarity of string with slang,
// a float number between 0 and 1.
func SlangSimilarity(str string, exceptions []string) float64 {
	var preprocessed string
	var similaritys []float64
	re := regexp.MustCompile("[^가-힣]")
	preprocessed = re.ReplaceAllString(str, "")
	for _, e := range exceptions {
		if Similarity([]rune(preprocessed), []rune(e)) > 0.4 {
			return 0
		}
	}
	for _, s := range slangs {
		similaritys = append(similaritys, Similarity([]rune(s), []rune(preprocessed)))
	}
	var max float64
	for _, s := range similaritys {
		if s >= max {
			max = s
		}
	}
	return max
}
