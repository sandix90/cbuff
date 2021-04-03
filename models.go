package cbuff

import (
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2/table"
	"time"
)

type Pict struct {
	CameraId       int        `db:"camera_id"`
	CreatedAt      time.Time  `db:"created_at"`
	Content        []byte     `db:"content"`
	PictId         gocql.UUID `db:"pict_id"`
	PreviousPictId gocql.UUID `db:"previous_pict_id"`
	Session        gocql.UUID `db:"session"`
}

func getPictTable() *table.Table {
	pictsMetadata := table.Metadata{
		Name:    "frames.picts",
		Columns: []string{"camera_id", "created_at", "content", "pict_id", "previous_pict_id", "session"},
		PartKey: []string{"camera_id"},
		SortKey: []string{"created_at"},
	}
	return table.New(pictsMetadata)
}
