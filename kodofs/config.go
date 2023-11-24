package kodofs

import (
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/zzqqw/gfs"
	"net/url"
)

type Config struct {
	CDN                  string
	AccessKey, SecretKey string
	Bucket               string

	Policy *storage.PutPolicy
	Config *storage.Config
}

func (c *Config) New() gfs.IAdapter {
	return NewKoDo(c)
}

func (c *Config) URL(path string) (*url.URL, error) {
	return gfs.PublicURLMake(c.CDN, path)
}
func (c *Config) UseBucket(bucket string) string {
	if bucket != "" {
		return bucket
	}
	return c.Bucket
}
