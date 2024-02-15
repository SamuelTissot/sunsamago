package sunsama

import (
	"testing"

	"github.com/shurcooL/graphql"
	"github.com/stretchr/testify/assert"
)

func Test_streams_contextsForIds(t *testing.T) {
	sts := streams{
		{
			ID:               "643e9937ceaa9e0e548d0164",
			StreamName:       "Admin",
			Personal:         false,
			CategoryStreamID: "616f5f463deaa43ea9da3b8d",
		},
		{
			ID:               "616f5f463deaa43ea9da3b8d",
			StreamName:       "The Parent",
			Personal:         false,
			CategoryStreamID: "",
		},
	}

	ecs := sts.contextsForIds([]graphql.String{"643e9937ceaa9e0e548d0164"})
	want := eventContext{
		name:    "The Parent",
		private: false,
		channel: "Admin",
	}

	assert.Len(t, ecs, 1)
	assert.Equal(t, want, ecs[0])
}
