package base36

import (
	"encoding/binary"
	"errors"
	"strconv"
)

// An Encoder implements indigo.Encoder interface by Base36.
type Encoder struct {
	encode    [36]byte
	decodeMap [256]int
}

// StdEncoding is Base36 Encoder.
var StdEncoding = MustNewEncoder("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")

// MustNewEncoder returns new base36.Encoder.
func MustNewEncoder(source string) *Encoder {
	enc, err := NewEncoder(source)
	if err != nil {
		panic(err)
	}
	return enc
}

// NewEncoder returns new base36.Encoder.
func NewEncoder(source string) (*Encoder, error) {
	if len(source) != 36 {
		return nil, errors.New("base36: encoding source is not 36-bytes long")
	}
	enc := new(Encoder)

	for i := range enc.decodeMap {
		enc.decodeMap[i] = -1
	}

	for i := range source {
		enc.encode[i] = byte(source[i])
		enc.decodeMap[enc.encode[i]] = i
	}

	return enc, nil
}

// Encode returns encoded string by Base36.
func (enc *Encoder) Encode(id uint64) []byte {
	if id == 0 {
		return enc.encode[:1]
	}

	bin := make([]byte, 0, binary.MaxVarintLen64)
	for id > 0 {
		bin = append(bin, enc.encode[id%36])
		id /= 36
	}

	for i, j := 0, len(bin)-1; i < j; i, j = i+1, j-1 {
		bin[i], bin[j] = bin[j], bin[i]
	}

	return bin
}

// Decode returns decoded unsigned int64 by Base36.
func (enc *Encoder) Decode(id []byte) (uint64, error) {
	if len(id) == 0 {
		return 0, errors.New("base36: id should not be empty")
	}

	var n uint64
	for i := range id {
		u := enc.decodeMap[id[i]]
		if u < 0 {
			return 0, errors.New("base36: invalid character - " + string(id[i]))
		}
		n = n*36 + uint64(u)
	}

	return n, nil
}

func ConverToBianry(n uint64) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.FormatUint(lsb, 36) + result
	}
	return result
}
