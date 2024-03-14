package cosfs

import (
	"github.com/pkg6/gfs"
	"net/url"
)

type Config struct {
	CDN string `json:"cdn" xml:"CDN" yaml:"CDN"`
	//  https://console.cloud.tencent.com/cos5/bucket
	BucketURL string `json:"bucket_url" xml:"BucketURL" yaml:"BucketURL"`
	//  https://cloud.tencent.com/document/product/598/37140
	SecretID  string `json:"secret_id" xml:"SecretID" yaml:"SecretID"`
	SecretKey string `json:"secret_key" xml:"SecretKey" yaml:"SecretKey"`
}

func (c *Config) NewAdapter() gfs.IAdapter {
	return NewCOS(c)
}

func (c *Config) URL(path string) (*url.URL, error) {
	bucketUrl, err := c.BucketUrl()
	if err != nil {
		return nil, err
	}
	return gfs.PublicURLMake(bucketUrl.String(), path)
}

func (c *Config) BucketUrl() (*url.URL, error) {
	if c.CDN == "" {
		c.CDN = c.BucketURL
	}
	return url.Parse(c.CDN)
}

func (c *Config) UseBucket(bucket string) string {
	if bucket != "" {
		return bucket
	}
	return c.BucketURL
}
