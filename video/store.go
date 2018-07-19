package video

type Store interface {
	ListVideos() ([]Video, error)
}
