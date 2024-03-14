package gfs

import (
	"net/url"
	"reflect"
	"testing"
)

func TestBucketURLMake(t *testing.T) {
	type args struct {
		cdn      string
		endpoint string
		bucket   string
	}
	tests := []struct {
		name    string
		args    args
		want    *url.URL
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BucketURLMake(tt.args.cdn, tt.args.endpoint, tt.args.bucket)
			if (err != nil) != tt.wantErr {
				t.Errorf("BucketURLMake() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BucketURLMake() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPublicURLMake(t *testing.T) {
	type args struct {
		domain string
		key    string
	}
	tests := []struct {
		name    string
		args    args
		want    *url.URL
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PublicURLMake(tt.args.domain, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("PublicURLMake() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PublicURLMake() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPublicURLRead(t *testing.T) {
	type args struct {
		publicURL string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PublicURLRead(tt.args.publicURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("PublicURLRead() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PublicURLRead() got = %v, want %v", got, tt.want)
			}
		})
	}
}
