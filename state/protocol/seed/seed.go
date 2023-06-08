package seed

import (
	"fmt"

	"github.com/onflow/flow-go/crypto"
	"github.com/onflow/flow-go/crypto/random"
	"golang.org/x/crypto/sha3"
)

const RandomSourceLength = crypto.SignatureLenBLSBLS12381

// PRGFromRandomSource returns a PRG seeded by the input source of randomness.
// The customizer is used to generate a task-specific PRG. A customizer can be any slice
// of up-to-12 bytes.
// The diversifier is used to further diversify the PRGs beyond the customizer. A diversifer
// can be a slice of any length. If no difersification is needed, `diversifier` can be `nil`.
//
// The function uses an extandable-output function (xof) to extract and expand the the input source,
// so that any source with enough entropy (at least 128 bits) can be used (no need to pre-hash).
// Current implementation generates a ChaCha20-based CSPRG.
func PRGFromRandomSource(source []byte, customizer []byte, diversifier []byte) (random.Rand, error) {
	seed, err := xof(source, diversifier, random.Chacha20SeedLen)
	if err != nil {
		return nil, fmt.Errorf("extendable output function failed: %w", err)
	}

	// create random number generator from the seed and customizer
	rng, err := random.NewChacha20PRG(seed, customizer)
	if err != nil {
		return nil, fmt.Errorf("could not create ChaCha20 PRG: %w", err)
	}
	return rng, nil
}

// xof (extendable output function) extracts and expands the input `source` of entropy into
// an output of length `outLen`.
// It also takes a `diversifier` slice as an input to create orthogonal outputs.
//
// Why this function is needed: this function abstracts the extraction and expansion of
// entropy source from the rest of PRG logic. The source doesn't necessarily have a uniformly
// distributed entropy (for instance a cryptographic signature), and hashing doesn't necessarily
// output the number of bytes required by the PRG (the code currently relies on ChaCha20 but this
// choice could change).
func xof(source []byte, diversifier []byte, outLen int) ([]byte, error) {
	// CShake is used in this case but any other primitive that acts as a xof
	// and accepts a diversifier can be used.
	shake := sha3.NewCShake128(nil, diversifier)
	_, _ = shake.Write(source) // cshake Write doesn't error
	out := make([]byte, outLen)
	_, _ = shake.Read(out) // cshake Read doesn't error
	return out, nil
}
