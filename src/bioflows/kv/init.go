package kv

import (
	api "github.com/hashicorp/consul/api"
)

type KVStore interface {

	Put(p *api.KVPair, q *api.WriteOptions) (*api.WriteMeta, error)
	Get(key string, q *api.QueryOptions) (*api.KVPair, *api.QueryMeta, error)
	Delete(key string, w *api.WriteOptions) (*api.WriteMeta, error)
	Keys(prefix, separator string, q *api.QueryOptions) ([]string, *api.QueryMeta, error)
}

