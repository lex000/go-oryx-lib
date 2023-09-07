package sei

import (
	"errors"
	"fmt"
)

const (
	SEIType_BufferingPeriod                   = 0
	SEIType_PicTiming                         = 1
	SEIType_PanScanRect                       = 2
	SEIType_Filler                            = 3
	SEIType_UserDataRegistered                = 4
	SEIType_UserDataUnregistered              = 5
	SEIType_RecoveryPoint                     = 6
	SEIType_DecRefPicMarkingRepetition        = 7
	SEIType_SparePic                          = 8
	SEIType_SceneInfo                         = 9
	SEIType_SubSeqInfo                        = 10
	SEIType_SubSeqLayerCharacteristics        = 11
	SEIType_SubSeqCharacteristics             = 12
	SEIType_FullFrameFreeze                   = 13
	SEIType_FullFrameFreezeRelease            = 14
	SEIType_FullFrameSnapshot                 = 15
	SEIType_ProgressiveRefinementSegmentStart = 16
	SEIType_ProgressiveRefinementSegmentEnd   = 17
	SEIType_MotionConstrainedSliceGroupSet    = 18
	SEIType_FilmGrainCharacteristics          = 19
	SEIType_DeblockingFilterDisplayPreference = 20
	SEIType_StereoVideoInfo                   = 21
	SEIType_PostFilterHint                    = 22
	SEIType_ToneMappingInfo                   = 23
	SEIType_ScalabilityInfo                   = 24
	SEIType_SubPicScalableLayer               = 25
	SEIType_NonRequiredLayerRep               = 26
	SEIType_PriorityLayerInfo                 = 27
	SEIType_LayersNotPresent                  = 28
	SEIType_LayerDependencyChange             = 29
	SEIType_ScalableNesting                   = 30
	SEIType_BaseLayerTemporalHrd              = 31
	SEIType_QualityLayerIntegrityCheck        = 32
	SEIType_RedundantPicProperty              = 33
	SEIType_TemporalLayerSwitchingPoint       = 34
)

type SEIHeader struct {
	SeiType uint8 //1 byte
	SeiSize uint8 // 1 byte
}

type SEI struct {
	*SEIHeader
	data []byte
}

type SEIUserData struct {
	SeiUUID       []byte // 16 bytes
	SliceID       uint32 // 4 bytes
	TagsNum       uint16 // 2 bytes
	SourceID      uint32 // 4 bytes
	UnixTimestamp uint32 // 4 bytes
}

func NewSEIHeader() *SEIHeader {
	return &SEIHeader{}
}

func (v *SEIHeader) String() string {
	return fmt.Sprintf("sei type=%v, size=%v", v.SeiType, v.SeiSize)
}

func NewSEI() *SEI {
	return &SEI{SEIHeader: NewSEIHeader()}
}

func (s *SEI) UnmarshalHeader(data []byte) error {
	if len(data) < 2 {
		return errors.New("invalid sei size")
	}
	s.SeiType = uint8(data[0])
	s.SeiSize = uint8(data[1])

	return nil
}

func (s *SEI) UnmarshalBinary(data []byte) error {
	if len(data) < 2 {
		return errors.New("invalid sei size")
	}
	realData := make([]byte, 0)

	// 如果连续三个字节为0x000003，跳过其中的0x03，同时i递增2
	for i := 0; i < len(data); i++ {
		if i+2 < len(data) && data[i] == 0x00 && data[i+1] == 0x00 && data[i+2] == 0x03 {
			realData = append(realData, data[i])
			realData = append(realData, data[i+1])
			i += 2
		} else {
			realData = append(realData, data[i])
		}
	}

	s.UnmarshalHeader(realData)

	s.data = realData[2:]

	return nil
}

func (s *SEI) UnmarshalUnregisteredUserData() (*SEIUserData, error) {
	if s.SeiType != 5 {
		return nil, errors.New("invalid sei type")
	}

	if len(s.data) < 30 {
		return nil, errors.New("invalid sei size")
	}

	ud := &SEIUserData{}
	ud.SeiUUID = s.data[0:16]
	ud.SliceID = uint32(s.data[16])<<24 | uint32(s.data[17])<<16 | uint32(s.data[18])<<8 | uint32(s.data[19])
	ud.TagsNum = uint16(s.data[20])<<8 | uint16(s.data[21])
	ud.SourceID = uint32(s.data[22])<<24 | uint32(s.data[23])<<16 | uint32(s.data[24])<<8 | uint32(s.data[25])
	ud.UnixTimestamp = uint32(s.data[26])<<24 | uint32(s.data[27])<<16 | uint32(s.data[28])<<8 | uint32(s.data[29])

	return ud, nil
}
