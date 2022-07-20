package main

import (
    "bytes"
    "crypto/sha256"
    "encoding/binary"
    "encoding/gob"
    "fmt"
)

type keyPair[K comparable, V any] struct {
    key   K
    value V
}

type bucket[K comparable, V any] struct {
    arr [8]keyPair[K, V]
}

type HashTable[K comparable, V any] struct {
    buckets       []bucket[K, V]
    b             uint8
    numOfElements uint64
    loadFactor    float64
}

func (h *HashTable[K, V]) addBuckets(newB uint8) {
    addB := (1 << newB) - (1 << h.b) - 1
    for i := 0; i < addB; i++ {
        h.buckets = append(h.buckets, bucket[K, V]{})
    }
    h.b = newB
    h.redistribute()
}

func (h *HashTable[K, V]) redistribute() {
    if h.b == 3 {
        return
    } else {

    }
}

//Find returns value of needed key, if element with such key does not exist, ir returns error
func (h *HashTable[K, V]) Find(key K) (V, error) {

}

//Delete returns error if element with such key does not exists
func (h *HashTable[K, V]) Delete(key K) error {
    return nil
}

func (h *HashTable[K, V]) Add(key K, value V) {
    if h.b == 0 {
        h.addBuckets(3)
        h.loadFactor = 0.0
    }
    if h.loadFactor > 0.5 {
        h.addBuckets(h.b + 1)
    }

    h.loadFactor = float64(h.numOfElements+1) / float64(1<<h.b)
}

func putMask(b uint8, hash [32]byte) uint64 {
    hashId, _ := binary.Uvarint(hash[:8])
    mask := uint64(18446744073709551615)
    mask = mask >> (64 - b)
    return mask & hashId
}

func computeHash(key any) (error, [32]byte) {
    var b bytes.Buffer
    enc := gob.NewEncoder(&b)
    err := enc.Encode(key)
    if err != nil {
        return err, [32]byte{}
    }
    h := sha256.Sum256(b.Bytes())
    return nil, h
}

func main() {
    var h HashTable[string, int64]
    fmt.Println(h)
}
