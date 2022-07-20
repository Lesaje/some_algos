//It does not working

package main

// The history of my struggle is quite interesting.
// I understand the reasons for my defeat, but I still do not understand what made the developers develop the language in this way.
//
// First of all, it's worth saying that I now understand quite well how hash tables are implemented, and I can implement them for specific data types.
// The only problem that may arise is writing a function to redistribute elements to new buckets, however, this should not cause serious problems.
//
// In fact, the problem arises from the fact that the interface comparable cannot act as a data type, only a contraint. 
// At the moment, I don't understand the reasons for this design decision.
// 
// When trying to make a bucket and a key pair through generics, I have problems when adding elements (I bet that in a bunch of other places)
// related to the fact that when creating a generic key pair inside of function Add (which looks similar to the one from this code)
// and trying to add it to the bucket, elements of slice []keyPair[K, V] and variable of type keyPair[K comparable, V any] do not match. 
// I don't exactly know why, but I have some thoughts.
//
// I leave this code for the sake of this comment, and for the future, it would be interesting to look at all this in a couple of months, 
// when my knowledge of the language improves and a couple of updates for Go come out.
// 

import (
    "bytes"
    "crypto/sha256"
    "encoding/binary"
    "encoding/gob"
    "fmt"
)

type keyPair struct {
    key   interface{ comparable }
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
    addB := (1 << newB) - (1 << h.b) - 1 //buckets to add: 2**newB - 2**b - 1
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
    var kP keyPair
    kP.key = key
    kP.value = value

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
    h.buckets[index].append(kP)
    h.loadFactor = float64(h.numOfElements+1) / float64(1<<h.b)
    h.numOfElements += 1
    return nil
}

func (b *bucket) append(kP keyPair) {
    for _, el := range b.slice {
        if el.key == kP.key {
            el.value = kP.value
            return
        }
    }
    b.slice = append(b.slice, kP)
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
