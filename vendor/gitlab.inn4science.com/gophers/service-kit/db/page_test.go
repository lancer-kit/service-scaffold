package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPageQuery_Offset(t *testing.T) {
	var pq PageQuery
	err := pq.Validate()

	assert.Equal(t, nil, err)
	assert.Equal(t, OrderAscending, pq.Order)
	assert.Equal(t, uint64(1), pq.Page)
	assert.Equal(t, DefaultPageSize, pq.PageSize)

	assert.Equal(t, uint64(0), pq.Offset())

	pq.Page = 2
	assert.Equal(t, uint64(20), pq.Offset())
	pq.Page = 3
	assert.Equal(t, uint64(40), pq.Offset())
	pq.Page = 4
	assert.Equal(t, uint64(60), pq.Offset())
}
