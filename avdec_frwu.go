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

type AvdecFrwu struct {
	Element *gst.Element
}

func NewAvdecFrwu() (*AvdecFrwu, error) {
	e, err := gst.NewElement("avdec_frwu")
	if err != nil {
		return nil, err
	}
	return &AvdecFrwu{Element: e}, nil
}

func NewAvdecFrwuWithName(name string) (*AvdecFrwu, error) {
	e, err := gst.NewElementWithName("avdec_frwu", name)
	if err != nil {
		return nil, err
	}
	return &AvdecFrwu{Element: e}, nil
}

// ----- Properties -----

// debug-mv (bool)
//
// Whether to print motion vectors on top of the image (deprecated, non-functional)

func (e *AvdecFrwu) GetDebugMv() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("debug-mv")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *AvdecFrwu) SetDebugMv(value bool) error {
	return e.Element.SetProperty("debug-mv", value)
}

// direct-rendering (bool)
//
// Enable direct rendering

func (e *AvdecFrwu) GetDirectRendering() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("direct-rendering")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *AvdecFrwu) SetDirectRendering(value bool) error {
	return e.Element.SetProperty("direct-rendering", value)
}

// lowres (GstLibAVVidDecLowres)
//
// At which resolution to decode images

func (e *AvdecFrwu) GetLowres() (interface{}, error) {
	return e.Element.GetProperty("lowres")
}

func (e *AvdecFrwu) SetLowres(value interface{}) error {
	return e.Element.SetProperty("lowres", value)
}

// output-corrupt (bool)
//
// Whether libav should output frames even if corrupted

func (e *AvdecFrwu) GetOutputCorrupt() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("output-corrupt")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *AvdecFrwu) SetOutputCorrupt(value bool) error {
	return e.Element.SetProperty("output-corrupt", value)
}

// skip-frame (GstLibAVVidDecSkipFrame)
//
// Which types of frames to skip during decoding

func (e *AvdecFrwu) GetSkipFrame() (interface{}, error) {
	return e.Element.GetProperty("skip-frame")
}

func (e *AvdecFrwu) SetSkipFrame(value interface{}) error {
	return e.Element.SetProperty("skip-frame", value)
}

