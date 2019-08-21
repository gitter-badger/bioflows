package services

import "bioflows/kv"

type ZooKeeperOrchestrator struct {

	kvStore *kv.ZookeeperKVStoreManager
}
