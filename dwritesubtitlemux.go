// Copyright 2023 LiveKit, Inc.
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

package gstplugins

import (
	"github.com/tinyzimmer/go-gst/gst"

	"github.com/livekit/gstplugins/base"
)

type Dwritesubtitlemux struct {
	*base.GstAggregator
}

func NewDwritesubtitlemux() (*Dwritesubtitlemux, error) {
	e, err := gst.NewElement("dwritesubtitlemux")
	if err != nil {
		return nil, err
	}
	return &Dwritesubtitlemux{GstAggregator: &base.GstAggregator{Element: e}}, nil
}

func NewDwritesubtitlemuxWithName(name string) (*Dwritesubtitlemux, error) {
	e, err := gst.NewElementWithName("dwritesubtitlemux", name)
	if err != nil {
		return nil, err
	}
	return &Dwritesubtitlemux{GstAggregator: &base.GstAggregator{Element: e}}, nil
}

