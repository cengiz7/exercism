package gigasecond

import "time"

func AddGigasecond(t time.Time) time.Time {
	t = t.Add(time.Second * time.Duration(1e9))
	return t
}