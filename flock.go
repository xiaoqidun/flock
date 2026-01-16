// Copyright 2025-2026 肖其顿 (XIAO QI DUN)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package flock 提供了一个简单、可靠、跨平台的 Go 语言进程级文件锁。
//
// 本包用于在多个独立进程之间安全地共享资源，以防止并发冲突。
// 它支持共享锁（读锁）和排它锁（写锁），并统一使用 Unlock 方法释放。
//
// 注意：本包不适用于单个进程内的 Goroutine 同步（请使用 sync.Mutex）。
package flock

import (
	"errors"
	"os"
)

// ErrUnsupportedPlatform 表示在当前操作系统上不支持文件锁功能。
var ErrUnsupportedPlatform = errors.New("file locking is not supported on this platform")

// Flock 代表一个文件锁实例，封装了跨平台的实现细节。
type Flock struct {
	path string
	file *os.File
}

// New 创建一个文件锁实例。
// 它接受锁文件的路径，但不会立即执行加锁操作。
// 如果锁文件不存在，会在首次加锁时自动创建。
func New(path string) *Flock {
	return &Flock{
		path: path,
	}
}
