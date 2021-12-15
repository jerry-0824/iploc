package main

import (
	"fmt"
	"kayon/iploc"
)

func main() {
	loc, err := iploc.Open("qqwry.dat")
	if err != nil {
		panic(err)
	}

	var ipFind string

	// **** 00 ****

	detail := loc.Find("8.8.8") // 补全为8.8.0.8, 参考 ping 工具
	fmt.Println("===== 8.8.8.8 =====")
	fmt.Printf("IP:%s; 网段:%s - %s; %s\n", detail.IP, detail.Start, detail.End, detail)

	detail2 := loc.Find("8.8.3.1")
	fmt.Printf("%t %t\n", detail.In(detail2.IP.String()), detail.String() == detail2.String())

	// output
	// IP:8.8.0.8; 网段: 8.7.245.0 - 8.8.3.255; 美国 科罗拉多州布隆菲尔德市Level 3通信股份有限公司
	// true true

	// **** 01 ****
	ipFind = "1.24.41.0"
	PrintDetail(loc, ipFind)
	// output
	// ===== 1.24.41.0 =====
	// 内蒙古锡林郭勒盟苏尼特右旗 联通
	// 中国 内蒙古 锡林郭勒盟 苏尼特右旗
	// 中国 联通 内蒙古 锡林郭勒盟 苏尼特右旗

	// **** 02 ****
	ipFind = "185.199.110.154"
	PrintDetail(loc, ipFind)
	// output
	// ===== 185.199.110.154 =====
	// 美国 GitHub+Fastly节点
	// 美国

	// **** 03 ****
	ipFind = "129.226.162.177"
	PrintDetail(loc, ipFind)
	// output
	// ===== 129.226.162.177 =====
	// 香港 腾讯云
	// 中国 香港
	// 中国 腾讯云 香港

	// **** 04 ****
	ipFind = "163.19.9.247"
	PrintDetail(loc, ipFind)
	// output
	// ===== 163.19.9.247 =====
	// 台湾省新竹县 TANet
	// 中国 台湾 新竹县
	// 中国 TANet 台湾 新竹县

	// **** 05 ****
	ipFind = "60.246.73.3"
	PrintDetail(loc, ipFind)
	// output
	// ===== 60.246.73.3 =====
	// 澳门 澳门电讯
	// 中国 澳门
	// 中国 澳门电讯 澳门

	// **** 06 ****
	ipFind = "35.187.224.178"
	PrintDetail(loc, ipFind)
	// output
	// ===== 35.187.224.178 =====
	// 新加坡 Google云计算数据中心
	// 新加坡
	// 新加坡 Google云计算数据中心

	// **** 07 ****
	ipFind = "127.0.0.1"
	PrintDetail(loc, ipFind)
	// output
	// ===== 127.0.0.1 =====
	// 本机地址 N/A
	// 本机地址
	// 本机地址 N/A

	// **** 08 ****
	ipFind = "192.168.0.1"
	PrintDetail(loc, ipFind)
	// output
	// ===== 192.168.0.1 =====
	// 局域网 对方和您在同一内部网
	// 局域网
	// 局域网 对方和您在同一内部网
}

func PrintDetail(loc *iploc.Locator, ipFind string) {
	detail := loc.Find(ipFind)
	fmt.Println("=====", ipFind, "=====")
	fmt.Println(detail.String())
	fmt.Println(detail.Country, detail.Province, detail.City, detail.County)
	fmt.Println(detail.Country, detail.Region, detail.Province, detail.City, detail.County)
}
