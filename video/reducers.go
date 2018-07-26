package video

import "strings"

type Reducer func([]Video) []Video

// todo: if user puts 'season X' or 'episode Y', filter by those fields
func NewSearchReducer(search string) Reducer {
	words := strings.Split(strings.ToLower(search), " ")
	return func(videos []Video) []Video {
		if search == "" {
			return []Video{}
		}

		results := make([]Video, 0, len(videos))
		for i := 0; i < len(videos); i++ {
			title := strings.ToLower(videos[i].Title)
			series := strings.ToLower(videos[i].Series)

			for _, word := range words {
				if strings.Contains(title, word) {
					results = append(results, videos[i])
					break
				} else if strings.Contains(series, word) {
					results = append(results, videos[i])
					break
				}
			}
		}

		return results
	}
}

func NewLimitReducer(limit int) Reducer {
	return func(videos []Video) []Video {
		switch {
		case limit < 0:
			return []Video{}
		case limit > len(videos):
			return videos
		default:
			return videos[:limit]
		}
	}
}

func NewStartReducer(start int) Reducer {
	return func(videos []Video) []Video {
		switch {
		case start < 0:
			return []Video{}
		case start > len(videos):
			return []Video{}
		default:
			return videos[start:]
		}
	}
}
