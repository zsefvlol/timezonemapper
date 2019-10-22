package timezonemapper

import (
	"testing"
	"time"
)

func TestLatLngToTimezoneString(t *testing.T) {
	var timezone string
	timezone = LatLngToTimezoneString(39.9254474,116.3870752)
	if timezone != "Asia/Shanghai" {
		t.Errorf("Wrong timezone result.")
	}
	timezone = LatLngToTimezoneString(-19.1635114,-75.5995365)
	if timezone != "America/Lima" {
		t.Errorf("Wrong timezone result.")
	}
	timezone = LatLngToTimezoneString(-100,116.3870752)
	if timezone != "" {
		t.Errorf("Wrong timezone result.")
	}
}


func TestLatLngToTimezoneStringBenchmark(t *testing.T) {
	lat := -90.0
	lng := -180.0
	step := 0.5
	totalLoops := int64((180 / step) * (360 / step))
	loops := totalLoops
	startTime := time.Now().UnixNano()
	for loops > 0{
		LatLngToTimezoneString(lat,lng)
		lat += step
		lng += step
		loops-=1
	}
	endTime := time.Now().UnixNano()
	t.Logf("\nLoop count: %d\nTime used: %d (nano sec)", totalLoops, endTime-startTime)
}
