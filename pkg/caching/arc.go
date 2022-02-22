package caching

import (
	"github.com/barklan/gotemplate/pkg/myapp/config"
	lru "github.com/hashicorp/golang-lru"
)

type ArcCache struct {
	cl *lru.ARCCache
}

func NewArc(conf *config.Config) (FastCache, error) {
	arc, err := lru.NewARC(conf.FastCacheSize)
	return &ArcCache{cl: arc}, err
}

func (a *ArcCache) Set(key interface{}, val interface{}) {
	a.cl.Add(key, val)
}

func (a *ArcCache) Get(key interface{}) (interface{}, bool) {
	return a.cl.Get(key)
}
