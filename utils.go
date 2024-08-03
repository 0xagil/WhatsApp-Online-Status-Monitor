package main

import "time"

func calculateOnlineRanges(logs []StatusLog) []OnlineRange {
	var ranges []OnlineRange
	var currentRange *OnlineRange

	for _, log := range logs {
		if log.Status == "Online" {
			if currentRange == nil {
				currentRange = &OnlineRange{Start: log.Time}
			}
		} else if log.Status == "Offline" && currentRange != nil {
			currentRange.End = log.Time
			ranges = append(ranges, *currentRange)
			currentRange = nil
		}
	}

	if currentRange != nil {
		currentRange.End = time.Now()
		ranges = append(ranges, *currentRange)
	}

	return ranges
}
