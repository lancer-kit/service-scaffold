package render

import (
	"math"
	"net/http"
)

type Page struct {
	Page     uint64      `json:"page"`
	PageSize uint64      `json:"pageSize"`
	Order    string      `json:"order"`
	Total    int64       `json:"total"`
	Records  interface{} `json:"records"`
}

func (page *Page) Render(w http.ResponseWriter) {
	WriteJSON(w, http.StatusOK, page)
}

func (page *Page) SetTotal(rowCount, pageSize uint64) {
	page.Total = int64(math.Ceil(float64(rowCount) / float64(pageSize)))
}

type BaseRow struct {
	RowCount int64 `db:"row_count" json:"-"`
}
