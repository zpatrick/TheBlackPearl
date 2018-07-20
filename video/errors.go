package video

import "fmt"

type DoesNotExistError struct {
	Message string
}

func NewDoesNotExistError(videoID string) *DoesNotExistError {
	return &DoesNotExistError{
		Message: fmt.Sprintf("Video '%s' does not exist", videoID),
	}
}

func (e *DoesNotExistError) Error() string {
	return e.Message
}
