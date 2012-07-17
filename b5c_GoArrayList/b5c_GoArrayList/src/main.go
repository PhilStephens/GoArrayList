// GoArrayList project unit tests
package main

import (
	//"bytes" //' was needed by a 'special test', may need again
	"fmt"
	. "goArrayList" //' library/package being tested by main.go
)

// main does a fairly simple set of tests of the implemented ArrayList features; done in the style
// of regression tests, but not exhaustive test coverage (nor tried with multiple OS, just WindowsXP)
func main() {

	fmt.Println("<test> Some tests in main to check Size etc")

	tstALn1 := ArrayListNew(0)
	FlagIfValNE("<test, ArrayListNew(0), IsEmpty>", true, tstALn1.IsEmpty())
	FlagIfValNE("<test, ArrayListNew(0), Size>", 0, tstALn1.Size())
	FlagIfValNE("<test, ArrayListNew(0), Cap>", 0, tstALn1.Cap())

	tstALn2 := ArrayListNew(10)
	FlagIfValNE("<test, ArrayListNew(10), IsEmpty>", true, tstALn2.IsEmpty())
	FlagIfValNE("<test, ArrayListNew(10), Size>", 0, tstALn2.Size())
	FlagIfValNE("<test, ArrayListNew(10), Cap>", 10, tstALn2.Cap())

	tstALn3 := ArrayListPreset([]Obj{1, 2, 3})
	sz, cp := tstALn3.SizeCap()
	FlagIfValNE("<test, ArrayListPreset, IsEmpty>", false, tstALn3.IsEmpty())
	FlagIfValNE("<test, ArrayListPreset, Size>", 3, sz)
	FlagIfValNE("<test, ArrayListPreset, Cap>", 3, cp)

	tstALn3.Clear()
	FlagIfValNE("<test, Clear, IsEmpty>", true, tstALn3.IsEmpty())
	FlagIfValNE("<test, Clear, Size>", 0, tstALn3.Size())
	FlagIfValNE("<test, Clear, Cap>", 3, tstALn3.Cap())

	tstALn4 := ArrayListPreset([]Obj{1, 2, 3})
	sz, cp = tstALn4.SizeCap()
	FlagIfValNE("<test, ArrayListPreset, IsEmpty>", false, tstALn4.IsEmpty())
	FlagIfValNE("<test, ArrayListPreset, Size>", 3, sz)
	FlagIfValNE("<test, ArrayListPreset, Cap>", 3, cp)

	tstALn4.EnsureCapacity(12)
	FlagIfValNE("<test, EnsureCapacity(12), IsEmpty>", false, tstALn4.IsEmpty())
	FlagIfValNE("<test, EnsureCapacity(12), Size>", 3, tstALn4.Size())
	FlagIfValNE("<test, EnsureCapacity(12), Cap>", 12, tstALn4.Cap())

	tstALn4.TrimToSize()
	FlagIfValNE("<test, TrimToSize, IsEmpty>", false, tstALn4.IsEmpty())
	FlagIfValNE("<test, TrimToSize, Size>", 3, tstALn4.Size())
	FlagIfValNE("<test, TrimToSize, Cap>", 3, tstALn4.Cap())

    //' 'mass-produced' set of calls to testSet, skipping type bool; see later for bool w/o testSet

    tmpAry := make([]Obj, 10)
    tmpAryString := []string{"1", "2", "3", "4", "5", "1", "2", "3", "4", "5"}
    for ix, elem := range tmpAryString {
        tmpAry[ix] = elem
    }
    testSet(tmpAry)

    tmpAryInt := []int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
    for ix, elem := range tmpAryInt {
        tmpAry[ix] = elem
    }
    testSet(tmpAry)

    tmpAryInt8 := []int8{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
    for ix, elem := range tmpAryInt8 {
        tmpAry[ix] = elem
    }
    testSet(tmpAry)

    tmpAryInt16 := []int16{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
    for ix, elem := range tmpAryInt16 { tmpAry[ix] = elem }; testSet(tmpAry)

    tmpAryRune := []rune{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
    for ix, elem := range tmpAryRune { tmpAry[ix] = elem }; testSet(tmpAry)

    tmpAryInt64 := []int64{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
    for ix, elem := range tmpAryInt64 { tmpAry[ix] = elem }; testSet(tmpAry)

    tmpAryUint := []uint{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
    for ix, elem := range tmpAryUint { tmpAry[ix] = elem }; testSet(tmpAry)

    tmpAryByte := []byte{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
    for ix, elem := range tmpAryByte { tmpAry[ix] = elem }; testSet(tmpAry)

    tmpAryUint16 := []uint16{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
    for ix, elem := range tmpAryUint16 { tmpAry[ix] = elem }; testSet(tmpAry)

    tmpAryUint32 := []uint32{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
    for ix, elem := range tmpAryUint32 { tmpAry[ix] = elem }; testSet(tmpAry)

    tmpAryUint64 := []uint64{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
    for ix, elem := range tmpAryUint64 { tmpAry[ix] = elem }; testSet(tmpAry)

    tmpAryUintptr := []uintptr{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
    for ix, elem := range tmpAryUintptr { tmpAry[ix] = elem }; testSet(tmpAry)

    tmpAryFloat32 := []float32{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
    for ix, elem := range tmpAryFloat32 { tmpAry[ix] = elem }; testSet(tmpAry)

    tmpAryFloat64 := []float64{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
    for ix, elem := range tmpAryFloat64 { tmpAry[ix] = elem }; testSet(tmpAry)

    tmpAryComplex64 := []complex64{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
    for ix, elem := range tmpAryComplex64 { tmpAry[ix] = elem }; testSet(tmpAry)

    tmpAryComplex128 := []complex128{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
    for ix, elem := range tmpAryComplex128 { tmpAry[ix] = elem }; testSet(tmpAry)
    //  */

	fmt.Printf("\nRunning tests with type bool, inline\n\n")

	//' will need several diffrent presets for different types of test; collected them here
	tttt := strToBoolAsObjAry("tttt")
	ffff := strToBoolAsObjAry("ffff")
	fftf := strToBoolAsObjAry("fftf")
	fftffftf := strToBoolAsObjAry("fftffftf")
	fffttttfff := strToBoolAsObjAry("fffttttfff")
	ffffff := strToBoolAsObjAry("ffffff")
	fftftfft := strToBoolAsObjAry("fftftfft")
	fffftttt := strToBoolAsObjAry("fffftttt")
	ffttttfftttt := strToBoolAsObjAry("ffttttfftttt")

	tstAL := ArrayListPreset(fftf)
	FlagIfValNE("<booltest, Get>", true, tstAL.Get(2))
	FlagIfValNE("<booltest, Contains>a", true, tstAL.Contains(true))

	tstAL = ArrayListPreset(ffff)
	FlagIfValNE("<booltest, Contains>b", false, tstAL.Contains(true))
	FlagIfValNE("<booltest, Contains>c", true, tstAL.Contains(false))

	tstAL = ArrayListPreset(fftffftf)
	FlagIfValNE("<booltest, IndexOf>", 2, tstAL.IndexOf(true))
	FlagIfValNE("<booltest, LastIndexOf>", 6, tstAL.LastIndexOf(true))
	FlagIfValNE("<booltest, Remove(2)>", true, tstAL.Remove(2))

	tstAL.RemoveObj(true)
	FlagIfObjAryNE("<booltest, RemoveObj(true)>", ffffff, tstAL.Ary)

	tstAL = ArrayListPreset(fffttttfff)
	FlagIfObjAryNE("<booltest, before RemoveRange(3,7)>", fffttttfff, tstAL.Ary)

	tstAL.RemoveRange(3, 7)
	FlagIfObjAryNE("<booltest, after RemoveRange(3,7)>", ffffff, tstAL.Ary)

	tstAL.Set(3, true)
	FlagIfValNE("<booltest, Set>", true, tstAL.Get(3))

	tstAL.Append(true)
	FlagIfValNE("<booltest, Set>", true, tstAL.Get(6))

	tstAL.Insert(2, true)
	FlagIfValNE("<booltest, Set>", true, tstAL.Get(2))
	FlagIfObjAryNE("<booltest, Set>", fftftfft, tstAL.Ary)

	tstAL = ArrayListPreset(ffff)
	tstAL.AppendAll(tttt)
	FlagIfObjAryNE("<booltest, ffff after AppendAll(tttt)>", fffftttt, tstAL.Ary)

	tstAL.InsertAll(2, tttt)
	FlagIfObjAryNE("<booltest, fffftttt after InsertAll(2, tttt)>", ffttttfftttt, tstAL.Ary)

} //' main (tests)

// ====================================== funcs for testing ======================================

// testSet is local func to streamline tests of most methods with all 17 array types except bool; 
// includes a series of method tests, output mostly only if error found.
func testSet(tstAry []Obj) {
    //tmpObj := []Obj{} //' for scope
    typestr := fmt.Sprintf("%T", tstAry[0])
    fmt.Println("\nRunning testSet with type ", typestr)
    n := len(tstAry)

    //' (((from ToArrayOmni after the sw-case, but mostly cmtd-out)))
    tstAL := ArrayListPreset(tstAry)
    //' ToArray tests relocated from various case stmts, now that ToArrayInt etc dropped
    tstObj := make([]Obj, n)
    tstObj = tstAL.ToArray(tstObj)
    FlagIfObjAryNE("<testSet, ToArray with "+typestr+"> ", tstAry, tstObj)
    tstObj = nil
    tstObj = tstAL.ToArrayNew()
    FlagIfObjAryNE("<testSet, ToArrayNew with "+typestr+"> ", tstAry, tstObj)

    dblObj := dblr(tstObj) //' special func for testing 'not found' result of search, all types but bool
    tstRef := tstAL.Copy() //' for occasionally restoring to original preset w/o creating a new ArrayList

    //' start of sub-tests
    tgt1 := tstAL.Get(3)
    FlagIfValNE("<testSet, Get & Contains t>", true, tstAL.Contains(tgt1))
    tgt2 := dblObj
    FlagIfValNE("<testSet, Get & Contains f>", false, tstAL.Contains(tgt2))
    FlagIfValNE("<testSet, IndexOf "+fmt.Sprint(tgt1)+">", 3, tstAL.IndexOf(tgt1))
    FlagIfValNE("<testSet, IndexOf "+fmt.Sprint(tgt2)+">", -1, tstAL.IndexOf(tgt2))
    FlagIfValNE("<testSet, LastIndexOf "+fmt.Sprint(tgt1)+">", 8, tstAL.LastIndexOf(tgt1))
    FlagIfValNE("<testSet, LastIndexOf "+fmt.Sprint(tgt2)+">", -1, tstAL.LastIndexOf(tgt2))

    v := tstAL.Remove(2)
    //' assemble expected from subsets of array tstObj; copy not reslice
    x := make([]Obj, n-1)
    copy(x[:2], tstObj[:2])
    copy(x[2:], tstObj[3:])
    FlagIfObjAryNE("<testSet, tstAL.Remove(2) "+fmt.Sprint(v)+">", x, tstAL.Ary)

    //' next 2 ck the bool (found or not), *and* the 'result' (same for both)
    b := tstAL.RemoveObj(tgt1)
    x = make([]Obj, n-2)
    copy(x[:2], tstObj[:2])
    copy(x[2:], tstObj[4:])
    FlagIfValNE("<testSet, RemoveObj "+fmt.Sprint(tgt1)+"; found?>", true, b)
    FlagIfObjAryNE("<testSet, RemoveObj "+fmt.Sprint(tgt1)+"; result>", x, tstAL.Ary)

    b = tstAL.RemoveObj(tgt2)
    FlagIfValNE("<testSet, RemoveObj "+fmt.Sprint(tgt2)+"; found?>", false, b)
    FlagIfObjAryNE("<testSet, RemoveObj "+fmt.Sprint(tgt2)+"; result>", x, tstAL.Ary)

    x = make([]Obj, n-4)
    copy(x[:2], tstObj[:2])
    copy(x[2:], tstObj[6:])
    tstAL.RemoveRange(2, 4)
    FlagIfObjAryNE("<testSet, RemoveRange(2,4), tgt 5,1>", x, tstAL.Ary)

    tstAL = tstRef.Copy()

    v = tstAL.Get(2)
    tstAL.Append(v)
    x = make([]Obj, n+1)
    copy(x[:n], tstObj)
    copy(x[n:], tstObj[2:3])
    FlagIfValNE("<testSet, tstAL.Append("+fmt.Sprint(v)+"> value appended", tstObj[2], v)
    FlagIfObjAryNE("<testSet, RemoveRange(2,4), tgt 5,1>", x, tstAL.Ary)

    x = make([]Obj, n+2)
    copy(x[:4], tstObj[:4])
    copy(x[4:5], tstObj[2:3])
    copy(x[5:n+1], tstObj[4:])
    copy(x[n+1:], tstObj[2:3])
    tstAL.Insert(4, v)
    FlagIfObjAryNE("<testSet, tstAL.Insert(4,"+fmt.Sprint(v)+">", x, tstAL.Ary)

    xx := make([]Obj, n+5)
    copy(xx[:n+2], x)
    copy(xx[n+2:], tstObj[:3])
    shortAry := tstObj[:3] //' reslice is safe here
    tstAL.AppendAll(shortAry)
    FlagIfObjAryNE("<testSet, tstAL.AppendAll("+fmt.Sprint(shortAry)+">", xx, tstAL.Ary)

    tstAL.Set(6, v)
    FlagIfValNE("<testSet, tstAL.tstAL.Set(6, "+fmt.Sprint(v)+")> ", tstObj[2], tstAL.Get(6))

    return
} //' testSet

// dblr is a local convenience utility called by testSet; doubles most types, flips bool,
// concatenates string to itself; complains if unknown type
func dblr(inary []Obj) (v Obj) {
	n := len(inary)
	v = inary[n-1]
	switch v.(type) {
	case bool:
		fmt.Println("Type bool should be tested separately w/o testSet & dblr")
	case string:
		v = v.(string) + v.(string)
	//' would have been nice to lump all the rest into one case, but did not find a way
	case int:
		v = 2 * v.(int)
	case int8:
		v = 2 * v.(int8)
	case int16:
		v = 2 * v.(int16)
	case rune:
		v = 2 * v.(rune)
	case int64:
		v = 2 * v.(int64)
	case uint:
		v = 2 * v.(uint)
	case byte:
		v = 2 * v.(byte)
	case uint16:
		v = 2 * v.(uint16)
	case uint32:
		v = 2 * v.(uint32)
	case uint64:
		v = 2 * v.(uint64)
	case uintptr:
		v = 2 * v.(uintptr)
	case float32:
		v = 2 * v.(float32)
	case float64:
		v = 2 * v.(float64)
	case complex64:
		v = 2 * v.(complex64)
	case complex128:
		v = 2 * v.(complex128)
		//  */
	default:
		fmt.Printf("dblr cannot handle type %T\n", v)
	} //' switch-case
	return
} //' dblr

// strToBoolAsObjAry is a local convenience util to convert a string of t's & f's to true's & false's
// in a []Obj with elements of type bool
func strToBoolAsObjAry(str string) []Obj {
	ba := []byte(str)
	bo := make([]Obj, len(ba))
	for ix, elem := range ba {
		bo[ix] = (elem == 't')
	}
	return bo
}

// FlagIfValNE takes place of assertEquals(msg, xpctd, actl) for a scalar
func FlagIfValNE(msg string, expected, actual interface{}) {
	if expected != actual {
		fmt.Printf("%s expected [%v], found [%v]\n", msg, expected, actual)
	}
}

// FlagIfObjAryNE takes place of assertArrayEquals(msg, ary1, ary2) for int arrays
func FlagIfObjAryNE(msg string, expected, actual []Obj) {
	FlagIfValNE(msg+"Array length mismatch, checking to shorter anyway", len(expected), len(actual))
	for i := 0; i < MinInt(len(actual), len(expected)); i++ {
		FlagIfValNE(msg+" Array mismatch at index "+fmt.Sprint(i)+":", expected[i], actual[i])
	}
}

// func MinInt is trivial, just to avoid multiple casts required by use of math.Min on pair of int
func MinInt(a, b int) (c int) {
	c = a
	if c > b {
		c = b
	}
	return
}
