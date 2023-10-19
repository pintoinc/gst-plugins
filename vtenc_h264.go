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

type VtencH264 struct {
	Element *gst.Element
}

func NewVtencH264() (*VtencH264, error) {
	e, err := gst.NewElement("vtenc_h264")
	if err != nil {
		return nil, err
	}
	return &VtencH264{Element: e}, nil
}

func NewVtencH264WithName(name string) (*VtencH264, error) {
	e, err := gst.NewElementWithName("vtenc_h264", name)
	if err != nil {
		return nil, err
	}
	return &VtencH264{Element: e}, nil
}

// ----- Properties -----

// allow-frame-reordering (bool)
//
// Whether to allow frame reordering or not

func (e *VtencH264) GetAllowFrameReordering() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("allow-frame-reordering")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *VtencH264) SetAllowFrameReordering(value bool) error {
	return e.Element.SetProperty("allow-frame-reordering", value)
}

// bitrate (uint)
//
// Target video bitrate in kbps (0 = auto)

func (e *VtencH264) GetBitrate() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *VtencH264) SetBitrate(value uint) error {
	return e.Element.SetProperty("bitrate", value)
}

// max-keyframe-interval (int)
//
// Maximum number of frames between keyframes (0 = auto)

func (e *VtencH264) GetMaxKeyframeInterval() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-keyframe-interval")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *VtencH264) SetMaxKeyframeInterval(value int) error {
	return e.Element.SetProperty("max-keyframe-interval", value)
}

// max-keyframe-interval-duration (uint64)
//
// Maximum number of nanoseconds between keyframes (0 = no limit)

func (e *VtencH264) GetMaxKeyframeIntervalDuration() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("max-keyframe-interval-duration")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *VtencH264) SetMaxKeyframeIntervalDuration(value uint64) error {
	return e.Element.SetProperty("max-keyframe-interval-duration", value)
}

// quality (float64)
//
// The desired compression quality

func (e *VtencH264) GetQuality() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("quality")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *VtencH264) SetQuality(value float64) error {
	return e.Element.SetProperty("quality", value)
}

// realtime (bool)
//
// Configure the encoder for realtime output

func (e *VtencH264) GetRealtime() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("realtime")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *VtencH264) SetRealtime(value bool) error {
	return e.Element.SetProperty("realtime", value)
}

