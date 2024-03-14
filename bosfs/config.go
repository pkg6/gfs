package bosfs

import (
	"github.com/pkg6/gfs"
	"net/url"
)

type Config struct {
	CDN string `json:"cdn" xml:"CDN" yaml:"CDN"`
	Ak  string `json:"ak" xml:"Ak" yaml:"Ak"`
	Sk  string `json:"sk" xml:"Sk" yaml:"Sk"`
	//https://cloud.baidu.com/doc/BOS/s/Ojwvyrpgd
	Endpoint         string `json:"endpoint" xml:"Endpoint" yaml:"Endpoint"`
	RedirectDisabled bool   `json:"redirect_disabled" xml:"RedirectDisabled" yaml:"RedirectDisabled"`
	Bucket           string `json:"bucket" xml:"Bucket" yaml:"Bucket"`
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
