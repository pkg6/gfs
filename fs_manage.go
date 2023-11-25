package gfs

import (
	"fmt"
	"io"
	"net/url"
	"reflect"
	"strings"
	"sync"
)

var (
	ConfigPtrTag                 = "gfs"
	ConfigPtrSplitTagDefaultDisk = "default"
)

type FsManage struct {
	disk         string
	disks        []string
	diskAdapters map[string]IAdapter
	l            *sync.Mutex
}

func NewConfig(config any) (*FsManage, error) {
	fs := &FsManage{l: &sync.Mutex{}}
	err := fs.ExtendConfigPtr(config)
	return fs, err
}

func New() *FsManage {
	return &FsManage{
		diskAdapters: make(map[string]IAdapter),
		l:            &sync.Mutex{},
	}
}

func (f *FsManage) ExtendConfigPtr(config any) error {
	v := reflect.ValueOf(config)
	t := reflect.TypeOf(config)
	if t.Kind() == reflect.Ptr {
		for i := 0; i < v.Elem().NumField(); i++ {
			e := v.Elem().Field(i)
			if !e.IsZero() {
				if fsConfig, ok := e.Interface().(IAdapterConfig); ok {
					var diskName string
					gfsName := t.Elem().Field(i).Tag.Get(ConfigPtrTag)
					if gfsName == "" {
						diskName = t.Elem().Field(i).Name
					} else {
						split := strings.Split(gfsName, ",")
						diskName = split[0]
						if len(split) > 2 && split[1] == ConfigPtrSplitTagDefaultDisk {
							if f.disk == "" {
								f.disk = diskName
							}
						}
					}
					f.Extend(fsConfig.NewAdapter(), diskName)
				}
			}
		}
		return nil
	}
	return fmt.Errorf("the data type is incorrect %v", config)
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

// Disks 获取注册所有的驱动
func (f *FsManage) Disks() []string {
	return f.disks
}

// DiskExist 判断驱动是否存在
func (f *FsManage) DiskExist(disk string) bool {
	_, ok := f.diskAdapters[disk]
	return ok
}
func (f *FsManage) Disk(disk string) string {
	if disk != "" {
		f.disk = disk
	}
	if f.disk == "" {
		f.disk = f.disks[0]
	}
	return f.disk
}

// Adapter 根据驱动名称找到适配器
func (f *FsManage) Adapter(disk string) (IAdapter, error) {
	if adapter, ok := f.diskAdapters[f.Disk(disk)]; ok {
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
