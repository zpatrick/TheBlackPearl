package video

type Store interface {
	ListVideos() ([]Video, error)
	GetVideo(videoID string) (Video, error)
}
