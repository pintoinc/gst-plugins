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

type AvdecVp4 struct {
	Element *gst.Element
}

func NewAvdecVp4() (*AvdecVp4, error) {
	e, err := gst.NewElement("avdec_vp4")
	if err != nil {
		return nil, err
	}
	return &AvdecVp4{Element: e}, nil
}

func NewAvdecVp4WithName(name string) (*AvdecVp4, error) {
	e, err := gst.NewElementWithName("avdec_vp4", name)
	if err != nil {
		return nil, err
	}
	return &AvdecVp4{Element: e}, nil
}

// ----- Properties -----

// debug-mv (bool)
//
// Whether to print motion vectors on top of the image (deprecated, non-functional)

func (e *AvdecVp4) GetDebugMv() (bool, error) {
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

func (e *AvdecVp4) SetDebugMv(value bool) error {
	return e.Element.SetProperty("debug-mv", value)
}

// direct-rendering (bool)
//
// Enable direct rendering

func (e *AvdecVp4) GetDirectRendering() (bool, error) {
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

func (e *AvdecVp4) SetDirectRendering(value bool) error {
	return e.Element.SetProperty("direct-rendering", value)
}

// lowres (GstLibAVVidDecLowres)
//
// At which resolution to decode images

func (e *AvdecVp4) GetLowres() (interface{}, error) {
	return e.Element.GetProperty("lowres")
}

func (e *AvdecVp4) SetLowres(value interface{}) error {
	return e.Element.SetProperty("lowres", value)
}

// max-threads (int)
//
// Maximum number of worker threads to spawn. (0 = auto)

func (e *AvdecVp4) GetMaxThreads() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-threads")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvdecVp4) SetMaxThreads(value int) error {
	return e.Element.SetProperty("max-threads", value)
}

// output-corrupt (bool)
//
// Whether libav should output frames even if corrupted

func (e *AvdecVp4) GetOutputCorrupt() (bool, error) {
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

func (e *AvdecVp4) SetOutputCorrupt(value bool) error {
	return e.Element.SetProperty("output-corrupt", value)
}

// skip-frame (GstLibAVVidDecSkipFrame)
//
// Which types of frames to skip during decoding

func (e *AvdecVp4) GetSkipFrame() (interface{}, error) {
	return e.Element.GetProperty("skip-frame")
}

func (e *AvdecVp4) SetSkipFrame(value interface{}) error {
	return e.Element.SetProperty("skip-frame", value)
}

// thread-type (GstLibAVVidDecThreadType)
//
// Multithreading methods to use

func (e *AvdecVp4) GetThreadType() (interface{}, error) {
	return e.Element.GetProperty("thread-type")
}

func (e *AvdecVp4) SetThreadType(value interface{}) error {
	return e.Element.SetProperty("thread-type", value)
}

