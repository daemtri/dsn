package dsn

import (
	"net/url"
	"strconv"
	"time"
)

// Config 表示一个DSN解析后的配置
type Config struct {
	url *url.URL
	*Values
}

// Parse 解析dsn字符串并返回DSN实例
// eg. nsq://user:pass@localhost:4161/test
func Parse(dataSourceName string) (*Config, error) {

	parsed, err := url.Parse(dataSourceName)
	if err != nil {
		return nil, err
	}
	c := &Config{
		parsed,
		&Values{parsed.Query()},
	}
	return c, nil
}

// Host 返回Host信息
func (c *Config) Host() string {
	return c.url.Host
}

// Scheme 返回资源协议
func (c *Config) Scheme() string {
	return c.url.Scheme
}

// Path 返回资源路径
func (c *Config) Path() string {
	return c.url.Path
}

// User 返回用户信息
func (c *Config) User() *url.Userinfo {
	return c.url.User
}

// Values 表示 DSN Values
type Values struct {
	url.Values
}

// Int 返回querystring参数param的int值
// 注意：当param的value不能解析为int的时候，将会直接返回默认值并且不会报错
func (v *Values) Int(param string, def int) int {
	value := v.Get(param)
	if value == "" {
		return def
	}
	if i, err := strconv.Atoi(value); err == nil {
		return i
	}
	return def
}

// String 返回querystring参数param的字符串的值
func (v *Values) String(param string, def string) string {
	value := v.Get(param)
	if value == "" {
		return def
	}
	return value
}

// Bool 返回querystring参数param的布尔值
// 注意：当参数不能解析为布尔值的时候，将直接返回默认值并且不会报错
func (v *Values) Bool(param string, def bool) bool {
	value := v.Get(param)
	if value == "" {
		return def
	}
	result, err := strconv.ParseBool(value)
	if err != nil {
		return def
	}
	return result
}

// Duration returns time.Duration value
// 注意：当param的value不合法时，将不会报告错误，并直接返回默认值
func (v *Values) Duration(param string, def time.Duration) time.Duration {
	value := v.Get(param)
	if value == "" {
		return def
	}
	if du, err := time.ParseDuration(value); err == nil {
		return du
	}

	return def
}
