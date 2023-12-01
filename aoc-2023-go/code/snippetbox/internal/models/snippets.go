package models

import (
	"database/sql"
	"errors"
	"time"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	query := `
		INSERT INTO snippets (title, content, created, expires) 
		values ($1, $2, now(), now() + ($3 || ' day')::interval)
		returning id
		`

	var ID int
	err := m.DB.QueryRow(query, title, content, expires).Scan(&ID)
	if err != nil {
		return 0, err
	}

	return ID, nil
}

func (m *SnippetModel) Get(ID int) (Snippet, error) {
	query := `select id, title, content, created, expires from snippets where expires > now() and id = $1`

	var snippet Snippet
	err := m.DB.QueryRow(query, ID).Scan(&snippet.ID, &snippet.Title, &snippet.Content, &snippet.Created, &snippet.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Snippet{}, ErrNoRecord
		} else {
			return Snippet{}, err
		}
	}

	return snippet, nil
}

func (m *SnippetModel) Latest() ([]Snippet, error) {
	query := `select id, title, content, created, expires from snippets where expires > now() order by id LIMIT 10`
	rows, err := m.DB.Query(query)
	if err != nil {
		return nil, err
	}

	var snippets []Snippet
	for rows.Next() {
		var snippet Snippet
		err := rows.Scan(&snippet.ID, &snippet.Title, &snippet.Content, &snippet.Created, &snippet.Expires)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, snippet)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}
