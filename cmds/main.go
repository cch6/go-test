package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	// ctx, cancel := context.WithCancel(context.Background())
	cmd := exec.Command("ffmpeg", "-i", "rtsp://10.58.123.230:554/SubStream", "-c", "copy", "-f", "flv", "rtmp://localhost:1935/live/movie/a")
	// cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.Stdout = os.Stdout
	cmd.Start()
	fmt.Println("程序开始运行...")

	time.Sleep(1000 * time.Second)
	fmt.Println("退出程序中...", cmd.Process.Pid)
	cmd.Process.Kill()
	cmd.Wait()
}
