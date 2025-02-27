package fdkaac

type TransportType int

var (
	TransportTypeUnknown       TransportType = -1 // Unknown format.
	TransportTypeMP4_RAW       TransportType = 0  // "as is" access units (packet based since there is obviously no sync layer)
	TransportTypeMP4_ADIF      TransportType = 1  // ADIF bitstream format.
	TransportTypeMP4_ADTS      TransportType = 2  // ADTS bitstream format.
	TransportTypeMP4_LATM_MCP1 TransportType = 6  // Audio Mux Elements with muxConfigPresent = 1
	TransportTypeMP4_LATM_MCP0 TransportType = 7  // Audio Mux Elements with muxConfigPresent = 0, out of band StreamMuxConfig
	TransportTypeMP4_LOAS      TransportType = 10 // Audio Sync Stream.
	TransportTypeDRM           TransportType = 12 // Digital Radio Mondial (DRM30/DRM+) bitstream format.
)
