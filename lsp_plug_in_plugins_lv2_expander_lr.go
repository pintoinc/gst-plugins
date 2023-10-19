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

type LspPlugInPluginsLv2ExpanderLr struct {
	*base.GstBaseTransform
}

func NewLspPlugInPluginsLv2ExpanderLr() (*LspPlugInPluginsLv2ExpanderLr, error) {
	e, err := gst.NewElement("lsp-plug-in-plugins-lv2-expander-lr")
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2ExpanderLr{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLspPlugInPluginsLv2ExpanderLrWithName(name string) (*LspPlugInPluginsLv2ExpanderLr, error) {
	e, err := gst.NewElementWithName("lsp-plug-in-plugins-lv2-expander-lr", name)
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2ExpanderLr{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// al-l (float32)
//
// Attack level Left

func (e *LspPlugInPluginsLv2ExpanderLr) GetAlL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ExpanderLr) SetAlL(value float32) error {
	return e.Element.SetProperty("al-l", value)
}

// al-r (float32)
//
// Attack level Right

func (e *LspPlugInPluginsLv2ExpanderLr) GetAlR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ExpanderLr) SetAlR(value float32) error {
	return e.Element.SetProperty("al-r", value)
}

// at-l (float32)
//
// Attack time Left

func (e *LspPlugInPluginsLv2ExpanderLr) GetAtL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ExpanderLr) SetAtL(value float32) error {
	return e.Element.SetProperty("at-l", value)
}

// at-r (float32)
//
// Attack time Right

func (e *LspPlugInPluginsLv2ExpanderLr) GetAtR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ExpanderLr) SetAtR(value float32) error {
	return e.Element.SetProperty("at-r", value)
}

// bypass (bool)
//
// Bypass

func (e *LspPlugInPluginsLv2ExpanderLr) GetBypass() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bypass")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ExpanderLr) SetBypass(value bool) error {
	return e.Element.SetProperty("bypass", value)
}

// cdr-l (float32)
//
// Dry gain Left

func (e *LspPlugInPluginsLv2ExpanderLr) GetCdrL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cdr-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ExpanderLr) SetCdrL(value float32) error {
	return e.Element.SetProperty("cdr-l", value)
}

// cdr-r (float32)
//
// Dry gain Right

func (e *LspPlugInPluginsLv2ExpanderLr) GetCdrR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cdr-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ExpanderLr) SetCdrR(value float32) error {
	return e.Element.SetProperty("cdr-r", value)
}

// clear (bool)
//
// Clear graph analysis

func (e *LspPlugInPluginsLv2ExpanderLr) GetClear() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("clear")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ExpanderLr) SetClear(value bool) error {
	return e.Element.SetProperty("clear", value)
}

// clm-l (float32)
//
// Curve level meter Left

func (e *LspPlugInPluginsLv2ExpanderLr) GetClmL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("clm-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// clm-r (float32)
//
// Curve level meter Right

func (e *LspPlugInPluginsLv2ExpanderLr) GetClmR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("clm-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// cr-l (float32)
//
// Ratio Left

func (e *LspPlugInPluginsLv2ExpanderLr) GetCrL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cr-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ExpanderLr) SetCrL(value float32) error {
	return e.Element.SetProperty("cr-l", value)
}

// cr-r (float32)
//
// Ratio Right

func (e *LspPlugInPluginsLv2ExpanderLr) GetCrR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cr-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ExpanderLr) SetCrR(value float32) error {
	return e.Element.SetProperty("cr-r", value)
}

// cwt-l (float32)
//
// Wet gain Left

func (e *LspPlugInPluginsLv2ExpanderLr) GetCwtL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cwt-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ExpanderLr) SetCwtL(value float32) error {
	return e.Element.SetProperty("cwt-l", value)
}

// cwt-r (float32)
//
// Wet gain Right

func (e *LspPlugInPluginsLv2ExpanderLr) GetCwtR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cwt-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ExpanderLr) SetCwtR(value float32) error {
	return e.Element.SetProperty("cwt-r", value)
}

// elm-l (float32)
//
// Envelope level meter Left

func (e *LspPlugInPluginsLv2ExpanderLr) GetElmL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("elm-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// elm-r (float32)
//
// Envelope level meter Right

func (e *LspPlugInPluginsLv2ExpanderLr) GetElmR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("elm-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// elv-l (bool)
//
// Envelope level visibility Left

func (e *LspPlugInPluginsLv2ExpanderLr) GetElvL() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("elv-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ExpanderLr) SetElvL(value bool) error {
	return e.Element.SetProperty("elv-l", value)
}

// elv-r (bool)
//
// Envelope level visibility Right

func (e *LspPlugInPluginsLv2ExpanderLr) GetElvR() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("elv-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ExpanderLr) SetElvR(value bool) error {
	return e.Element.SetProperty("elv-r", value)
}

// em-l (GstLspPlugInPluginsLv2ExpanderLremL)
//
// Expander mode Left

func (e *LspPlugInPluginsLv2ExpanderLr) GetEmL() (interface{}, error) {
	return e.Element.GetProperty("em-l")
}

func (e *LspPlugInPluginsLv2ExpanderLr) SetEmL(value interface{}) error {
	return e.Element.SetProperty("em-l", value)
}

// em-r (GstLspPlugInPluginsLv2ExpanderLremR)
//
// Expander mode Right

func (e *LspPlugInPluginsLv2ExpanderLr) GetEmR() (interface{}, error) {
	return e.Element.GetProperty("em-r")
}

func (e *LspPlugInPluginsLv2ExpanderLr) SetEmR(value interface{}) error {
	return e.Element.SetProperty("em-r", value)
}

// g-in (float32)
//
// Input gain

func (e *LspPlugInPluginsLv2ExpanderLr) GetGIn() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-in")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ExpanderLr) SetGIn(value float32) error {
	return e.Element.SetProperty("g-in", value)
}

// g-out (float32)
//
// Output gain

func (e *LspPlugInPluginsLv2ExpanderLr) GetGOut() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-out")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ExpanderLr) SetGOut(value float32) error {
	return e.Element.SetProperty("g-out", value)
}

// grv-l (bool)
//
// Gain reduction visibility Left

func (e *LspPlugInPluginsLv2ExpanderLr) GetGrvL() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("grv-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ExpanderLr) SetGrvL(value bool) error {
	return e.Element.SetProperty("grv-l", value)
}

// grv-r (bool)
//
// Gain reduction visibility Right

func (e *LspPlugInPluginsLv2ExpanderLr) GetGrvR() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("grv-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ExpanderLr) SetGrvR(value bool) error {
	return e.Element.SetProperty("grv-r", value)
}

// ilm-l (float32)
//
// Input level meter Left

func (e *LspPlugInPluginsLv2ExpanderLr) GetIlmL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ilm-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// ilm-r (float32)
//
// Input level meter Right

func (e *LspPlugInPluginsLv2ExpanderLr) GetIlmR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ilm-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// ilv-l (bool)
//
// Input level visibility Left

func (e *LspPlugInPluginsLv2ExpanderLr) GetIlvL() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ilv-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ExpanderLr) SetIlvL(value bool) error {
	return e.Element.SetProperty("ilv-l", value)
}

// ilv-r (bool)
//
// Input level visibility Right

func (e *LspPlugInPluginsLv2ExpanderLr) GetIlvR() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ilv-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ExpanderLr) SetIlvR(value bool) error {
	return e.Element.SetProperty("ilv-r", value)
}

// kn-l (float32)
//
// Knee Left

func (e *LspPlugInPluginsLv2ExpanderLr) GetKnL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ExpanderLr) SetKnL(value float32) error {
	return e.Element.SetProperty("kn-l", value)
}

// kn-r (float32)
//
// Knee Right

func (e *LspPlugInPluginsLv2ExpanderLr) GetKnR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ExpanderLr) SetKnR(value float32) error {
	return e.Element.SetProperty("kn-r", value)
}

// mk-l (float32)
//
// Makeup gain Left

func (e *LspPlugInPluginsLv2ExpanderLr) GetMkL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("mk-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ExpanderLr) SetMkL(value float32) error {
	return e.Element.SetProperty("mk-l", value)
}

// mk-r (float32)
//
// Makeup gain Right

func (e *LspPlugInPluginsLv2ExpanderLr) GetMkR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("mk-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ExpanderLr) SetMkR(value float32) error {
	return e.Element.SetProperty("mk-r", value)
}

// olm-l (float32)
//
// Output level meter Left

func (e *LspPlugInPluginsLv2ExpanderLr) GetOlmL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("olm-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// olm-r (float32)
//
// Output level meter Right

func (e *LspPlugInPluginsLv2ExpanderLr) GetOlmR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("olm-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// olv-l (bool)
//
// Output level visibility Left

func (e *LspPlugInPluginsLv2ExpanderLr) GetOlvL() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("olv-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ExpanderLr) SetOlvL(value bool) error {
	return e.Element.SetProperty("olv-l", value)
}

// olv-r (bool)
//
// Output level visibility Right

func (e *LspPlugInPluginsLv2ExpanderLr) GetOlvR() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("olv-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ExpanderLr) SetOlvR(value bool) error {
	return e.Element.SetProperty("olv-r", value)
}

// out-latency (int)
//
// Latency OUT

func (e *LspPlugInPluginsLv2ExpanderLr) GetOutLatency() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("out-latency")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// pause (bool)
//
// Pause graph analysis

func (e *LspPlugInPluginsLv2ExpanderLr) GetPause() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("pause")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ExpanderLr) SetPause(value bool) error {
	return e.Element.SetProperty("pause", value)
}

// rl-l (float32)
//
// Release level Left

func (e *LspPlugInPluginsLv2ExpanderLr) GetRlL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rl-r (float32)
//
// Release level Right

func (e *LspPlugInPluginsLv2ExpanderLr) GetRlR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rlm-l (float32)
//
// Reduction level meter Left

func (e *LspPlugInPluginsLv2ExpanderLr) GetRlmL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rlm-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rlm-r (float32)
//
// Reduction level meter Right

func (e *LspPlugInPluginsLv2ExpanderLr) GetRlmR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rlm-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rrl-l (float32)
//
// Relative release level Left

func (e *LspPlugInPluginsLv2ExpanderLr) GetRrlL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rrl-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ExpanderLr) SetRrlL(value float32) error {
	return e.Element.SetProperty("rrl-l", value)
}

// rrl-r (float32)
//
// Relative release level Right

func (e *LspPlugInPluginsLv2ExpanderLr) GetRrlR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rrl-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ExpanderLr) SetRrlR(value float32) error {
	return e.Element.SetProperty("rrl-r", value)
}

// rt-l (float32)
//
// Release time Left

func (e *LspPlugInPluginsLv2ExpanderLr) GetRtL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ExpanderLr) SetRtL(value float32) error {
	return e.Element.SetProperty("rt-l", value)
}

// rt-r (float32)
//
// Release time Right

func (e *LspPlugInPluginsLv2ExpanderLr) GetRtR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ExpanderLr) SetRtR(value float32) error {
	return e.Element.SetProperty("rt-r", value)
}

// scl-l (bool)
//
// Sidechain listen Left

func (e *LspPlugInPluginsLv2ExpanderLr) GetSclL() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("scl-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ExpanderLr) SetSclL(value bool) error {
	return e.Element.SetProperty("scl-l", value)
}

// scl-r (bool)
//
// Sidechain listen Right

func (e *LspPlugInPluginsLv2ExpanderLr) GetSclR() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("scl-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ExpanderLr) SetSclR(value bool) error {
	return e.Element.SetProperty("scl-r", value)
}

// scm-l (GstLspPlugInPluginsLv2ExpanderLrscmL)
//
// Sidechain mode Left

func (e *LspPlugInPluginsLv2ExpanderLr) GetScmL() (interface{}, error) {
	return e.Element.GetProperty("scm-l")
}

func (e *LspPlugInPluginsLv2ExpanderLr) SetScmL(value interface{}) error {
	return e.Element.SetProperty("scm-l", value)
}

// scm-r (GstLspPlugInPluginsLv2ExpanderLrscmR)
//
// Sidechain mode Right

func (e *LspPlugInPluginsLv2ExpanderLr) GetScmR() (interface{}, error) {
	return e.Element.GetProperty("scm-r")
}

func (e *LspPlugInPluginsLv2ExpanderLr) SetScmR(value interface{}) error {
	return e.Element.SetProperty("scm-r", value)
}

// scp-l (float32)
//
// Sidechain preamp Left

func (e *LspPlugInPluginsLv2ExpanderLr) GetScpL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scp-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ExpanderLr) SetScpL(value float32) error {
	return e.Element.SetProperty("scp-l", value)
}

// scp-r (float32)
//
// Sidechain preamp Right

func (e *LspPlugInPluginsLv2ExpanderLr) GetScpR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scp-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ExpanderLr) SetScpR(value float32) error {
	return e.Element.SetProperty("scp-r", value)
}

// scr-l (float32)
//
// Sidechain reactivity Left

func (e *LspPlugInPluginsLv2ExpanderLr) GetScrL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scr-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ExpanderLr) SetScrL(value float32) error {
	return e.Element.SetProperty("scr-l", value)
}

// scr-r (float32)
//
// Sidechain reactivity Right

func (e *LspPlugInPluginsLv2ExpanderLr) GetScrR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scr-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ExpanderLr) SetScrR(value float32) error {
	return e.Element.SetProperty("scr-r", value)
}

// scs-l (GstLspPlugInPluginsLv2ExpanderLrscsL)
//
// Sidechain source Left

func (e *LspPlugInPluginsLv2ExpanderLr) GetScsL() (interface{}, error) {
	return e.Element.GetProperty("scs-l")
}

func (e *LspPlugInPluginsLv2ExpanderLr) SetScsL(value interface{}) error {
	return e.Element.SetProperty("scs-l", value)
}

// scs-r (GstLspPlugInPluginsLv2ExpanderLrscsR)
//
// Sidechain source Right

func (e *LspPlugInPluginsLv2ExpanderLr) GetScsR() (interface{}, error) {
	return e.Element.GetProperty("scs-r")
}

func (e *LspPlugInPluginsLv2ExpanderLr) SetScsR(value interface{}) error {
	return e.Element.SetProperty("scs-r", value)
}

// sla-l (float32)
//
// Sidechain lookahead Left

func (e *LspPlugInPluginsLv2ExpanderLr) GetSlaL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sla-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ExpanderLr) SetSlaL(value float32) error {
	return e.Element.SetProperty("sla-l", value)
}

// sla-r (float32)
//
// Sidechain lookahead Right

func (e *LspPlugInPluginsLv2ExpanderLr) GetSlaR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sla-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ExpanderLr) SetSlaR(value float32) error {
	return e.Element.SetProperty("sla-r", value)
}

// slm-l (float32)
//
// Sidechain level meter Left

func (e *LspPlugInPluginsLv2ExpanderLr) GetSlmL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("slm-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// slm-r (float32)
//
// Sidechain level meter Right

func (e *LspPlugInPluginsLv2ExpanderLr) GetSlmR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("slm-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// slv-l (bool)
//
// Sidechain level visibility Left

func (e *LspPlugInPluginsLv2ExpanderLr) GetSlvL() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("slv-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ExpanderLr) SetSlvL(value bool) error {
	return e.Element.SetProperty("slv-l", value)
}

// slv-r (bool)
//
// Sidechain level visibility Right

func (e *LspPlugInPluginsLv2ExpanderLr) GetSlvR() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("slv-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2ExpanderLr) SetSlvR(value bool) error {
	return e.Element.SetProperty("slv-r", value)
}

// ----- Constants -----

type LspPlugInPluginsLv2ExpanderLremR string

const (
	LspPlugInPluginsLv2ExpanderLremRDown LspPlugInPluginsLv2ExpanderLremR = "Down" // Down (0) – Down
	LspPlugInPluginsLv2ExpanderLremRUp LspPlugInPluginsLv2ExpanderLremR = "Up" // Up (1) – Up
)

type LspPlugInPluginsLv2ExpanderLrscmL string

const (
	LspPlugInPluginsLv2ExpanderLrscmLPeak LspPlugInPluginsLv2ExpanderLrscmL = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2ExpanderLrscmLRms LspPlugInPluginsLv2ExpanderLrscmL = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2ExpanderLrscmLLowPass LspPlugInPluginsLv2ExpanderLrscmL = "Low-Pass" // Low-Pass (2) – Low-Pass
	LspPlugInPluginsLv2ExpanderLrscmLUniform LspPlugInPluginsLv2ExpanderLrscmL = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2ExpanderLrscmR string

const (
	LspPlugInPluginsLv2ExpanderLrscmRPeak LspPlugInPluginsLv2ExpanderLrscmR = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2ExpanderLrscmRRms LspPlugInPluginsLv2ExpanderLrscmR = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2ExpanderLrscmRLowPass LspPlugInPluginsLv2ExpanderLrscmR = "Low-Pass" // Low-Pass (2) – Low-Pass
	LspPlugInPluginsLv2ExpanderLrscmRUniform LspPlugInPluginsLv2ExpanderLrscmR = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2ExpanderLrscsL string

const (
	LspPlugInPluginsLv2ExpanderLrscsLMiddle LspPlugInPluginsLv2ExpanderLrscsL = "Middle" // Middle (0) – Middle
	LspPlugInPluginsLv2ExpanderLrscsLSide LspPlugInPluginsLv2ExpanderLrscsL = "Side" // Side (1) – Side
	LspPlugInPluginsLv2ExpanderLrscsLLeft LspPlugInPluginsLv2ExpanderLrscsL = "Left" // Left (2) – Left
	LspPlugInPluginsLv2ExpanderLrscsLRight LspPlugInPluginsLv2ExpanderLrscsL = "Right" // Right (3) – Right
)

type LspPlugInPluginsLv2ExpanderLrscsR string

const (
	LspPlugInPluginsLv2ExpanderLrscsRMiddle LspPlugInPluginsLv2ExpanderLrscsR = "Middle" // Middle (0) – Middle
	LspPlugInPluginsLv2ExpanderLrscsRSide LspPlugInPluginsLv2ExpanderLrscsR = "Side" // Side (1) – Side
	LspPlugInPluginsLv2ExpanderLrscsRLeft LspPlugInPluginsLv2ExpanderLrscsR = "Left" // Left (2) – Left
	LspPlugInPluginsLv2ExpanderLrscsRRight LspPlugInPluginsLv2ExpanderLrscsR = "Right" // Right (3) – Right
)

type LspPlugInPluginsLv2ExpanderLremL string

const (
	LspPlugInPluginsLv2ExpanderLremLDown LspPlugInPluginsLv2ExpanderLremL = "Down" // Down (0) – Down
	LspPlugInPluginsLv2ExpanderLremLUp LspPlugInPluginsLv2ExpanderLremL = "Up" // Up (1) – Up
)
