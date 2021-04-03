package cbuff

import (
	"fmt"
	"github.com/gocql/gocql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/cassandra"
	"gitlab.itnap.ru/industrial_safety/cbuff/migrations"
	bindata "github.com/golang-migrate/migrate/v4/source/go_bindata"
	"github.com/scylladb/gocqlx/v2"
	log "github.com/sirupsen/logrus"
)

type cHandler func(session *gocqlx.Session) error

type CassandraConn struct {
	cluster *gocql.ClusterConfig
	session gocqlx.Session
}

func NewCassandraConn(cassandraConfig *CassandraConfig) (*CassandraConn, error) {
	err := migrateDb(cassandraConfig.ConnectionString())
	if err != nil {
		return nil, err
	}

	cluster := gocql.NewCluster(cassandraConfig.Host)
	cluster.Keyspace = cassandraConfig.Keyspace
	cluster.Consistency = gocql.One
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: cassandraConfig.Username,
		Password: cassandraConfig.Password,
	}
	session, err := gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		return nil, err
	}

	return &CassandraConn{cluster: cluster, session: session}, nil
}

func (conn *CassandraConn) runSession(handler cHandler) error {
	return handler(&conn.session)
}

func (conn *CassandraConn) CloseSession() {
	conn.session.Close()
}

func migrateDb(connStr string) error {

	s := bindata.Resource(migrations.AssetNames(),
		func(name string) ([]byte, error) {
			return migrations.Asset(name)
		})

	d, err := bindata.WithInstance(s)
	if err != nil {
		return err
	}

	m, err := migrate.NewWithSourceInstance("go-bindata", d, connStr)
	if err != nil {
		return err
	}

	err = m.Up()
	switch err {
	case nil:
		log.Info("Migrations completed")
	case migrate.ErrNoChange:
		log.Info("Database migrations is up to date")
	default:
		return fmt.Errorf("error migrate database. OError: %v", err)
	}
	return nil
}
