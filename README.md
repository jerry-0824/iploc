# iploc

[![Build Status](https://travis-ci.org/kayon/iploc.svg?branch=master)](https://travis-ci.org/kayon/iploc)

使用纯真IP库 `qqwry.dat`，高性能，线程安全，并对国内数据格式化到省、市、县

> 需要 go 1.9 或更高

> <del>附带的 `qqwry.dat` 为 `UTF-8` 编码 `2018-05-15版本`</del>

> 不再提供 `qqwry.dat`, 新增命令行工具 `iploc-fetch`, 可在线获取官方最新版本的 `qqwry.dat`


## 安装

```
go get -u github.com/kayon/iploc/...
```

#### 无法安装 `golang.org/x/text` 包，没有梯子使用下面方法

```
$ mkdir -P $GOPATH/src/golang.org/x
cd $GOPATH/src/golang.org/x
git clone https://github.com/golang/text.git
```

## 获取&更新 qqwry.dat

#### 1. 下载 `qqwry.dat`

##### 方法一：使用命令行工具 [iploc-fetch](#iploc-fetch)

由于服务器限制国外IP，只能使用国内网络。

下载到当前目录，保存为 `qqwry.gbk.dat`

```
$ iploc-fetch qqwry.gbk.dat
```

##### 方法二：手动下载

在[纯真官网下载](http://www.cz88.net/fox/ipdat.shtml)并安装，复制安装目录中的 `qqwry.dat`

#### 2. 转换为 `UTF-8`

使用命令行工具 [iploc-conv](#iploc-conv) 将刚刚下载的 `qqwry.gbk.dat` 转换为 `UTF-8` 保存为 `qqwry.dat`

```
iploc-conv -s qqwry.gbk.dat -d qqwry.dat
```


## Benchmarks

```
// 缓存索引
BenchmarkFind-8            	 2000000	       735 ns/op               136万/秒
// 无索引
BenchmarkFindUnIndexed-8   	   20000	     78221 ns/op               1.2万/秒
// 缓存索引，并发
BenchmarkFindParallel-8    	10000000	       205 ns/op               487万/秒
```

## 使用

```
func main() {
	loc, err := iploc.Open("qqwry.dat")
	if err != nil {
		panic(err)
	}
	detail := loc.Find("8.8.8") // 补全为8.8.0.8, 参考 ping 工具
	fmt.Printf("IP:%s; 网段:%s - %s; %s\n", detail.IP, detail.Start, detail.End, detail)
	
	detail2 := loc.Find("8.8.3.1")
	fmt.Printf("%t %t\n", detail.In(detail2.IP.String()), detail.String() == detail2.String())

	// output
	// IP:8.8.0.8; 网段: 8.7.245.0 - 8.8.3.255; 美国 科罗拉多州布隆菲尔德市Level 3通信股份有限公司
	// true true
	
	detail = loc.Find("1.24.41.0")
	fmt.Println(detail.String())
	fmt.Println(detail.Country, detail.Province, detail.City, detail.County)
	
	// output
	// 内蒙古锡林郭勒盟苏尼特右旗 联通
	// 中国 内蒙古 锡林郭勒盟 苏尼特右旗
	
}	
```

#### 快捷方法
##### Find(qqwrySrc, ip string) (*Detail, error)
`iploc.Find` 使用 `OpenWithoutIndexes`

#### 初始化
##### Open(qqwrySrc string) (*Locator, error)

`iploc.Open` 缓存并索引，生成索引需要耗时500毫秒左右，但会带来更高的查询性能

##### OpenWithoutIndexes(qqwrySrc string) (*Locator, error)

`iploc.OpenWithoutIndexes` 只读取文件头做简单检查，无索引

#### 查询

```
(*Locator).Find(ip string) *Detail

```
> 如果IP不合法，返回 `nil`


## 命令行工具

#### <a name="iploc-conv"></a>iploc-conv

将原版 `qqwry.dat` 由 `GBK` 转换为 `UTF-8`

```
$ iploc-conv -s src.gbk.dat -d dst.utf8.dat
```

> 新生成的DAT文件保留原数据结构，由于编码原因，文件会增大一点

> 修正原 qqwry.dat 中几处错误的重定向 (qqwry.dat 2018-05-10)，并将 "CZ88.NET" 替换为 "N/A"

#### <a name="iploc-fetch"></a>iploc-fetch

从纯真官网下载最新 `qqwry.dat`

由于服务器限制国外IP，只能使用国内网络。

```
$ iploc-fetch qqwry.gbk.dat
```

> 下载后别忘了使用 `iploc-conv` 转换为 `UTF-8`

#### iploc-gen

创建静态版本的 **iploc** 集成到你的项目中

`iploc-gen` 会在当前目录创建 iploc-binary.go 文件，拷贝到你的项目中，通过变量名 *IPLoc* 直接可以使用

```
$ iploc-gen path/qqwry.dat
```

> `--pkg` 指定 package name, 默认 main

> `-n` 使用 `OpenWithoutIndexes` 初始化，无索引


## 静态编译 iploc 和 qqwry.dat 并集成到你的项目中

编译后的二进制没有 `qqwry.dat` 依赖，不需要再带着 `qqwry.dat` 一起打包了

##### 示例

到项目目录 `$GOPATH/src/myproject/` 中

```
$ mkdir myloc && cd myloc
$ iploc-gen path/qqwry.dat --pkg myloc
```

> $GOPATH/src/myproject/main.go

```
package main
	
import (
	"fmt"
	
	"myproject/myloc"
)
	
func main() {
	fmt.Println(myloc.IPLoc.Find("8.8.8.8"))
}
```

```
===== 8.8.8.8 =====
IP:8.8.0.8; 网段:8.7.245.0 - 8.8.3.255; 美国 科罗拉多州布隆菲尔德市Level 3通信股份有限公司
true true
===== 1.24.41.0 =====
内蒙古锡林郭勒盟苏尼特右旗 联通
中国 内蒙古 锡林郭勒盟 苏尼特右旗
中国 联通 内蒙古 锡林郭勒盟 苏尼特右旗
===== 185.199.110.154 =====
美国 GitHub+Fastly节点
美国   
美国 GitHub+Fastly节点   
===== 129.226.162.177 =====
香港 腾讯云
中国 香港  
中国 腾讯云 香港  
===== 163.19.9.247 =====
台湾省新竹县 TANet
中国 台湾 新竹县 
中国 TANet 台湾 新竹县 
===== 60.246.73.3 =====
澳门 澳门电讯
中国 澳门  
中国 澳门电讯 澳门  
===== 35.187.224.178 =====
新加坡 Google云计算数据中心
新加坡   
新加坡 Google云计算数据中心   
===== 127.0.0.1 =====
本机地址 N/A
本机地址   
本机地址 N/A   
===== 192.168.0.1 =====
局域网 对方和您在同一内部网
局域网   
局域网 对方和您在同一内部网  
```

