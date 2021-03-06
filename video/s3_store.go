package video

import (
	"encoding/base64"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

const (
	SeriesTag       = "Series"
	SeasonTag       = "Season"
	EpisodeTag      = "Episode"
	PosterTag       = "Poster"
)

type S3Store struct {
	bucket string
	path string
	s3     s3iface.S3API
}

func NewS3Store(bucket, path string, s3 s3iface.S3API) *S3Store {
	return &S3Store{
		bucket: bucket,
		path: path,
		s3:     s3,
	}
}

func (s *S3Store) ListVideos() ([]Video, error) {
	keys, err := s.listVideoKeys()
	if err != nil {
		return nil, err
	}

	videos := make([]Video, len(keys))
	for i, key := range keys {
		tags, err := s.getObjectTags(key)
		if err != nil {
			return nil, err
		}

		split := strings.Split(key, "/")
		title := strings.TrimSuffix(split[len(split)-1], ".mp4")

		videos[i] = Video{
			ID:      base64.StdEncoding.EncodeToString([]byte(key)),
			Title:   title,
			Path:    key,
			Poster:  tags[PosterTag],
			Series:  tags[SeriesTag],
			Season:  tags[SeasonTag],
			Episode: tags[EpisodeTag],
		}
	}

	return videos, nil
}

func (s *S3Store) listVideoKeys() ([]string, error) {
	input := &s3.ListObjectsV2Input{
		Bucket: aws.String(s.bucket),
		Prefix: aws.String(s.path),
	}

	if err := input.Validate(); err != nil {
		return nil, err
	}

	output, err := s.s3.ListObjectsV2(input)
	if err != nil {
		return nil, err
	}

	keys := make([]string, 0, int(aws.Int64Value(output.KeyCount)))
	for _, object := range output.Contents {
		if key := aws.StringValue(object.Key); strings.HasSuffix(key, ".mp4") {
			keys = append(keys, key)
		}
	}

	return keys, nil
}

func (s *S3Store) getObjectTags(key string) (map[string]string, error) {
	input := &s3.GetObjectTaggingInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	}

	if err := input.Validate(); err != nil {
		return nil, err
	}

	output, err := s.s3.GetObjectTagging(input)
	if err != nil {
		return nil, err
	}

	tags := map[string]string{}
	for _, tag := range output.TagSet {
		tags[aws.StringValue(tag.Key)] = aws.StringValue(tag.Value)
	}

	return tags, nil
}

func (s *S3Store) GetVideo(videoID string) (Video, error) {
	videos, err := s.ListVideos()
	if err != nil {
		return Video{}, err
	}

	for _, video := range videos {
		if video.ID == videoID {
			return video, nil
		}
	}

	return Video{}, NewDoesNotExistError(videoID)
}
