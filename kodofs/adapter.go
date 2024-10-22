package kodofs

import (
	"bytes"
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/pkg6/gfs"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

type Adapter struct {
	bucket string
	Config *Config
	lock   *sync.Mutex
}

func New(config gfs.IAdapterConfig) gfs.IAdapter {
	return config.NewAdapter()
}

func NewKoDo(config *Config) *Adapter {
	a := &Adapter{Config: config}
	a.lock = &sync.Mutex{}
	return a
}

func (a *Adapter) Mac() *qbox.Mac {
	return qbox.NewMac(a.Config.AccessKey, a.Config.SecretKey)
}
func (a *Adapter) Bucket(bucket string) gfs.IAdapter {
	a.bucket = bucket
	return a
}

func (a *Adapter) UploadToken() string {
	if a.Config.Policy == nil {
		a.Config.Policy = &storage.PutPolicy{}
	}
	a.Config.Policy.Scope = a.Config.UseBucket(a.bucket)
	return a.Config.Policy.UploadToken(a.Mac())
}
func (a *Adapter) StorageConfig() *storage.Config {
	if a.Config.Config == nil {
		a.Config.Config = &storage.Config{}
	}
	if a.Config.Config.Region == nil {
		a.Config.Config.Region = &storage.ZoneHuadong
	}
	return a.Config.Config
}
func (a *Adapter) BucketManager() *storage.BucketManager {
	return storage.NewBucketManager(a.Mac(), a.StorageConfig())
}

func (a *Adapter) BucketManagerBatch(operations []string) error {
	rets, err := a.BucketManager().Batch(operations)
	if err != nil {
		return err
	}
	for _, ret := range rets {
		if ret.Code != http.StatusOK {
			return fmt.Errorf("BucketManagerBatch err=%v", ret.Data.Error)
		}
	}
	return nil
}

func (a *Adapter) Stat(path string) (info storage.FileInfo, err error) {
	return a.BucketManager().Stat(a.Config.UseBucket(a.bucket), path)
}
func (a *Adapter) URL(path string) (*url.URL, error) {
	return a.Config.URL(path)
}
func (a *Adapter) Exist(path string) (bool, error) {
	stat, err := a.BucketManager().Stat(a.Config.UseBucket(a.bucket), path)
	if stat.Md5 != "" && err == nil {
		return true, nil
	}
	return false, err
}

func (a *Adapter) WriteReader(path string, reader io.Reader) error {
	contents, err := io.ReadAll(reader)
	if err != nil {
		return err
	}
	return a.Write(path, contents)
}

func (a *Adapter) Write(path string, contents []byte) error {
	formUploader := storage.NewFormUploader(a.StorageConfig())
	dataLen := int64(len(contents))
	return formUploader.Put(context.Background(),
		&storage.PutRet{},
		a.UploadToken(),
		path,
		bytes.NewReader(contents),
		dataLen,
		&storage.PutExtra{},
	)
}

func (a *Adapter) WriteStream(path, resource string) error {
	var err error
	if strings.HasPrefix(resource, "http") {
		_, err = a.BucketManager().Fetch(resource, a.Config.UseBucket(a.bucket), path)
	} else {
		formUploader := storage.NewFormUploader(a.StorageConfig())
		err = formUploader.PutFile(context.Background(),
			&storage.PutRet{},
			a.UploadToken(),
			path,
			resource,
			&storage.PutExtra{},
		)
	}
	return err
}

func (a *Adapter) Read(path string) ([]byte, error) {
	u, err := a.URL(path)
	if err != nil {
		return nil, err
	}
	return gfs.PublicURLRead(u.String())
}

func (a *Adapter) Delete(path string) (int64, error) {
	err := a.BucketManagerBatch([]string{storage.URIDelete(a.Config.UseBucket(a.bucket), path)})
	if err != nil {
		return 0, err
	}
	return 1, err
}

func (a *Adapter) Size(path string) (int64, error) {
	stat, err := a.Stat(path)
	if err != nil {
		return 0, err
	}
	return stat.Fsize, nil
}

func (a *Adapter) Update(path string, contents []byte) error {
	return fmt.Errorf("update implement me")
}

func (a *Adapter) UpdateStream(path, resource string) error {
	return fmt.Errorf("updateStream implement me")
}

func (a *Adapter) MimeType(path string) (string, error) {
	stat, err := a.Stat(path)
	if err != nil {
		return "", err
	}
	return stat.MimeType, nil
}

func (a *Adapter) Move(source, destination string) (bool, error) {
	err := a.BucketManagerBatch([]string{storage.URIMove(a.Config.UseBucket(a.bucket), source, a.Config.UseBucket(a.bucket), destination, true)})
	if err != nil {
		return false, err
	}
	return true, err
}

func (a *Adapter) Copy(source, destination string) (bool, error) {
	err := a.BucketManagerBatch([]string{storage.URICopy(a.Config.UseBucket(a.bucket), source, a.Config.UseBucket(a.bucket), destination, true)})
	if err != nil {
		return false, err
	}
	return true, err
}

func (a *Adapter) DiskName() string {
	return gfs.DiskNameQiNiuKoDo
}
