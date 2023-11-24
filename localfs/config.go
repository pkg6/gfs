package localfs

import (
	"github.com/zzqqw/gfs"
	"net/url"
)

type Config struct {
	CDN string
}

func (c *Config) New() gfs.IAdapter {
	return NewLocal(c)
}
func (c *Config) URL(path string) (*url.URL, error) {
	return gfs.PublicURLMake(c.CDN, path)
}
func (c *Config) UseBucket(bucket string) string {
	return bucket
}
