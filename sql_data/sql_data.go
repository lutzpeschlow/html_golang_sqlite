package sql_data

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"github.com/lutzpeschlow/html_golang_sqlite/objects"
)

func PrepareSQLData(in *objects.InputData, sql *objects.SqlData) int {
	status := 0
	sql.Date = in.Date
	sql.Option2 = in.Option2
	sql.Option3 = in.Option3
	sql.Option4 = in.Option4
	sql.Option5 = in.Option5

	sql.Score1 = in.Scores[0]
	sql.Score2 = in.Scores[1]
	sql.Score3 = in.Scores[2]
	sql.Score4 = in.Scores[3]
	sql.Score5 = in.Scores[4]
	sql.Score6 = in.Scores[5]
	sql.Score7 = in.Scores[6]
	sql.Score8 = in.Scores[7]
	sql.Score9 = in.Scores[8]
	sql.Score10 = in.Scores[9]
	sql.Score11 = in.Scores[10]
	sql.Score12 = in.Scores[11]
	sql.Score13 = in.Scores[12]
	sql.Score14 = in.Scores[13]
	sql.Score15 = in.Scores[14]
	sql.Score16 = in.Scores[15]
	sql.Score17 = in.Scores[16]
	sql.Score18 = in.Scores[17]
	sql.CreatedAt = time.Now()
	return status
}

func CreateDB(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

func InitSchema(db *sql.DB) error {
	stmt := `
	CREATE TABLE IF NOT EXISTS rounds (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		date TEXT,
		option2 TEXT,
		option3 TEXT,
		option4 TEXT,
		option5 TEXT,
		score1 INTEGER,
		score2 INTEGER,
		score3 INTEGER,
		score4 INTEGER,
		score5 INTEGER,
		score6 INTEGER,
		score7 INTEGER,
		score8 INTEGER,
		score9 INTEGER,
		score10 INTEGER,
		score11 INTEGER,
		score12 INTEGER,
		score13 INTEGER,
		score14 INTEGER,
		score15 INTEGER,
		score16 INTEGER,
		score17 INTEGER,
		score18 INTEGER,
		created_at TEXT
	);`
	_, err := db.Exec(stmt)
	return err
}

func InsertSQLData(db *sql.DB, s objects.SqlData) (int64, error) {
	query := `
		INSERT INTO rounds (
			date, option2, option3, option4, option5,
			score1, score2, score3, score4, score5, score6, score7, score8, score9,
			score10, score11, score12, score13, score14, score15, score16, score17, score18,
			created_at
		) VALUES (?, ?, ?, ?, ?,
		          ?, ?, ?, ?, ?, ?, ?, ?, ?,
		          ?, ?, ?, ?, ?, ?, ?, ?, ?,
		          ?)
	`

	res, err := db.Exec(
		query,
		s.Date, s.Option2, s.Option3, s.Option4, s.Option5,
		s.Score1, s.Score2, s.Score3, s.Score4, s.Score5, s.Score6, s.Score7, s.Score8, s.Score9,
		s.Score10, s.Score11, s.Score12, s.Score13, s.Score14, s.Score15, s.Score16, s.Score17, s.Score18,
		s.CreatedAt,
	)
	if err != nil {
		return 0, fmt.Errorf("insert sql data: %w", err)
	}
	return res.LastInsertId()
}
