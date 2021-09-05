package cache

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var GoCache *cache.Cache  //定义全局变量
func Init(){
	GoCache = cache.New(5*time.Minute, 60*time.Second)
}
