package tmdb

import (
	"net/url"
	"strconv"
)

type FetchDiscoverMovieData = fetchListData[ListItemMovie]

type DiscoverMovieParams struct {
	Ctx
	IncludeAdult  bool
	Page          int
	SortBy        string
	WithCompanies string // can be a comma (AND) or pipe (OR) separated query
}

func (c APIClient) DiscoverMovie(params *DiscoverMovieParams) (APIResponse[FetchDiscoverMovieData], error) {
	query := url.Values{}
	if params.Page > 0 {
		query.Set("page", strconv.Itoa(params.Page))
	}
	if params.SortBy != "" {
		query.Set("sort_by", params.SortBy)
	}
	if params.IncludeAdult {
		query.Set("include_adult", "true")
	}
	if params.WithCompanies != "" {
		query.Set("with_companies", params.WithCompanies)
	}
	params.Query = &query

	response := FetchDiscoverMovieData{}
	res, err := c.Request("GET", "/3/discover/movie", params, &response)
	return newAPIResponse(res, response), err
}

type FetchDiscoverTVData = fetchListData[ListItemShow]

type DiscoverTVParams struct {
	Ctx
	IncludeAdult  bool
	Page          int
	SortBy        string
	WithCompanies string // can be a comma (AND) or pipe (OR) separated query
	    WithNetworks string // can be a comma (AND) or pipe (OR) separated query
}

func (c APIClient) DiscoverTV(params *DiscoverTVParams) (APIResponse[FetchDiscoverTVData], error) {
	query := url.Values{}
	if params.Page > 0 {
		query.Set("page", strconv.Itoa(params.Page))
	}
	if params.SortBy != "" {
		query.Set("sort_by", params.SortBy)
	}
	if params.IncludeAdult {
		query.Set("include_adult", "true")
	}
	        if params.WithNetworks != "" {
            query.Set("with_networks", params.WithNetworks)
        }
	if params.WithCompanies != "" {
		query.Set("with_companies", params.WithCompanies)
	}
	params.Query = &query

	response := FetchDiscoverTVData{}
	res, err := c.Request("GET", "/3/discover/tv", params, &response)
	return newAPIResponse(res, response), err
}
