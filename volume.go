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

type Volume struct {
	*base.GstBaseTransform
}

func NewVolume() (*Volume, error) {
	e, err := gst.NewElement("volume")
	if err != nil {
		return nil, err
	}
	return &Volume{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewVolumeWithName(name string) (*Volume, error) {
	e, err := gst.NewElementWithName("volume", name)
	if err != nil {
		return nil, err
	}
	return &Volume{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// mute (bool)
//
// mute channel

func (e *Volume) GetMute() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("mute")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Volume) SetMute(value bool) error {
	return e.Element.SetProperty("mute", value)
}

// volume (float64)
//
// volume factor, 1.0=100%%

func (e *Volume) GetVolume() (float64, error) {
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

func (e *Volume) SetVolume(value float64) error {
	return e.Element.SetProperty("volume", value)
}

