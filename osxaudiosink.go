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

type Osxaudiosink struct {
	*base.GstBaseSink
}

func NewOsxaudiosink() (*Osxaudiosink, error) {
	e, err := gst.NewElement("osxaudiosink")
	if err != nil {
		return nil, err
	}
	return &Osxaudiosink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewOsxaudiosinkWithName(name string) (*Osxaudiosink, error) {
	e, err := gst.NewElementWithName("osxaudiosink", name)
	if err != nil {
		return nil, err
	}
	return &Osxaudiosink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// ----- Properties -----

// device (int)
//
// Device ID of output device

func (e *Osxaudiosink) GetDevice() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("device")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Osxaudiosink) SetDevice(value int) error {
	return e.Element.SetProperty("device", value)
}

// volume (float64)
//
// Volume of this stream

func (e *Osxaudiosink) GetVolume() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("volume")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Osxaudiosink) SetVolume(value float64) error {
	return e.Element.SetProperty("volume", value)
}

