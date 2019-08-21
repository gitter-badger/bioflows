package kv

import "github.com/hashicorp/consul/api"

type ConsulKVStoreManager struct{

}

func (kv *ConsulKVStoreManager) Put(p *api.KVPair, q *api.WriteOptions) (*api.WriteMeta, error){
	return nil , nil
}

func (kv *ConsulKVStoreManager) Get(key string, q *api.QueryOptions) (*api.KVPair, *api.QueryMeta, error){
	return nil , nil , nil
}

func (kv *ConsulKVStoreManager) Delete(key string, w *api.WriteOptions) (*api.WriteMeta, error){
	return nil , nil
}

func (kv *ConsulKVStoreManager) Keys(prefix, separator string, q *api.QueryOptions) ([]string, *api.QueryMeta, error){
	return nil , nil , nil
}