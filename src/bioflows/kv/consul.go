package kv

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

type ConsulKVStoreManager struct{
	client *api.Client
}
func (kv *ConsulKVStoreManager) GetClient() interface{} {
	return kv.client
}
func (kv *ConsulKVStoreManager) Setup(creds Credentials) error {

	config := api.DefaultConfig()
	config.Address = fmt.Sprintf("%s:%d",creds.Address,creds.Port)
	if creds.Username != "" && creds.Password != "" {
		config.HttpAuth = &api.HttpBasicAuth{
			Username:creds.Username,
			Password:creds.Password,
		}
	}
	client , err := api.NewClient(config)
	if err != nil {
		return err
	}
	kv.client = client
	return nil
}

func (kv *ConsulKVStoreManager) List(prefix string, q *api.QueryOptions) (api.KVPairs, *api.QueryMeta, error){
	store := kv.client.KV()
	return store.List(prefix,q)
}

func (kv *ConsulKVStoreManager) Put(p *api.KVPair, q *api.WriteOptions) (*api.WriteMeta, error){

	store := kv.client.KV()
	return store.Put(p,q)
}

func (kv *ConsulKVStoreManager) Get(key string, q *api.QueryOptions) (*api.KVPair, *api.QueryMeta, error){
	store := kv.client.KV()
	return store.Get(key,q)
}

func (kv *ConsulKVStoreManager) Delete(key string, w *api.WriteOptions) (*api.WriteMeta, error){
	store := kv.client.KV()
	return store.Delete(key,w)
}

func (kv *ConsulKVStoreManager) Keys(prefix, separator string, q *api.QueryOptions) ([]string, *api.QueryMeta, error){
	store := kv.client.KV()
	return store.Keys(prefix,separator,q)
}