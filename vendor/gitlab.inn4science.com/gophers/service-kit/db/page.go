package db

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

const (
	DefaultPageSize uint64 = 20
	MaxPageSize     uint64 = 500

	OrderAscending  = "asc"
	OrderDescending = "desc"
)

var (
	ErrInvalidOrder = func(val string) error {
		return fmt.Errorf("order(%s): accept only %s|%s",
			val, OrderAscending, OrderDescending)
	}
	ErrTooBigPage = func(val uint64) error {
		return fmt.Errorf("pageSize(%d): shoud be less or equal %d", val, MaxPageSize)
	}
)

// PageQuery is the structure for
// building query with pagination.
type PageQuery struct {
	Order    string `json:"order"`
	Page     uint64 `json:"page"`
	PageSize uint64 `json:"pageSize"`
}

// ParsePageQuery extracts `PageQuery` from the url Query Values.
func ParsePageQuery(values url.Values) (pq PageQuery, err error) {
	err = pq.FromRQuery(values)
	return
}

// FromRQuery extracts `PageQuery` from the url Query Values and validate.
func (pq *PageQuery) FromRQuery(query url.Values) error {
	page := query.Get("page")
	if page == "" {
		page = "0"
	}
	var err error
	pq.Page, err = strconv.ParseUint(page, 10, 64)
	if err != nil {
		return errors.Wrap(err, "page")
	}
	pageSize := query.Get("pageSize")
	if pageSize == "" {
		pageSize = "0"
	}

	pq.PageSize, err = strconv.ParseUint(pageSize, 10, 64)
	if err != nil {
		return errors.Wrap(err, "pageSize")
	}
	pq.Order = query.Get("order")

	return pq.Validate()
}

// Validate checks is correct values and set default values if `PageQuery` empty.
func (pq *PageQuery) Validate() error {
	switch strings.ToLower(pq.Order) {
	case "":
		pq.Order = OrderAscending
	case OrderAscending, OrderDescending:
		break
	default:
		return ErrInvalidOrder(pq.Order)
	}

	if pq.Page == 0 {
		pq.Page = 1
	}

	if pq.PageSize == 0 {
		pq.PageSize = DefaultPageSize
	}

	if pq.PageSize > MaxPageSize {
		return ErrTooBigPage(pq.PageSize)
	}

	return nil
}

// Offset calculates select offset.
func (pq *PageQuery) Offset() uint64 {
	return (pq.Page - 1) * pq.PageSize
}

// Apply sets limit and ordering params to SelectBuilder.
func (pq *PageQuery) Apply(query sq.SelectBuilder, orderColumn string) sq.SelectBuilder {
	query = query.Limit(pq.PageSize).Offset(pq.Offset())
	if pq.Order != "" && orderColumn != "" {
		query = query.OrderBy(orderColumn + " " + pq.Order)
	}

	return query
}
