package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("请提供一个mp4文件作为参数")
	}

	mp4File := os.Args[1]

	// 获取文件名
	paths := strings.Split(mp4File, "/")
	filename := paths[len(paths)-1]
	uri := fmt.Sprintf("rtsp://localhost:8554/live/%s", filename)

	log.Println("[RTSP流地址]:", uri)

	// 获取本机器所有网络接口
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("获取网络接口失败:", err)
		return
	}

	// 遍历所有网络接口
	for _, iface := range interfaces {
		// 获取网络接口的所有地址
		addrs, err := iface.Addrs()
		if err != nil {
			fmt.Println("获取地址失败:", err)
			continue
		}

		// 遍历地址列表
		for _, addr := range addrs {
			// 检查地址是否为IPv4
			if ipnet, ok := addr.(*net.IPNet); ok && ipnet.IP.To4() != nil {
				// 输出 IPv4 地址，包括 0.0.0.0 和 localhost
				log.Printf("[RTSP流地址]: rtsp://%s:8554/live/%s\n", ipnet.IP, filename)
			}
		}
	}

	pushStream(mp4File, uri)
}

func pushStream(mp4File, uri string) {
	cmd := exec.Command("ffmpeg",
		"-re",
		"-stream_loop", "-1",
		"-i", mp4File,
		"-rtsp_transport", "udp",
		"-c", "copy",
		"-f", "rtsp",
		uri)

	cmd.Stdout = nil
	cmd.Stderr = nil
	err := cmd.Run()

	if err != nil {
		log.Printf("[RTSP流地址]失败: %v", err)
	}
}
