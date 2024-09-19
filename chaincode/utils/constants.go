package utils

const (
	SIGNATURE_LENGTH = 65
	V_LOWER_BOUND    = 27
	V_UPPER_BOUND    = 28
	BIT_SIZE         = 64
	// The special precision -1 uses the smallest number of digits necessary such that ParseFloat will return f exactly.
	PRECISION                            = -1
	WORKERPOOL_ALLOCATION                = 0.3
	VERIFY_DATA_HASH_CHANNEL_BUFFER_SIZE = 100
)
