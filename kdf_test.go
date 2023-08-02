package main

import (
	"crypto/rand"
	"testing"

	"github.com/alexedwards/argon2id"
	scrypt "github.com/elithrar/simple-scrypt"
	"golang.org/x/crypto/bcrypt"
)

var result []byte
var password []byte

func init() {
	password, _ = GenerateRandomBytes(32)
}

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// err == nil only if len(b) == n
	if err != nil {
		return nil, err
	}

	return b, nil
}

func benchmarkBcrypt(cost int, b *testing.B) {
	var h []byte
	for i := 0; i < b.N; i++ {
		h = Bcrypt(password, cost)
	}
	result = h
}

// m: memory (MB)
// t: iterations
// p: parallelism
func benchmarkArgon(m, t uint32, p uint8, b *testing.B) {
	var h []byte
	for i := 0; i < b.N; i++ {
		h = Argon2(password, &argon2id.Params{
			Memory:      m << 10,
			Iterations:  t,
			Parallelism: p,
			SaltLength:  16,
			KeyLength:   32,
		})
	}
	result = h
}

// N: work factor
// R: memory cost
// P: parallelism
func benchmarkScrypt(n, r, p int, b *testing.B) {
	var h []byte
	for i := 0; i < b.N; i++ {
		h = Scrypt(password, scrypt.Params{
			N:       n,
			R:       r,
			P:       p,
			SaltLen: 16,
			DKLen:   32,
		})
	}
	result = h
}

// Parameters from https://tobtu.com/minimum-password-settings/

func BenchmarkBcryptMin(b *testing.B)     { benchmarkBcrypt(9, b) }
func BenchmarkBcryptDefault(b *testing.B) { benchmarkBcrypt(bcrypt.DefaultCost, b) }
func BenchmarkBcrypt16(b *testing.B)      { benchmarkBcrypt(16, b) }

// Memory usage of scrypt: (128 * r * N).
// BenchmarkScryptMin* is the minimum safe configuration
// BenchmarkScryptDefault is the default configuration

// Memory usage: ~128MB
func BenchmarkScryptMin128(b *testing.B) { benchmarkScrypt(128<<10, 8, 1, b) }

// Memory usage: ~64MB
func BenchmarkScryptMin64(b *testing.B) { benchmarkScrypt(64<<10, 8, 2, b) }

// Memory usage: ~32MB
func BenchmarkScryptMin32(b *testing.B) { benchmarkScrypt(32<<10, 8, 3, b) }

// Memory usage: ~16MB
func BenchmarkScryptMin16(b *testing.B) { benchmarkScrypt(16<<10, 8, 5, b) }

// Memory usage: ~8MB
func BenchmarkScryptMin8(b *testing.B) { benchmarkScrypt(8<<10, 8, 10, b) }

// Memory usage: ~32MB
func BenchmarkScryptDefault(b *testing.B) { benchmarkScrypt(32<<10, 8, 1, b) }

// Memory usage for Argon is set by the memory parameter (m).
// BenchmarkArgonMin* is the minimum safe configuration
// BenchmarkArgonMem* has its memory usage set to the same as scrypt
// BenchmarkArgonDefault is the default configuration
func BenchmarkArgonMin46(b *testing.B)   { benchmarkArgon(46, 1, 1, b) }
func BenchmarkArgonMin19(b *testing.B)   { benchmarkArgon(19, 2, 1, b) }
func BenchmarkArgonMin12(b *testing.B)   { benchmarkArgon(12, 3, 1, b) }
func BenchmarkArgonMin9(b *testing.B)    { benchmarkArgon(9, 4, 1, b) }
func BenchmarkArgonMin7(b *testing.B)    { benchmarkArgon(7, 5, 1, b) }
func BenchmarkArgonMem64(b *testing.B)   { benchmarkArgon(64, 1, 1, b) }
func BenchmarkArgonMem32(b *testing.B)   { benchmarkArgon(32, 2, 1, b) }
func BenchmarkArgonMem16(b *testing.B)   { benchmarkArgon(16, 3, 1, b) }
func BenchmarkArgonMem8(b *testing.B)    { benchmarkArgon(8, 5, 1, b) }
func BenchmarkArgonDefault(b *testing.B) { benchmarkArgon(64, 1, 2, b) }

// https://datatracker.ietf.org/doc/html/rfc9106#section-7.4
// The Argon2id variant with t=1 and 2 GiB memory is the FIRST RECOMMENDED option and is suggested as a default setting for all environments.
// This setting is secure against side-channel attacks and maximizes adversarial costs on dedicated brute-force hardware.
// The Argon2id variant with t=3 and 64 MiB memory is the SECOND RECOMMENDED option and is suggested as a default setting for memory-constrained environments.
//
// Here we try to vary the parallelism to see how if it can be made faster for the defender.
func BenchmarkArgonRecommendedFirst(b *testing.B)   { benchmarkArgon(2*1024, 1, 1, b) }
func BenchmarkArgonRecommendedFirst2(b *testing.B)  { benchmarkArgon(2*1024, 1, 2, b) }
func BenchmarkArgonRecommendedFirst8(b *testing.B)  { benchmarkArgon(2*1024, 1, 8, b) }
func BenchmarkArgonRecommendedSecond(b *testing.B)  { benchmarkArgon(64, 3, 1, b) }
func BenchmarkArgonRecommendedSecond2(b *testing.B) { benchmarkArgon(64, 3, 2, b) }
func BenchmarkArgonRecommendedSecond8(b *testing.B) { benchmarkArgon(64, 3, 8, b) }
