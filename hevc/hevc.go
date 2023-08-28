package hevc

import (
	"bytes"
	"fmt"

	"github.com/ossrs/go-oryx-lib/errors"
)

// @doc ISO_IEC_14496-10-AVC-2003.pdf at page 44, 7.3.1 NAL unit syntax
type NALRefIDC uint8

// @doc ISO_IEC_14496-10-AVC-2003.pdf at page 44, 7.3.1 NAL unit syntax
type NALUType uint8

const (
	NALUType_TRAIL_N NALUType = 0 // Coded slice segment of a non-TSA, non_STSA trailing picture
	NALUType_TRAIL_R NALUType = 1 // slice_segment_layer_rbsp()

	NALUType_TSA_N NALUType = 2 // Coded slice segment of a TSA picture
	NALUType_TSA_R NALUType = 3 // slice_segment_layer_rbsp()

	NALUType_STSA_N NALUType = 4 // Coded slice segment of an STSA picture
	NALUType_STSA_R NALUType = 5 // slice_segment_layer_rbsp()

	NALUType_RADL_N NALUType = 6 // Coded slice segment of a RADL picture
	NALUType_RADL_R NALUType = 7 // slice_segment_layer_rbsp()

	NALUType_RASL_N NALUType = 8 // Coded slice segment of a RASL picture
	NALUType_RASL_R NALUType = 9 // slice_segment_layer_rbsp()

	NALUType_RSV_VCL_N10 = 10 // Reserved non-IRAP SLNR VCL NAL unit types
	NALUType_RSV_VCL_R11 = 11 // Reserved non-IRAP sub-layer reference VCL NAL unit types
	NALUType_RSV_VCL_N12 = 12 // Reserved non-IRAP SLNR VCL NAL unit types
	NALUType_RSV_VCL_R13 = 13 // Reserved non-IRAP sub-layer reference VCL NAL unit types
	NALUType_RSV_VCL_N14 = 14 // Reserved non-IRAP SLNR VCL NAL unit types
	NALUType_RSV_VCL_R15 = 15 // Reserved non-IRAP sub-layer reference VCL NAL unit types

	NALUType_BLA_W_LP   NALUType = 16 // Coded slice segment of a BLA picture
	NALUType_BLA_W_RADL NALUType = 17 // slice_segment_layer_rbsp()
	NALUType_BLA_N_LP   NALUType = 18

	NALUType_IDR_W_RADL NALUType = 19 // Coded slice segment of an IDR picture
	NALUType_IDR_N_LP   NALUType = 20 // slice_segment_layer_rbsp()

	NALUType_CRA_NUT NALUType = 21 // Coded slice segment of a CRA picture slice_segment_layer_rbsp()

	NALUType_RSV_IRAP_VCL22 = 22 // Reserved IRAP VCL NAL unit types
	NALUType_RSV_IRAP_VCL23 = 23 // Reserved IRAP VCL NAL unit types

	NALUType_RSV_VCL24 = 24 // Reserved non-IRAP VCL NAL unit types
	NALUType_RSV_VCL25 = 25 // Reserved non-IRAP VCL NAL unit types
	NALUType_RSV_VCL26 = 26 // Reserved non-IRAP VCL NAL unit types
	NALUType_RSV_VCL27 = 27 // Reserved non-IRAP VCL NAL unit types
	NALUType_RSV_VCL28 = 28 // Reserved non-IRAP VCL NAL unit types
	NALUType_RSV_VCL29 = 29 // Reserved non-IRAP VCL NAL unit types
	NALUType_RSV_VCL30 = 30 // Reserved non-IRAP VCL NAL unit types
	NALUType_RSV_VCL31 = 31 // Reserved non-IRAP VCL NAL unit types

	NALUType_VPS_NUT NALUType = 32 // Video parameter set video_parameter_set_rbsp()
	NALUType_SPS_NUT NALUType = 33 // Sequence parameter set seq_parameter_set_rbsp()
	NALUType_PPS_NUT NALUType = 34 // Picture parameter set pic_parameter_set_rbsp()
	NALUType_AUD_NUT NALUType = 35 // Access unit delimiter access_unit_delimiter_rbsp()
	NALUType_EOS_NUT NALUType = 36 // End of sequence end_of_seq_rbsp()
	NALUType_EOB_NUT NALUType = 37 // End of bitstream end_of_bitstream_rbsp()
	NALUType_FD_NUT  NALUType = 38 // Filler data filler_data_rbsp()

	NALUType_Prefix_SEI_NUT NALUType = 39 // Supplemental enhancement information sei_rbsp()
	NALUType_Suffix_SEI_NUT NALUType = 40 // Supplemental enhancement information sei_rbsp()

	NALUType_RSV_NVCL41 = 41 // Reserved
	NALUType_RSV_NVCL42 = 42 // Reserved
	NALUType_RSV_NVCL43 = 43 // Reserved
	NALUType_RSV_NVCL44 = 44 // Reserved
	NALUType_RSV_NVCL45 = 45 // Reserved
	NALUType_RSV_NVCL46 = 46 // Reserved
	NALUType_RSV_NVCL47 = 47 // Reserved

	NALUType_UNSPEC48 = 48 // Unspecified
	NALUType_UNSPEC49 = 49 // Unspecified
	NALUType_UNSPEC50 = 50 // Unspecified
	NALUType_UNSPEC51 = 51 // Unspecified
	NALUType_UNSPEC52 = 52 // Unspecified
	NALUType_UNSPEC53 = 53 // Unspecified
	NALUType_UNSPEC54 = 54 // Unspecified
	NALUType_UNSPEC55 = 55 // Unspecified
	NALUType_UNSPEC56 = 56 // Unspecified
	NALUType_UNSPEC57 = 57 // Unspecified
	NALUType_UNSPEC58 = 58 // Unspecified
	NALUType_UNSPEC59 = 59 // Unspecified
	NALUType_UNSPEC60 = 60 // Unspecified
	NALUType_UNSPEC61 = 61 // Unspecified
	NALUType_UNSPEC62 = 62 // Unspecified
	NALUType_UNSPEC63 = 63 // Unspecified
)

func (v NALUType) String() string {
	switch v {
	case NALUType_TRAIL_N:
		return "TRAIL_N"
	case NALUType_TRAIL_R:
		return "TRAIL_R"
	case NALUType_TSA_N:
		return "TSA_N"
	case NALUType_TSA_R:
		return "TSA_R"
	case NALUType_STSA_N:
		return "STSA_N"
	case NALUType_STSA_R:
		return "STSA_R"
	case NALUType_RADL_N:
		return "RADL_N"
	case NALUType_RADL_R:
		return "RADL_R"
	case NALUType_RASL_N:
		return "RASL_N"
	case NALUType_RASL_R:
		return "RASL_R"
	case NALUType_RSV_VCL_N10:
		return "RSV_VCL_N10"
	case NALUType_RSV_VCL_R11:
		return "RSV_VCL_R11"
	case NALUType_RSV_VCL_N12:
		return "RSV_VCL_N12"
	case NALUType_RSV_VCL_R13:
		return "RSV_VCL_R13"
	case NALUType_RSV_VCL_N14:
		return "RSV_VCL_N14"
	case NALUType_RSV_VCL_R15:
		return "RSV_VCL_R15"
	case NALUType_BLA_W_LP:
		return "BLA_W_LP"
	case NALUType_BLA_W_RADL:
		return "BLA_W_RADL"
	case NALUType_BLA_N_LP:
		return "BLA_N_LP"
	case NALUType_IDR_W_RADL:
		return "IDR_W_RADL"
	case NALUType_IDR_N_LP:
		return "IDR_N_LP"
	case NALUType_CRA_NUT:
		return "CRA_NUT"
	case NALUType_RSV_IRAP_VCL22:
	case NALUType_RSV_IRAP_VCL23:
		return "RSV_IRAP_VCL"
	case NALUType_RSV_VCL24:
	case NALUType_RSV_VCL25:
	case NALUType_RSV_VCL26:
	case NALUType_RSV_VCL27:
	case NALUType_RSV_VCL28:
	case NALUType_RSV_VCL29:
	case NALUType_RSV_VCL30:
	case NALUType_RSV_VCL31:
		return "RSV_VCL"
	case NALUType_VPS_NUT:
		return "VPS_NUT"
	case NALUType_SPS_NUT:
		return "SPS_NUT"
	case NALUType_PPS_NUT:
		return "PPS_NUT"
	case NALUType_AUD_NUT:
		return "AUD_NUT"
	case NALUType_EOS_NUT:
		return "EOS_NUT"
	case NALUType_EOB_NUT:
		return "EOB_NUT"
	case NALUType_FD_NUT:
		return "FD_NUT"
	case NALUType_Prefix_SEI_NUT:
		return "PREFIX_SEI_NUT"
	case NALUType_Suffix_SEI_NUT:
		return "SUFFIX_SEI_NUT"
	case NALUType_RSV_NVCL41:
	case NALUType_RSV_NVCL42:
	case NALUType_RSV_NVCL43:
	case NALUType_RSV_NVCL44:
	case NALUType_RSV_NVCL45:
	case NALUType_RSV_NVCL46:
	case NALUType_RSV_NVCL47:
		return "RSV_NVCL"
	case NALUType_UNSPEC48:
	case NALUType_UNSPEC49:
	case NALUType_UNSPEC50:
	case NALUType_UNSPEC51:
	case NALUType_UNSPEC52:
	case NALUType_UNSPEC53:
	case NALUType_UNSPEC54:
	case NALUType_UNSPEC55:
	case NALUType_UNSPEC56:
	case NALUType_UNSPEC57:
	case NALUType_UNSPEC58:
	case NALUType_UNSPEC59:
	case NALUType_UNSPEC60:
	case NALUType_UNSPEC61:
	case NALUType_UNSPEC62:
	case NALUType_UNSPEC63:
		return "UNSPEC"
	default:
		return fmt.Sprintf("NALU/%v", uint8(v))
	}
	return "Forbbidden"
}

// H.265 NAL Header
type NALUHeader struct {
	// The 1-bit forbidden_bit
	// @remark It's 1 bit.
	Forbidden uint8
	// The 6-bits nal_unit_type.
	// @remark It's 6 bits.
	NALUType NALUType
	// The 6-bits nuh_layer_id.
	// @remark It's 6 bits.
	NUHLayerID uint8
	// The 3-bits nuh_temporal_id_plus1.
	// @remark It's 3 bits.
	NUHTemporalIDPlus1 uint8
}

func NewNALUHeader() *NALUHeader {
	return &NALUHeader{}
}

func (v *NALUHeader) String() string {
	return fmt.Sprintf("%v, LayerID=%v, TemporalIDPlus1=%v", v.NALUType, v.NUHLayerID, v.NUHTemporalIDPlus1)
}

func (v *NALUHeader) Size() int {
	return 2
}

// Unmarshal H.265 NAL Header from bytes.
// @remark user must ensure the bytes left is at least 2.
func (v *NALUHeader) UnmarshalBinary(data []byte) error {
	if len(data) < 2 {
		return errors.New("empty NALU")
	}

	v.Forbidden = uint8(uint8(data[0]>>7) & 0x01)                  // 1 bit
	v.NALUType = NALUType(uint8(data[0]>>1) & 0x3f)                // 6 bits
	v.NUHLayerID = uint8(data[0]&0x01)<<5 | uint8(data[1]>>3)&0x1f // 6 bits
	v.NUHTemporalIDPlus1 = uint8(data[1] & 0x07)                   // 3 bits

	return nil
}

// func (v *NALUHeader) UnmarshalBinary(data []byte) error {
// 	if len(data) < 2 {
// 		return errors.New("empty NALU")
// 	}

//  v.NALRefIDC = NALRefIDC(uint8(data[0]>>5) & 0x03)
// 	v.NALUType = NALUType(uint8(data[0]) & 0x1f)
// 	return nil
// }

// Marshal H.265 NAL Header to bytes.
// @remark user must ensure the bytes left is at least 2.
func (v *NALUHeader) MarshalBinary() ([]byte, error) {
	return []byte{
		byte(v.Forbidden)<<7 | byte(v.NALUType)<<1 | byte(v.NUHLayerID&0x01),
		byte(v.NUHLayerID&0x1f)<<3 | byte(v.NUHTemporalIDPlus1),
	}, nil
}

// func (v *NALUHeader) MarshalBinary() ([]byte, error) {
// 	return []byte{
// 		byte(v.NALRefIDC)<<5 | byte(v.NALUType),
// 	}, nil
// }

// @doc ISO_IEC_14496-10-AVC-2003.pdf at page 60, 7.4.1 NAL unit semantics
type NALU struct {
	*NALUHeader
	Data []byte
}

func NewNALU() *NALU {
	return &NALU{NALUHeader: NewNALUHeader()}
}

func (v *NALU) String() string {
	return fmt.Sprintf("%v, size=%vB", v.NALUHeader, len(v.Data))
}

func (v *NALU) Size() int {
	return 2 + len(v.Data)
}

func (v *NALU) UnmarshalBinary(data []byte) error {
	if err := v.NALUHeader.UnmarshalBinary(data); err != nil {
		return errors.WithMessage(err, "unmarshal")
	}

	v.Data = data[2:]
	return nil
}

func (v *NALU) MarshalBinary() ([]byte, error) {
	b, err := v.NALUHeader.MarshalBinary()
	if err != nil {
		return nil, errors.WithMessage(err, "marshal")
	}

	if len(v.Data) == 0 {
		return b, nil
	}
	return append(b, v.Data...), nil
}

type HEVCProfile uint8

// hevc sequence header
type HEVCDecoderConfigurationRecord struct {
	// It contains the profile code as defined in ISO/IEC 14496-10.
	configurationVersion uint8
	// It is a 2 bits profile_space
	profileSpace uint8
	// It is a byte defined tier_flag
	tierFlag uint8
	// It is a 5 bits profile_idc
	// @remark It's 5 bits.
	HEVCProfileIndication HEVCProfile
	// It is a 32 bits profile_compatibility_indication
	// @remark It's 32 bits.
	profileCompatibilityFlags uint32
	// It is a 48 bits constraint_indicator_flags
	// @remark It's 48 bits.
	constraintIndicatorFlags uint64
	// It is a 8 bits level_idc
	// @remark It's 8 bits.
	levelIndication uint8
	// It is a 4 bits reserved 4 bits
	// @remark It's 4 bits.
	reserved4bits uint8
	// It is a 12 bits min_spatial_segmentation_idc
	// @remark It's 12 bits.
	minSpatialSegmentationIDC uint16
	// It is a 6 bits reserved 6 bits
	// @remark It's 6 bits.
	reserved6bits uint8
	// It is a 2 bits parallelismType
	parallelismType uint8
	// It is a 6 bits reserved 6 bits
	// @remark It's 6 bits.
	reserved6bits2 uint8
	// It is a 2 bits chromaFormat
	chromaFormat uint8
	// It is a 5 bits reserved 5 bits
	// @remark It's 5 bits.
	reserved5bits uint8
	// It is a 3 bits bitDepthLumaMinus8
	bitDepthLumaMinus8 uint8
	// It is a 5 bits reserved 5 bits
	// @remark It's 5 bits.
	reserved5bits2 uint8
	// It is a 3 bits bitDepthChromaMinus8
	bitDepthChromaMinus8 uint8
	// It is a 16 bits avgFrameRate
	avgFrameRate uint16
	// It is a 2 bits constantFrameRate
	constantFrameRate uint8
	// It is a 3 bits numTemporalLayers
	numTemporalLayers uint8
	// It is a 1 bits temporalIdNested
	temporalIdNested uint8
	// It is a 2 bits lengthSizeMinusOne
	LengthSizeMinusOne uint8
	// It is a 8 bits numOfNaluArrays
	numOfNaluArrays uint8
	// It contains a VPS NAL unit, as specified in ISO/IEC 14496-10. VPSs shall occur in
	// order of ascending parameter set identifier with gaps being allowed.
	VideoParameterSetNALUnits []*NALU
	// It contains a SPS NAL unit, as specified in ISO/IEC 14496-10. SPSs shall occur in
	// order of ascending parameter set identifier with gaps being allowed.
	SequenceParameterSetNALUnits []*NALU
	// It contains a PPS NAL unit, as specified in ISO/IEC 14496-10. PPSs shall occur in
	// order of ascending parameter set identifier with gaps being allowed.
	PictureParameterSetNALUnits []*NALU
	// @remark We ignore the sequenceParameterSetExtNALUnit.
}

func NewHEVCDecoderConfigurationRecord() *HEVCDecoderConfigurationRecord {
	v := &HEVCDecoderConfigurationRecord{}
	v.configurationVersion = 0x01
	return v
}

// func (v *HEVCDecoderConfigurationRecord) MarshalBinary() ([]byte, error) {
// 	var buf bytes.Buffer
// 	buf.WriteByte(byte(v.configurationVersion))
// 	buf.WriteByte(byte(v.HEVCProfileIndication))
// 	buf.WriteByte(byte(v.profileCompatibility))
// 	buf.WriteByte(byte(v.AVCLevelIndication))
// 	buf.WriteByte(byte(v.LengthSizeMinusOne))

// 	// numOfSequenceParameterSets
// 	buf.WriteByte(byte(len(v.SequenceParameterSetNALUnits)))
// 	for _, sps := range v.SequenceParameterSetNALUnits {
// 		b, err := sps.MarshalBinary()
// 		if err != nil {
// 			return nil, errors.WithMessage(err, "sps")
// 		}

// 		sequenceParameterSetLength := uint16(len(b))
// 		buf.WriteByte(byte(sequenceParameterSetLength >> 8))
// 		buf.WriteByte(byte(sequenceParameterSetLength))
// 		buf.Write(b)
// 	}

// 	// numOfPictureParameterSets
// 	buf.WriteByte(byte(len(v.PictureParameterSetNALUnits)))
// 	for _, pps := range v.PictureParameterSetNALUnits {
// 		b, err := pps.MarshalBinary()
// 		if err != nil {
// 			return nil, errors.WithMessage(err, "pps")
// 		}

// 		pictureParameterSetLength := uint16(len(b))
// 		buf.WriteByte(byte(pictureParameterSetLength >> 8))
// 		buf.WriteByte(byte(pictureParameterSetLength))
// 		buf.Write(b)
// 	}

// 	return buf.Bytes(), nil
// }

// Unmarshal H.265 sequence header with HEVCDecoderConfigurationRecord from bytes.
// @remark user must ensure the bytes left is at least 23.
func (v *HEVCDecoderConfigurationRecord) UnmarshalBinary(data []byte) error {
	resLen := len(data)
	b := data
	if len(b) < 24 {
		return errors.Errorf("requires 24+ only %v bytes", len(b))
	}

	v.configurationVersion = uint8(b[0])
	v.profileSpace = uint8(uint8(b[1]>>6) & 0x03)
	v.tierFlag = uint8(b[1] >> 5 & 0x01)
	v.HEVCProfileIndication = HEVCProfile(uint8(b[1]) & 0x1f)
	v.profileCompatibilityFlags = uint32(b[2])<<24 | uint32(b[3])<<16 | uint32(b[4])<<8 | uint32(b[5])
	v.constraintIndicatorFlags = uint64(b[6])<<40 | uint64(b[7])<<32 | uint64(b[8])<<24 | uint64(b[9])<<16 | uint64(b[10])<<8 | uint64(b[11])
	v.levelIndication = uint8(b[12])
	v.reserved4bits = uint8(b[13] >> 4 & 0x0f)
	v.minSpatialSegmentationIDC = uint16(b[13]&0x0f)<<8 | uint16(b[14])
	v.reserved6bits = uint8(b[15] >> 2 & 0x3f)
	v.parallelismType = uint8(b[15] & 0x03)
	v.reserved6bits2 = uint8(b[16] >> 2 & 0x3f)
	v.chromaFormat = uint8(b[16] & 0x03)
	v.reserved5bits = uint8(b[17] >> 3 & 0x1f)
	v.bitDepthLumaMinus8 = uint8(b[17] & 0x07)
	v.reserved5bits2 = uint8(b[18] >> 3 & 0x1f)
	v.bitDepthChromaMinus8 = uint8(b[18] & 0x07)
	v.avgFrameRate = uint16(b[19])<<8 | uint16(b[20])
	v.constantFrameRate = uint8(b[21] >> 6 & 0x03)
	v.numTemporalLayers = uint8(b[21] >> 3 & 0x07)
	v.temporalIdNested = uint8(b[21] >> 2 & 0x01)
	v.LengthSizeMinusOne = uint8(b[21] & 0x03)
	v.numOfNaluArrays = uint8(b[22])

	b = b[23:]
	resLen -= 23
	for {
		var (
			offset  = 0
			err     error
			nalType = NALUType(b[0])
		)
		switch nalType {
		case NALUType_VPS_NUT:
			offset, err = v.parseVPS_SPS_PPS(b[1:], NALUType_VPS_NUT)
		case NALUType_SPS_NUT:
			offset, err = v.parseVPS_SPS_PPS(b[1:], NALUType_SPS_NUT)
		case NALUType_PPS_NUT:
			offset, err = v.parseVPS_SPS_PPS(b[1:], NALUType_PPS_NUT)
		default:
			return errors.New("error NALU type")
		}
		if err != nil {
			return err
		}
		b = b[offset+1:]
		resLen -= offset + 1
		if resLen <= 0 {
			break
		}
	}

	return nil
}

func (v *HEVCDecoderConfigurationRecord) parseVPS_SPS_PPS(b []byte, nalType NALUType) (int, error) {
	offset := 0
	numOfVideoParameterSets := uint16(b[0])<<8 | uint16(b[1])
	b = b[2:]
	offset += 2
	for i := 0; i < int(numOfVideoParameterSets); i++ {
		if len(b) < 2 {
			return 0, errors.Errorf("requires 2+ only %v bytes", len(b))
		}
		videoParameterSetLength := int(uint16(b[0])<<8 | uint16(b[1]))
		b = b[2:]
		offset += 2

		if len(b) < videoParameterSetLength {
			return 0, errors.Errorf("requires %v only %v bytes", videoParameterSetLength, len(b))
		}
		nalu := NewNALU()
		if err := nalu.UnmarshalBinary(b[:videoParameterSetLength]); err != nil {
			return 0, errors.WithMessage(err, "unmarshal")
		}
		b = b[videoParameterSetLength:]
		offset += videoParameterSetLength

		switch nalType {
		case NALUType_VPS_NUT:
			v.VideoParameterSetNALUnits = append(v.VideoParameterSetNALUnits, nalu)
		case NALUType_SPS_NUT:
			v.SequenceParameterSetNALUnits = append(v.SequenceParameterSetNALUnits, nalu)
		case NALUType_PPS_NUT:
			v.PictureParameterSetNALUnits = append(v.PictureParameterSetNALUnits, nalu)
		}
	}

	return offset, nil
}

// func (v *HEVCDecoderConfigurationRecord) UnmarshalBinary(data []byte) error {
// 	b := data
// 	if len(b) < 23 {
// 		return errors.Errorf("requires 23+ only %v bytes", len(b))
// 	}

// 	v.configurationVersion = uint8(b[0])
// 	v.HEVCProfileIndication = HEVCProfile(uint8(b[1]))
// 	v.profileCompatibility = uint8(b[2])
// 	v.AVCLevelIndication = AVCLevel(uint8(b[3]))
// 	v.LengthSizeMinusOne = uint8(b[4]) & 0x03
// 	b = b[5:]

// 	numOfSequenceParameterSets := uint8(b[0]) & 0x1f
// 	b = b[1:]
// 	for i := 0; i < int(numOfSequenceParameterSets); i++ {
// 		if len(b) < 2 {
// 			return errors.Errorf("requires 2+ only %v bytes", len(b))
// 		}
// 		sequenceParameterSetLength := int(uint16(b[0])<<8 | uint16(b[1]))
// 		b = b[2:]

// 		if len(b) < sequenceParameterSetLength {
// 			return errors.Errorf("requires %v only %v bytes", sequenceParameterSetLength, len(b))
// 		}
// 		sps := NewNALU()
// 		if err := sps.UnmarshalBinary(b[:sequenceParameterSetLength]); err != nil {
// 			return errors.WithMessage(err, "unmarshal")
// 		}
// 		b = b[sequenceParameterSetLength:]

// 		v.SequenceParameterSetNALUnits = append(v.SequenceParameterSetNALUnits, sps)
// 	}

// 	if len(b) < 1 {
// 		return errors.New("no PPS length")
// 	}
// 	numOfPictureParameterSets := uint8(b[0])
// 	b = b[1:]
// 	for i := 0; i < int(numOfPictureParameterSets); i++ {
// 		if len(b) < 2 {
// 			return errors.Errorf("requiers 2+ only %v bytes", len(b))
// 		}

// 		pictureParameterSetLength := int(uint16(b[0])<<8 | uint16(b[1]))
// 		b = b[2:]

// 		if len(b) < pictureParameterSetLength {
// 			return errors.Errorf("requires %v only %v bytes", pictureParameterSetLength, len(b))
// 		}
// 		pps := NewNALU()
// 		if err := pps.UnmarshalBinary(b[:pictureParameterSetLength]); err != nil {
// 			return errors.WithMessage(err, "unmarshal")
// 		}
// 		b = b[pictureParameterSetLength:]

// 		v.PictureParameterSetNALUnits = append(v.PictureParameterSetNALUnits, pps)
// 	}
// 	return nil
// }

// @doc ISO_IEC_14496-15-AVC-format-2012.pdf at page 20, 5.3.4.2 Sample format
type HEVCSample struct {
	lengthSizeMinusOne uint8
	NALUs              []*NALU
}

func NewHEVCSample(lengthSizeMinusOne uint8) *HEVCSample {
	return &HEVCSample{lengthSizeMinusOne: lengthSizeMinusOne}
}

func (v *HEVCSample) MarshalBinary() ([]byte, error) {
	sizeOfNALU := int(v.lengthSizeMinusOne) + 1

	var buf bytes.Buffer
	for _, nalu := range v.NALUs {
		b, err := nalu.MarshalBinary()
		if err != nil {
			return nil, errors.WithMessage(err, "write")
		}

		length := uint64(len(b))
		for i := 0; i < sizeOfNALU; i++ {
			buf.WriteByte(byte(length >> uint8(8*(sizeOfNALU-1-i))))
		}
		buf.Write(b)
	}

	return buf.Bytes(), nil
}

func (v *HEVCSample) UnmarshalBinary(data []byte) error {
	sizeOfNALU := int(v.lengthSizeMinusOne) + 1
	for b := data; len(b) > 0; {
		if len(b) < sizeOfNALU {
			return errors.Errorf("requires %v+ only %v bytes", sizeOfNALU, len(b))
		}

		var length uint64
		for i := 0; i < sizeOfNALU; i++ {
			length |= uint64(b[i]) << uint8(8*(sizeOfNALU-1-i))
		}
		b = b[sizeOfNALU:]

		if len(b) < int(length) {
			return errors.Errorf("requires %v only %v bytes", length, len(b))
		}

		nalu := NewNALU()
		if err := nalu.UnmarshalBinary(b[:length]); err != nil {
			return errors.WithMessage(err, "unmarshal")
		}
		b = b[length:]

		v.NALUs = append(v.NALUs, nalu)
	}
	return nil
}
