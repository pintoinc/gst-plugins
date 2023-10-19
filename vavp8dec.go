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
)

type Vavp8Dec struct {
	Element *gst.Element
}

func NewVavp8Dec() (*Vavp8Dec, error) {
	e, err := gst.NewElement("vavp8dec")
	if err != nil {
		return nil, err
	}
	return &Vavp8Dec{Element: e}, nil
}

func NewVavp8DecWithName(name string) (*Vavp8Dec, error) {
	e, err := gst.NewElementWithName("vavp8dec", name)
	if err != nil {
		return nil, err
	}
	return &Vavp8Dec{Element: e}, nil
}

// ----- Properties -----

// device-path (string)
//
// It shows the DRM device path used for the VA operation, if any.

func (e *Vavp8Dec) GetDevicePath() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("device-path")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}
