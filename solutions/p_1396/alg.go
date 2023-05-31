package main

import "fmt"

type UndergroundSystem struct {
	checkIns map[int]checkIn
	records  map[string][]int
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{
		checkIns: map[int]checkIn{},
		records:  map[string][]int{},
	}
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	this.checkIns[id] = checkIn{
		id:        id,
		station:   stationName,
		timestamp: t,
	}
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	key := fmt.Sprintf("%s->%s", this.checkIns[id].station, stationName)
	this.records[key] = append(this.records[key], t-this.checkIns[id].timestamp)
	delete(this.checkIns, id)
}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	key := fmt.Sprintf("%s->%s", startStation, endStation)
	records := this.records[key]
	if len(records) < 1 {
		return 0
	}
	sum := 0
	for _, r := range records {
		sum += r
	}
	return float64(sum) / float64(len(records))
}

type checkIn struct {
	id        int
	station   string
	timestamp int
}
