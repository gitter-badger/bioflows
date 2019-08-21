package kv

import "github.com/hashicorp/consul/api"

type ZookeeperKVStoreManager struct{

}


func (kv *ZookeeperKVStoreManager) Put(p *api.KVPair, q *api.WriteOptions) (*api.WriteMeta, error){
	return nil , nil
}

func (kv *ZookeeperKVStoreManager) Get(key string, q *api.QueryOptions) (*api.KVPair, *api.QueryMeta, error){
	return nil , nil , nil
}

func (kv *ZookeeperKVStoreManager) Delete(key string, w *api.WriteOptions) (*api.WriteMeta, error){
	return nil , nil
}

func (kv *ZookeeperKVStoreManager) Keys(prefix, separator string, q *api.QueryOptions) ([]string, *api.QueryMeta, error){
	return nil , nil , nil
}