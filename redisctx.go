package redisctx

import (
    "context"
    "encoding/json"
    "fmt"
    "strings"

    "github.com/redis/go-redis/v9"
)

type RedisContext struct {
    Client *redis.Client
    Key    string
}

func (ctx *RedisContext) callLuaFunction(args ...interface{}) func() ([]byte, error) {
    return func() ([]byte, error) {
        result, err := ctx.Client.Do(context.Background(), args...).Result()
        if err != nil {
            fmt.Printf("error executing FCALL for Lua function: %s\n", err)
            return nil, err
        }

        response := LuaResponse{
            LuaResponse: fmt.Sprintf("%v", result),
        }

        jsonData, err := json.Marshal(response)
        if err != nil {
            fmt.Printf("error marshalling JSON: %v\n", err)
            return nil, err
        }
        return jsonData, nil
    }
}

func (ctx *RedisContext) parseKey() (string, string) {
    parts := strings.SplitN(ctx.key, ":", 3)
    if len(parts) != 3 {
        return "", ""
    }
    asubID := parts[1]
    channelName := parts[2]
    return asubID, channelName
}
