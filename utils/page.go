package utils

import (
	"net/http"
	"strconv"
)

// PPToLF is the func for transform page and page size to limit and offset which
// prevalent in database.
// if the params is invalid, the result will shink to the least.
func PPToLF(page int, pageSize int) (limit int, offset int) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 1
	}
	limit = pageSize
	offset = (page - 1) * pageSize
	return
}

// PPInQuery is to extract the page and page size from query
func PPInQuery(r *http.Request) (page int, pageSize int, ok bool) {
	m := r.URL.Query()
	rawPage := m.Get("page")
	if rawPage == "" {
		return
	}
	rawPageSize := m.Get("page_size")
	if rawPageSize == "" {
		return
	}
	page, err := strconv.Atoi(rawPage)
	if err != nil {
		return
	}
	pageSize, err = strconv.Atoi(rawPageSize)
	if err != nil {
		return
	}
	ok = true
	return
}
