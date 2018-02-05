package model

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

var (
	Lunches *LunchesModel
)

type Food struct {
	Name      string
	delicious bool
}

type Modeler interface {
	Get(...interface{}) (interface{}, error)
	GetDelicious(...interface{}) (interface{}, error)
	New(...interface{}) interface{}
}

type Lunch struct {
	Date  time.Time
	Foods []Food
}

type LunchesModel struct {
	Value []Lunch
	DB    *sql.DB
}

func (l *LunchesModel) New(db *sql.DB) *LunchesModel {
	lunches := &LunchesModel{
		Value: []Lunch{},
		DB:    db,
	}
	return lunches
}

func (l *LunchesModel) Get(startTime, endTime time.Time) ([]Lunch, error) {
	var lunches LunchesModel

	if startTime.Format("20060102") != endTime.Format("20060102") {
		for d := startTime; d.Before(endTime) || d.Equal(endTime); d = d.AddDate(0, 0, 1) {
			lunch, err := l.getADay(d)
			if err != nil {
				continue
			}
			lunches.Value = append(lunches.Value, lunch)
		}
		return lunches.Value, nil
	}
	lunch, err := l.getADay(startTime)
	if err != nil {
		return lunches.Value, err
	}
	lunches.Value = append(lunches.Value, lunch)
	return lunches.Value, nil
}

func (l *LunchesModel) GetDelicious(startTime, endTime time.Time) ([]Lunch, error) {
	var deliciousLunches []Lunch

	lunches, err := l.Get(startTime, endTime)

	if err != nil {
		return deliciousLunches, err
	}

	for _, lunch := range lunches {
		var deliciousLunch Lunch
		for _, food := range lunch.Foods {
			if food.delicious {
				deliciousLunch.Foods = append(deliciousLunch.Foods, food)
			}
		}
		if len(deliciousLunch.Foods) == 0 {
			continue
		}
		deliciousLunch.Date = lunch.Date
		deliciousLunches = append(deliciousLunches, deliciousLunch)
	}
	return deliciousLunches, nil
}

func (l *LunchesModel) getADay(d time.Time) (Lunch, error) {
	var (
		lunch Lunch
		foods []Food
	)
	var lunchID int
	err := l.DB.QueryRow("SELECT lunch_id FROM lunches WHERE date=$1", d.Format("20060102")).Scan(&lunchID)

	if err != nil {
		return lunch, err
	}

	query := `
	SELECT f.food_name, f.delicious
	FROM foods as f
	NATURAL JOIN
	(SELECT food_id FROM lunches_foods
	WHERE lunch_id=$1) as food_ids
	`
	rows, err := l.DB.Query(query, lunchID)

	if err != nil {
		return lunch, err
	}

	defer rows.Close()

	for rows.Next() {
		var (
			food      Food
			foodName  string
			delicious bool
		)

		err := rows.Scan(&foodName, &delicious)

		if err != nil {
			return lunch, err
		}

		food.Name = foodName
		food.delicious = delicious
		foods = append(foods, food)
	}

	err = rows.Err()
	if err != nil {
		return lunch, err
	}

	lunch.Date = d
	lunch.Foods = foods

	return lunch, nil
}
