package service

import (
	"errors"
	"sync"

	"github.com/Mrmengqiushisan/go_test/repository"
)

type PageInfo struct {
	Topic    *repository.Topic
	PostList []*repository.Post
}

type QueryPageInfoFlow struct {
	topicId  int64
	pageInfo *PageInfo
	topic    *repository.Topic
	posts    []*repository.Post
}

func NewQueryPageInfoFlow(topID int64) *QueryPageInfoFlow {
	return &QueryPageInfoFlow{
		topicId: topID,
	}
}

func (f *QueryPageInfoFlow) Do() (*PageInfo, error) {
	if err := f.CheckParam(); err != nil {
		return nil, err
	}
	if err := f.PrepareInfo(); err != nil {
		return nil, err
	}
	if err := f.PackPageInfo(); err != nil {
		return nil, err
	}
	return f.pageInfo, nil
}

func (f *QueryPageInfoFlow) CheckParam() error {
	if f.topicId <= 0 {
		return errors.New("topci id must be larger than 0")
	}
	return nil
}

func (f *QueryPageInfoFlow) PrepareInfo() error {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		topic := repository.NewTopicDaoInstance().QueryTopicById(f.topicId)
		f.topic = topic
	}()
	go func() {
		defer wg.Done()
		posts := repository.NewPostDaoInstance().QueryPostsByParentId(f.topicId)
		f.posts = posts
	}()
	wg.Wait()
	return nil
}

func (f *QueryPageInfoFlow) PackPageInfo() error {
	f.pageInfo = &PageInfo{
		Topic:    f.topic,
		PostList: f.posts,
	}
	return nil
}
func QueryPageInfo(topicId int64) (*PageInfo, error) {
	return NewQueryPageInfoFlow(topicId).Do()
}
