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

type Frei0RFilterIirBlur struct {
	*base.GstBaseTransform
}

func NewFrei0RFilterIirBlur() (*Frei0RFilterIirBlur, error) {
	e, err := gst.NewElement("frei0r-filter-iir-blur")
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterIirBlur{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewFrei0RFilterIirBlurWithName(name string) (*Frei0RFilterIirBlur, error) {
	e, err := gst.NewElementWithName("frei0r-filter-iir-blur", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterIirBlur{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// amount (float64)
//
// Amount of blur

func (e *Frei0RFilterIirBlur) GetAmount() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("amount")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterIirBlur) SetAmount(value float64) error {
	return e.Element.SetProperty("amount", value)
}

// edge (bool)
//
// Edge compensation

func (e *Frei0RFilterIirBlur) GetEdge() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("edge")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Frei0RFilterIirBlur) SetEdge(value bool) error {
	return e.Element.SetProperty("edge", value)
}

// type (float64)
//
// Blur type

func (e *Frei0RFilterIirBlur) GetType() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("type")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterIirBlur) SetType(value float64) error {
	return e.Element.SetProperty("type", value)
}
