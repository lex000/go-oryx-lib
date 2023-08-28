// The MIT License (MIT)
//
// Copyright (c) 2013-2017 Oryx(ossrs)
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package hevc_test

import (
	"go-oryx-lib/flv"
	"go-oryx-lib/hevc"
	"io"
	"testing"
)

func TestAvcDecoderAndSample(t *testing.T) {
	// To open a flv file, or http flv stream.
	var r io.Reader
	// r := io.("./h265.flv")
	flvr, _ := flv.NewDemuxer(r)
	tagType, tagSize, tagTS, err := flvr.ReadTagHeader()
	if err != nil {
		t.Errorf("read tag header failed, %v", err.Error())
		return
	}
	t.Logf("tag type=%v, size=%v, ts=%v", tagType, tagSize, tagTS)

	payload, err := flvr.ReadTag(tagSize)
	if err != nil {
		t.Errorf("read tag failed, %v", err.Error())
		return
	}

	// To parse the flv tag.
	vp, _ := flv.NewVideoPackager()
	vf, err := vp.Decode(payload)
	if err != nil {
		t.Errorf("decode video failed, %v", err.Error())
		return
	}

	var lengthSizeMinusOne uint8

	if tagType == flv.TagTypeVideo {
		if vf.Trait == flv.VideoFrameTraitSequenceHeader {
			if vf.CodecID == flv.VideoCodecHEVC {
				hevcCR := hevc.NewHEVCDecoderConfigurationRecord()
				err := hevcCR.UnmarshalBinary(vf.Raw)
				if err != nil {
					t.Errorf("AVCDecoderConfigurationRecord UnmarshalBinary failed, %v", err.Error())
					return
				} else {
					lengthSizeMinusOne = hevcCR.LengthSizeMinusOne
				}
			}
		} else if vf.CodecID == flv.VideoCodecHEVC {
			hevcSample := hevc.NewHEVCSample(lengthSizeMinusOne)
			err := hevcSample.UnmarshalBinary(vf.Raw)
			if err != nil {
				t.Errorf("avcSample UnmarshalBinary failed, %v", err.Error())
			} else {
				for i, nalu := range hevcSample.NALUs {
					if nalu.NALUType == hevc.NALUType_Prefix_SEI_NUT {
						t.Logf("avcSample UnmarshalBinary ok, nalu %v, NALUHeader %+v,  NALUType %v, Data %v", i, nalu.NALUHeader, nalu.NALUType.String(), nalu.Data)
					} else {
						t.Logf("avcSample UnmarshalBinary ok, nalu %v, NALUHeader %+v,  NALUType %v, Data %v", i, nalu.NALUHeader, nalu.NALUType.String(), len(nalu.Data))
					}
				}
			}
		}
	}
}
