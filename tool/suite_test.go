package tool

import (
	"github.com/stretchr/testify/suite"
	"net/url"
	"os"
	"testing"
	"xorm.io/xorm"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/lib/pq"
)

type dbSuite struct {
	suite.Suite

	db    *xorm.Engine
	dbCfg *DB
}

func (suite *dbSuite) getDsn() string {
	dsn := os.Getenv("TEST_DSN")
	if dsn == "" {
		return "postgres://postgres:password@localhost:5432/mock?sslmode=disable"
	}

	// modify database name
	u, err := url.Parse(dsn)
	if err != nil {
		panic(err)
	}
	u.Path = "mock"

	return u.String()
}

func (suite *dbSuite) SetupSuite() {
	dbCfg := DB{
		Dsn: suite.getDsn(),
		//MigrationPath: migrationPath,
		MaxIdleConns: 100,
		MaxOpenConns: 100,
	}

	db, err := NewDB(dbCfg)
	suite.Require().NoError(err)

	suite.db = db
}

func TestDBSuite(t *testing.T) {
	suite.Run(t, new(dbSuite))
}

func (suite *dbSuite) Test_QueryMaker() {
	session := suite.db.NewSession()
	defer session.Close()

	req := QueryMakerReq{}
	_, sql := QueryMaker(nil, req)
	suite.Assert().Equal(sql, "")

	req = QueryMakerReq{
		Title: "title",
	}
	_, sql = QueryMaker(session, req)
	suite.Assert().Equal(sql, "title ilike '%title%'")
}
