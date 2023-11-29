package leetcode

import "math/rand"

// https://leetcode.com/problems/insert-delete-getrandom-o1/?envType=study-plan-v2&envId=top-interview-150

type RandomizedSet struct {
	set      map[int]int
	setSlice []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		set:      make(map[int]int),
		setSlice: []int{},
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, exists := this.set[val]
	if exists {
		return false
	}

	this.set[val] = len(this.setSlice)
	this.setSlice = append(this.setSlice, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	index, exists := this.set[val]
	if !exists {
		return false
	}

	lastIndex := len(this.setSlice) - 1
	lastValue := this.setSlice[lastIndex]
	this.setSlice[index] = lastValue
	this.set[lastValue] = index

	this.setSlice = this.setSlice[:lastIndex]

	delete(this.set, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	length := len(this.setSlice)
	if length == 0 {
		return 0
	}
	randomIndex := rand.Intn(length)
	return this.setSlice[randomIndex]
}

/**
* Your RandomizedSet object will be instantiated and called as such:
* obj := Constructor();
* param_1 := obj.Insert(val);
* param_2 := obj.Remove(val);
* param_3 := obj.GetRandom();
 */
