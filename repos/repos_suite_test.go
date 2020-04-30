package repos_test

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/archi-chester/learning-go-grpc/repos"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	err error
	db *xorm.Engine
	dbSql *sql.DB
	mock sqlmock.Sqlmock

	gr repos.GlobalRepository

	truncateUsers = func() {
		//	test with mock DB (2 lines)
		//mock.ExpectQuery("TRUNCATE users").
		//	WillReturnRows(sqlmock.NewRows([]string{}))

		_, err = db.Query("TRUNCATE users")
		Ω(err).To(BeNil())
	}

	clearDatabase = func() {
		if db == nil {
			Fail("unable to run test because database is missing")
		}

		truncateUsers()

		return
	}
)

var _ = BeforeSuite(func () {
	//	connection string - root:pass@tcp(localhost:3306)/grpc
	//	test with mock DB
	//db, err = xorm.NewEngine("mysql", "")
	db, err = xorm.NewEngine("mysql", "root:Qaz123456@tcp(localhost:3306)/grpc")
	Ω(err).To(BeNil())
	//	test without mock db (2 lines)
	_, err = db.DBMetas()
	Ω(err).To(BeNil())
	//	test with mock DB (3 lines)
	//dbSql, mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	//Ω(err).To(BeNil())
	//db.DB().DB = dbSql

	gr = repos.GlobalRepo(db)
})

var _ = AfterSuite(func() {
	//	test with mock DB (2 lines)
	//err = mock.ExpectationsWereMet()
	//Ω(err).To(BeNil())
})

func TestRepos(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Repos Suite")
}
