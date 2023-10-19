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

type Ristrtxreceive struct {
	Element *gst.Element
}

func NewRistrtxreceive() (*Ristrtxreceive, error) {
	e, err := gst.NewElement("ristrtxreceive")
	if err != nil {
		return nil, err
	}
	return &Ristrtxreceive{Element: e}, nil
}

func NewRistrtxreceiveWithName(name string) (*Ristrtxreceive, error) {
	e, err := gst.NewElementWithName("ristrtxreceive", name)
	if err != nil {
		return nil, err
	}
	return &Ristrtxreceive{Element: e}, nil
}

// ----- Properties -----

// num-rtx-packets (uint)
//
// Number of retransmission packets received

func (e *Ristrtxreceive) GetNumRtxPackets() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("num-rtx-packets")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// num-rtx-requests (uint)
//
// Number of retransmission events received

func (e *Ristrtxreceive) GetNumRtxRequests() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("num-rtx-requests")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

