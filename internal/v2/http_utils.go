package v2

import (
	"fmt"
	"net/url"
)

type PokeClientPagination struct {
	Limit  int
	Offset int
}

func AddPaginationToURL(u *url.URL, pagination PokeClientPagination) {
	query := u.Query()
	query.Add("limit", fmt.Sprintf("%d", pagination.Limit))
	query.Add("offset", fmt.Sprintf("%d", pagination.Offset))
	u.RawQuery = query.Encode()
}
