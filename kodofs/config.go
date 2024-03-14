package kodofs

import (
	"github.com/pkg6/gfs"
	"github.com/qiniu/go-sdk/v7/storage"
	"net/url"
)

type Config struct {
	CDN       string             `json:"cdn" xml:"CDN" yaml:"CDN"`
	AccessKey string             `json:"access_key" xml:"AccessKey" yaml:"AccessKey"`
	SecretKey string             `json:"secret_key" xml:"SecretKey" yaml:"SecretKey"`
	Bucket    string             `json:"bucket" xml:"Bucket" yaml:"Bucket"`
	Policy    *storage.PutPolicy `json:"policy" xml:"Policy" yaml:"Policy"`
	Config    *storage.Config    `json:"config" xml:"Config" yaml:"Config"`
}

func (c *Config) NewAdapter() gfs.IAdapter {
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
