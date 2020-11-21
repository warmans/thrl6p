package store

//go:generate go-bindata -modtime 1 -mode 420 -pkg store -o ./migrations_gen.go ./migrations

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/teris-io/shortid"
	"github.com/warmans/thrl6p/server/pkg/flag"
	"github.com/warmans/thrl6p/server/pkg/store/model"
	"github.com/warmans/thrl6p/server/pkg/util"
	"strings"
	"time"
)

type Config struct {
	DSN string
}

func (c *Config) RegisterFlags(prefix string) {
	flag.StringVarEnv(&c.DSN, prefix, "db-dsn", "", "DB connection string")
}

func NewConn(cfg *Config) (*Conn, error) {
	db, err := sqlx.Connect("postgres", cfg.DSN)
	if err != nil {
		return nil, err
	}
	return &Conn{db: db}, nil
}

type Conn struct {
	db *sqlx.DB
}

func (c *Conn) Migrate() error {

	_, err := c.db.Exec(`
		CREATE TABLE IF NOT EXISTS migration_log (
		  file_name TEXT PRIMARY KEY
		);
	`)
	if err != nil {
		return fmt.Errorf("failed to create metadata table: %w", err)
	}

	appliedMigrations := []string{}
	err = c.WithTx(func(tx *sqlx.Tx) error {
		rows, err := tx.Query("SELECT file_name FROM migration_log ORDER BY file_name DESC")
		if err != nil {
			return fmt.Errorf("failed to get migrations: %w", err)
		}
		defer rows.Close()

		for rows.Next() {
			var name string
			if err := rows.Scan(&name); err != nil {
				return err
			}
			appliedMigrations = append(appliedMigrations, name)
		}
		return nil
	})
	if err != nil {
		return err
	}
	if err = c.WithTx(func(tx *sqlx.Tx) error {
		for _, path := range AssetNames() {
			if !strings.HasSuffix(path, ".sql") {
				continue
			}
			if util.InStrings(path, appliedMigrations...) {
				continue
			}
			data, err := Asset(path)
			if err != nil {
				return fmt.Errorf("failed to read file %s: %w", path, err)
			}

			fileName := util.LastString(strings.Split(path, "/"))

			if _, err := tx.Exec(string(data)); err != nil {
				return fmt.Errorf("failed to apply migration %s: %w", fileName, err)
			}
			if _, err := tx.Exec("INSERT INTO migration_log (file_name) VALUES ($1)", fileName); err != nil {
				return fmt.Errorf("failed to update migration log: %w", err)
			}
		}
		return nil
	}); err != nil {
		return fmt.Errorf("failed to apply migraions: %w", err)
	}
	return nil
}

func (c *Conn) WithTx(f func(tx *sqlx.Tx) error) error {
	tx, err := c.db.Beginx()
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}
	if err := f(tx); err != nil {
		if err2 := tx.Rollback(); err2 != nil {
			return fmt.Errorf("failed to rollback (%s) from error : %w)", err2.Error(), err)
		}
		return err
	}
	return tx.Commit()
}

func (c *Conn) WithStore(f func(s *Store) error) error {
	return c.WithTx(func(tx *sqlx.Tx) error {
		return f(&Store{tx: tx})
	})
}

func (c *Conn) Store() *Store {

}

type Store struct {
	tx *sqlx.Tx
}

func (s *Store) CreatePatch(patch *model.Patch) error {
	patch.Id = shortid.MustGenerate()
	patch.CreatedAt = util.TimeP(time.Now())
	_, err := s.tx.NamedExec(
		`INSERT INTO "patch" (id, name, description, data, created_at) VALUES (:id, :name, :description, :data, :created_at)`,
		patch,
	)
	return err
}

func (s *Store) GetPatch(id string) (resp *model.Patch, err error) {
	return resp, s.tx.Get(resp, `SELECT * FROM "patch" WHERE id=?`, id)
}

func (s *Store) ListPatches() ([]*model.Patch, error) {

}
