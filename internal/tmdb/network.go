package tmdb

import "strconv"

// Network represents a TV network in TMDB.
type Network struct {
    Id            int    `json:"id"`
    Name          string `json:"name"`
    LogoPath      string `json:"logo_path"`
    OriginCountry string `json:"origin_country"`
}

// FetchNetworkParams contains parameters for fetching a network.
type FetchNetworkParams struct {
    Ctx
    Id int
}

// FetchNetworkData is the response payload for fetching a network.
type FetchNetworkData struct {
    ResponseError
    Network
}

// FetchNetwork fetches a network by ID from TMDB.
func (c APIClient) FetchNetwork(params *FetchNetworkParams) (APIResponse[Network], error) {
    response := FetchNetworkData{}
    res, err := c.Request("GET", "/3/network/"+strconv.Itoa(params.Id), params, &response)
    return newAPIResponse(res, response.Network), err
}
