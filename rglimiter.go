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

type Rglimiter struct {
	*base.GstBaseTransform
}

func NewRglimiter() (*Rglimiter, error) {
	e, err := gst.NewElement("rglimiter")
	if err != nil {
		return nil, err
	}
	return &Rglimiter{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewRglimiterWithName(name string) (*Rglimiter, error) {
	e, err := gst.NewElementWithName("rglimiter", name)
	if err != nil {
		return nil, err
	}
	return &Rglimiter{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// enabled (bool)
//
// Enable processing

func (e *Rglimiter) GetEnabled() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("enabled")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Rglimiter) SetEnabled(value bool) error {
	return e.Element.SetProperty("enabled", value)
}

