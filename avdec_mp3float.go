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

type AvdecMp3Float struct {
	Element *gst.Element
}

func NewAvdecMp3Float() (*AvdecMp3Float, error) {
	e, err := gst.NewElement("avdec_mp3float")
	if err != nil {
		return nil, err
	}
	return &AvdecMp3Float{Element: e}, nil
}

func NewAvdecMp3FloatWithName(name string) (*AvdecMp3Float, error) {
	e, err := gst.NewElementWithName("avdec_mp3float", name)
	if err != nil {
		return nil, err
	}
	return &AvdecMp3Float{Element: e}, nil
}

