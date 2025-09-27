package tmdb

import (
    "github.com/MunifTanjim/stremthru/internal/util"
)



// This file defines functions to fetch network-based lists for TV shows.
// It complements the existing list.go by providing support for network URLs.

// fetchNetworkListItems fetches a single page of TV shows for a given network ID.
// It uses the DiscoverTV API with the WithNetworks parameter to filter by network.
// It always returns TV shows (MediaTypeTVShow).
func fetchNetworkListItems(client *APIClient, networkId string, page int) (*(*t*tmdbListData, error
    data := tmdbListData{}
    // Fetch TV results for the network
    res, err := tmdbListData.DiscoverTV(&DiscoverTVParams{
        Page:        page,
        WithNetworks: networkId,
    })
    if err != nil {
        return nil, err
    }
    // Populate paging data
    data.Page = res.Data.Page
    data.TotalPages = res.Data.TotalPages
    data.TotalResults = res.Data.TotalResults
    // Append each result as a ListItem with TV media type
    for i := range res.Data.Results {
        data.Results = append(data.Results, ListItem{
            MediaType: MediaTypeTVShow,
            data:      res.Data.Results[i],
        })
    }
    return &data, nil
}

// fetchNetworkList fetches all pages of TV shows for a given network ID and returns a List.
// It uses FetchNetwork to retrieve the network's name for labeling the list.
func fetchNetworkList(client *APIClient, networkId string) (*List, error) {
    l := List{
        Name:      "",
        Public:    true,
        ItemCount: 0,
    }
    // Retrieve network details for the name
    network, err := client.FetchNetwork(&FetchNetworkParams{
        Id: util.SafeParseInt(networkId, -1),
    })
    if err != nil {
        return nil, err
    }
    l.Name = network.Data.Name

    page := 0
    for {
        page++
        res, err := fetchNetworkListItems(client, networkId, page)
        if err != nil {
            return nil, err
        }
        // Update item count and accumulate results
        l.ItemCount = res.TotalResults
        l.Results = append(l.Results, res.Results...)
        if res.Page >= res.TotalPages {
            break
        }
    }
    return &l, nil
}
