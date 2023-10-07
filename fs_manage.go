package gfs

import (
	"fmt"
	"io"
	"net/url"
	"sync"
)

type FsManage struct {
	disk         string
	disks        []string
	diskAdapters map[string]IAdapter
	l            *sync.Mutex
}

func NewConfig(config IConfig) *FsManage {
	fs := &FsManage{disk: config.Disk(), l: &sync.Mutex{}}
	adapters := config.Adapters()
	for s, adapter := range adapters {
		fs.Extend(adapter, s)
	}
	return fs
}

func New() *FsManage {
	return &FsManage{
		diskAdapters: make(map[string]IAdapter),
		l:            &sync.Mutex{},
	}
}

// Extend 扩展
func (f *FsManage) Extend(adapter IAdapter, names ...string) *FsManage {
	f.l.Lock()
	defer f.l.Unlock()
	var name string
	if len(names) > 0 {
		name = names[0]
	} else {
		name = adapter.DiskName()
	}
	f.disks = append(f.disks, name)
	f.diskAdapters[name] = adapter
	if f.disk == "" {
		f.disk = name
	}
	return f
}

// DiskGet 获取注册所有的驱动
func (f *FsManage) DiskGet() []string {
	return f.disks
}

// DiskExist 判断驱动是否存在
func (f *FsManage) DiskExist(disk string) bool {
	_, ok := f.diskAdapters[disk]
	return ok
}

// Adapter 根据驱动名称找到适配器
func (f *FsManage) Adapter(disk string) (IAdapter, error) {
	if disk != "" {
		f.disk = disk
	} else {
		f.disk = f.disks[0]
	}
	if adapter, ok := f.diskAdapters[f.disk]; ok {
		return adapter, nil
	}
	return nil, fmt.Errorf("unable to find %s disk", f.disk)
}

func (f *FsManage) URL(path string) (*url.URL, error) {
	adapter, err := f.Adapter("")
	if err != nil {
		return nil, err
	}
	return adapter.URL(path)
}

func (f *FsManage) Exist(path string) (bool, error) {
	adapter, err := f.Adapter("")
	if err != nil {
		return false, err
	}
	return adapter.Exist(path)
}

func (f *FsManage) WriteReader(path string, reader io.Reader) error {
	adapter, err := f.Adapter("")
	if err != nil {
		return err
	}
	return adapter.WriteReader(path, reader)
}

func (f *FsManage) Write(path string, contents []byte) error {
	adapter, err := f.Adapter("")
	if err != nil {
		return err
	}
	return adapter.Write(path, contents)
}

func (f *FsManage) WriteStream(path, resource string) error {
	adapter, err := f.Adapter("")
	if err != nil {
		return err
	}
	return adapter.WriteStream(path, resource)
}

func (f *FsManage) Update(path string, contents []byte) error {
	adapter, err := f.Adapter("")
	if err != nil {
		return err
	}
	return adapter.Update(path, contents)
}

func (f *FsManage) UpdateStream(path, resource string) error {
	adapter, err := f.Adapter("")
	if err != nil {
		return err
	}
	return adapter.UpdateStream(path, resource)
}

func (f *FsManage) Read(path string) ([]byte, error) {
	adapter, err := f.Adapter("")
	if err != nil {
		return nil, err
	}
	return adapter.Read(path)
}

func (f *FsManage) Delete(path string) (int64, error) {
	adapter, err := f.Adapter("")
	if err != nil {
		return 0, err
	}
	return adapter.Delete(path)
}

func (f *FsManage) MimeType(path string) (string, error) {
	adapter, err := f.Adapter("")
	if err != nil {
		return "", err
	}
	return adapter.MimeType(path)
}

func (f *FsManage) Size(path string) (int64, error) {
	adapter, err := f.Adapter("")
	if err != nil {
		return 0, err
	}
	return adapter.Size(path)
}

func (f *FsManage) Move(source, destination string) (bool, error) {
	adapter, err := f.Adapter("")
	if err != nil {
		return false, err
	}
	return adapter.Move(source, destination)
}

func (f *FsManage) Copy(source, destination string) (bool, error) {
	adapter, err := f.Adapter("")
	if err != nil {
		return false, err
	}
	return adapter.Copy(source, destination)
}
