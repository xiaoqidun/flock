# flock [![PkgGoDev](https://pkg.go.dev/badge/github.com/xiaoqidun/flock)](https://pkg.go.dev/github.com/xiaoqidun/flock)
一个简单、可靠、跨平台的 Go 语言进程级文件锁

# 安装指南
```shell
go get -u github.com/xiaoqidun/flock
```

# 快速开始
```go
package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/xiaoqidun/flock"
)

func main() {
	// 1. 创建 flock 实例，指定锁文件路径
	//    使用系统临时目录，确保路径明确且跨平台兼容
	lockPath := filepath.Join(os.TempDir(), "file.lock")
	f := flock.New(lockPath)
	// 2. 获取排它锁（写锁）
	//    如果锁已被其他进程持有，此调用将阻塞直到获取成功
	fmt.Println("尝试获取文件锁")
	err := f.Lock()
	if err != nil {
		fmt.Printf("获取文件锁失败: %v\n", err)
		return
	}
	// 3. 确保在函数退出时释放锁
	//    Unlock 方法只会释放锁，并不会删除文件
	defer f.Unlock()
	fmt.Println("获取文件锁成功")
}
```

# 授权协议
本项目使用 [Apache License 2.0](https://github.com/xiaoqidun/flock/blob/main/LICENSE) 授权协议