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

type AvencWmav2 struct {
	Element *gst.Element
}

func NewAvencWmav2() (*AvencWmav2, error) {
	e, err := gst.NewElement("avenc_wmav2")
	if err != nil {
		return nil, err
	}
	return &AvencWmav2{Element: e}, nil
}

func NewAvencWmav2WithName(name string) (*AvencWmav2, error) {
	e, err := gst.NewElementWithName("avenc_wmav2", name)
	if err != nil {
		return nil, err
	}
	return &AvencWmav2{Element: e}, nil
}

// ----- Properties -----

// ac (int)
//
// set number of audio channels (Generic codec option, might have no effect)

func (e *AvencWmav2) GetAc() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ac")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencWmav2) SetAc(value int) error {
	return e.Element.SetProperty("ac", value)
}

// ar (int)
//
// set audio sampling rate (in Hz) (Generic codec option, might have no effect)

func (e *AvencWmav2) GetAr() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ar")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencWmav2) SetAr(value int) error {
	return e.Element.SetProperty("ar", value)
}

// audio-service-type (GstAvcodeccontextAudioServiceType)
//
// audio service type (Generic codec option, might have no effect)

func (e *AvencWmav2) GetAudioServiceType() (interface{}, error) {
	return e.Element.GetProperty("audio-service-type")
}

func (e *AvencWmav2) SetAudioServiceType(value interface{}) error {
	return e.Element.SetProperty("audio-service-type", value)
}

// bitrate (int)
//
// set bitrate (in bits/s) (Generic codec option, might have no effect)

func (e *AvencWmav2) GetBitrate() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencWmav2) SetBitrate(value int) error {
	return e.Element.SetProperty("bitrate", value)
}

// bufsize (int)
//
// set ratecontrol buffer size (in bits) (Generic codec option, might have no effect)

func (e *AvencWmav2) GetBufsize() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("bufsize")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencWmav2) SetBufsize(value int) error {
	return e.Element.SetProperty("bufsize", value)
}

// channel-layout (uint64)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencWmav2) GetChannelLayout() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("channel-layout")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *AvencWmav2) SetChannelLayout(value uint64) error {
	return e.Element.SetProperty("channel-layout", value)
}

// compression-level (int)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencWmav2) GetCompressionLevel() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("compression-level")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencWmav2) SetCompressionLevel(value int) error {
	return e.Element.SetProperty("compression-level", value)
}

// cutoff (int)
//
// set cutoff bandwidth (Generic codec option, might have no effect)

func (e *AvencWmav2) GetCutoff() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("cutoff")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencWmav2) SetCutoff(value int) error {
	return e.Element.SetProperty("cutoff", value)
}

// debug (GstAvcodeccontextDebug)
//
// print specific debug info (Generic codec option, might have no effect)

func (e *AvencWmav2) GetDebug() (interface{}, error) {
	return e.Element.GetProperty("debug")
}

func (e *AvencWmav2) SetDebug(value interface{}) error {
	return e.Element.SetProperty("debug", value)
}

// dump-separator (string)
//
// set information dump field separator (Generic codec option, might have no effect)

func (e *AvencWmav2) GetDumpSeparator() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("dump-separator")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *AvencWmav2) SetDumpSeparator(value string) error {
	return e.Element.SetProperty("dump-separator", value)
}

// err-detect (GstAvcodeccontextErrDetect)
//
// set error detection flags (Generic codec option, might have no effect)

func (e *AvencWmav2) GetErrDetect() (interface{}, error) {
	return e.Element.GetProperty("err-detect")
}

func (e *AvencWmav2) SetErrDetect(value interface{}) error {
	return e.Element.SetProperty("err-detect", value)
}

// export-side-data (GstAvcodeccontextExportSideData)
//
// Export metadata as side data (Generic codec option, might have no effect)

func (e *AvencWmav2) GetExportSideData() (interface{}, error) {
	return e.Element.GetProperty("export-side-data")
}

func (e *AvencWmav2) SetExportSideData(value interface{}) error {
	return e.Element.SetProperty("export-side-data", value)
}

// flags (GstAvcodeccontextFlags)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencWmav2) GetFlags() (interface{}, error) {
	return e.Element.GetProperty("flags")
}

func (e *AvencWmav2) SetFlags(value interface{}) error {
	return e.Element.SetProperty("flags", value)
}

// flags2 (GstAvcodeccontextFlags2)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencWmav2) GetFlags2() (interface{}, error) {
	return e.Element.GetProperty("flags2")
}

func (e *AvencWmav2) SetFlags2(value interface{}) error {
	return e.Element.SetProperty("flags2", value)
}

// frame-size (int)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencWmav2) GetFrameSize() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("frame-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencWmav2) SetFrameSize(value int) error {
	return e.Element.SetProperty("frame-size", value)
}

// global-quality (int)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencWmav2) GetGlobalQuality() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("global-quality")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencWmav2) SetGlobalQuality(value int) error {
	return e.Element.SetProperty("global-quality", value)
}

// max-pixels (int64)
//
// Maximum number of pixels (Generic codec option, might have no effect)

func (e *AvencWmav2) GetMaxPixels() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("max-pixels")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

func (e *AvencWmav2) SetMaxPixels(value int64) error {
	return e.Element.SetProperty("max-pixels", value)
}

// max-prediction-order (int)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencWmav2) GetMaxPredictionOrder() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-prediction-order")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencWmav2) SetMaxPredictionOrder(value int) error {
	return e.Element.SetProperty("max-prediction-order", value)
}

// max-samples (int64)
//
// Maximum number of samples (Generic codec option, might have no effect)

func (e *AvencWmav2) GetMaxSamples() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("max-samples")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

func (e *AvencWmav2) SetMaxSamples(value int64) error {
	return e.Element.SetProperty("max-samples", value)
}

// maxrate (int64)
//
// maximum bitrate (in bits/s). Used for VBV together with bufsize. (Generic codec option, might have no effect)

func (e *AvencWmav2) GetMaxrate() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("maxrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

func (e *AvencWmav2) SetMaxrate(value int64) error {
	return e.Element.SetProperty("maxrate", value)
}

// min-prediction-order (int)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencWmav2) GetMinPredictionOrder() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("min-prediction-order")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencWmav2) SetMinPredictionOrder(value int) error {
	return e.Element.SetProperty("min-prediction-order", value)
}

// minrate (int64)
//
// minimum bitrate (in bits/s). Most useful in setting up a CBR encode. It is of little use otherwise. (Generic codec option, might have no effect)

func (e *AvencWmav2) GetMinrate() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("minrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

func (e *AvencWmav2) SetMinrate(value int64) error {
	return e.Element.SetProperty("minrate", value)
}

// side-data-only-packets (bool)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencWmav2) GetSideDataOnlyPackets() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("side-data-only-packets")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *AvencWmav2) SetSideDataOnlyPackets(value bool) error {
	return e.Element.SetProperty("side-data-only-packets", value)
}

// strict (GstAvcodeccontextStrict)
//
// how strictly to follow the standards (Generic codec option, might have no effect)

func (e *AvencWmav2) GetStrict() (interface{}, error) {
	return e.Element.GetProperty("strict")
}

func (e *AvencWmav2) SetStrict(value interface{}) error {
	return e.Element.SetProperty("strict", value)
}

// thread-type (GstAvcodeccontextThreadType)
//
// select multithreading type (Generic codec option, might have no effect)

func (e *AvencWmav2) GetThreadType() (interface{}, error) {
	return e.Element.GetProperty("thread-type")
}

func (e *AvencWmav2) SetThreadType(value interface{}) error {
	return e.Element.SetProperty("thread-type", value)
}

// threads (GstAvcodeccontextThreads)
//
// set the number of threads (Generic codec option, might have no effect)

func (e *AvencWmav2) GetThreads() (interface{}, error) {
	return e.Element.GetProperty("threads")
}

func (e *AvencWmav2) SetThreads(value interface{}) error {
	return e.Element.SetProperty("threads", value)
}

// ticks-per-frame (int)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencWmav2) GetTicksPerFrame() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ticks-per-frame")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencWmav2) SetTicksPerFrame(value int) error {
	return e.Element.SetProperty("ticks-per-frame", value)
}

// trellis (int)
//
// rate-distortion optimal quantization (Generic codec option, might have no effect)

func (e *AvencWmav2) GetTrellis() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("trellis")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencWmav2) SetTrellis(value int) error {
	return e.Element.SetProperty("trellis", value)
}

// ----- Constants -----

type AvcodeccontextAudioServiceType string

const (
	AvcodeccontextAudioServiceTypeMa AvcodeccontextAudioServiceType = "ma" // ma (0) – Main Audio Service
	AvcodeccontextAudioServiceTypeEf AvcodeccontextAudioServiceType = "ef" // ef (1) – Effects
	AvcodeccontextAudioServiceTypeVi AvcodeccontextAudioServiceType = "vi" // vi (2) – Visually Impaired
	AvcodeccontextAudioServiceTypeHi AvcodeccontextAudioServiceType = "hi" // hi (3) – Hearing Impaired
	AvcodeccontextAudioServiceTypeDi AvcodeccontextAudioServiceType = "di" // di (4) – Dialogue
	AvcodeccontextAudioServiceTypeCo AvcodeccontextAudioServiceType = "co" // co (5) – Commentary
	AvcodeccontextAudioServiceTypeEm AvcodeccontextAudioServiceType = "em" // em (6) – Emergency
	AvcodeccontextAudioServiceTypeVo AvcodeccontextAudioServiceType = "vo" // vo (7) – Voice Over
	AvcodeccontextAudioServiceTypeKa AvcodeccontextAudioServiceType = "ka" // ka (8) – Karaoke
)
