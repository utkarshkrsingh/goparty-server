// Package db stores room-code and data about number of active members in a room
package db

import "time"

type Room struct {
	ID        string    `db:"id" json:"id"`
	RoomCode  string    `db:"rome_code" json:"roomCode"`
	HostID    string    `db:"host_id" json:"hostId"`
	IsActive  bool      `db:"is_active" json:"isActive"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
}
