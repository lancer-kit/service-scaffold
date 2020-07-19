package models

// //go:generate forge model --type BuzzFeed --tmpl ./templates/pg_repo.tmpl --suffix _repo
type BuzzFeed struct {
	ID          int64       `db:"id" json:"id"`
	Name        string      `db:"name" json:"name"`
	BuzzType    ExampleType `db:"buzz_type" json:"buzzType"`
	Description string      `db:"description" json:"description"`
	Details     Feed        `db:"details" json:"details"`
	CreatedAt   int64       `db:"created_at" json:"createdAt"`
	UpdatedAt   int64       `db:"updated_at" json:"updatedAt"`
}
