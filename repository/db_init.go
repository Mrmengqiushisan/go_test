package repository

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

var (
	topicIndexMap map[int64]*Topic
	topicId       int64
	postIndexMap  map[int64][]*Post
	postId        int64
)

func GetPostCount() *int64 {
	return &postId
}
func GetTopicCount() *int64 {
	return &topicId
}

func InitTopicIndexMap(filepath string) error {
	if len(topicIndexMap) > 0 {
		topicIndexMap = make(map[int64]*Topic)
		topicId = 0
	}
	open, err := os.Open(filepath + "topic") //open os流handle
	if err != nil {
		return err
	}
	defer open.Close()
	scanner := bufio.NewScanner(open)
	topicTmpMap := make(map[int64]*Topic)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
		var topic Topic
		if err := json.Unmarshal([]byte(text), &topic); err != nil {
			return err
		}
		buf, err := json.MarshalIndent(topic, "", "\t")
		if err != nil {
			return err
		}
		fmt.Println(string(buf))
		topicTmpMap[topic.Id] = &topic
		topicId++
	}
	topicIndexMap = topicTmpMap
	return nil
}

func InitPostIndexMap(filepath string) error {
	if len(postIndexMap) > 0 {
		postIndexMap = make(map[int64][]*Post)
		postId = 0
	}
	open, err := os.Open(filepath + "post")
	if err != nil {
		return err
	}
	defer open.Close()
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
			postId++
			continue
		}
		post = append(post, &pos)
		posTmpMap[pos.ParentId] = post
		postId++
	}
	postIndexMap = posTmpMap
	return nil
}

func Init(filepath string) error {
	topicId = 0
	postId = 0
	if err := InitTopicIndexMap(filepath); err != nil {
		return err
	}
	if err := InitPostIndexMap(filepath); err != nil {
		return err
	}
	return nil
}
