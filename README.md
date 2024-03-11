[![Go Report Card](https://goreportcard.com/badge/github.com/pkg6/gfs)](https://goreportcard.com/report/github.com/pkg6/gfs)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/pkg6/gfs?tab=doc)
[![Sourcegraph](https://sourcegraph.com/github.com/pkg6/gfs/-/badge.svg)](https://sourcegraph.com/github.com/pkg6/gfs?badge)
[![Release](https://img.shields.io/github/release/pkg6/gfs.svg?style=flat-square)](https://github.com/pkg6/gfs/releases)

## gfs
gfs is a file storage library for Golang. It provides one interface to interact with many types of fs. When you use gfs, you're not only protected from vendor lock-in, you'll also have a consistent experience for which ever storage is right for you.

## Install

~~~
go get github.com/pkg6/gfs
~~~

## Officially supported adapters

- **Local**
- **[ALiYun OSS](https://help.aliyun.com/product/31815.html)**
- **[Google Cloud Storage](https://cloud.google.com/storage/docs/introduction?hl=zh-CN)**
- **[QiNiu KoDo](https://www.qiniu.com/products/kodo)**
- **[Tencent COS](https://cloud.tencent.com/product/cos)**
- **[BaiDu BOS](https://cloud.baidu.com/product/bos.html)**

## adapter Interface
[You can always create an adapter yourself.](https://github.com/pkg6/gfs/blob/main/ifs.go#L77)
