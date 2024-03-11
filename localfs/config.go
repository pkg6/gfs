package localfs

import (
	"github.com/pkg6/gfs"
	"net/url"
)

type Config struct {
	CDN string
}

func (c *Config) NewAdapter() gfs.IAdapter {
	return NewLocal(c)
}
func (c *Config) URL(path string) (*url.URL, error) {
	return gfs.PublicURLMake(c.CDN, path)
}
func (c *Config) UseBucket(bucket string) string {
	return bucket
}
