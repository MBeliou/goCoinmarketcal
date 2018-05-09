# goCoinmarketcal
Unofficial Golang API wrapper for Coinmarketcal

## TO DO:  
- [ ] Docs
- [ ] Comment in code
- [ ] More control over timeouts


## How to use:

- Create a new session:
```go
key := "userKey"
secret := "userSecret"
// s will contain a session containing the access token necessary for other queries.
s, _ := gocoinmarketcal.NewSession(key, secret)
```

- Retrieve coins
```
/*
    Doesn't take any argument.
    Returns an array of Coin structs:
    [{
        ID string
        Name string
        Symbol string
    },
    ...]

*/
s.QueryCoins()
```

- Retrieve categories
```
/*
    Doesn't take any argument.
    Returns an array of Category structs:
    [{
        ID int
        Name string
    },
    ...]
*/
s.QueryCategories()
```

- Retrieve events
```
/*
    Accepts the following arguments
	starting page number -  integer
	max page number - integer
	dateRangeStart - string - (default is today) (format 01/01/2018)
	dateRangeEnd - string -  (default is furthest event) (format 10/05/2018)
	coins string - coins id
	categories strings - Categories
	sortBy - string (created_desc, hot_events)
	showOnly - string (hot_events)
	showMetaData - string (true)
        

    Returns an array of Event structs:
    [{
    	ID                int
	    Title             string
	    CoinsConcerned    Coins
	    DateEvent         string
	    CreatedDate       string
	    Description       string
	    Proof             string
	    Source            string
	    IsHot             string
	    VoteCount         int
	    PositiveVoteCount int
	    Percentage        int
	    Categories        Categories
	    TipSymbol         string
	    TipAddress        string
	    TwitterAccount    string
        CanOccurBefore    bool
    },
    ...]
*/
// Arguments need to be passed as a map[string]interface{}
funcArgs := make(map[string]interface{})
funcArgs["page"] = 1
funcArgs["max"] = 15
resp, _ := s.QueryEvents(funcArgs)
```

## Tips:  
**BTC:** 1D3TYWgSF343oHe4kwD72DoRHUNneHWCdS