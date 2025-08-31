// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	ICache interface {
		UnmarshalJson(ctx context.Context, value string, interfaceType interface{}) interface{}
		GetCacheWithPrefix(ctx context.Context, key string, interfaceType interface{}) interface{}
		SetCacheWithPrefix(ctx context.Context, key string, value string) error
		SetCacheWithPrefixByInterface(ctx context.Context, key string, value interface{}) error
		RemoveCacheWithPrefix(ctx context.Context, key string) error
		RemoveMulCachesWithPrefix(ctx context.Context, keys []string) int
		RemoveCacheMatchedPattern(ctx context.Context, pattern string) error
	}
)

var (
	localCache ICache
)

func Cache() ICache {
	if localCache == nil {
		panic("implement not found for interface ICache, forgot register?")
	}
	return localCache
}

func RegisterCache(i ICache) {
	localCache = i
}
