package cloudstoragefs

import (
	"cloud.google.com/go/storage"
	"context"
	"github.com/pkg6/gfs"
	"io"
	"net/url"
	"reflect"
	"sync"
	"testing"
)

func TestAdapter_Bucket(t *testing.T) {
	type fields struct {
		bucket             string
		Config             *Config
		ctx                context.Context
		lock               *sync.Mutex
		closeTimeoutCancel []func()
		storageClient      *storage.Client
	}
	type args struct {
		bucket string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   gfs.IAdapter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Adapter{
				bucket:             tt.fields.bucket,
				Config:             tt.fields.Config,
				ctx:                tt.fields.ctx,
				lock:               tt.fields.lock,
				closeTimeoutCancel: tt.fields.closeTimeoutCancel,
				storageClient:      tt.fields.storageClient,
			}
			if got := a.Bucket(tt.args.bucket); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bucket() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdapter_Client(t *testing.T) {
	type fields struct {
		bucket             string
		Config             *Config
		ctx                context.Context
		lock               *sync.Mutex
		closeTimeoutCancel []func()
		storageClient      *storage.Client
	}
	tests := []struct {
		name    string
		fields  fields
		want    *storage.Client
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Adapter{
				bucket:             tt.fields.bucket,
				Config:             tt.fields.Config,
				ctx:                tt.fields.ctx,
				lock:               tt.fields.lock,
				closeTimeoutCancel: tt.fields.closeTimeoutCancel,
				storageClient:      tt.fields.storageClient,
			}
			got, err := a.Client()
			if (err != nil) != tt.wantErr {
				t.Errorf("Client() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdapter_Close(t *testing.T) {
	type fields struct {
		bucket             string
		Config             *Config
		ctx                context.Context
		lock               *sync.Mutex
		closeTimeoutCancel []func()
		storageClient      *storage.Client
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Adapter{
				bucket:             tt.fields.bucket,
				Config:             tt.fields.Config,
				ctx:                tt.fields.ctx,
				lock:               tt.fields.lock,
				closeTimeoutCancel: tt.fields.closeTimeoutCancel,
				storageClient:      tt.fields.storageClient,
			}
			a.Close()
		})
	}
}

func TestAdapter_Copy(t *testing.T) {
	type fields struct {
		bucket             string
		Config             *Config
		ctx                context.Context
		lock               *sync.Mutex
		closeTimeoutCancel []func()
		storageClient      *storage.Client
	}
	type args struct {
		source      string
		destination string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Adapter{
				bucket:             tt.fields.bucket,
				Config:             tt.fields.Config,
				ctx:                tt.fields.ctx,
				lock:               tt.fields.lock,
				closeTimeoutCancel: tt.fields.closeTimeoutCancel,
				storageClient:      tt.fields.storageClient,
			}
			got, err := a.Copy(tt.args.source, tt.args.destination)
			if (err != nil) != tt.wantErr {
				t.Errorf("Copy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Copy() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdapter_CopyObject(t *testing.T) {
	type fields struct {
		bucket             string
		Config             *Config
		ctx                context.Context
		lock               *sync.Mutex
		closeTimeoutCancel []func()
		storageClient      *storage.Client
	}
	type args struct {
		source       string
		destination  string
		deleteSource bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Adapter{
				bucket:             tt.fields.bucket,
				Config:             tt.fields.Config,
				ctx:                tt.fields.ctx,
				lock:               tt.fields.lock,
				closeTimeoutCancel: tt.fields.closeTimeoutCancel,
				storageClient:      tt.fields.storageClient,
			}
			got, err := a.CopyObject(tt.args.source, tt.args.destination, tt.args.deleteSource)
			if (err != nil) != tt.wantErr {
				t.Errorf("CopyObject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CopyObject() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdapter_Delete(t *testing.T) {
	type fields struct {
		bucket             string
		Config             *Config
		ctx                context.Context
		lock               *sync.Mutex
		closeTimeoutCancel []func()
		storageClient      *storage.Client
	}
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Adapter{
				bucket:             tt.fields.bucket,
				Config:             tt.fields.Config,
				ctx:                tt.fields.ctx,
				lock:               tt.fields.lock,
				closeTimeoutCancel: tt.fields.closeTimeoutCancel,
				storageClient:      tt.fields.storageClient,
			}
			got, err := a.Delete(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Delete() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdapter_DiskName(t *testing.T) {
	type fields struct {
		bucket             string
		Config             *Config
		ctx                context.Context
		lock               *sync.Mutex
		closeTimeoutCancel []func()
		storageClient      *storage.Client
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Adapter{
				bucket:             tt.fields.bucket,
				Config:             tt.fields.Config,
				ctx:                tt.fields.ctx,
				lock:               tt.fields.lock,
				closeTimeoutCancel: tt.fields.closeTimeoutCancel,
				storageClient:      tt.fields.storageClient,
			}
			if got := a.DiskName(); got != tt.want {
				t.Errorf("DiskName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdapter_Exist(t *testing.T) {
	type fields struct {
		bucket             string
		Config             *Config
		ctx                context.Context
		lock               *sync.Mutex
		closeTimeoutCancel []func()
		storageClient      *storage.Client
	}
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Adapter{
				bucket:             tt.fields.bucket,
				Config:             tt.fields.Config,
				ctx:                tt.fields.ctx,
				lock:               tt.fields.lock,
				closeTimeoutCancel: tt.fields.closeTimeoutCancel,
				storageClient:      tt.fields.storageClient,
			}
			got, err := a.Exist(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("Exist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Exist() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdapter_GetMetadata(t *testing.T) {
	type fields struct {
		bucket             string
		Config             *Config
		ctx                context.Context
		lock               *sync.Mutex
		closeTimeoutCancel []func()
		storageClient      *storage.Client
	}
	type args struct {
		object string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *storage.ObjectAttrs
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Adapter{
				bucket:             tt.fields.bucket,
				Config:             tt.fields.Config,
				ctx:                tt.fields.ctx,
				lock:               tt.fields.lock,
				closeTimeoutCancel: tt.fields.closeTimeoutCancel,
				storageClient:      tt.fields.storageClient,
			}
			got, err := a.GetMetadata(tt.args.object)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMetadata() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMetadata() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdapter_MimeType(t *testing.T) {
	type fields struct {
		bucket             string
		Config             *Config
		ctx                context.Context
		lock               *sync.Mutex
		closeTimeoutCancel []func()
		storageClient      *storage.Client
	}
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Adapter{
				bucket:             tt.fields.bucket,
				Config:             tt.fields.Config,
				ctx:                tt.fields.ctx,
				lock:               tt.fields.lock,
				closeTimeoutCancel: tt.fields.closeTimeoutCancel,
				storageClient:      tt.fields.storageClient,
			}
			got, err := a.MimeType(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("MimeType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MimeType() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdapter_Move(t *testing.T) {
	type fields struct {
		bucket             string
		Config             *Config
		ctx                context.Context
		lock               *sync.Mutex
		closeTimeoutCancel []func()
		storageClient      *storage.Client
	}
	type args struct {
		source      string
		destination string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Adapter{
				bucket:             tt.fields.bucket,
				Config:             tt.fields.Config,
				ctx:                tt.fields.ctx,
				lock:               tt.fields.lock,
				closeTimeoutCancel: tt.fields.closeTimeoutCancel,
				storageClient:      tt.fields.storageClient,
			}
			got, err := a.Move(tt.args.source, tt.args.destination)
			if (err != nil) != tt.wantErr {
				t.Errorf("Move() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Move() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdapter_Read(t *testing.T) {
	type fields struct {
		bucket             string
		Config             *Config
		ctx                context.Context
		lock               *sync.Mutex
		closeTimeoutCancel []func()
		storageClient      *storage.Client
	}
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Adapter{
				bucket:             tt.fields.bucket,
				Config:             tt.fields.Config,
				ctx:                tt.fields.ctx,
				lock:               tt.fields.lock,
				closeTimeoutCancel: tt.fields.closeTimeoutCancel,
				storageClient:      tt.fields.storageClient,
			}
			got, err := a.Read(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Read() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdapter_Size(t *testing.T) {
	type fields struct {
		bucket             string
		Config             *Config
		ctx                context.Context
		lock               *sync.Mutex
		closeTimeoutCancel []func()
		storageClient      *storage.Client
	}
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Adapter{
				bucket:             tt.fields.bucket,
				Config:             tt.fields.Config,
				ctx:                tt.fields.ctx,
				lock:               tt.fields.lock,
				closeTimeoutCancel: tt.fields.closeTimeoutCancel,
				storageClient:      tt.fields.storageClient,
			}
			got, err := a.Size(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("Size() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Size() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdapter_StorageObject(t *testing.T) {
	type fields struct {
		bucket             string
		Config             *Config
		ctx                context.Context
		lock               *sync.Mutex
		closeTimeoutCancel []func()
		storageClient      *storage.Client
	}
	type args struct {
		object string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *storage.ObjectHandle
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Adapter{
				bucket:             tt.fields.bucket,
				Config:             tt.fields.Config,
				ctx:                tt.fields.ctx,
				lock:               tt.fields.lock,
				closeTimeoutCancel: tt.fields.closeTimeoutCancel,
				storageClient:      tt.fields.storageClient,
			}
			got, err := a.StorageObject(tt.args.object)
			if (err != nil) != tt.wantErr {
				t.Errorf("StorageObject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StorageObject() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdapter_StorageObjectTimeout(t *testing.T) {
	type fields struct {
		bucket             string
		Config             *Config
		ctx                context.Context
		lock               *sync.Mutex
		closeTimeoutCancel []func()
		storageClient      *storage.Client
	}
	type args struct {
		object string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *storage.ObjectHandle
		want1   context.Context
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Adapter{
				bucket:             tt.fields.bucket,
				Config:             tt.fields.Config,
				ctx:                tt.fields.ctx,
				lock:               tt.fields.lock,
				closeTimeoutCancel: tt.fields.closeTimeoutCancel,
				storageClient:      tt.fields.storageClient,
			}
			got, got1, err := a.StorageObjectTimeout(tt.args.object)
			if (err != nil) != tt.wantErr {
				t.Errorf("StorageObjectTimeout() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StorageObjectTimeout() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("StorageObjectTimeout() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestAdapter_URL(t *testing.T) {
	type fields struct {
		bucket             string
		Config             *Config
		ctx                context.Context
		lock               *sync.Mutex
		closeTimeoutCancel []func()
		storageClient      *storage.Client
	}
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *url.URL
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Adapter{
				bucket:             tt.fields.bucket,
				Config:             tt.fields.Config,
				ctx:                tt.fields.ctx,
				lock:               tt.fields.lock,
				closeTimeoutCancel: tt.fields.closeTimeoutCancel,
				storageClient:      tt.fields.storageClient,
			}
			got, err := a.URL(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("URL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("URL() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdapter_Update(t *testing.T) {
	type fields struct {
		bucket             string
		Config             *Config
		ctx                context.Context
		lock               *sync.Mutex
		closeTimeoutCancel []func()
		storageClient      *storage.Client
	}
	type args struct {
		path     string
		contents []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Adapter{
				bucket:             tt.fields.bucket,
				Config:             tt.fields.Config,
				ctx:                tt.fields.ctx,
				lock:               tt.fields.lock,
				closeTimeoutCancel: tt.fields.closeTimeoutCancel,
				storageClient:      tt.fields.storageClient,
			}
			if err := a.Update(tt.args.path, tt.args.contents); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAdapter_UpdateStream(t *testing.T) {
	type fields struct {
		bucket             string
		Config             *Config
		ctx                context.Context
		lock               *sync.Mutex
		closeTimeoutCancel []func()
		storageClient      *storage.Client
	}
	type args struct {
		path     string
		resource string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Adapter{
				bucket:             tt.fields.bucket,
				Config:             tt.fields.Config,
				ctx:                tt.fields.ctx,
				lock:               tt.fields.lock,
				closeTimeoutCancel: tt.fields.closeTimeoutCancel,
				storageClient:      tt.fields.storageClient,
			}
			if err := a.UpdateStream(tt.args.path, tt.args.resource); (err != nil) != tt.wantErr {
				t.Errorf("UpdateStream() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAdapter_Write(t *testing.T) {
	type fields struct {
		bucket             string
		Config             *Config
		ctx                context.Context
		lock               *sync.Mutex
		closeTimeoutCancel []func()
		storageClient      *storage.Client
	}
	type args struct {
		path     string
		contents []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Adapter{
				bucket:             tt.fields.bucket,
				Config:             tt.fields.Config,
				ctx:                tt.fields.ctx,
				lock:               tt.fields.lock,
				closeTimeoutCancel: tt.fields.closeTimeoutCancel,
				storageClient:      tt.fields.storageClient,
			}
			if err := a.Write(tt.args.path, tt.args.contents); (err != nil) != tt.wantErr {
				t.Errorf("Write() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAdapter_WriteReader(t *testing.T) {
	type fields struct {
		bucket             string
		Config             *Config
		ctx                context.Context
		lock               *sync.Mutex
		closeTimeoutCancel []func()
		storageClient      *storage.Client
	}
	type args struct {
		path   string
		reader io.Reader
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Adapter{
				bucket:             tt.fields.bucket,
				Config:             tt.fields.Config,
				ctx:                tt.fields.ctx,
				lock:               tt.fields.lock,
				closeTimeoutCancel: tt.fields.closeTimeoutCancel,
				storageClient:      tt.fields.storageClient,
			}
			if err := a.WriteReader(tt.args.path, tt.args.reader); (err != nil) != tt.wantErr {
				t.Errorf("WriteReader() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAdapter_WriteStream(t *testing.T) {
	type fields struct {
		bucket             string
		Config             *Config
		ctx                context.Context
		lock               *sync.Mutex
		closeTimeoutCancel []func()
		storageClient      *storage.Client
	}
	type args struct {
		path     string
		resource string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Adapter{
				bucket:             tt.fields.bucket,
				Config:             tt.fields.Config,
				ctx:                tt.fields.ctx,
				lock:               tt.fields.lock,
				closeTimeoutCancel: tt.fields.closeTimeoutCancel,
				storageClient:      tt.fields.storageClient,
			}
			if err := a.WriteStream(tt.args.path, tt.args.resource); (err != nil) != tt.wantErr {
				t.Errorf("WriteStream() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		config gfs.IAdapterConfig
	}
	tests := []struct {
		name string
		args args
		want gfs.IAdapter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewGCS(t *testing.T) {
	type args struct {
		config *Config
	}
	tests := []struct {
		name string
		args args
		want *Adapter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGCS(tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGCS() = %v, want %v", got, tt.want)
			}
		})
	}
}
