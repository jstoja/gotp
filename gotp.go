package gotp

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"hash"
	"strings"
	"time"
)

func CreateHmacBuffer(password string) hash.Hash {
	var key, _ = base32.StdEncoding.DecodeString(strings.ToUpper(password))
	return (hmac.New(sha1.New, []byte(key)))
}

func GetDigest(hmac_buffer hash.Hash, interval_no int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(interval_no))
	hmac_buffer.Write(buf)
	return (hmac_buffer.Sum(nil))
}

func HotpDigest(dig []byte) uint32 {
	offset := dig[19] & 15
	dt := make([]byte, 4)
	copy(dt, dig[offset:offset+4])
	var ui = binary.BigEndian.Uint32(dt)
	otp := (ui & 0x7fffffff) % 1000000
	return (otp)
}

func Totp(hmac_buffer string) uint32 {
	var interval_no = time.Now().Unix() / 30
	return (Hotp(hmac_buffer, interval_no))
}

func Hotp(hmac_buffer string, interval_no int64) uint32 {
	return (HotpDigest(GetDigest(CreateHmacBuffer(hmac_buffer), interval_no)))
}
