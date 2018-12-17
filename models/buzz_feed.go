package models
//go:generate goplater model --type BuzzFeed --tmpl /Users/mike/go/src/gitlab.inn4science.com/gophers/service-kit/db/q.tmpl --suffix _q
type BuzzFeed struct {
	ID          int64       `db:"id" json:"id"`
	Name        string      `db:"name" json:"name"`
	BuzzType    ExampleType `db:"buzz_type" json:"buzzType"`
	Description string      `db:"description" json:"description"`
	Details     Feed        `db:"details" json:"details"`
	CreatedAt   int64       `db:"created_at" json:"createdAt"`
	UpdatedAt   int64       `db:"updated_at" json:"updatedAt"`
}
