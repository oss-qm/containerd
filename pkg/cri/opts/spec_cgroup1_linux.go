//go:build linux && !no_cgroup1
/*
   Copyright The containerd Authors.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package opts

import (
	"os"
	"sync"
)

var (
	_cgroupv1HasHugetlbOnce sync.Once
	_cgroupv1HasHugetlb     bool
)

// cgroupv1HasHugetlb returns whether the hugetlb controller is present on
// cgroup v1.
func cgroupv1HasHugetlb() bool {
	_cgroupv1HasHugetlbOnce.Do(func() {
		if _, err := os.ReadDir("/sys/fs/cgroup/hugetlb"); err != nil {
			_cgroupv1HasHugetlb = false
		} else {
			_cgroupv1HasHugetlb = true
		}
	})
	return _cgroupv1HasHugetlb
}
