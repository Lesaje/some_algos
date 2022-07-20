package main

import (
    "bytes"
    "crypto/sha256"
    "encoding/binary"
    "encoding/gob"
    "fmt"
)

type keyPair struct {
    key   interface{comparable}
    value any
}

type bucket struct {
    slice []keyPair
}

type HashTable[K comparable, V any] struct {
    buckets       []bucket
    b             uint8
    numOfElements uint64
    loadFactor    float64
}

func (h *HashTable[K, V]) addBuckets(newB uint8) {
    addB := (1 << newB) - (1 << h.b) - 1
    for i := 0; i < addB; i++ {
        h.buckets = append(h.buckets, bucket{})
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

func (h *HashTable[K, V]) Add(key K, value V) error {
    if h.b == 0 {
        h.addBuckets(3)
        h.loadFactor = 0.0
    }
    if h.loadFactor > 0.5 {
        h.addBuckets(h.b + 1)
    }
    hsh, err := computeHash(key)
    if err != nil {
        return err
    }
    index := putMask(h.b, hsh)
    var kP keyPair
    kP.key = key
    kP.value = value
    h.buckets[index].slice = append(h.buckets[index].slice, kP)
    h.loadFactor = float64(h.numOfElements+1) / float64(1<<h.b)
    h.numOfElements += 1
    return nil
}

func putMask(b uint8, hash [32]byte) uint64 {
    hashId, _ := binary.Uvarint(hash[:8])
    mask := uint64(18446744073709551615)
    mask = mask >> (64 - b)
    return mask & hashId
}

func computeHash(key any) ([32]byte, error) {
    var b bytes.Buffer
    enc := gob.NewEncoder(&b)
    err := enc.Encode(key)
    if err != nil {
        return [32]byte{}, err
    }
    h := sha256.Sum256(b.Bytes())
    return h, nil
}

func main() {
    var h HashTable[string, int64]
    fmt.Println(h)
}
