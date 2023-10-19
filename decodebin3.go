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
)

type Decodebin3 struct {
	Element *gst.Element
}

func NewDecodebin3() (*Decodebin3, error) {
	e, err := gst.NewElement("decodebin3")
	if err != nil {
		return nil, err
	}
	return &Decodebin3{Element: e}, nil
}

func NewDecodebin3WithName(name string) (*Decodebin3, error) {
	e, err := gst.NewElementWithName("decodebin3", name)
	if err != nil {
		return nil, err
	}
	return &Decodebin3{Element: e}, nil
}

// ----- Properties -----

// caps (GstCaps)
//
// The caps on which to stop decoding. (NULL = default)

func (e *Decodebin3) GetCaps() (*gst.Caps, error) {
	var value *gst.Caps
	var ok bool
	v, err := e.Element.GetProperty("caps")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Caps)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Caps")
	}
	return value, nil
}

func (e *Decodebin3) SetCaps(value *gst.Caps) error {
	return e.Element.SetProperty("caps", value)
}
