package uuid7

import (
	"crypto/rand"
	"encoding/base32"
	"encoding/hex"
	"math/big"
	"time"
)

type UUID [16]byte

func MustNew() UUID {
	return must(NewWithTime(time.Now()))
}

func MustNewWithTime(unixMilli time.Time) UUID {
	return must(NewWithTime(unixMilli))
}

func New() (UUID, error) {
	return NewWithTime(time.Now())
}

func NewWithTime(t time.Time) (UUID, error) {
	var bytes [16]byte
	if _, err := rand.Read(bytes[:]); err != nil {
		return bytes, err
	}
	timestamp := big.NewInt(t.UnixMilli())
	timestamp.FillBytes(bytes[0:6])
	bytes[6] = (bytes[6] & 0x0f) | 0x70 // version 7
	bytes[8] = (bytes[8] & 0x3f) | 0x80 // variant 10
	return bytes, nil
}

func (u UUID) String() string {
	return u.DashedHexString()
}

func (u UUID) HexString() string {
	return hex.EncodeToString(u[:])
}

func (u UUID) DashedHexString() string {
	var buf [36]byte
	encodeToDashedHex(buf[:], u)
	return string(buf[:])
}

func (u UUID) B32HexString() string {
	return base32.HexEncoding.WithPadding(base32.NoPadding).EncodeToString(u[:])
}

func (u UUID) B32StdString() string {
	return base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(u[:])
}

func (u UUID) Time() time.Time {
	return time.UnixMilli(int64(big.NewInt(0).SetBytes(u[0:6]).Uint64()))
}

func encodeToDashedHex(dst []byte, uuid UUID) {
	hex.Encode(dst, uuid[:4])
	dst[8] = '-'
	hex.Encode(dst[9:13], uuid[4:6])
	dst[13] = '-'
	hex.Encode(dst[14:18], uuid[6:8])
	dst[18] = '-'
	hex.Encode(dst[19:23], uuid[8:10])
	dst[23] = '-'
	hex.Encode(dst[24:], uuid[10:])
}

func must[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}
	return val
}
