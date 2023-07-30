package service

import (
	"os"
	"testing"

	"github.com/Mrmengqiushisan/go_test/repository"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	repository.Init("../data/")
	os.Exit(m.Run())
}

func TestQueryPageInfo(t *testing.T) {
	PageInfo, _ := QueryPageInfo(1)
	assert.NotEqual(t, nil, PageInfo)
	assert.Equal(t, 5, len(PageInfo.PostList))
}
