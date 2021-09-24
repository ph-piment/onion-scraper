package dao

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"github.com/jmoiron/sqlx"
	"time"
)

// News represents a row from 'public.news'.
type News struct {
	ID          int       `db:"id"`          // id
	Title       string    `db:"title"`       // title
	Description string    `db:"description"` // description
	CreatedAt   time.Time `db:"created_at"`  // created_at
	UpdatedAt   time.Time `db:"updated_at"`  // updated_at
}

func NewNews(
	ID int,
	Title string,
	Description string,
	CreatedAt time.Time,
	UpdatedAt time.Time,
) *News {
	return &News{
		ID:          ID,
		Title:       Title,
		Description: Description,
		CreatedAt:   CreatedAt,
		UpdatedAt:   UpdatedAt,
	}
}

// Insert inserts the News to the database.
func (n *News) Insert(ctx context.Context, db *sqlx.DB, now time.Time) error {
	n.CreatedAt = now
	n.UpdatedAt = now

	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO public.news (` +
		`title, description, created_at, updated_at` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`) RETURNING id`
	// run
	logf(sqlstr, n.Title, n.Description, n.CreatedAt, n.UpdatedAt)
	if err := db.QueryRowContext(ctx, sqlstr, n.Title, n.Description, n.CreatedAt, n.UpdatedAt).Scan(&n.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Update updates a News in the database.
func (n *News) Update(ctx context.Context, db *sqlx.DB, now time.Time) error {
	n.UpdatedAt = now

	// update with composite primary key
	const sqlstr = `UPDATE public.news SET ` +
		`title = $1, description = $2, created_at = $3, updated_at = $4 ` +
		`WHERE id = $5`
	// run
	logf(sqlstr, n.Title, n.Description, n.CreatedAt, n.UpdatedAt, n.ID)
	if _, err := db.ExecContext(ctx, sqlstr, n.Title, n.Description, n.CreatedAt, n.UpdatedAt, n.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Upsert performs an upsert for News.
func (n *News) Upsert(ctx context.Context, db *sqlx.DB, now time.Time) error {
	// upsert
	const sqlstr = `INSERT INTO public.news (` +
		`id, title, description, created_at, updated_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5` +
		`)` +
		` ON CONFLICT (id) DO ` +
		`UPDATE SET ` +
		`title = EXCLUDED.title, description = EXCLUDED.description, created_at = EXCLUDED.created_at, updated_at = EXCLUDED.updated_at `
	// run
	logf(sqlstr, n.ID, n.Title, n.Description, n.CreatedAt, n.UpdatedAt)
	if _, err := db.ExecContext(ctx, sqlstr, n.ID, n.Title, n.Description, n.CreatedAt, n.UpdatedAt); err != nil {
		return logerror(err)
	}
	return nil
}

// Delete deletes the News from the database.
func (n *News) Delete(ctx context.Context, db *sqlx.DB, now time.Time) error {
	// delete with single primary key
	const sqlstr = `DELETE FROM public.news ` +
		`WHERE id = $1`
	// run
	logf(sqlstr, n.ID)
	if _, err := db.ExecContext(ctx, sqlstr, n.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// NewsByID retrieves a row from 'public.news' as a News.
//
// Generated from index 'news_pkey'.
func NewsByID(ctx context.Context, db *sqlx.DB, id int) (*News, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, title, description, created_at, updated_at ` +
		`FROM public.news ` +
		`WHERE id = $1`
	// run
	logf(sqlstr, id)
	n := News{}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&n.ID, &n.Title, &n.Description, &n.CreatedAt, &n.UpdatedAt); err != nil {
		return nil, logerror(err)
	}
	return &n, nil
}
