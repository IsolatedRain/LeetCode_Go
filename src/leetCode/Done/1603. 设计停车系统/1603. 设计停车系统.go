// https://leetcode-cn.com/problems/design-parking-system/
package main

import "sync"

type ParkingSystem struct {
	big    int
	medium int
	small  int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{big: big, medium: medium, small: small}
}

var lock sync.Mutex

func (p *ParkingSystem) AddCar(carType int) bool {
	switch carType {
	case 1:
		if p.big > 0 {
			lock.Lock()
			p.big--
			lock.Unlock()
			return true
		}
	case 2:
		if p.medium > 0 {
			lock.Lock()
			p.medium--
			lock.Unlock()
			return true
		}
	case 3:
		if p.small > 0 {
			lock.Lock()
			p.small--
			lock.Unlock()
			return true
		}
	}
	return false
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
