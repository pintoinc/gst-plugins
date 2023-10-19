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

type Rtpvorbispay struct {
	Element *gst.Element
}

func NewRtpvorbispay() (*Rtpvorbispay, error) {
	e, err := gst.NewElement("rtpvorbispay")
	if err != nil {
		return nil, err
	}
	return &Rtpvorbispay{Element: e}, nil
}

func NewRtpvorbispayWithName(name string) (*Rtpvorbispay, error) {
	e, err := gst.NewElementWithName("rtpvorbispay", name)
	if err != nil {
		return nil, err
	}
	return &Rtpvorbispay{Element: e}, nil
}

// ----- Properties -----

// config-interval (uint)
//
// Send Config Insertion Interval in seconds (configuration headers will be multiplexed in the data stream when detected.) (0 = disabled)

func (e *Rtpvorbispay) GetConfigInterval() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("config-interval")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Rtpvorbispay) SetConfigInterval(value uint) error {
	return e.Element.SetProperty("config-interval", value)
}

