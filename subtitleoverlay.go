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

type Subtitleoverlay struct {
	Element *gst.Element
}

func NewSubtitleoverlay() (*Subtitleoverlay, error) {
	e, err := gst.NewElement("subtitleoverlay")
	if err != nil {
		return nil, err
	}
	return &Subtitleoverlay{Element: e}, nil
}

func NewSubtitleoverlayWithName(name string) (*Subtitleoverlay, error) {
	e, err := gst.NewElementWithName("subtitleoverlay", name)
	if err != nil {
		return nil, err
	}
	return &Subtitleoverlay{Element: e}, nil
}

// ----- Properties -----

// font-desc (string)
//
// Pango font description of font to be used for subtitle rendering

func (e *Subtitleoverlay) GetFontDesc() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("font-desc")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Subtitleoverlay) SetFontDesc(value string) error {
	return e.Element.SetProperty("font-desc", value)
}

// silent (bool)
//
// Whether to show subtitles

func (e *Subtitleoverlay) GetSilent() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("silent")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Subtitleoverlay) SetSilent(value bool) error {
	return e.Element.SetProperty("silent", value)
}

// subtitle-encoding (string)
//
// Encoding to assume if input subtitles are not in UTF-8 encoding. If not set, the GST_SUBTITLE_ENCODING environment variable will be checked for an encoding to use. If that is not set either, ISO-8859-15 will be assumed.

func (e *Subtitleoverlay) GetSubtitleEncoding() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("subtitle-encoding")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Subtitleoverlay) SetSubtitleEncoding(value string) error {
	return e.Element.SetProperty("subtitle-encoding", value)
}

// subtitle-ts-offset (int64)
//
// The synchronisation offset between text and video in nanoseconds

func (e *Subtitleoverlay) GetSubtitleTsOffset() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("subtitle-ts-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

func (e *Subtitleoverlay) SetSubtitleTsOffset(value int64) error {
	return e.Element.SetProperty("subtitle-ts-offset", value)
}
