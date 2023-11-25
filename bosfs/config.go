package bosfs

import (
	"github.com/zzqqw/gfs"
	"net/url"
)

type Config struct {
	CDN string
	Ak  string
	Sk  string
	//https://cloud.baidu.com/doc/BOS/s/Ojwvyrpgd
	Endpoint         string
	RedirectDisabled bool
	Bucket           string
}

func (c *Config) NewAdapter() gfs.IAdapter {
	return NewBOS(c)
}

func (c *Config) URL(path string) (*url.URL, error) {
	bucketUrl, err := gfs.BucketURLMake(c.CDN, c.Endpoint, c.Bucket)
	if err != nil {
		return nil, err
	}
	return gfs.PublicURLMake(bucketUrl.String(), path)
}

func (c *Config) UseBucket(bucket string) string {
	if bucket != "" {
		return bucket
	}
	return c.Bucket
}
