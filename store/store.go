package store

import (
	"database/sql"

	"maunium.net/go/mautrix"
	mid "maunium.net/go/mautrix/id"
)

type StateStore struct {
	DB                  *sql.DB
	Client              *mautrix.Client
	UserConfigRoomCache map[mid.UserID]mid.RoomID
	UserTimezoneCache   map[mid.UserID]string
	UserNotifyTimeCache map[mid.UserID]int
	UserSendRoomCache   map[mid.UserID]mid.RoomID
}

func NewStateStore(db *sql.DB) *StateStore {
	return &StateStore{
		DB:                  db,
		UserConfigRoomCache: map[mid.UserID]mid.RoomID{},
		UserTimezoneCache:   map[mid.UserID]string{},
		UserNotifyTimeCache: map[mid.UserID]int{},
		UserSendRoomCache:   map[mid.UserID]mid.RoomID{},
	}
}

func (store *StateStore) CreateTables() error {
	tx, err := store.DB.Begin()
	if err != nil {
		return err
	}

	queries := []string{
		`
		CREATE TABLE IF NOT EXISTS standupbot_meta (
			meta_id       INTEGER PRIMARY KEY,
			access_token  VARCHAR(255)
		)
		`,
		`
		CREATE TABLE IF NOT EXISTS user_filter_ids (
			user_id    VARCHAR(255) PRIMARY KEY,
			filter_id  VARCHAR(255)
		)
		`,
		`
		CREATE TABLE IF NOT EXISTS user_batch_tokens (
			user_id           VARCHAR(255) PRIMARY KEY,
			next_batch_token  VARCHAR(255)
		)
		`,
		`
		CREATE TABLE IF NOT EXISTS rooms (
			room_id           VARCHAR(255) PRIMARY KEY,
			encryption_event  VARCHAR(65535) NULL
		)
		`,
		`
		CREATE TABLE IF NOT EXISTS room_members (
			room_id  VARCHAR(255),
			user_id  VARCHAR(255),
			PRIMARY KEY (room_id, user_id)
		)
		`,
		`
		DROP TABLE IF EXISTS user_config_room
		`,
	}

	for _, query := range queries {
		if _, err := tx.Exec(query); err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
