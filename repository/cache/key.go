package cache

import (
	"fmt"
	"strconv"
)

const (
	RankKey = "rank"
)

// TaskViewKey 点击数的key
func PostViewKey(id uint) string {
	return fmt.Sprintf("view:task:%s", strconv.Itoa(int(id)))
}
