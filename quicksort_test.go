package genericshit

import "testing"

func TestTypes(t *testing.T) {
	listInt8 := []int8{1, 2, 3}
	listUint8 := []uint8{1, 2, 3}
	listInt16 := []int16{1, 2, 3}
	listUint16 := []uint16{1, 2, 3}
	listInt32 := []int32{1, 2, 3}
	listUint32 := []uint32{1, 2, 3}
	listInt64 := []int64{1, 2, 3}
	listUint64 := []uint64{1, 2, 3}
	listInt := []int{1, 2, 3}
	listUint := []uint{1, 2, 3}
	listFloat32 := []float32{1, 2, 3}
	listFloat64 := []float64{1, 2, 3}

	Quicksort(listInt8)
	Quicksort(listUint8)
	Quicksort(listInt16)
	Quicksort(listUint16)
	Quicksort(listInt32)
	Quicksort(listUint32)
	Quicksort(listInt64)
	Quicksort(listUint64)
	Quicksort(listInt)
	Quicksort(listUint)
	Quicksort(listFloat32)
	Quicksort(listFloat64)

}

func TestFunction(t *testing.T) {
	vals := []int32{-34123, 8520, -34, 37585, 4737, 29395, 17, 29, 39349, -1234, 589, 39381, 10101, 2383, 1238483, -13431234}

	Quicksort(vals)

	for i := range vals[1:] {
		if vals[i] > vals[i+1] {
			t.Log(vals)
			t.FailNow()
		}
	}
}
