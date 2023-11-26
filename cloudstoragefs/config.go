package cloudstoragefs

import (
	"github.com/zzqqw/gfs"
	"google.golang.org/api/option"
	"net/url"
	"time"
)

var (
	DefaultDNS         = "https://storage.googleapis.com/"
	DefaultWithTimeout = time.Second * 50
)

type Config struct {
	CDN             string
	Bucket          string
	WithTimeout     time.Duration
	CredentialsFile string
	Option          []option.ClientOption
}

func (c *Config) NewAdapter() gfs.IAdapter {
	return NewGCS(c)
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
		c.CDN = DefaultDNS + c.Bucket
	}
	return url.Parse(c.CDN)
}
func (c *Config) UseBucket(bucket string) string {
	if bucket != "" {
		return bucket
	}
	return c.Bucket
}
