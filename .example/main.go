package main

import (
	"fmt"
	"github.com/zzqqw/gfs"
	"github.com/zzqqw/gfs/bosfs"
	"github.com/zzqqw/gfs/cloudstoragefs"
	"github.com/zzqqw/gfs/cosfs"
	"github.com/zzqqw/gfs/kodofs"
	"github.com/zzqqw/gfs/localfs"
	"github.com/zzqqw/gfs/ossfs"
	"google.golang.org/api/option"
	"strings"
)

var (
	root   = "./.test_data/"
	local  gfs.IAdapter
	oss    gfs.IAdapter
	google gfs.IAdapter
	kodo   gfs.IAdapter
	cos    gfs.IAdapter
	bos    gfs.IAdapter
)

func init() {
	local = localfs.New(&localfs.Config{})
	oss = ossfs.New(&ossfs.Config{
		Bucket:          "test",
		Endpoint:        "oss-cn-hangzhou.aliyuncs.com",
		AccessKeyID:     "*******************",
		AccessKeySecret: "**************",
	})
	google = cloudstoragefs.New(&cloudstoragefs.Config{
		Bucket: "test bucket",
		Option: []option.ClientOption{
			option.WithCredentialsFile("CredentialsFile.json"),
		},
	})
	kodo = kodofs.New(&kodofs.Config{
		AccessKey: "AccessKey",
		SecretKey: "SecretKey",
		Bucket:    "test bucket",
	})
	//Create a bucket automatically generated URL
	cos = cosfs.New(&cosfs.Config{
		BucketURL: "https://bucket-id.cos.ap-beijing.myqcloud.com",
		SecretID:  "SecretID",
		SecretKey: "SecretKey",
	})
	bos = bosfs.New(&bosfs.Config{
		Endpoint: bosfs.BJEndpoint,
		Ak:       "Ak",
		Sk:       "Sk",
		Bucket:   "test bucket",
	})
}

func main() {
	adapters := gfs.New()
	adapters.Extend(local)
	var err error
	err = adapters.WriteReader(root+"4.txt", strings.NewReader("test"))
	fmt.Println(err)
	err = adapters.Disk(gfs.DiskNameLocal).WriteReader(root+"5.txt", strings.NewReader("test"))
	fmt.Println(err)
	//Write file
	err = adapters.Write(root+"1.txt", []byte("test data"))
	fmt.Println(err)
	//Write data from resource file
	err = adapters.WriteStream(root+"2.txt", root+"/1.txt")
	fmt.Println(err)
	//Update file
	err = adapters.Update(root+"1.txt", []byte("test update data"))
	fmt.Println(err)
	//Update data from resource file
	err = adapters.UpdateStream(root+"2.txt", root+"/1.txt")
	fmt.Println(err)
	exists, err := adapters.Exist(root + "2.txt")
	if err != nil {
		return
	}
	fmt.Println(exists)
	//Read file
	read, err := adapters.Read(root + "2.txt")
	fmt.Println(read, err)
	//Get file mime type
	mimeType, err := adapters.MimeType(root + "2.txt")
	fmt.Println(mimeType, err)
	//Get file size
	size, err := adapters.Size(root + "2.txt")
	fmt.Println(size, err)
	//Move file
	_, err = adapters.Move(root+"1.txt", root+"4.txt")
	fmt.Println(err)
	//Copy file
	_, err = adapters.Copy(root+"2.txt", root+"5.txt")
	fmt.Println(err)
}
