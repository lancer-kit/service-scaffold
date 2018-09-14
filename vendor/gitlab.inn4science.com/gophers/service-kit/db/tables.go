package db

import (
	sq "github.com/Masterminds/squirrel"
)

// Table is the basis for implementing
// Querier for some model or table.
type Table struct {
	Name  string
	Alias string

	DB        sq.BaseRunner
	QBuilder  sq.SelectBuilder
	GQBuilder sq.SelectBuilder
	IQBuilder sq.InsertBuilder
	UQBuilder sq.UpdateBuilder
	DQBuilder sq.DeleteBuilder
	Page      *PageQuery
}

// AliasedName returns table name with the alias postfix.
func (t Table) AliasedName() string {
	return t.Name + " " + t.Alias
}

// SetPage is a setter for Page field.
func (t *Table) SetPage(pq *PageQuery) {
	t.Page = pq
}

// ApplyPage adds limit/offset and/or order to the queryBuilder.
func (t *Table) ApplyPage(orderColumn string) {
	if t.Page != nil {
		t.QBuilder = t.Page.Apply(t.QBuilder, orderColumn)
		return
	}

	t.QBuilder = t.QBuilder.OrderBy(orderColumn)
}

// WithCount adds a column with the total number of records.
// ATTENTION! The model must have a destination for this `row_count` column.
func (t *Table) WithCount() {
	t.QBuilder = t.QBuilder.Column("count(*) OVER() AS row_count")
}
