// Copyright 2015 The rkt Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build darwin

package stage0

// #include <sys/param.h>
// #include <sys/mount.h>
import "C"

import (
	"unsafe"
	"fmt"
)

func mountfs(source string, target string, fstype string, flags uintptr, data string) error {
	ret := C.mount(C.CString(fstype), C.CString(target), C.int(flags), unsafe.Pointer(&source))
	if ret != 0 {
		fmt.Errorf("Couldn't mount %s to %s", source, target)
	}
	return nil
}

// not implemented
func remountPrivate(target string) error {
	return nil
}

