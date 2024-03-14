package cloudstoragefs

import (
	"github.com/pkg6/gfs"
	"google.golang.org/api/option"
	"net/url"
	"time"
)

var (
	DefaultDNS         = "https://storage.googleapis.com/"
	DefaultWithTimeout = time.Second * 50
)

type Config struct {
	CDN             string        `json:"cdn" xml:"CDN" yaml:"CDN"`
	Bucket          string        `json:"bucket" xml:"Bucket" yaml:"Bucket"`
	WithTimeout     time.Duration `json:"with_timeout" xml:"WithTimeout" yaml:"WithTimeout"`
	CredentialsFile string        `json:"credentials_file" xml:"CredentialsFile" yaml:"CredentialsFile"`
	CredentialsJSON string        `json:"credentials_json" xml:"CredentialsJSON" yaml:"CredentialsJSON"`
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
