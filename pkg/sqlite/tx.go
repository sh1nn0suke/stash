package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/stashapp/stash/pkg/logger"
)

const (
	slowLogTime = time.Millisecond * 200
)

type dbReader interface {
	Get(dest any, query string, args ...any) error
	GetContext(ctx context.Context, dest any, query string, args ...any) error
	SelectContext(ctx context.Context, dest any, query string, args ...any) error
	QueryxContext(ctx context.Context, query string, args ...any) (*sqlx.Rows, error)
}

type stmt struct {
	*sql.Stmt
	query string
}

func logSQL(start time.Time, query string, args ...any) {
	since := time.Since(start)
	if since >= slowLogTime {
		logger.Debugf("SLOW SQL [%v]: %s, args: %v", since, query, args)
	} else {
		logger.Tracef("SQL [%v]: %s, args: %v", since, query, args)
	}
}

type dbWrapperType struct{}

var dbWrapper = dbWrapperType{}

func sqlError(err error, sql string, args ...any) error {
	if err == nil {
		return nil
	}

	return fmt.Errorf("error executing `%s` [%v]: %w", sql, args, err)
}

func (*dbWrapperType) Get(ctx context.Context, dest any, query string, args ...any) error {
	tx, err := getDBReader(ctx)
	if err != nil {
		return sqlError(err, query, args...)
	}

	start := time.Now()
	err = tx.GetContext(ctx, dest, query, args...)
	logSQL(start, query, args...)

	return sqlError(err, query, args...)
}

func (*dbWrapperType) Select(ctx context.Context, dest any, query string, args ...any) error {
	tx, err := getDBReader(ctx)
	if err != nil {
		return sqlError(err, query, args...)
	}

	start := time.Now()
	err = tx.SelectContext(ctx, dest, query, args...)
	logSQL(start, query, args...)

	return sqlError(err, query, args...)
}

func (*dbWrapperType) Queryx(ctx context.Context, query string, args ...any) (*sqlx.Rows, error) {
	tx, err := getDBReader(ctx)
	if err != nil {
		return nil, sqlError(err, query, args...)
	}

	start := time.Now()
	ret, err := tx.QueryxContext(ctx, query, args...)
	logSQL(start, query, args...)

	return ret, sqlError(err, query, args...)
}

func (*dbWrapperType) QueryxContext(ctx context.Context, query string, args ...any) (*sqlx.Rows, error) {
	return dbWrapper.Queryx(ctx, query, args...)
}

func (*dbWrapperType) NamedExec(ctx context.Context, query string, arg any) (sql.Result, error) {
	tx, err := getTx(ctx)
	if err != nil {
		return nil, sqlError(err, query, arg)
	}

	start := time.Now()
	ret, err := tx.NamedExecContext(ctx, query, arg)
	logSQL(start, query, arg)

	return ret, sqlError(err, query, arg)
}

func (*dbWrapperType) Exec(ctx context.Context, query string, args ...any) (sql.Result, error) {
	tx, err := getTx(ctx)
	if err != nil {
		return nil, sqlError(err, query, args...)
	}

	start := time.Now()
	ret, err := tx.ExecContext(ctx, query, args...)
	logSQL(start, query, args...)

	return ret, sqlError(err, query, args...)
}

// Prepare creates a prepared statement.
func (*dbWrapperType) Prepare(ctx context.Context, query string, args ...any) (*stmt, error) {
	tx, err := getTx(ctx)
	if err != nil {
		return nil, sqlError(err, query, args...)
	}

	// nolint:sqlclosecheck
	ret, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return nil, sqlError(err, query, args...)
	}

	return &stmt{
		query: query,
		Stmt:  ret,
	}, nil
}

func (*dbWrapperType) ExecStmt(ctx context.Context, stmt *stmt, args ...any) (sql.Result, error) {
	_, err := getTx(ctx)
	if err != nil {
		return nil, sqlError(err, stmt.query, args...)
	}

	start := time.Now()
	ret, err := stmt.ExecContext(ctx, args...)
	logSQL(start, stmt.query, args...)

	return ret, sqlError(err, stmt.query, args...)
}
