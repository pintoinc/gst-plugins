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

type Autovideosink struct {
	Element *gst.Element
}

func NewAutovideosink() (*Autovideosink, error) {
	e, err := gst.NewElement("autovideosink")
	if err != nil {
		return nil, err
	}
	return &Autovideosink{Element: e}, nil
}

func NewAutovideosinkWithName(name string) (*Autovideosink, error) {
	e, err := gst.NewElementWithName("autovideosink", name)
	if err != nil {
		return nil, err
	}
	return &Autovideosink{Element: e}, nil
}

// ----- Properties -----

// ts-offset (int64)
//
// Timestamp offset in nanoseconds

func (e *Autovideosink) GetTsOffset() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("ts-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

func (e *Autovideosink) SetTsOffset(value int64) error {
	return e.Element.SetProperty("ts-offset", value)
}
