package ossfs

import (
	"bytes"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/pkg6/gfs"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"sync"
)

var (
	DefaultEndpoint = "oss-cn-hangzhou.aliyuncs.com"
)

type Adapter struct {
	Config *Config
	bucket string
	lock   *sync.Mutex
}

func New(config gfs.IAdapterConfig) gfs.IAdapter {
	return config.NewAdapter()
}

func NewOSS(config *Config) *Adapter {
	a := &Adapter{Config: config}
	if a.Config.Endpoint == "" {
		a.Config.Endpoint = DefaultEndpoint
	}
	a.lock = &sync.Mutex{}
	return a
}

func (a *Adapter) DiskName() string {
	return gfs.DiskNameOSS
}

func (a *Adapter) Client() (*oss.Client, error) {
	return oss.New(a.Config.Endpoint, a.Config.AccessKeyID, a.Config.AccessKeySecret, func(client *oss.Client) {
		if a.Config.Config != nil {
			client.Config = a.Config.Config
		}
	})
}

func (a *Adapter) OSSBucket() (*oss.Bucket, error) {
	client, err := a.Client()
	if err != nil {
		return nil, err
	}
	return client.Bucket(a.Config.UseBucket(a.bucket))
}

func (a *Adapter) CopyObject(srcObjectKey, destObjectKey string, isDelete bool) (bool, error) {
	a.lock.Lock()
	defer a.lock.Unlock()
	bucket, err := a.OSSBucket()
	if err != nil {
		return false, err
	}
	_, err = bucket.CopyObject(srcObjectKey, destObjectKey)
	if err != nil {
		return false, err
	}
	if isDelete {
		defer func() {
			_ = bucket.DeleteObject(srcObjectKey)
		}()
	}
	return true, nil
}

func (a *Adapter) Meta(path string) (header http.Header, err error) {
	bucket, err := a.OSSBucket()
	if err != nil {
		return header, err
	}
	return bucket.GetObjectMeta(path)
}

func (a *Adapter) Bucket(bucket string) gfs.IAdapter {
	a.bucket = bucket
	return a
}
func (a *Adapter) URL(path string) (*url.URL, error) {
	return a.Config.URL(path)
}
func (a *Adapter) Exist(path string) (bool, error) {
	a.lock.Lock()
	defer a.lock.Unlock()
	bucket, err := a.OSSBucket()
	if err != nil {
		return false, err
	}
	return bucket.IsObjectExist(path)
}
func (a *Adapter) WriteReader(path string, reader io.Reader) error {
	a.lock.Lock()
	defer a.lock.Unlock()
	bucket, err := a.OSSBucket()
	if err != nil {
		return err
	}
	return bucket.PutObject(path, reader)
}

func (a *Adapter) Write(path string, contents []byte) error {
	return a.WriteReader(path, bytes.NewReader(contents))
}

func (a *Adapter) WriteStream(path, resource string) error {
	a.lock.Lock()
	defer a.lock.Unlock()
	bucket, err := a.OSSBucket()
	if err != nil {
		return err
	}
	return bucket.PutObjectFromFile(path, resource)
}
func (a *Adapter) Update(path string, contents []byte) error {
	return a.Write(path, contents)
}

func (a *Adapter) UpdateStream(path, resource string) error {
	return a.WriteStream(path, resource)
}
func (a *Adapter) Read(path string) ([]byte, error) {
	a.lock.Lock()
	defer a.lock.Unlock()
	bucket, err := a.OSSBucket()
	if err != nil {
		return nil, err
	}
	object, err := bucket.GetObject(path)
	if err != nil {
		return nil, err
	}
	defer object.Close()
	contents, err := io.ReadAll(object)
	if err != nil {
		return nil, err
	}
	return contents, err
}

func (a *Adapter) Delete(path string) (int64, error) {
	a.lock.Lock()
	defer a.lock.Unlock()
	bucket, err := a.OSSBucket()
	if err != nil {
		return 0, err
	}
	if err = bucket.DeleteObject(path); err != nil {
		return 0, err
	}
	return 1, nil
}

func (a *Adapter) MimeType(path string) (string, error) {
	a.lock.Lock()
	defer a.lock.Unlock()
	meta, err := a.Meta(path)
	if err != nil {
		return "", err
	}
	return meta.Get(gfs.HeaderGetContentType), nil
}

func (a *Adapter) Size(path string) (int64, error) {
	a.lock.Lock()
	defer a.lock.Unlock()
	meta, err := a.Meta(path)
	if err != nil {
		return 0, err
	}
	i, err := strconv.ParseInt(meta.Get(gfs.HeaderGetLength), 10, 64)
	if err != nil {
		return 0, err
	}
	return i, nil
}
func (a *Adapter) Move(source, destination string) (bool, error) {
	return a.CopyObject(source, destination, true)
}

func (a *Adapter) Copy(source, destination string) (bool, error) {
	return a.CopyObject(source, destination, false)
}
