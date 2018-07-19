package main


/*
Design:
  Routes:
    /           -> Main screen that has big search bar, maybe some 'most-recent'
    /videos     -> Shows all videos in a table, alphabetical order
      ?search=string  -> Only return results that match the search
      ?limit=int      -> Only returns N results
      ?start=int      -> Only start after N results


reducers := []Reducer{}
if search := r.Query("search"); search != "" {
    reducers = append(reducers, NewSearchReducer(search))
}

if start := r.Query("start"); start != 0 {
    reducers = append(reducers, NewStartReducer(start))
}

if limit := r.Query("limit"); limit != 0 {
    reducers = append(reducers, NewLimitReducer(limit))
}

for _, reduce := range reducers {
	reduce(results)
}
*/
