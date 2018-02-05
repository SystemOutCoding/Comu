package handler

import (
	"regexp"
	"strings"
	"time"

	"github.com/seeeturtle/chopstick"
)

func parseContent(str string) (ok, delicious, similar, slang bool, scope Scope) {
	splitted := strings.Split(str, " ")
	re := regexp.MustCompile(`[\d]+월[\d]+일`)
	scope = NilScope{}
	exceptions := make([]string, 0)
	for _, s := range Scopes {
		exceptions = append(exceptions, s.Name())
	}
	for _, w := range splitted {
		if w == "" {
			continue
		}
		d := re.FindString(w)
		spaceRemoved := strings.Join(strings.Fields(w), "")
		var isNilScope bool
		switch scope.(type) {
		case NilScope:
			isNilScope = true
		}
		switch {
		case d != "":
			if isNilScope {
				t, _ := time.Parse("2006년1월2일", time.Now().In(loc).Format("2006년")+d)
				scope = Day{
					name: "날짜",
					date: t,
				}
			}
		case spaceRemoved == "ㅇㄴ":
			if isNilScope {
				scope = Today{
					name: "오늘",
				}
			}
			ok = true
		case spaceRemoved == "ㄴㅇ":
			if isNilScope {
				scope = Tomorrow{
					name: "내일",
				}
			}
			ok = true
		case spaceRemoved == "ㅁㄹ":
			if isNilScope {
				scope = Nextomorrow{
					name: "모레",
				}
			}
			ok = true
		case spaceRemoved == "ㄱㅍ":
			if isNilScope {
				scope = Threemorrow{
					name: "글피",
				}
			}
			ok = true
		case spaceRemoved == "ㅇㅂㅈ":
			if isNilScope {
				scope = ThisWeek{
					name: "이번주",
				}
			}
			ok = true
		case spaceRemoved == "ㄷㅇㅈ":
			if isNilScope {
				scope = NextWeek{
					name: "다음주",
				}
			}
			ok = true
		case strings.Contains(w, "오늘"):
			if isNilScope {
				scope = Today{
					name: "오늘",
				}
			}
		case strings.Contains(w, "내일"):
			if isNilScope {
				scope = Tomorrow{
					name: "내일",
				}
			}
		case strings.Contains(w, "모레"):
			if isNilScope {
				scope = Nextomorrow{
					name: "모레",
				}
			}
		case strings.Contains(w, "글피"):
			if isNilScope {
				scope = Threemorrow{
					name: "글피",
				}
			}
		case strings.Contains(w, "이번주"):
			if isNilScope {
				scope = ThisWeek{
					name: "이번주",
				}
			}
		case strings.Contains(w, "다다음주"):
			if isNilScope {
				scope = WeekAfterNext{
					name: "다다음주",
				}
			}
		case strings.Contains(w, "다음주"):
			if isNilScope {
				scope = NextWeek{
					name: "다음주",
				}
			}
		case strings.Contains(w, "이번달"):
			if isNilScope {
				scope = ThisMonth{
					name: "이번달",
				}
			}
		case strings.Contains(w, "다음달"):
			if isNilScope {
				scope = NextMonth{
					name: "다음달",
				}
			}
		case chopstick.Similarity([]rune("오늘"), []rune(w)) >= 0.5:
			if isNilScope {
				scope = Today{
					name: "오늘",
				}
				similar = true
			}
		case chopstick.Similarity([]rune("내일"), []rune(w)) >= 0.5:
			if isNilScope {
				scope = Tomorrow{
					name: "내일",
				}
				similar = true
			}
		case chopstick.Similarity([]rune("모레"), []rune(w)) >= 0.42:
			if isNilScope {
				scope = Nextomorrow{
					name: "모레",
				}
				similar = true
			}
		case chopstick.Similarity([]rune("글피"), []rune(w)) >= 0.5:
			if isNilScope {
				scope = Threemorrow{
					name: "글피",
				}
				similar = true
			}
		case chopstick.Similarity([]rune("이번주"), []rune(w)) >= 0.5:
			if isNilScope {
				scope = ThisWeek{
					name: "이번주",
				}
				similar = true
			}
		case chopstick.Similarity([]rune("다음주"), []rune(w)) >= 0.5:
			if isNilScope {
				scope = NextWeek{
					name: "다음주",
				}
				similar = true
			}
		case chopstick.Similarity([]rune("다다음주"), []rune(w)) >= 0.5:
			if isNilScope {
				scope = WeekAfterNext{
					name: "다다음주",
				}
				similar = true
			}
		case chopstick.Similarity([]rune("이번달"), []rune(w)) >= 0.5:
			if isNilScope {
				scope = ThisMonth{
					name: "이번달",
				}
				similar = true
			}
		case chopstick.Similarity([]rune("다음달"), []rune(w)) >= 0.5:
			if isNilScope {
				scope = NextMonth{
					name: "다음달",
				}
				similar = true
			}
		}
		switch {
		case strings.Contains(w, "급식"):
			ok = true
		case strings.Contains(w, "점심"):
			ok = true
		}
		if strings.Contains(w, "맛있") {
			delicious = true
		}
		if chopstick.SlangSimilarity(w, exceptions) >= 0.3 {
			slang = true
		}
	}
	if scope == (NilScope{}) {
		scope = UndefinedScope{}
	}
	return
}
