package main

func minimumTime(time []int, totalTrips int) int64 {
	l := int64(0)
	r := int64(totalTrips * time[len(time)-1])
	var mid, trip, last int64
	for l <= r {
		mid = (l + r) / 2
		trip = calTrip(time, mid)
		if trip < int64(totalTrips) {
			l = mid + 1
		} else {
			last = mid
			r = mid - 1
		}
	}
	return last
}

func calTrip(time []int, t int64) int64 {
	trip := int64(0)
	for i := 0; i < len(time); i++ {
		trip += t / int64(time[i])
	}
	return trip
}
