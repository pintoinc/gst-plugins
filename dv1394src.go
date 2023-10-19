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

type Dv1394Src struct {
	*base.GstPushSrc
}

func NewDv1394Src() (*Dv1394Src, error) {
	e, err := gst.NewElement("dv1394src")
	if err != nil {
		return nil, err
	}
	return &Dv1394Src{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

func NewDv1394SrcWithName(name string) (*Dv1394Src, error) {
	e, err := gst.NewElementWithName("dv1394src", name)
	if err != nil {
		return nil, err
	}
	return &Dv1394Src{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

// ----- Properties -----

// channel (int)
//
// Channel number for listening

func (e *Dv1394Src) GetChannel() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("channel")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Dv1394Src) SetChannel(value int) error {
	return e.Element.SetProperty("channel", value)
}

// consecutive (int)
//
// send n consecutive frames after skipping

func (e *Dv1394Src) GetConsecutive() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("consecutive")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Dv1394Src) SetConsecutive(value int) error {
	return e.Element.SetProperty("consecutive", value)
}

// device-name (string)
//
// Descriptive name of the currently opened device

func (e *Dv1394Src) GetDeviceName() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("device-name")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// drop-incomplete (bool)
//
// drop incomplete frames

func (e *Dv1394Src) GetDropIncomplete() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("drop-incomplete")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Dv1394Src) SetDropIncomplete(value bool) error {
	return e.Element.SetProperty("drop-incomplete", value)
}

// guid (uint64)
//
// select one of multiple DV devices by its GUID. use a hexadecimal like 0xhhhhhhhhhhhhhhhh. (0 = no guid)

func (e *Dv1394Src) GetGuid() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("guid")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Dv1394Src) SetGuid(value uint64) error {
	return e.Element.SetProperty("guid", value)
}

// port (int)
//
// Port number (-1 automatic)

func (e *Dv1394Src) GetPort() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("port")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Dv1394Src) SetPort(value int) error {
	return e.Element.SetProperty("port", value)
}

// skip (int)
//
// skip n frames

func (e *Dv1394Src) GetSkip() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("skip")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Dv1394Src) SetSkip(value int) error {
	return e.Element.SetProperty("skip", value)
}

// use-avc (bool)
//
// Use AV/C VTR control

func (e *Dv1394Src) GetUseAvc() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("use-avc")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Dv1394Src) SetUseAvc(value bool) error {
	return e.Element.SetProperty("use-avc", value)
}
