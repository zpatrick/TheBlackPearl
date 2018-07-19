package video

import "strings"

type Reducer func([]Video) []Video

func NewSearchReducer(search string) Reducer {
	return func(videos []Video) []Video {
		if search == "" {
			return []Video{}
		}

		results := make([]Video, 0, len(videos))
		for i := 0; i < len(videos); i++ {
			if strings.Contains(videos[i].Name, search) {
				results = append(results, videos[i])
			}
		}

		return results
	}
}

func NewLimitReducer(limit int) Reducer {
	return func(videos []Video) []Video {
		if limit <= 0 {
			return []Video{}
		}

		results := make([]Video, 0, len(videos))
		for i := 0; i < limit && i < len(videos); i++ {
			results = append(results, videos[i])
		}

		return results
	}
}

func NewStartReducer(start int) Reducer {
	return func(videos []Video) []Video {
		if start < 0 || start >= len(videos) {
			return []Video{}
		}

		results := make([]Video, 0, len(videos)-start)
		for i := 0; i < len(videos[start:]); i++ {
			results = append(results, videos[start+i])
		}

		return results
	}
}
