package middleware

import (
	"context"
	"errors"
	"net/http"
	"time"

	"go.uber.org/zap"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
)

type LimitConfig struct {
	// GenerationKey generates key according to the business. The following CheckOrMark query generates it.
	GenerationKey func(c *gin.Context) string
	// Check function, the user can modify the specific logic, which is more flexible
	CheckOrMark func(key string, expire int, limit int) error
	// Expire key expiration time
	Expire int
	// Limit cycle time
	Limit int
}

func (l LimitConfig) LimitWithTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := l.CheckOrMark(l.GenerationKey(c), l.Expire, l.Limit); err != nil {
			c.JSON(http.StatusOK, gin.H{"code": response.ERROR, "msg": err})
			c.Abort()
			return
		} else {
			c.Next()
		}
	}
}

// DefaultGenerationKey generates key by default
func DefaultGenerationKey(c *gin.Context) string {
	return "GVA_Limit" + c.ClientIP()
}

func DefaultCheckOrMark(key string, expire int, limit int) (err error) {
	// Determine whether to open redis
	if global.GVA_REDIS == nil {
		return err
	}
	if err = SetLimitWithTime(key, limit, time.Duration(expire)*time.Second); err != nil {
		global.GVA_LOG.Error("limit", zap.Error(err))
	}
	return err
}

func DefaultLimit() gin.HandlerFunc {
	return LimitConfig{
		GenerationKey: DefaultGenerationKey,
		CheckOrMark:   DefaultCheckOrMark,
		Expire:        global.GVA_CONFIG.System.LimitTimeIP,
		Limit:         global.GVA_CONFIG.System.LimitCountIP,
	}.LimitWithTime()
}

// SetLimitWithTime sets the number of visits
func SetLimitWithTime(key string, limit int, expiration time.Duration) error {
	count, err := global.GVA_REDIS.Exists(context.Background(), key).Result()
	if err != nil {
		return err
	}
	if count == 0 {
		pipe := global.GVA_REDIS.TxPipeline()
		pipe.Incr(context.Background(), key)
		pipe.Expire(context.Background(), key, expiration)
		_, err = pipe.Exec(context.Background())
		return err
	} else {
		// times
		if times, err := global.GVA_REDIS.Get(context.Background(), key).Int(); err != nil {
			return err
		} else {
			if times >= limit {
				if t, err := global.GVA_REDIS.PTTL(context.Background(), key).Result(); err != nil {
					return errors.New("The request is too frequent, please try again later")
				} else {
					return errors.New("The request is too frequent, please try " + t.String() + " seconds later")
				}
			} else {
				return global.GVA_REDIS.Incr(context.Background(), key).Err()
			}
		}
	}
}
