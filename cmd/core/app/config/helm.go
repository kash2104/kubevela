/*
Copyright 2026 The KubeVela Authors.

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

package config

import (
	"context"
	"time"

	"github.com/spf13/pflag"

	"github.com/oam-dev/kubevela/pkg/cue/cuex/providers/helm"
)

// HelmConfig contains configuration for the Helm chart cache.
type HelmConfig struct {
	ChartCacheMaxBytes      int64
	ChartCacheSweepInterval time.Duration
}

// NewHelmConfig creates a new HelmConfig with defaults.
func NewHelmConfig() *HelmConfig {
	return &HelmConfig{
		ChartCacheMaxBytes:      helm.DefaultChartCacheMaxBytes,
		ChartCacheSweepInterval: helm.DefaultChartCacheSweepInterval,
	}
}

// AddFlags registers Helm configuration flags.
func (c *HelmConfig) AddFlags(fs *pflag.FlagSet) {
	fs.Int64Var(&c.ChartCacheMaxBytes,
		"helm-cache-max-bytes",
		c.ChartCacheMaxBytes,
		"Maximum number of bytes to keep in the Helm chart cache. Set to 0 to use the default.")
	fs.DurationVar(&c.ChartCacheSweepInterval,
		"helm-cache-sweep-interval",
		c.ChartCacheSweepInterval,
		"How often the Helm chart cache sweeps expired entries. Set to 0 to use the default.")
}

// SyncToHelmGlobals syncs the parsed configuration values to the Helm provider
// package globals. This should be called after flag parsing and before any Helm
// provider is constructed.
//
// ctx SHOULD be the controller's root context so the cache sweeper goroutine is
// tied to its lifetime.
func (c *HelmConfig) SyncToHelmGlobals(ctx context.Context) {
	helm.InitChartCache(ctx, c.ChartCacheMaxBytes, c.ChartCacheSweepInterval)
}
