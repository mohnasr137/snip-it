package models

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// ErrNoRecord is returned when a requested record does not exist.
var ErrNoRecord = errors.New("models: no matching record found")

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// SnippetModel wraps a PostgreSQL connection pool.
type SnippetModel struct {
	DB *pgxpool.Pool
}

// Insert adds a new snippet to the database.
func (m *SnippetModel) Insert(
	ctx context.Context,
	title string,
	content string,
	expires int,
) (int, error) {
	var id int

	query := `
		INSERT INTO snippets (title, content, created, expires)
		VALUES (
			$1,
			$2,
			CURRENT_TIMESTAMP,
			CURRENT_TIMESTAMP + ($3 * INTERVAL '1 day')
		)
		RETURNING id
	`

	err := m.DB.QueryRow(
		ctx,
		query,
		title,
		content,
		expires,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

// Get returns a specific snippet based on its ID.
func (m *SnippetModel) Get(
	ctx context.Context,
	id int,
) (*Snippet, error) {
	var snippet Snippet

	query := `
		SELECT id, title, content, created, expires
		FROM snippets
		WHERE id = $1
	`

	err := m.DB.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&snippet.ID,
		&snippet.Title,
		&snippet.Content,
		&snippet.Created,
		&snippet.Expires,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrNoRecord
	}

	if err != nil {
		return nil, err
	}

	return &snippet, nil
}

// Latest returns the 10 most recently created snippets.
func (m *SnippetModel) Latest(
	ctx context.Context,
) ([]*Snippet, error) {
	query := `
		SELECT id, title, content, created, expires
		FROM snippets
		ORDER BY created DESC
		LIMIT 10
	`

	rows, err := m.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var snippets []*Snippet

	for rows.Next() {
		var snippet Snippet

		err := rows.Scan(
			&snippet.ID,
			&snippet.Title,
			&snippet.Content,
			&snippet.Created,
			&snippet.Expires,
		)
		if err != nil {
			return nil, err
		}

		snippets = append(snippets, &snippet)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}
