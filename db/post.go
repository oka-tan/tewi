// Package db has models for db access
package db

import (
	"encoding/base64"
	"encoding/json"
	"time"

	"github.com/uptrace/bun"
)

// MediaHash has different JSON marshalling behavior from standard bytes
type MediaHash []byte

// MarshalJSON marshals a MediaHash as JSON.
// Default for []byte is basically base64.StdEncoding.EncodeToString(m)
func (m MediaHash) MarshalJSON() ([]byte, error) {
	return json.Marshal(base64.URLEncoding.EncodeToString(m))
}

// Post is a post in the db
type Post struct {
	bun.BaseModel `bun:"table:post,alias:post"`

	Board                 string     `bun:"board,pk" json:"-"`
	PostNumber            int64      `bun:"post_number,pk" json:"postNumber"`
	ThreadNumber          int64      `bun:"thread_number" json:"threadNumber"`
	Op                    bool       `bun:"op" json:"op"`
	Deleted               bool       `bun:"deleted" json:"deleted"`
	Hidden                bool       `bun:"hidden" json:"-"`
	TimePosted            time.Time  `bun:"time_posted" json:"timePosted"`
	LastModified          time.Time  `bun:"last_modified" json:"lastModified"`
	CreatedAt             time.Time  `bun:"created_at" json:"createdAt"`
	Name                  *string    `bun:"name" json:"name,omitempty"`
	Tripcode              *string    `bun:"tripcode" json:"tripcode,omitempty"`
	Capcode               *string    `bun:"capcode" json:"capcode,omitempty"`
	PosterID              *string    `bun:"poster_id" json:"posterId,omitempty"`
	Country               *string    `bun:"country" json:"country,omitempty"`
	Flag                  *string    `bun:"flag" json:"flag,omitempty"`
	Email                 *string    `bun:"email" json:"email,omitempty"`
	Subject               *string    `bun:"subject" json:"subject,omitempty"`
	Comment               *string    `bun:"comment" json:"comment,omitempty"`
	HasMedia              bool       `bun:"has_media" json:"hasMedia"`
	MediaDeleted          *bool      `bun:"media_deleted" json:"mediaDeleted,omitempty"`
	TimeMediaDeleted      *time.Time `bun:"time_media_deleted" json:"timeMediaDeleted,omitempty"`
	MediaTimestamp        *int64     `bun:"media_timestamp" json:"mediaTimestamp,omitempty"`
	Media4chanHash        *MediaHash `bun:"media_4chan_hash" json:"media4chanHash,omitempty"`
	MediaInternalHash     *MediaHash `bun:"media_internal_hash" json:"mediaInternalHash,omitempty"`
	ThumbnailInternalHash *MediaHash `bun:"thumbnail_internal_hash" json:"thumbnailInternalHash,omitempty"`
	MediaExtension        *string    `bun:"media_extension" json:"mediaExtension,omitempty"`
	MediaFileName         *string    `bun:"media_file_name" json:"mediaFileName,omitempty"`
	MediaSize             *int       `bun:"media_size" json:"mediaSize,omitempty"`
	MediaHeight           *int16     `bun:"media_height" json:"mediaHeight,omitempty"`
	MediaWidth            *int16     `bun:"media_width" json:"mediaWidth,omitempty"`
	ThumbnailHeight       *int16     `bun:"thumbnail_height" json:"thumbnailHeight,omitempty"`
	ThumbnailWidth        *int16     `bun:"thumbnail_width" json:"thumbnailWidth,omitempty"`
	Spoiler               *bool      `bun:"spoiler" json:"spoiler,omitempty"`
	CustomSpoiler         *int16     `bun:"custom_spoiler" json:"customSpoiler,omitempty"`
	Sticky                *bool      `bun:"sticky" json:"sticky,omitempty"`
	Closed                *bool      `bun:"closed" json:"closed,omitempty"`
	Posters               *int16     `bun:"posters" json:"posters,omitempty"`
	Replies               *int16     `bun:"replies" json:"replies,omitempty"`
	Since4Pass            *int16     `bun:"since4pass" json:"since4pass,omitempty"`
}
