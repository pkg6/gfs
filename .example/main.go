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
		Bucket: "test",
		//Endpoint:        "oss-cn-hangzhou.aliyuncs.com",
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
		Endpoint: bosfs.DefaultEndpoint,
		Ak:       "Ak",
		Sk:       "Sk",
		Bucket:   "test bucket",
	})
}

func main() {
	gf := gfs.New()
	gf.Extend(local)
	var err error
	err = gf.WriteReader(root+"4.txt", strings.NewReader("test"))
	adapter, err := gf.Adapter(gfs.DiskNameLocal)
	err = adapter.WriteReader(root+"5.txt", strings.NewReader("test"))
	fmt.Println(err)
	//Write file
	err = gf.Write(root+"1.txt", []byte("test data"))
	fmt.Println(err)
	//Write data from resource file
	err = gf.WriteStream(root+"2.txt", root+"/1.txt")
	fmt.Println(err)
	//Update file
	err = gf.Update(root+"1.txt", []byte("test update data"))
	fmt.Println(err)
	//Update data from resource file
	err = gf.UpdateStream(root+"2.txt", root+"/1.txt")
	fmt.Println(err)
	exists, err := gf.Exist(root + "2.txt")
	fmt.Println(exists)
	//Read file
	read, err := gf.Read(root + "2.txt")
	fmt.Println(read, err)
	//Get file mime type
	mimeType, err := gf.MimeType(root + "2.txt")
	fmt.Println(mimeType, err)
	//Get file size
	size, err := gf.Size(root + "2.txt")
	fmt.Println(size, err)
	//Move file
	_, err = gf.Move(root+"1.txt", root+"4.txt")
	fmt.Println(err)
	//Copy file
	_, err = gf.Copy(root+"2.txt", root+"5.txt")
	fmt.Println(err)
}
