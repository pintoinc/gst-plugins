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
	"github.com/tinyzimmer/go-gst/gst"
)

type AvdecAdpcmEaR3 struct {
	Element *gst.Element
}

func NewAvdecAdpcmEaR3() (*AvdecAdpcmEaR3, error) {
	e, err := gst.NewElement("avdec_adpcm_ea_r3")
	if err != nil {
		return nil, err
	}
	return &AvdecAdpcmEaR3{Element: e}, nil
}

func NewAvdecAdpcmEaR3WithName(name string) (*AvdecAdpcmEaR3, error) {
	e, err := gst.NewElementWithName("avdec_adpcm_ea_r3", name)
	if err != nil {
		return nil, err
	}
	return &AvdecAdpcmEaR3{Element: e}, nil
}

