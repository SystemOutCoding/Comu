package handler

import (
	"fmt"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/jinzhu/now"
	"github.com/seeeturtle/Fork/app/model"
)

type Scope interface {
	Beginning() time.Time
	End() time.Time
	Name() string
	FoodMessage([]model.Lunch) string
	DeliciousFoodMessage([]model.Lunch) string
}

type NilScope struct{}

type UndefinedScope struct{}

type Day struct {
	name string
	date time.Time
}

type Today struct {
	name string
}

type Tomorrow struct {
	name string
}

type Nextomorrow struct {
	name string
}

type Threemorrow struct {
	name string
}

type ThisWeek struct {
	name string
}

type WeekAfterNext struct {
	name string
}

type NextWeek struct {
	name string
}

type ThisMonth struct {
	name string
}

type NextMonth struct {
	name string
}

var Scopes = []Scope{
	Today{name: "오늘"},
	Tomorrow{name: "내일"},
	Nextomorrow{name: "모레"},
	Threemorrow{name: "글피"},
	WeekAfterNext{name: "다다음주"},
	NextWeek{name: "다음주"},
	ThisWeek{name: "이번주"},
	ThisMonth{name: "이번달"},
	NextMonth{name: "다음달"},
}

func (n NilScope) Beginning() time.Time {
	return time.Time{}
}

func (n NilScope) End() time.Time {
	return time.Time{}
}

func (n NilScope) Name() string {
	return ""
}

func (n NilScope) FoodMessage(lunches []model.Lunch) string {
	return ""
}

func (n NilScope) DeliciousFoodMessage(lunches []model.Lunch) string {
	return ""
}

func (u UndefinedScope) Beginning() time.Time {
	return time.Time{}
}

func (u UndefinedScope) End() time.Time {
	return time.Time{}
}

func (u UndefinedScope) Name() string {
	return ""
}

func (u UndefinedScope) FoodMessage(lunches []model.Lunch) string {
	return ""
}

func (u UndefinedScope) DeliciousFoodMessage(lunches []model.Lunch) string {
	return ""
}
func (d Day) Beginning() time.Time {
	return d.date
}

func (d Day) End() time.Time {
	return d.date
}

func (d Day) Name() string {
	return d.name
}

func (d Day) FoodMessage(lunches []model.Lunch) string {
	f := JoinWithComma(lunches)
	return f + "나온다"
}

func (d Day) DeliciousFoodMessage(lunches []model.Lunch) string {
	f := JoinWithComma(lunches)
	return f + "나온다!!!"
}

func (t Today) Beginning() time.Time {
	return time.Now().In(loc)
}

func (t Today) End() time.Time {
	return time.Now().In(loc)
}

func (t Today) Name() string {
	return t.name
}

func (t Today) FoodMessage(lunches []model.Lunch) string {
	f := JoinWithComma(lunches)
	return f + "나온다"
}

func (t Today) DeliciousFoodMessage(lunches []model.Lunch) string {
	f := JoinWithComma(lunches)
	return f + "나온다!!!!"
}

func (to Tomorrow) Beginning() time.Time {
	return time.Now().In(loc).AddDate(0, 0, 1)
}

func (to Tomorrow) End() time.Time {
	return time.Now().In(loc).AddDate(0, 0, 1)
}

func (to Tomorrow) Name() string {
	return to.name
}

func (to Tomorrow) FoodMessage(lunches []model.Lunch) string {
	f := JoinWithComma(lunches)
	return f + "나온다"
}

func (to Tomorrow) DeliciousFoodMessage(lunches []model.Lunch) string {
	f := JoinWithComma(lunches)
	return f + "나온다!!!!"
}

func (nt Nextomorrow) Beginning() time.Time {
	return time.Now().In(loc).AddDate(0, 0, 2)
}

func (nt Nextomorrow) End() time.Time {
	return time.Now().In(loc).AddDate(0, 0, 2)
}

func (nt Nextomorrow) Name() string {
	return nt.name
}

func (nt Nextomorrow) FoodMessage(lunches []model.Lunch) string {
	f := JoinWithComma(lunches)
	return f + "나온다"
}

func (nt Nextomorrow) DeliciousFoodMessage(lunches []model.Lunch) string {
	f := JoinWithComma(lunches)
	return f + "나온다!!!"
}

func (tm Threemorrow) Beginning() time.Time {
	return time.Now().In(loc).AddDate(0, 0, 3)
}

func (tm Threemorrow) End() time.Time {
	return time.Now().In(loc).AddDate(0, 0, 3)
}

func (tm Threemorrow) Name() string {
	return tm.name
}

func (tm Threemorrow) FoodMessage(lunches []model.Lunch) string {
	f := JoinWithComma(lunches)
	return f + "나온다"
}

func (tm Threemorrow) DeliciousFoodMessage(lunches []model.Lunch) string {
	f := JoinWithComma(lunches)
	return f + "나온다!!!"
}

func (tw ThisWeek) Beginning() time.Time {
	return now.New(time.Now().In(loc)).BeginningOfWeek()
}

func (tw ThisWeek) End() time.Time {
	return now.New(time.Now().In(loc)).EndOfWeek()
}

func (tw ThisWeek) Name() string {
	return tw.name
}

func (tw ThisWeek) FoodMessage(lunches []model.Lunch) string {
	f := JoinWithComma(lunches)
	return f + "나온다. 이번주에는 그럭저럭 하네."
}

func (tw ThisWeek) DeliciousFoodMessage(lunches []model.Lunch) string {
	f := JoinWithComma(lunches)
	return f + "나온다!!!! 이번주에 맛있는게 많다!"
}

func (wn WeekAfterNext) Beginning() time.Time {
	n := now.New(time.Now().In(loc).AddDate(0, 0, 14))
	return n.BeginningOfWeek()
}

func (wn WeekAfterNext) End() time.Time {
	n := now.New(time.Now().In(loc).AddDate(0, 0, 14))
	return n.EndOfWeek()
}

func (wn WeekAfterNext) Name() string {
	return wn.name
}

func (wn WeekAfterNext) FoodMessage(lunches []model.Lunch) string {
	f := JoinWithComma(lunches)
	return f + "나온다"
}

func (wn WeekAfterNext) DeliciousFoodMessage(lunches []model.Lunch) string {
	f := JoinWithComma(lunches)
	return f + "나온다!!!"
}

func (nw NextWeek) Beginning() time.Time {
	n := now.New(time.Now().In(loc).AddDate(0, 0, 7))
	return n.BeginningOfWeek()
}

func (nw NextWeek) End() time.Time {
	n := now.New(time.Now().In(loc).AddDate(0, 0, 7))
	return n.EndOfWeek()
}

func (nw NextWeek) Name() string {
	return nw.name
}

func (nw NextWeek) FoodMessage(lunches []model.Lunch) string {
	f := JoinWithComma(lunches)
	return f + "나온다. 괜찮은데?"
}

func (nm NextWeek) DeliciousFoodMessage(lunches []model.Lunch) string {
	f := JoinWithComma(lunches)
	return f + "나온다!!!!"
}

func (tm ThisMonth) Beginning() time.Time {
	return now.New(time.Now().In(loc)).BeginningOfMonth()
}

func (tm ThisMonth) End() time.Time {
	return now.New(time.Now().In(loc)).EndOfWeek()
}

func (tm ThisMonth) Name() string {
	return tm.name
}

func (tm ThisMonth) FoodMessage(lunches []model.Lunch) string {
	f := JoinWithComma(lunches)
	return f + "나온다.\n이번달 급식임"
}

func (tm ThisMonth) DeliciousFoodMessage(lunches []model.Lunch) string {
	f := JoinWithComma(lunches)
	return f + "나온다.\n기억해둬"
}

func (nm NextMonth) Beginning() time.Time {
	n := now.New(time.Now().In(loc).AddDate(0, 1, 0))
	return n.BeginningOfMonth()
}

func (nm NextMonth) End() time.Time {
	n := now.New(time.Now().In(loc).AddDate(0, 1, 0))
	return n.EndOfMonth()
}

func (nm NextMonth) Name() string {
	return nm.name
}

func (nm NextMonth) FoodMessage(lunches []model.Lunch) string {
	f := JoinWithComma(lunches)
	return f + "나온다.\n다음달 급식임."
}

func (nm NextMonth) DeliciousFoodMessage(lunches []model.Lunch) string {
	f := JoinWithComma(lunches)
	return f + "나온다.\n다음달 급식도 괜찮은 듯."
}

func JoinWithComma(lunches []model.Lunch) string {
	var str string
	if len(lunches) > 1 {
		for i, lunch := range lunches {
			names := []string{}
			for _, food := range lunch.Foods {
				names = append(names, food.Name)
			}
			dateStr := getDatesStr(lunch.Date)
			str += dateStr + strings.Join(names, ", ") + getPostposition(names[len(names)-1])
			if i != len(lunches)-1 {
				str += ",\n"
			} else {
				str += " "
			}
		}
	} else if len(lunches) == 1 {
		names := []string{}
		for _, food := range lunches[0].Foods {
			names = append(names, food.Name)
		}
		dateStr := getDateStr(lunches[0].Date)
		str = dateStr + strings.Join(names, ", ") + getPostposition(names[len(names)-1])
	}
	return str
}

func getDatesStr(date time.Time) string {
	var weekDays [7]string = [7]string{
		"일요일",
		"월요일",
		"화요일",
		"수요일",
		"목요일",
		"금요일",
		"토요일",
	}
	dateTime := date
	n := time.Now().In(loc)
	diffWeeks := int(now.New(dateTime).BeginningOfWeek().Sub(now.New(n).BeginningOfWeek())) / (int(time.Hour) * 24 * 7)
	weekDay := weekDays[int(dateTime.Weekday())]
	switch {
	case diffWeeks == -3:
		return "3주전 " + weekDay + "은 "
	case diffWeeks == -2:
		return "저저번주 " + weekDay + "은 "
	case diffWeeks == -1:
		return "저번주 " + weekDay + "은 "
	case diffWeeks == 0:
		return "이번주 " + weekDay + "은 "
	case diffWeeks == 1:
		return "다음주 " + weekDay + "은 "
	case diffWeeks == 2:
		return "다다음주 " + weekDay + "은 "
	case diffWeeks == 3:
		return "3주후 " + weekDay + "은 "
	default:
		return fmt.Sprintf("%d월 %d일은 ", dateTime.Month(), dateTime.Day())
	}
}

func getDateStr(date time.Time) string {
	dateTime := roundTime(date)
	now := roundTime(time.Now().In(loc))
	duration := dateTime.Sub(now)
	diffDays := int(duration.Hours() / 24)
	switch diffDays {
	case 0:
		return "오늘은 "
	case 1:
		return "내일은 "
	case 2:
		return "모레는 "
	case 3:
		return "글피는 "
	default:
		return fmt.Sprintf("%d월 %d일은 ", dateTime.Month(), dateTime.Day())
	}
}

func roundTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func getPostposition(str string) string {
	defaultStr := "가"
	s := []rune(str)
	lastCharacter := string(s[len(s)-1])
	r, _ := utf8.DecodeRuneInString(lastCharacter)
	jongSeongCode := (int(r) - 44032) % 28
	if jongSeongCode != 0 {
		defaultStr = "이"
	}
	return defaultStr
}
