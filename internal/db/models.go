// Package db stores room-code and data about number of active members in a room
package db

type Room struct {
	ID        string `schema:"id,required"`
	UserCount int    `schema:"user_count"`
	CreatedAt string `schema:"create_at"`
}
