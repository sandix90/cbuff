package cbuff

import (
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
	"time"
)

func (conn *CassandraConn) InsertPict(sessionId string, cameraId int, pictId string, previousPictId string, createdAt time.Time, b []byte) error {

	return conn.runSession(func(session *gocqlx.Session) error {

		insertPictQ := getPictTable().InsertQuery(*session)

		p := &Pict{
			CameraId:       cameraId,
			CreatedAt:      createdAt,
			Content:        b,
			PictId:         mustParseUUID(pictId),
			PreviousPictId: mustParseUUID(previousPictId),
			Session:        mustParseUUID(sessionId),
		}
		insertPictQ.BindStruct(p)
		if err := insertPictQ.ExecRelease(); err != nil {
			return err
		}
		return nil
	})
}

func (conn *CassandraConn) GetPictById(cameraId int, createdAt time.Time, id string) (*Pict, error) {
	pict := new(Pict)

	err := conn.runSession(
		func(session *gocqlx.Session) error {

			q := qb.Select("frames.picts").Columns(
				"camera_id", "created_at", "content", "pict_id", "previous_pict_id", "session",
			).Where(qb.Eq("pict_id"), qb.Eq("camera_id"), qb.Eq("created_at"))

			getQ := q.Query(*session).BindStruct(&Pict{
				PictId: mustParseUUID(id),
				CameraId: cameraId,
				CreatedAt: createdAt,
			})

			if err := getQ.Get(pict); err != nil {
				return err
			}
			return nil
		},
	)
	if err != nil {
		return nil, err
	}

	return pict, nil
}

func (conn *CassandraConn) DeletePictById(cameraId int, createdAt time.Time, id string) error {
	return conn.runSession(
		func(session *gocqlx.Session) error {
			q := qb.Delete("frames.picts").Where(qb.Eq("pict_id"), qb.Eq("camera_id"), qb.Eq("created_at"))

			deleteQ := q.Query(*session).BindStruct(&Pict{
				PictId:   mustParseUUID(id),
				CameraId: cameraId,
				CreatedAt: createdAt,
			})

			if err := deleteQ.Exec(); err != nil {
				return err
			}
			return nil
		},
	)
}

func mustParseUUID(s string) gocql.UUID {
	u, err := gocql.ParseUUID(s)
	if err != nil {
		return gocql.UUID{}
	}
	return u
}
