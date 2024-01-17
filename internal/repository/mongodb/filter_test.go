package mongodb_test

import (
	"testing"

	"github.com/ZecretBone/ips-bff/internal/repository/mongodb"
	"github.com/stretchr/testify/assert"
)

func Test_AddFilter_FirstTime(t *testing.T) {
	filter := mongodb.Filter{"object_id": "ABCD"}
	filter, err := mongodb.AddFilter(filter, mongodb.AddPaginationFilter(1, 10))
	assert.NoError(t, err)

	expect := mongodb.Filter{
		"$and": []mongodb.Filter{{"$skip": 10, "$limit": 10}, {"object_id": "ABCD"}},
	}
	assert.Equal(t, expect, filter)
}

func Test_AddFilter_Not_First(t *testing.T) {
	filter := mongodb.Filter{"$and": []mongodb.Filter{{"object_id": "ABCD"}}}
	filter, err := mongodb.AddFilter(filter, mongodb.AddPaginationFilter(1, 10))
	assert.NoError(t, err)

	expect := mongodb.Filter{
		"$and": []mongodb.Filter{{"$skip": 10, "$limit": 10}, {"object_id": "ABCD"}},
	}
	assert.Equal(t, expect, filter)
}

func Test_AddFilter_Invalid_Type(t *testing.T) {
	filter := mongodb.Filter{"$and": "ABCD"}
	filter, err := mongodb.AddFilter(filter, mongodb.AddPaginationFilter(1, 10))
	assert.Error(t, err, "input invalid structure")
	assert.Nil(t, filter)
}
