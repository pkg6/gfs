package config

import (
	"github.com/zzqqw/gfs"
	"github.com/zzqqw/gfs/bosfs"
	"github.com/zzqqw/gfs/cloudstoragefs"
	"github.com/zzqqw/gfs/cosfs"
	"github.com/zzqqw/gfs/kodofs"
	"github.com/zzqqw/gfs/localfs"
	"github.com/zzqqw/gfs/ossfs"
	"reflect"
)

type AdapterConfig struct {
	Default string
	Adapter *Adapter
}

type Adapter struct {
	LOCAL        *localfs.Config
	OSS          *ossfs.Config
	BOS          *bosfs.Config
	COS          *cosfs.Config
	KODO         *kodofs.Config
	CloudStorage *cloudstoragefs.Config
}

func (a *AdapterConfig) Disk() string {
	return a.Default
}

func (a *AdapterConfig) Adapters() map[string]gfs.IAdapter {
	sortIAdapters := map[string]gfs.IAdapter{}
	adapters := map[string]gfs.IAdapter{}
	var sortNames []string
	if a.Adapter == nil {
		return sortIAdapters
	}
	v := reflect.ValueOf(a.Adapter)
	for i := 0; i < v.Elem().NumField(); i++ {
		e := v.Elem().Field(i)
		if !e.IsZero() {
			if c, ok := e.Interface().(gfs.IAdapterConfig); ok {
				adapter := c.New()
				name := adapter.DiskName()
				adapters[name] = adapter
				sortNames = append(sortNames, name)
			}
		}
	}
	for _, name := range sortNames {
		if ad, ok := adapters[name]; ok {
			sortIAdapters[name] = ad
		}
	}
	return sortIAdapters
}
