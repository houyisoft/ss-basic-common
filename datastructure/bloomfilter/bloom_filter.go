
package main

import (
	"github.com/willf/bitset"
	"fmt"
)
const DEFAULT_SIZE = 2<<24
var seeds = []uint{7, 11, 13, 31, 37, 61}

type BloomFilter struct {
	set *bitset.BitSet
	funcs [6]SimpleHash
}

func NewBloomFilter() *BloomFilter {
	bf := new(BloomFilter)
	for i:=0;i< len(bf.funcs);i++{
		bf.funcs[i] = SimpleHash{DEFAULT_SIZE,seeds[i]}
	}
	bf.set = bitset.New(DEFAULT_SIZE)
	return bf
}

func (bf BloomFilter) add(value string){
	for _,f:=range(bf.funcs){
		bf.set.Set(f.hash(value))
	}
}

func (bf BloomFilter) contains(value string) bool  {
	if(value == ""){
		return false
	}
	ret := true
	for _,f:=range(bf.funcs){
		ret = ret && bf.set.Test(f.hash(value))
	}
	return ret
}


type SimpleHash struct{
	cap uint
	seed uint
}

func (s SimpleHash) hash(value string) uint{
	var result uint = 0
	for i:=0;i< len(value);i++{
		result = result*s.seed+uint(value[i])
	}
	return (s.cap-1)&result
}

func main(){
	fmt.Println(2<<24)
	filter := NewBloomFilter()
	fmt.Println(filter.funcs[1].seed)
	str1 := "hello,bloom filter!"
	filter.add(str1)
	str2 := "A happy day"
	filter.add(str2)
	str3 := "Greate wall"
	filter.add(str3)
	fmt.Println(filter.contains(str1))
	fmt.Println(filter.contains(str2))
	fmt.Println(filter.contains(str3))
	fmt.Println(filter.contains("blockchain technology"))
}