// Copyright 2025 肖其顿
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

//go:build darwin || dragonfly || freebsd || illumos || ios || linux || netbsd || openbsd || solaris

package flock

import (
	"os"
	"syscall"
)

// Lock 获取一个排它锁（写锁）。
// 如果锁已被其他进程持有，则调用会阻塞，直到可以获取锁为止。
func (f *Flock) Lock() error {
	file, err := os.OpenFile(f.path, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	f.file = file
	err = syscall.Flock(int(f.file.Fd()), syscall.LOCK_EX)
	if err != nil {
		f.file.Close()
		f.file = nil
		return err
	}
	return nil
}

// RLock 获取一个共享锁（读锁）。
// 多个进程可以同时持有读锁。如果锁被其他进程以写锁方式持有，
// 则调用会阻塞，直到可以获取锁为止。
func (f *Flock) RLock() error {
	file, err := os.OpenFile(f.path, os.O_CREATE|os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	f.file = file
	err = syscall.Flock(int(f.file.Fd()), syscall.LOCK_SH)
	if err != nil {
		f.file.Close()
		f.file = nil
		return err
	}
	return nil
}

// Unlock 释放文件锁。它可用于释放通过 Lock 或 RLock 获取的锁。
// 此方法是幂等的，可以安全地多次调用。
func (f *Flock) Unlock() error {
	if f.file == nil {
		return nil
	}
	defer func() {
		f.file.Close()
		f.file = nil
	}()
	return syscall.Flock(int(f.file.Fd()), syscall.LOCK_UN)
}

// TryLock 尝试获取一个排它锁（写锁），但不会阻塞。
// 如果成功获取锁，返回 true, nil。
// 如果锁已被其他进程持有，立即返回 false, nil。
func (f *Flock) TryLock() (bool, error) {
	file, err := os.OpenFile(f.path, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return false, err
	}
	f.file = file
	err = syscall.Flock(int(f.file.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
	if err == nil {
		return true, nil
	}
	f.file.Close()
	f.file = nil
	if err == syscall.EWOULDBLOCK || err == syscall.EAGAIN {
		return false, nil
	}
	return false, err
}

// TryRLock 尝试获取一个共享锁（读锁），但不会阻塞。
// 如果成功获取锁，返回 true, nil。
// 如果锁已被其他进程以写锁方式持有，立即返回 false, nil。
func (f *Flock) TryRLock() (bool, error) {
	file, err := os.OpenFile(f.path, os.O_CREATE|os.O_RDONLY, 0666)
	if err != nil {
		return false, err
	}
	f.file = file
	err = syscall.Flock(int(f.file.Fd()), syscall.LOCK_SH|syscall.LOCK_NB)
	if err == nil {
		return true, nil
	}
	f.file.Close()
	f.file = nil
	if err == syscall.EWOULDBLOCK || err == syscall.EAGAIN {
		return false, nil
	}
	return false, err
}
