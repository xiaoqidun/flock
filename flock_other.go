// Copyright 2025 肖其顿 (XIAO QI DUN)
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

//go:build !darwin && !dragonfly && !freebsd && !illumos && !ios && !linux && !netbsd && !openbsd && !solaris && !windows

package flock

// Lock 在不支持的平台上总是返回 ErrUnsupportedPlatform 错误。
func (f *Flock) Lock() error {
	return ErrUnsupportedPlatform
}

// RLock 在不支持的平台上总是返回 ErrUnsupportedPlatform 错误。
func (f *Flock) RLock() error {
	return ErrUnsupportedPlatform
}

// Unlock 在不支持的平台上总是返回 ErrUnsupportedPlatform 错误。
func (f *Flock) Unlock() error {
	return ErrUnsupportedPlatform
}

// TryLock 在不支持的平台上总是返回 false 和 ErrUnsupportedPlatform 错误。
func (f *Flock) TryLock() (bool, error) {
	return false, ErrUnsupportedPlatform
}

// TryRLock 在不支持的平台上总是返回 false 和 ErrUnsupportedPlatform 错误。
func (f *Flock) TryRLock() (bool, error) {
	return false, ErrUnsupportedPlatform
}
