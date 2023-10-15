package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"hash"
)

type Bloom struct {
	bitArr []uint64
	hash1  hash.Hash
	hash2  hash.Hash
}

// Insert implements Ibloom.
func (b *Bloom) Insert(name string) *Bloom {
	//implement a hash function to hash the incoming string to index in the bloom
	// in bloom filter
	_, err := b.hash1.Write([]byte(name))
	h1 := b.hash1.Sum(nil)
	if err != nil {
		panic("error while SHA-256")
	}

	_, err = b.hash2.Write([]byte(name))
	h2 := b.hash2.Sum(nil)
	if err != nil {
		panic("error while SHA-256")
	}

	b.bitArr[int(h2[0])] = 1
	b.bitArr[int(h1[0])] = 1
	return b
}

func newBloom() *Bloom {
	return &Bloom{
		bitArr: make([]uint64, 256),
		hash1:  sha256.New(),
		hash2:  sha512.New(),
	}
}

// IsPresent implements Ibloom.
func (b *Bloom) IsPresent(name string) bool {
	//check if the hash exists in the bloom
	b.hash1 = sha256.New()
	b.hash2 = sha256.New()
	_, err := b.hash1.Write([]byte(name))
	h1 := b.hash1.Sum(nil)
	if err != nil {
		panic("error while SHA-256")
	}

	_, err = b.hash2.Write([]byte(name))
	h2 := b.hash2.Sum(nil)
	if err != nil {
		panic("error while SHA-256")
	}

	if b.bitArr[int(h2[0])] != 1 && b.bitArr[int(h1[0])] != 1 {
		return false
	}

	return true
}

// Len implements Ibloom.
func (b *Bloom) Len() {

}

type Ibloom interface {
	IsPresent(name string) bool
	Len()
	Insert(name string) *Bloom
}

var _ Ibloom = &Bloom{}
