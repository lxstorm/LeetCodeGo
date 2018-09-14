package MergeIntevals

import "sort"

type Interval struct {
	Start int
	End   int
}

func merge(intervals []Interval) []Interval {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i].Start < intervals[j].Start })
	ret := []Interval{}

	for i := 0; i < len(intervals); i++ {
		curInterval := intervals[i]
		if i == 0 {
			ret = append(ret, curInterval)
		} else {
			lastInterval := ret[len(ret)-1]
			if curInterval.Start > lastInterval.End {
				ret = append(ret, curInterval)
			} else {
				if curInterval.End > lastInterval.End {
					ret[len(ret)-1].End = curInterval.End
				}
			}
		}
	}

	return ret

}
