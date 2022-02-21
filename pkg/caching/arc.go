package caching

import lru "github.com/hashicorp/golang-lru"

type ArcCache struct {
	cl *lru.ARCCache
}

func NewArc(size int) (*ArcCache, error) {
	arc, err := lru.NewARC(size)
	return &ArcCache{cl: arc}, err
}

func (a *ArcCache) Set(key interface{}, val interface{}) {
	a.cl.Add(key, val)
}

func (a *ArcCache) Get(key interface{}) (interface{}, bool) {
	return a.cl.Get(key)
}
