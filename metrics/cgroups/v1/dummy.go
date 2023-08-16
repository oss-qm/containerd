//go:build linux && no_cgroup1

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

package v1

import (
	"context"
	"errors"

	"github.com/containerd/containerd/events"
	"github.com/containerd/containerd/runtime"
	"github.com/docker/go-metrics"
)

func NewTaskMonitor(ctx context.Context, publisher events.Publisher, ns *metrics.Namespace) (runtime.TaskMonitor, error) {
	return nil, errors.New("no cgroup1 support")
}
