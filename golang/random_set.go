package main

import (
	"time"
)

type RandomizedSet struct {
	indexes map[int]int
	values  []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		indexes: make(map[int]int, 10),
		values:  make([]int, 0, 10),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.indexes[val]; ok {
		return false
	}

	this.values = append(this.values, val)
	this.indexes[val] = len(this.values) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	idx, ok := this.indexes[val]
	if !ok {
		return false
	}
	lastIdx := len(this.values) - 1
	this.values[idx] = this.values[lastIdx]
	this.indexes[this.values[idx]] = idx
	delete(this.indexes, val)
	this.values = this.values[:lastIdx]
	return true
}

var seed uint64 = uint64(time.Now().UnixNano())

func fastRand() uint64 {
	// Простой LCG (Linear Congruential Generator)
	seed = seed*1103515245 + 12345
	return seed
}

func (this *RandomizedSet) GetRandom() int {
	return this.values[fastRand()%uint64(len(this.values))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
