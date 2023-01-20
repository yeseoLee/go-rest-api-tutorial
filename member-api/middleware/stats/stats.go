package stats

import (
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type (
	Stats struct {
		Uptime       time.Time      `json:"uptime"`
		RequestCount uint64         `json:"requestCount"`
		Statuses     map[string]int `json:"statuses"`
		APIs         map[string]int `json:"apis"`
		mutex        sync.RWMutex
	}
)

type StatsConfig struct {
	Skipper middleware.Skipper
}

var once sync.Once
var instance *Stats

var DefaultStatsConfig = StatsConfig{
	Skipper: middleware.DefaultSkipper,
}

func newStats() *Stats {
	return &Stats{
		Uptime:   time.Now(),
		Statuses: map[string]int{},
		APIs:     map[string]int{},
	}
}

func GetInstance() *Stats {
	once.Do(func() {
		instance = newStats()
	})
	return instance
}

func (s *Stats) Middleware() echo.MiddlewareFunc {
	return s.MiddlewareWithConfig(DefaultStatsConfig)
}

func (s *Stats) MiddlewareWithConfig(config StatsConfig) echo.MiddlewareFunc {
	if config.Skipper == nil {
		config.Skipper = DefaultStatsConfig.Skipper
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// skip
			if config.Skipper(c) {
				return next(c)
			}

			// next는 락 잡기전에 호출
			err := next(c)

			s.mutex.Lock()
			defer s.mutex.Unlock()
			s.RequestCount++
			status := strconv.Itoa(c.Response().Status)
			api := c.Path()
			s.Statuses[status]++ // 상태코드 종류
			s.APIs[api]++        // API 종류
			return err
		}
	}
}

// Handle is the endpoint to get stats.
func (s *Stats) Handle(c echo.Context) error {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return c.JSON(http.StatusOK, s)
}
