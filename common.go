package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
 * path:bucket/group/year/month/day/microunixtime_rand(1000).jpg
 *
 */
func generationPath(group, suffix string) string {
	now := time.Now()
	year, month, day := now.Date()
	microunixtime := now.UnixNano()
	r := rand.Int31()
	path := fmt.Sprintf("/%s/%d_%d_%d_%v%v.%s", group, year, month, day, microunixtime, r, suffix)
	return path
}
