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

type Rtphdrextrepairedstreamid struct {
	Element *gst.Element
}

func NewRtphdrextrepairedstreamid() (*Rtphdrextrepairedstreamid, error) {
	e, err := gst.NewElement("rtphdrextrepairedstreamid")
	if err != nil {
		return nil, err
	}
	return &Rtphdrextrepairedstreamid{Element: e}, nil
}

func NewRtphdrextrepairedstreamidWithName(name string) (*Rtphdrextrepairedstreamid, error) {
	e, err := gst.NewElementWithName("rtphdrextrepairedstreamid", name)
	if err != nil {
		return nil, err
	}
	return &Rtphdrextrepairedstreamid{Element: e}, nil
}

// ----- Properties -----

// rid (string)
//
// The RepairedRtpStreamID (RID) value either last retrieved from the RTP
// Header extension, or to set on outgoing RTP packets.

func (e *Rtphdrextrepairedstreamid) GetRid() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("rid")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Rtphdrextrepairedstreamid) SetRid(value string) error {
	return e.Element.SetProperty("rid", value)
}

