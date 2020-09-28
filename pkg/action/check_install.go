/**
 * Copyright © 2014-2020 The SiteWhere Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package action

import "github.com/sitewhere/swctl/pkg/install"

// CheckInstall is the action for check SiteWhere installation
type CheckInstall struct {
	cfg *Configuration

	// Use verbose mode
	Verbose bool
}

// NewCheckInstall constructs a new *Install
func NewCheckInstall(cfg *Configuration) *CheckInstall {
	return &CheckInstall{
		cfg:     cfg,
		Verbose: false,
	}
}

// Run executes the list command, returning a set of matches.
func (i *CheckInstall) Run() (*install.SiteWhereInstall, error) {
	if err := i.cfg.KubeClient.IsReachable(); err != nil {
		return nil, err
	}

	return &install.SiteWhereInstall{}, nil
}