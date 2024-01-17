package mongodb

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson"
)

type Filter bson.M

// add filter but filter must be structure as Filter["$and"] = []Filter{}
func AddFilter(f, additionalFilter Filter) (Filter, error) {
	filterList := []Filter{additionalFilter}
	if v, found := f["$and"]; found {
		if _, ok := v.([]Filter); !ok {
			return nil, errors.New("input invalid structure")
		}
		filterList = append(filterList, v.([]Filter)...)
	} else {
		filterList = append(filterList, f)
	}
	f = Filter{"$and": filterList}
	return f, nil
}

func FilterByKeyList(key string, list []string) Filter {
	return Filter{key: bson.M{"$in": list}}
}

func AddPaginationFilter(currentPage, limit int) Filter {
	return Filter{"$skip": currentPage * limit, "$limit": limit}
}
