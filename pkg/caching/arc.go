package caching

import (
	"fmt"

	"github.com/barklan/gotemplate/pkg/myapp/config"
	lru "github.com/hashicorp/golang-lru"
)

type ArcCache struct {
	cl *lru.ARCCache
}

func NewArc(conf *config.Config) (*ArcCache, error) {
	size := conf.FastCacheSize
	if size <= 0 {
		size = 10
	}
	arc, err := lru.NewARC(size)

	return &ArcCache{cl: arc}, fmt.Errorf("failed to init arc cache: %w", err)
}

func (a *ArcCache) Set(key interface{}, val interface{}) {
	a.cl.Add(key, val)
}

func (a *ArcCache) Get(key interface{}) (interface{}, bool) {
	return a.cl.Get(key)
}
