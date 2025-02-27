package fdkaac

/*
#cgo CFLAGS: -I./fdk-aac/libAACdec/include -I./fdk-aac/libSYS/include
#cgo LDFLAGS: ./fdk-aac/libs/libfdk-aac.a -lm
#include "aacdecoder_lib.h"
*/
import "C"

import (
	"fmt"
	"unsafe"
)

type Decoder struct {
	handle     C.HANDLE_AACDECODER
	streamInfo *StreamInfo
}

type DecoderOptions struct {
	TransportType TransportType
}

func NewDecoder(options *DecoderOptions) (*Decoder, error) {
	handle := C.aacDecoder_Open(C.TRANSPORT_TYPE(options.TransportType), C.uint(1))

	if handle == nil {
		return nil, fmt.Errorf("Failed to open decoder")
	}

	return &Decoder{
		handle: handle,
	}, nil
}

func (d *Decoder) Fill(data []byte) (int, error) {
	if len(data) == 0 {
		// if we send no data in it's a nop
		return 0, nil
	}

	if len(data) > 1024*32 {
		panic("fill data must be a maximum of 32k")
	}

	dataPointer := (*C.uchar)(unsafe.Pointer(&data[0]))

	inLen := C.uint(len(data))
	bytesValid := inLen

	err := C.aacDecoder_Fill(d.handle, &dataPointer, &inLen, &bytesValid)
	if err != C.AAC_DEC_OK {
		return 0, fmt.Errorf("failed to fill internal buffer %v", err)
	}

	return int(inLen - bytesValid), nil
}

func (d *Decoder) Decode(target []byte) (int, error) {
	targetPtr := (*C.INT_PCM)(unsafe.Pointer(&target[0]))
	targetLen := C.INT(len(target) / 2)
	errNo := C.aacDecoder_DecodeFrame(d.handle, targetPtr, targetLen, 0)
	if errNo == C.AAC_DEC_NOT_ENOUGH_BITS {
		return 0, nil
	}

	if errNo != C.AAC_DEC_OK {
		return 0, fmt.Errorf("error no %v", errNo)
	}

	if d.streamInfo == nil {
		d.streamInfo = d.GetStreamInfo()
	}

	return d.streamInfo.FrameSize * 2 * d.streamInfo.NumChannels, nil
}

type StreamInfo struct {
	SampleRate  int
	FrameSize   int
	BitRate     int
	OutputDelay int
	NumChannels int
}

func (d *Decoder) GetStreamInfo() *StreamInfo {
	si := C.aacDecoder_GetStreamInfo(d.handle)

	return &StreamInfo{
		SampleRate:  int(si.sampleRate),
		FrameSize:   int(si.frameSize),
		BitRate:     int(si.bitRate),
		OutputDelay: int(si.outputDelay),
		NumChannels: int(si.numChannels),
	}
}

func (d *Decoder) GetFreeBytes() int {
	fb := C.UINT(0)
	C.aacDecoder_GetFreeBytes(d.handle, &fb)

	return int(fb)
}

func (d *Decoder) Close() error {
	C.aacDecoder_Close(d.handle)
	d.handle = nil
	return nil
}
