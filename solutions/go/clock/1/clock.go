package clock

import "strconv"

type Clock int

func New(hour,minute int) (clock Clock) {
	if minute < 0 {
		hour += minute / 60
		minute = minute % 60
		if minute < 0 {
			hour -= 1
			minute += 60
		}
	}
	if hour < 0 {
		hour = hour % 24
		hour += 24
	}
	clock = Clock(Converter(hour, minute))
	return
}

func (clock Clock) Add(minute int) (newClock Clock){
	newClock = Clock( Converter(0,int(clock) + minute))
	return
}

func (clock Clock) Subtract(minute int) (newClock Clock) {
	minute = int(Converter(0,minute))
	if int( Converter(0,int(clock)) ) < minute {
		newClock = Converter(0,1440 - ( minute - int(Converter(0,int(clock)))))
	}else {
		newClock = Converter(0,int(Converter(0,int(clock) - minute)))
	}
	return
}

func Converter(hours,minutes int) (clock Clock){
	hours += minutes/60
	hours = hours % 24
	minutes = minutes % 60
	clock = Clock (hours*60 + minutes)
	return
}

func (clock Clock) String() string {
	hour := strconv.Itoa(int(clock)/60)
	minute := strconv.Itoa(int(clock) % 60)
	return prepareShow(hour) + ":" + prepareShow(minute)
}

func prepareShow(s string) string{
	if len(s) < 2{
		s = "0" + s }
	return  s
}