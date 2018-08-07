package video

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/zpatrick/TheBlackPearl/mocks"
)

func TestS3StoreListVideos(t *testing.T) {
	t.Skip("TODO")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockS3 := mocks.NewMockS3API(ctrl)
	store := NewS3Store("bucket", "/path", mockS3)

	result, err := store.ListVideos()
	if err != nil {
		t.Fatal(err)
	}

	expected := []Video{}
	assert.Equal(t, expected, result)
}
