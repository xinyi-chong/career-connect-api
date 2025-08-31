package cache

import (
	"context"
	"encoding/json"
	"gf_demo/internal/service"
	"time"

	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type (
	sCache struct{}
)

func init() {
	service.RegisterCache(New())
}

func New() service.ICache {
	return &sCache{}
}

func (s *sCache) UnmarshalJson(ctx context.Context, value string, interfaceType interface{}) (interface{}) {
	// Unmarshal JSON string into interfaceType
	err := json.Unmarshal([]byte(value), interfaceType)
	if err != nil {
		g.Log().Error(ctx, "Failed to unmarshal cache value: ", err)
	}
	return interfaceType
}

func (s *sCache) GetCacheWithPrefix(ctx context.Context, key string, interfaceType interface{}) (interface{}) {
	value, err := g.Redis().Get(ctx, "SelectCache:" + key)
	if err != nil {
		g.Log().Error(ctx, "Failed to get cache by key: ", key, err)
		return nil
	} else if value.Interface() == nil {
		g.Log().Error(ctx, "Cache Not Found")
		return nil
	}

	jsonValue := s.UnmarshalJson(ctx, value.String(), interfaceType)

	return jsonValue
}

func (s *sCache) SetCacheWithPrefix(ctx context.Context, key string, value string) error {
	ttl := int64((time.Hour * 24) / time.Second)
	err := g.Redis().SetEX(ctx, "SelectCache:" + key, value, ttl)
	// g.Redis().Do(ctx, "EXPIRE", key, time.Hour * 24)
	if err != nil {
		g.Log().Error(ctx, "Failed to Set Cache: ", key, ":", value, err)
		err = gerror.NewCode(gcode.CodeOperationFailed, "Failed to Set Cache: " + err.Error())
	}
	
	return err
}

func (s *sCache) SetCacheWithPrefixByInterface(ctx context.Context, key string, value interface{}) error {
	valueJson, err := json.Marshal(value)
	if err != nil {
		g.Log().Error(ctx, "Failed to marshal json: ", value, err)
		err = gerror.NewCode(gcode.CodeOperationFailed, "Failed to marshal json: " + err.Error())
		return err
	}
	
	err = s.SetCacheWithPrefix(ctx, key, string(valueJson))

	return err
}

func (s *sCache) RemoveCacheWithPrefix(ctx context.Context, key string) error {
	_, err := g.Redis().Del(ctx, "SelectCache:" + key)
	if err != nil {
		g.Log().Error(ctx, "Failed to Remove Cache from Redis By Key: ", key, err)
		err = gerror.NewCode(gcode.CodeOperationFailed, "Failed to Remove Cache from Redis By Key: " + err.Error())
	}
	
	return err
}

func (s *sCache) RemoveMulCachesWithPrefix(ctx context.Context, keys []string) int {
	countSuccess := 0
	for _, key := range keys {
		err := s.RemoveCacheWithPrefix(ctx, key)
		if err != nil {
			g.Log().Error(ctx, "Failed to Remove Cache from Redis By Key: ", key, err)
			continue
		}
		countSuccess++
	}
	
	g.Log().Info(ctx, "Success Removed ", countSuccess, " keys.")
	return countSuccess
}

func (s *sCache) RemoveCacheMatchedPattern(ctx context.Context, pattern string) error {
	// Retrieve and delete keys matching a pattern
	scanOption := gredis.ScanOption{
		Match: "*" + pattern + "*",
	}

	_, keys, err := g.Redis().Scan(ctx, 0, scanOption)
	if err != nil {
		g.Log().Error(ctx, "Failed to Scan Redis Cache By Pattern: ", scanOption.Match, err)
		err = gerror.NewCode(gcode.CodeOperationFailed, "Failed to Scan Redis Cache By Pattern: " + err.Error())
		return err
	}

	num, err := g.Redis().Del(ctx, keys...)
	if err != nil {
		g.Log().Error(ctx, "Failed to Remove Redis Caches By Keys: ", keys, err)
		err = gerror.NewCode(gcode.CodeOperationFailed, "Failed to Remove Redis Caches By Keys: " + err.Error())
		return err
	}

	g.Log().Info(ctx, "Removed Caches :", num)
	return nil
}