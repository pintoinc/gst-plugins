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
	"fmt"

	"github.com/tinyzimmer/go-gst/gst"

	"github.com/livekit/gstplugins/base"
)

type Mj2Mux struct {
	*base.GstAggregator
}

func NewMj2Mux() (*Mj2Mux, error) {
	e, err := gst.NewElement("mj2mux")
	if err != nil {
		return nil, err
	}
	return &Mj2Mux{GstAggregator: &base.GstAggregator{Element: e}}, nil
}

func NewMj2MuxWithName(name string) (*Mj2Mux, error) {
	e, err := gst.NewElementWithName("mj2mux", name)
	if err != nil {
		return nil, err
	}
	return &Mj2Mux{GstAggregator: &base.GstAggregator{Element: e}}, nil
}

// ----- Properties -----

// streamable (bool)
//
// If set to true, the output should be as if it is to be streamed and hence no indexes written or duration written. (DEPRECATED, only valid for fragmented MP4)

func (e *Mj2Mux) GetStreamable() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("streamable")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Mj2Mux) SetStreamable(value bool) error {
	return e.Element.SetProperty("streamable", value)
}

