package bosfs

import (
	"github.com/pkg6/gfs"
	"net/url"
	"reflect"
	"testing"
)

func TestConfig_NewAdapter(t *testing.T) {
	type fields struct {
		CDN              string
		Ak               string
		Sk               string
		Endpoint         string
		RedirectDisabled bool
		Bucket           string
	}
	tests := []struct {
		name   string
		fields fields
		want   gfs.IAdapter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				CDN:              tt.fields.CDN,
				Ak:               tt.fields.Ak,
				Sk:               tt.fields.Sk,
				Endpoint:         tt.fields.Endpoint,
				RedirectDisabled: tt.fields.RedirectDisabled,
				Bucket:           tt.fields.Bucket,
			}
			if got := c.NewAdapter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAdapter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_URL(t *testing.T) {
	type fields struct {
		CDN              string
		Ak               string
		Sk               string
		Endpoint         string
		RedirectDisabled bool
		Bucket           string
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
			c := &Config{
				CDN:              tt.fields.CDN,
				Ak:               tt.fields.Ak,
				Sk:               tt.fields.Sk,
				Endpoint:         tt.fields.Endpoint,
				RedirectDisabled: tt.fields.RedirectDisabled,
				Bucket:           tt.fields.Bucket,
			}
			got, err := c.URL(tt.args.path)
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

func TestConfig_UseBucket(t *testing.T) {
	type fields struct {
		CDN              string
		Ak               string
		Sk               string
		Endpoint         string
		RedirectDisabled bool
		Bucket           string
	}
	type args struct {
		bucket string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				CDN:              tt.fields.CDN,
				Ak:               tt.fields.Ak,
				Sk:               tt.fields.Sk,
				Endpoint:         tt.fields.Endpoint,
				RedirectDisabled: tt.fields.RedirectDisabled,
				Bucket:           tt.fields.Bucket,
			}
			if got := c.UseBucket(tt.args.bucket); got != tt.want {
				t.Errorf("UseBucket() = %v, want %v", got, tt.want)
			}
		})
	}
}
