package repository

import (
	"bufio"
	"encoding/json"
	"os"
)

var (
	topicIndexMap map[int64]*Topic
	postIndexMap  map[int64][]*Post
)

func InitTopicIndexMap(filepath string) error {
	open, err := os.Open(filepath + "topic") //open os流handle
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(open)
	topicTmpMap := make(map[int64]*Topic)
	for scanner.Scan() {
		text := scanner.Text()
		var topic Topic
		if err := json.Unmarshal([]byte(text), &topic); err != nil {
			return err
		}
		topicTmpMap[topic.Id] = &topic
	}
	topicIndexMap = topicTmpMap
	return nil
}

func InitPostIndexMap(filepath string) error {
	open, err := os.Open(filepath + "post")
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(open)
	posTmpMap := make(map[int64][]*Post)
	for scanner.Scan() {
		text := scanner.Text()
		var pos Post
		if err := json.Unmarshal([]byte(text), &pos); err != nil {
			return err
		}
		post, ok := posTmpMap[pos.ParentId]
		if !ok {
			posTmpMap[pos.ParentId] = []*Post{&pos}
			continue
		}
		post = append(post, &pos)
		posTmpMap[pos.ParentId] = post
	}
	postIndexMap = posTmpMap
	return nil
}

func Init(filepath string) error {
	if err := InitTopicIndexMap(filepath); err != nil {
		return err
	}
	if err := InitPostIndexMap(filepath); err != nil {
		return err
	}
	return nil
}
