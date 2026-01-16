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

//go:build windows

package flock

import (
	"errors"
	"os"

	"golang.org/x/sys/windows"
)

// Lock 获取一个排它锁（写锁）。
// 如果锁已被其他进程持有，则调用会阻塞，直到可以获取锁为止。
func (f *Flock) Lock() error {
	file, err := os.OpenFile(f.path, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	f.file = file
	handle := windows.Handle(f.file.Fd())
	err = windows.LockFileEx(handle, windows.LOCKFILE_EXCLUSIVE_LOCK, 0, 1, 0, &windows.Overlapped{})
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
	file, err := os.OpenFile(f.path, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	f.file = file
	handle := windows.Handle(f.file.Fd())
	err = windows.LockFileEx(handle, 0, 0, 1, 0, &windows.Overlapped{})
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
	handle := windows.Handle(f.file.Fd())
	return windows.UnlockFileEx(handle, 0, 1, 0, &windows.Overlapped{})
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
	handle := windows.Handle(f.file.Fd())
	err = windows.LockFileEx(handle, windows.LOCKFILE_EXCLUSIVE_LOCK|windows.LOCKFILE_FAIL_IMMEDIATELY, 0, 1, 0, &windows.Overlapped{})
	if err == nil {
		return true, nil
	}
	f.file.Close()
	f.file = nil
	if errors.Is(err, windows.ERROR_LOCK_VIOLATION) {
		return false, nil
	}
	return false, err
}

// TryRLock 尝试获取一个共享锁（读锁），但不会阻塞。
// 如果成功获取锁，返回 true, nil。
// 如果锁已被其他进程以写锁方式持有，立即返回 false, nil。
func (f *Flock) TryRLock() (bool, error) {
	file, err := os.OpenFile(f.path, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return false, err
	}
	f.file = file
	handle := windows.Handle(f.file.Fd())
	err = windows.LockFileEx(handle, windows.LOCKFILE_FAIL_IMMEDIATELY, 0, 1, 0, &windows.Overlapped{})
	if err == nil {
		return true, nil
	}
	f.file.Close()
	f.file = nil
	if errors.Is(err, windows.ERROR_LOCK_VIOLATION) {
		return false, nil
	}
	return false, err
}
