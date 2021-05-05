package hash

import (
	"encoding/binary"
)

// All functions are copied and modified from golang.org/x/crypto/sha3
// This is a specific version of sha3 optimized only for the functions in
// this package and must not be used elsewhere
//
// Copyright (c) 2009 The Go Authors. All rights reserved.

// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:

//    * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//    * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//    * Neither the name of Google Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.

// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

const (
	// rate is size of the internal buffer.
	rate = 136

	// dsbyte contains the "domain separation" bits and the first bit of
	// the padding. Sections 6.1 and 6.2 of [1] separate the outputs of the
	// SHA-3 and SHAKE functions by appending bitstrings to the message.
	// Using a little-endian bit-ordering convention, these are "01" for SHA-3
	// and "1111" for SHAKE, or 00000010b and 00001111b, respectively. Then the
	// padding rule from section 5.1 is applied to pad the message to a multiple
	// of the rate, which involves adding a "1" bit, zero or more "0" bits, and
	// a final "1" bit. We merge the first "1" bit from the padding into dsbyte,
	// giving 00000110b (0x06) and 00011111b (0x1f).
	// [1] http://csrc.nist.gov/publications/drafts/fips-202/fips_202_draft.pdf
	//     "Draft FIPS 202: SHA-3 Standard: Permutation-Based Hash and
	//      Extendable-Output Functions (May 2014)"
	dsbyte     = byte(0x06)
	paddingEnd = uint64(1 << 63)
)

type state struct {
	a [25]uint64 // main state of the hash
}

// New256 creates a new SHA3-256 hash.
// Its generic security strength is 256 bits against preimage attacks,
// and 128 bits against collision attacks.
func new256() *state {
	d := &state{}
	return d
}

// copyOut copies ulint64s to a byte buffer.
func (d *state) copyOut() Hash {
	var out Hash
	for i := 0; i < 4; i++ {
		binary.LittleEndian.PutUint64(out[i<<3:], d.a[i])
	}
	return out
}

func xorInAtIndex(d *state, buf []byte, index int) {
	n := len(buf) >> 3
	aAtIndex := d.a[index:]

	for i := 0; i < n; i++ {
		a := binary.LittleEndian.Uint64(buf)
		aAtIndex[i] ^= a
		buf = buf[8:]
	}
}

func (d *state) hash256Plus(p1 Hash, p2 []byte) Hash {
	//xorIn since p1 length is a multiple of 8
	xorInAtIndex(d, p1[:], 0)
	written := 32 // written uint64s in the state

	for len(p2)+written >= rate {
		xorInAtIndex(d, p2[:rate-written], written>>3)
		keccakF1600(&d.a)
		p2 = p2[rate-written:]
		written = 0 // to avoid
	}

	// xorIn the left over of p2, 64 bits at a time
	for len(p2) >= 8 {
		a := binary.LittleEndian.Uint64(p2[:8])
		d.a[written>>3] ^= a
		p2 = p2[8:]
		written += 8
	}

	var tmp [8]byte
	copy(tmp[:], p2)
	tmp[len(p2)] = dsbyte
	a := binary.LittleEndian.Uint64(tmp[:])
	d.a[written>>3] ^= a

	// the last padding
	d.a[16] ^= paddingEnd

	// permute
	keccakF1600(&d.a)

	// reverse and output
	return d.copyOut()
}

// hash256plus256 absorbs two 256 bits slices of data into the hash's state
// applies the permutation, and outpute the result in out
func (d *state) hash256plus256(p1, p2 Hash) Hash {
	xorIn512(d, p1, p2)
	// permute
	keccakF1600(&d.a)
	// reverse the endianess to the output
	return d.copyOut()
}

// xorIn256 xors two 32 bytes slices into the state; it
// makes no non-portable assumptions about memory layout
// or alignment.
func xorIn512(d *state, buf1, buf2 Hash) {
	sliceBuf1, sliceBuf2 := buf1[:], buf2[:]

	var i int
	for ; i < 4; i++ {
		d.a[i] = binary.LittleEndian.Uint64(sliceBuf1)
		sliceBuf1 = sliceBuf1[8:]
	}
	for ; i < 8; i++ {
		d.a[i] = binary.LittleEndian.Uint64(sliceBuf2)
		sliceBuf2 = sliceBuf2[8:]
	}
	// xor with the dsbyte
	// dsbyte also contains the first one bit for the padding.
	d.a[8] = 0x6
	// xor the last padding bit
	d.a[16] = paddingEnd
}
