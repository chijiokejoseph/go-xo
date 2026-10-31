package files

import "errors"

type Record struct {
	Name      string `json:"name"`
	Losses    int    `json:"losses"`
	Draws     int    `json:"draws"`
	Wins      int    `json:"wins"`
	Rounds    int    `json:"rounds"`
	Score     int    `json:"score"`
	HighScore int    `json:"high_score"`
}

var ErrNotFound = errors.New("record not found in database")

type Database struct {
	records []Record
}

func (db *Database) Add(r Record) {
	// check if current record matches existing record
	i, found := db.Search(r.Name)

	if !found {
		// if not found, append new record to the list of records
		db.records = append(db.records, r)
		return
	}

	db.records[i].Draws += r.Draws
	db.records[i].Wins += r.Wins
	db.records[i].Losses += r.Losses
	db.records[i].Rounds += r.Rounds
	db.records[i].Score = r.Wins - (r.Draws + r.Losses)

	if r.Score > db.records[i].HighScore {
		db.records[i].HighScore = r.Score
	}
}

func (db *Database) Get(name string) (Record, error) {
	i, found := db.Search(name)

	if !found {
		return Record{}, ErrNotFound
	}

	return db.records[i], nil

}

func (db *Database) Search(name string) (int, bool) {
	for i, record := range db.records {
		if record.Name == name {
			return i, true
		}
	}

	return 0, false
}

func (db *Database) MaxHighScore() int {
	// if db is empty, high score defaults to 0
	if len(db.records) == 0 {
		return 0
	}
	// set `max` to first record highscore
	max := db.records[0].HighScore

	// update `max` to the maximum highscore
	for _, r := range db.records {
		if r.HighScore > max {
			max = r.HighScore
		}
	}

	return max
}
