package examples

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"gitlab.inn4science.com/gophers/service-kit/db"
)

type User struct {
	ID    int64  `db:"id" json:"id"`
	Name  string `db:"name" json:"name"`
	Email string `db:"email" json:"email"`
	Age   int    `db:"age" json:"age"`

	RowCount int64 `db:"row_count" json:"-"`
}

// UserQ is a interface for
// interacting with the `users` table.
type UserQ struct {
	*db.SQLConn
	db.Table
}

// NewUserQ returns the new instance of the `UserQ`.
func NewUserQ(conn *db.SQLConn) *UserQ {
	return &UserQ{
		SQLConn: conn.Clone(),
		Table: db.Table{
			Name:     "users",
			QBuilder: sq.Select("*").From("users"),
		},
	}
}

// Insert adds new row into the `users` table.
func (q *UserQ) Insert(user *User) error {
	query := sq.Insert(q.Name).SetMap(map[string]interface{}{
		"name":  user.Name,
		"email": user.Email,
		"age":   user.Age,
	})

	idi, err := q.SQLConn.Insert(query)
	user.ID = idi.(int64)
	return err
}

// ByAge adds in the query filter by the `age` column.
func (q *UserQ) ByAge(age int) *UserQ {
	q.QBuilder = q.QBuilder.Where("age = ?", age)
	return q
}

// SetPage sets the limitation of select
// by the parameters from `db.PageQuery`.
func (q *UserQ) SetPage(pq *db.PageQuery) *UserQ {
	q.Table.SetPage(pq)
	q.WithCount()
	return q
}

// Select gets all records
func (q *UserQ) Select() ([]User, error) {
	dest := make([]User, 0, 1)
	q.ApplyPage("id")

	err := q.SQLConn.Select(q.QBuilder, &dest)
	if err == sql.ErrNoRows {
		return dest, nil
	}

	return dest, err
}

func main() {
	// initialize SQLConn singleton
	err := db.Init("postgres://postgres:postgres@localhost/postgres?sslmode=disable", nil)
	if err != nil {
		panic(err)
	}

	sqlConn := db.GetConnector()
	err = sqlConn.ExecRaw(`CREATE TABLE IF NOT EXIST users(
    id SERIAL, name VARCHAR(64), email VARCHAR(64), age INTEGER)`, nil)
	if err != nil {
		panic(err)
	}
	user := &User{
		Name:  "Mike",
		Email: "mike@example.com",
		Age:   42,
	}
	q := NewUserQ(sqlConn)
	err = q.Insert(user)
	if err != nil {
		panic(err)
	}
}
