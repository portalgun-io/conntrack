// Code generated by "stringer -type=eventType"; DO NOT EDIT.

package conntrack

import "strconv"

const _eventType_name = "EventUnknownEventNewEventUpdateEventDestroyEventExpNewEventExpDestroy"

var _eventType_index = [...]uint8{0, 12, 20, 31, 43, 54, 69}

func (i eventType) String() string {
	if i >= eventType(len(_eventType_index)-1) {
		return "eventType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _eventType_name[_eventType_index[i]:_eventType_index[i+1]]
}
