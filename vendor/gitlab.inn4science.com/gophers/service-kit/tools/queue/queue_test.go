package queue

import (
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestQueue(t *testing.T) {
	testCases := map[string]struct {
		UID   string
		Value interface{}
	}{
		"test42": {"test42", 42},
		"foo":    {"foo", true},
		"bar":    {"bar", []int{1, 2, 3, 4}},
		"buzz":   {"buzz", "hello, johnny"},
	}

	Convey("Put some data to the Queue", t, func() {
		for _, val := range testCases {
			Put(val.UID, val)
		}

		Convey("The values should be in the Queue index", func() {
			for _, val := range testCases {
				So(IsInQueue(val.UID), ShouldEqual, true)
			}
		})

		Convey("The values should be obtained by uid from the Queue", func() {
			for _, value := range testCases {
				act := GetByID(value.UID)
				So(reflect.DeepEqual(act, value), ShouldEqual, true)
			}
		})

		Convey("The all values should be obtained from the Queue", func() {
			counter := 0
			for i := 0; i < len(testCases)+1; i++ {
				val := GetFirst()
				if val == nil {
					break
				}

				data, ok := val.(struct {
					UID   string
					Value interface{}
				})
				So(ok, ShouldEqual, true)
				So(reflect.DeepEqual(data, testCases[data.UID]), ShouldEqual, true)
				counter++
			}
		})

		Convey("When value was deleted", func() {
			for _, value := range testCases {
				Delete(value.UID)
			}

			Convey("In should be not in the Queue index", func() {
				for _, val := range testCases {
					So(IsInQueue(val.UID), ShouldEqual, false)
				}
			})

			Convey("The values should be not obtained by uid from the Queue", func() {
				for _, value := range testCases {
					So(GetByID(value.UID), ShouldEqual, nil)
				}
			})
		})
	})
}
