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

type Tcpclientsrc struct {
	*base.GstPushSrc
}

func NewTcpclientsrc() (*Tcpclientsrc, error) {
	e, err := gst.NewElement("tcpclientsrc")
	if err != nil {
		return nil, err
	}
	return &Tcpclientsrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

func NewTcpclientsrcWithName(name string) (*Tcpclientsrc, error) {
	e, err := gst.NewElementWithName("tcpclientsrc", name)
	if err != nil {
		return nil, err
	}
	return &Tcpclientsrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

// ----- Properties -----

// host (string)
//
// The host IP address to receive packets from

func (e *Tcpclientsrc) GetHost() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("host")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Tcpclientsrc) SetHost(value string) error {
	return e.Element.SetProperty("host", value)
}

// port (int)
//
// The port to receive packets from

func (e *Tcpclientsrc) GetPort() (int, error) {
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

func (e *Tcpclientsrc) SetPort(value int) error {
	return e.Element.SetProperty("port", value)
}

// stats (GstStructure)
//
// Retrieve a statistics structure

func (e *Tcpclientsrc) GetStats() (interface{}, error) {
	return e.Element.GetProperty("stats")
}

// timeout (uint)
//
// Value in seconds to timeout a blocking I/O. 0 = No timeout.

func (e *Tcpclientsrc) GetTimeout() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("timeout")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Tcpclientsrc) SetTimeout(value uint) error {
	return e.Element.SetProperty("timeout", value)
}
