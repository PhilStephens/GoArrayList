package main

import (
	//"bytes"
	"fmt"
	//"strings"
	//"reflect"
)

/* Created 2012.06.21 to develop more general version of Java ArrayList emulated in Go, separate from
   the project w/i which I had started to develop a smaller subset that I actually needed for that 
   project (int elements only, and about a third of the Java features); related notes file has notes 
   from nearly a week earlier, well before got multi-type interface working.
   As of this date, don't expect to get all features of Java ArrayList perfectly emulated in Go, but 
   intend to get as many as possible implemented as faithfully as possible prior to making this package 
   available to other Go programmers. (Before then, I'll need to learn how to make a distribution, 
   where to upload to, etc).

   2012.07.03 
   Nearly all features of Java ArrayList now emulated; imperfectly but in useful form.
   Skipped field modCount as it appears to be a lot of work for a feature of which I could not find 
   usage examples.  Replaced Java ArrayList.clone (which does essentially what '=' or ':=' would do in
   Go, a new slice pointing to the same underlying array) with .Copy, returns a new ArrayList with 
   elements copied (changes to elements of the copy will NOT change elements of the original).  Did not
   find a way to return arrays with conveniently selected types, but provided a way to convert them;
   ToArrayObj returns a new array of type Obj (like object in Java, is an empty interface in Go) with
   elements in some type such as int; ToArray returns updated input array (of wrapper type Omni) if
   large enough, or a new array of same type if ArrayList is larger.
   This file includes my acceptance tests; not perfect, but what I considered adequate before releasing
   to other Go programmers.  I did not use gotest because I have not found a way to use it in goclipse
   for Windows; not sure whether I'm missing something that would allow it.  I also did not arrange for
   all tests to have convenient assertion-like pass/fail, as I have done in other projects for the sake
   of regression testing, mainly because this project is simpler and less subject to dynamic change.
   IMPORTANT: If others have corrections or improvements to either the ArrayList or its test, they are 
   welcome to share them directly or to funnel such changes through me.  I am not claiming future 
   control of this package, but willing to do the work of coordinating if that is appropriate.
   NOT ready to release today; this comment is a rough draft for later release notes.
   2012.07.05 
   Thought of a way to make the tests 'go-nogo', no need for visual inspection; decided to go ahead 
   with that prior to attempting git/github etc


*/

//' main does a fairly simple set of tests of the implemented ArrayList features
func main() {

	//' new tests as part of an over-all test-plan prior to possible release to others

	//fmt.Println("<test> Expect next 4 lines to show t 0 0, t 0 10, f 3 3, t 0 3")
	fmt.Println("<test> Some tests in main to check Size etc")

	tstALn1 := ArrayListNew(0)
	//fmt.Println("<test,ArrayListNew(0)> IsEmpty, Size, Cap", tstALn1.IsEmpty(), tstALn1.Size(), tstALn1.Cap())
	FlagIfValNE("<test, ArrayListNew(0), IsEmpty>", true, tstALn1.IsEmpty())
	FlagIfValNE("<test, ArrayListNew(0), Size>", 0, tstALn1.Size())
	FlagIfValNE("<test, ArrayListNew(0), Cap>", 0, tstALn1.Cap())

	tstALn2 := ArrayListNew(10)
	//fmt.Println("<test,ArrayListNew(10)> IsEmpty, Size, Cap", tstALn2.IsEmpty(), tstALn2.Size(), tstALn2.Cap())
	FlagIfValNE("<test, ArrayListNew(10), IsEmpty>", true, tstALn2.IsEmpty())
	FlagIfValNE("<test, ArrayListNew(10), Size>", 0, tstALn2.Size())
	FlagIfValNE("<test, ArrayListNew(10), Cap>", 10, tstALn2.Cap())

	tstALn3 := ArrayListPreset([]Obj{1, 2, 3})
	sz, cp := tstALn3.SizeCap()
	//fmt.Println("<test,ArrayListPreset> IsEmpty, Size, Cap", tstALn3.IsEmpty(), sz, cp)
	FlagIfValNE("<test, ArrayListPreset, IsEmpty>", false, tstALn3.IsEmpty())
	FlagIfValNE("<test, ArrayListPreset, Size>", 3, sz)
	FlagIfValNE("<test, ArrayListPreset, Cap>", 3, cp)

	tstALn3.Clear()
	//fmt.Println("<test,Clear> IsEmpty, Size, Cap", tstALn3.IsEmpty(), tstALn3.Size(), tstALn3.Cap())
	FlagIfValNE("<test, Clear, IsEmpty>", true, tstALn3.IsEmpty())
	FlagIfValNE("<test, Clear, Size>", 0, tstALn3.Size())
	FlagIfValNE("<test, Clear, Cap>", 3, tstALn3.Cap())

	//fmt.Println("<test> Expect next 3 lines to show f 3 3, f 3 12, f 3 3")
	tstALn4 := ArrayListPreset([]Obj{1, 2, 3})
	sz, cp = tstALn4.SizeCap()
	//fmt.Println("<test,ArrayListPreset> IsEmpty, Size, Cap", tstALn4.IsEmpty(), sz, cp)
	FlagIfValNE("<test, ArrayListPreset, IsEmpty>", false, tstALn4.IsEmpty())
	FlagIfValNE("<test, ArrayListPreset, Size>", 3, sz)
	FlagIfValNE("<test, ArrayListPreset, Cap>", 3, cp)

	tstALn4.EnsureCapacity(12)
	//fmt.Println("<test,EnsureCapacity(12)> IsEmpty, Size, Cap", tstALn4.IsEmpty(), tstALn4.Size(), tstALn4.Cap())
	FlagIfValNE("<test, EnsureCapacity(12), IsEmpty>", false, tstALn4.IsEmpty())
	FlagIfValNE("<test, EnsureCapacity(12), Size>", 3, tstALn4.Size())
	FlagIfValNE("<test, EnsureCapacity(12), Cap>", 12, tstALn4.Cap())

	tstALn4.TrimToSize()
	//fmt.Println("<test,TrimToSize> IsEmpty, Size, Cap", tstALn4.IsEmpty(), tstALn4.Size(), tstALn4.Cap())
	FlagIfValNE("<test, TrimToSize, IsEmpty>", false, tstALn4.IsEmpty())
	FlagIfValNE("<test, TrimToSize, Size>", 3, tstALn4.Size())
	FlagIfValNE("<test, TrimToSize, Cap>", 3, tstALn4.Cap())

	//' now on to the testSet calls

	tstWrap := Omni{}
	//fmtString := "<diag.tstAL.toAryT> input %T; output %T, output element %T; content: %v\n"
	//...
	//tstWrap.aryInt = []int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5} //' Omni to skip copy array before call test func
	//testSet(tstWrap, tstWrap.aryInt[0])

	//' tmp fake print just to avoid 'not used' msg
	//fmt.Println(fmtString)

	//' 'mass-produced' set of calls to testSet, skipping type bool; see later

	tstWrap.aryString = []string{"1", "2", "3", "4", "5", "1", "2", "3", "4", "5"}
	testSet(tstWrap, tstWrap.aryString[0])

	tstWrap.aryInt = []int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
	testSet(tstWrap, tstWrap.aryInt[0])

	tstWrap.aryInt8 = []int8{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
	testSet(tstWrap, tstWrap.aryInt8[0])

	tstWrap.aryInt16 = []int16{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
	testSet(tstWrap, tstWrap.aryInt16[0])

	//tstWrap.aryInt32 = []Int32{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
	//testSet(tstWrap, tstWrap.aryInt32[0])

	//' byte and uint8 or rune and int32 not allowed in same switch-case or struct,
	//' one alias of other so would count as redundant
	tstWrap.aryRune = []rune{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
	testSet(tstWrap, tstWrap.aryRune[0])

	tstWrap.aryInt64 = []int64{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
	testSet(tstWrap, tstWrap.aryInt64[0])

	tstWrap.aryUint = []uint{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
	testSet(tstWrap, tstWrap.aryUint[0])

	//tstWrap.aryUint8 = []uint8{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
	//testSet(tstWrap, tstWrap.aryUint8[0])

	//' byte and uint8 or rune and int32 not allowed in same switch-case or struct,
	//' one alias of other so would count as redundant
	tstWrap.aryByte = []byte{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
	testSet(tstWrap, tstWrap.aryByte[0])

	tstWrap.aryUint16 = []uint16{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
	testSet(tstWrap, tstWrap.aryUint16[0])

	tstWrap.aryUint32 = []uint32{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
	testSet(tstWrap, tstWrap.aryUint32[0])

	tstWrap.aryUint64 = []uint64{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
	testSet(tstWrap, tstWrap.aryUint64[0])

	tstWrap.aryUintptr = []uintptr{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
	testSet(tstWrap, tstWrap.aryUintptr[0])

	tstWrap.aryFloat32 = []float32{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
	testSet(tstWrap, tstWrap.aryFloat32[0])

	tstWrap.aryFloat64 = []float64{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
	testSet(tstWrap, tstWrap.aryFloat64[0])

	tstWrap.aryComplex64 = []complex64{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
	testSet(tstWrap, tstWrap.aryComplex64[0])

	tstWrap.aryComplex128 = []complex128{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
	testSet(tstWrap, tstWrap.aryComplex128[0])

	/*
	    - general plan for testing bool
	      - do {Get Contains IndexOf LastIndexOf Remove RemoveObj}
	        on the one true in a list otherwise all false, and vice versa;  
	      - set up a list with a contrasting patch in the middle, and do RemoveRange on that
	      - do {Set Append Insert AppendAll InsertAll} with similar contrasts
	      - also check ToArray
	//  */

	//fmt.Printf("\nRunning tests with type bool; cannot use testSet, so modified & did inline\n\n")
	fmt.Printf("\nRunning tests with type bool, inline\n\n")

	//' will need several diffrent presets for different types of test

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
	//fmt.Println("<booltest, Get> expect true:", tstAL.Get(2), ", on AL:", tstAL)
	FlagIfValNE("<booltest, Get>", true, tstAL.Get(2))
	//fmt.Println("<booltest, Contains> expect true:", tstAL.Contains(true), ", on AL:", tstAL)
	FlagIfValNE("<booltest, Contains>a", true, tstAL.Contains(true))

	tstAL = ArrayListPreset(ffff)
	//fmt.Println("<booltest, Contains> expect false:", tstAL.Contains(true), ", on AL:", tstAL)
	FlagIfValNE("<booltest, Contains>b", false, tstAL.Contains(true))
	//fmt.Println("<booltest, Contains> expect true:", tstAL.Contains(false), ", on AL:", tstAL)
	FlagIfValNE("<booltest, Contains>c", true, tstAL.Contains(false))

	tstAL = ArrayListPreset(fftffftf)
	//fmt.Println("<booltest, IndexOf> expect 2:", tstAL.IndexOf(true), ", on AL:", tstAL)
	//fmt.Println("<booltest, LastIndexOf> expect 6:", tstAL.LastIndexOf(true), ", on AL:", tstAL)
	//fmt.Println("<booltest, Remove(2)> expect rmvd value true:", tstAL.Remove(2), ", on AL:", tstAL)
	FlagIfValNE("<booltest, IndexOf>", 2, tstAL.IndexOf(true))
	FlagIfValNE("<booltest, LastIndexOf>", 6, tstAL.LastIndexOf(true))
	FlagIfValNE("<booltest, Remove(2)>", true, tstAL.Remove(2))

	tstAL.RemoveObj(true)
	//fmt.Println("<booltest, RemoveObj(true)> expect ffffff in AL:", tstAL)
	FlagIfObjAryNE("<booltest, RemoveObj(true)>", ffffff, tstAL.ary)

	tstAL = ArrayListPreset(fffttttfff)
	//fmt.Println("<booltest, before RemoveRange(3,7)> fffttttfff in AL:", tstAL)
	FlagIfObjAryNE("<booltest, before RemoveRange(3,7)>", fffttttfff, tstAL.ary)

	tstAL.RemoveRange(3, 7)
	//fmt.Println("<booltest, after RemoveRange(3,7)> expect ffffff in AL:", tstAL)
	FlagIfObjAryNE("<booltest, after RemoveRange(3,7)>", ffffff, tstAL.ary)

	tstAL.Set(3, true)
	//fmt.Println("<booltest, after Set(3,true)> expect ffftff in AL:", tstAL)
	FlagIfValNE("<booltest, Set>", true, tstAL.Get(3))

	tstAL.Append(true)
	//fmt.Println("<booltest, after Append(true)> expect ffftfft in AL:", tstAL)
	FlagIfValNE("<booltest, Set>", true, tstAL.Get(6))

	tstAL.Insert(2, true)
	//fmt.Println("<booltest, after Insert(2,true)> expect fftftfft in AL:", tstAL)
	FlagIfValNE("<booltest, Set>", true, tstAL.Get(2))
	FlagIfObjAryNE("<booltest, Set>", fftftfft, tstAL.ary)

	tstAL = ArrayListPreset(ffff)
	tstAL.AppendAll(tttt)
	//fmt.Println("<booltest, ffff after AppendAll(tttt)> expect fffftttt in AL:", tstAL)
	FlagIfObjAryNE("<booltest, ffff after AppendAll(tttt)>", fffftttt, tstAL.ary)

	tstAL.InsertAll(2, tttt)
	//fmt.Println("<booltest, fffftttt after InsertAll(2, tttt)> expect ffttttfftttt in AL:", tstAL)
	FlagIfObjAryNE("<booltest, fffftttt after InsertAll(2, tttt)>", ffttttfftttt, tstAL.ary)

    //' inline test of ToArrayBool
    tmpObj := ffttttfftttt
    n := len(tmpObj)
    tstBoolAx := make([]bool, n) //' expected array
    tstBoolAr := make([]bool, n) //' result array
    for ix, elem := range tmpObj { tstBoolAx[ix] = elem.(bool) }
    tstAL = ArrayListPreset(tmpObj)
    tstBoolAr = tstAL.ToArrayBool(tstBoolAr)
    FlagIfBoolAryNE("<testSet, ToArrayInt>", tstBoolAx, tstBoolAr)



	//' --> all passed

	/*
	    //' special test, research rather than acceptance test: are nils allowed?  Prefer 'yes'
	    //' "TBD: check whether nil allowed in an element, as null is allowed in Java ArrayList";
	    //' also check allowed in []Obj, as that is input to ArrayListPreset, and in Obj as that is
	    //' input to various other methods and funcs

	    //tmpO := (Obj)nil //' unexpected name...
	    //tmpO := nil.(Obj) //' use of untyped nil
	    tmpOA := []Obj{nil, nil}
	    tmpO := tmpOA[0]
	    tstAL = ArrayListPreset(tmpOA)
	    tstAL.Append(tmpO)
	    fmt.Println("<test nil> expect 3 nil:", tstAL)

	    //' tmp fake print just to avoid 'not used' msg
	    //fmt.Println(tmpO,tmpOA)
	    //fmt.Println(tmpOA)
	//  */

	/*
	    //' special test, research rather than acceptance test: how to deal with 'aliases' byte & rune?
	    //'     byte is alias for uint8 and rune is alias for int32

	    //' these 'still exist', here for reference only
	    //tstWrap.aryUint8 = []uint8{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
	    //tstWrap.aryInt32 = []int32{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}

	    //tmpAui8 := []uint8{1, 2, 3}
	    tmpAb := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0} //' must be at least as large as tstWrap.aryUint8
	    tmpAo :=  []Obj{0, 0, 0, 0, 0, 0, 0, 0, 0, 0} //' must be at least as large as tstWrap.aryUint8
	    for ix, elem := range tstWrap.aryUint8 { tmpAo[ix] = elem }
	    tstAL = ArrayListPreset(tmpAo)
	    fmt.Println("<diag.tstAL.before ToArray> tstAL", tstAL)
	    //tmpAui8 = tstAL.ToArray(tmpAui8)
	    tmpAo = tstAL.ToArrayObj()
	    //for ix, elem := range tmpAo { tmpAb[ix] = (byte)elem } //' 'unexpected name'
	    for ix, elem := range tmpAo { tmpAb[ix] = elem.(byte) } //' fails to force type byte
	    //for ix, elem := range tmpAo { tmpAb[ix] = elem } //' worse, 'need type assertion'
	    fmt.Printf("<diag.tstAL.after ToArray and conversion to byte> tmpAb: %v, type: %t\n", tmpAb, tmpAb[0])

	    //' perhaps the conclusion is: "no need to convert, is really identical"? ...can test that,
	    //' eg with almost any method or func from pkg "bytes"
	    fmt.Println("<diag.bytes> bytes.Contains(tmpAb, tmpAb[1:2]): ", bytes.Contains(tmpAb, tmpAb[1:2]))
	    //' ...yes, works fine, so change the comment w/i Omni
	//  */


} //' main (tests)

//' Obj interface matches any type; named to reflect similarity to Java 'object' in ArrayList etc
type Obj interface{}

/*
- from 'tour of go' slide 23 sidebar, list of 'basic types'
----------------------------------------------------
bool
string
int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr
byte // alias for uint8
rune // alias for int32
     // represents a Unicode code point
float32 float64
complex64 complex128
----------------------------------------------------
*/

//' struct Omni contains named arrays of all basic types, eg aryint []int, aryfloat64 []float64, etc; 
//' in a sw-case, ToArray loads the appropriate one with values from base.ary
//' Note: no need for workaround for 2 excluded aliases byte and rune, as they can be used as-is
type Omni struct {
	//' from slide 23 in the 'Tour of Go', paraphrased slightly: "Go's basic types are bool string int 
	//' int8 int16 int32 int64 uint uint8 uint16 uint32 uint64 uintptr float32 float64 complex64 complex128" 
	//' plus (2 aliases that compiler considers redundant here) "byte //alias for uint8" & 
	//' "rune //alias for int32, represents a Unicode code point"
	//' ...later decided to go with the 'alias' versions byte and rune rather than uint8 & int32
	aryObj []Obj //' for default case in ToArray() sw-case

	aryBool   []bool
	aryString []string
	aryInt    []int
	aryInt8   []int8
	aryInt16  []int16
	//' byte and uint8 or rune and int32 not allowed in same switch-case or struct,
	//' one alias of other so would count as redundant
	//aryInt32      []int32
	aryInt64 []int64
	aryUint  []uint
	//aryUint8      []uint8
	aryUint16     []uint16
	aryUint32     []uint32
	aryUint64     []uint64
	aryUintptr    []uintptr
	aryFloat32    []float32
	aryFloat64    []float64
	aryComplex64  []complex64
	aryComplex128 []complex128
	//' byte and rune not allowed in ToArray switch-case, as are duplicates of uint8 & int32;
	//' Note: no need for workaround for 2 excluded aliases byte and rune, as they can be used as-is
	aryByte []byte
	aryRune []rune
}

//' Java documents describe ArrayList as: "Resizable-array implementation of the List interface".
//' ArrayList struct in Go contains an array of type Obj, can contain any type of elements (this is
//' comparable to Java spec, which states that ArrayList can only contain objects, not primatives);
//' for conversion to an array of same type as elements, see ToArray method.
//' Note: nil is allowed in an element, much as null is allowed in Java ArrayList
//' Skipped: field modCount not implemented; could be added if there is interest
type ArrayList struct {
	ary []Obj
}

//' ArrayListPreset creates new ArrayList and preloads it with contents of preset, preserving their
//' type (per entry, not per array/slice; see ToArray method for array/slice matching entry type)
//' Comparable to Java constructor "ArrayList(Collection<? extends E> c)"
func ArrayListPreset(preset []Obj) (AL *ArrayList) { //' func, not method, so no 'base'
	//psz := len(preset)
	//AL = &ArrayList{preset, psz, psz}
	AL = &ArrayList{preset}
	return
} //' ArrayListPreset

//' func ArrayListNew returns (creates) a new empty ArrayList with specified capacity; emulates Java
//' constructor "ArrayList(int initialCapacity)"; to emulate Java constructor "ArrayList()", call
//' ArrayListNew with acap param 10.
func ArrayListNew(acap int) (AL *ArrayList) {
	//tmp := make([]Obj, acap)
	tmp := make([]Obj, 0, acap) //' create with size 0 and capacity acap
	AL = &ArrayList{tmp}
	return
}

//' Clears slice not array, via re-slice; does not clear values, but this is usually harmless;
//' if want to actually destroy old values, use New instead)
func (base *ArrayList) Clear() {
	base.ary = base.ary[0:0] //' Note: '[0:0]' looks like a slice of 1, but is actually empty
	return
}

//' Method Append emulates the append version of Java ArrayList.add; uses Go built-in 'append', which
//' handles cap increase if needed; always returns 'true', since that is the behaviour in Java
func (base *ArrayList) Append(val Obj) bool {
	base.ary = append(base.ary, val)
	return true //' (as per the general contract of Collection.add)
}

//' Method Insert emulates the insert version of Java ArrayList.add; uses Go built-in 'append', which
//' handles cap increase if needed; inserts 'before' entry already in position pos
func (base *ArrayList) Insert(pos int, val Obj) {
	oldlen := len(base.ary)
	base.ary = append(base.ary, 0) //' expand first with dummy zero value
	if pos < oldlen {              //' shift following entries before put new one
		copy(base.ary[pos+1:], base.ary[pos:])
	}
	base.ary[pos] = val
	return
}

//' Get returns element at position pos (type determined by element)
func (base *ArrayList) Get(pos int) (val Obj) {
	val = base.ary[pos]
	return
}

//' Set overwrites element at position pos (type determined by param; consistency is left to
//' programmer); returns value removed, can ignore if not needed
func (base *ArrayList) Set(pos int, val Obj) Obj {
	//' for better emulation of Java ArrayList
	oldval := base.ary[pos]
	base.ary[pos] = val
	return oldval
}

//' Remove emulates indexed version of Java ArrayList.remove; returns value removed, can ignore if not needed
func (base *ArrayList) Remove(pos int) Obj {
	oldlen := len(base.ary)
	val := base.ary[pos]
	copy(base.ary[pos:], base.ary[pos+1:])
	//' needs to return reduced len, so 're-slice' base.ary
	base.ary = base.ary[:oldlen-1]
	return val
}

//' EnsureCapacity emulates Java ArrayList.ensureCapacity; chose to copy last element and append it
//' to guarentee type matches, then reslice to old length
func (base *ArrayList) EnsureCapacity(minCapacity int) {
	oldlen := len(base.ary)
	delta := minCapacity - oldlen
	if minCapacity > cap(base.ary) { //' need to increase size by amt delta, then reslice to oldlen
		elem := base.ary[len(base.ary)-1] //' copy last element for sake of type match, value ignored
		for a := 0; a < delta; a++ {      //' kludgey, I know; suggestions welcome
			base.ary = append(base.ary, elem)
		}
		base.ary = base.ary[:oldlen]
	}
}

//' Returns true if this list contains no elements
func (base *ArrayList) IsEmpty() bool {
	if len(base.ary) < 1 {
		return true
	}
	return false
}

//' Size emulates Java ArrayList.size, returns slice size; for Go version I chose to also implement
//' Cap to return slice capacity and SizeCap to return both at once
func (base *ArrayList) Size() int {
	return len(base.ary)
}

//' for Go version I chose to implement Cap to return slice capacity
func (base *ArrayList) Cap() int { //' not in Java, but makes sense for Go
	return cap(base.ary)
}

//' for Go version I chose to implement SizeCap to return both slice size and slice cap
func (base *ArrayList) SizeCap() (int, int) { //' convenience method
	return len(base.ary), cap(base.ary)
}

//' TrimToSize emulates Java ArrayList.trimToSize, trims slice cap to slice len; actually saves old
//' values to a tmp array, makes a new one as base.ary, and copies old values to new -- use sparingly
func (base *ArrayList) TrimToSize() {
	oldlen := len(base.ary)
	tmpa := make([]Obj, oldlen)
	copy(tmpa, base.ary)
	base.ary = make([]Obj, oldlen, oldlen) //' two copies; open to suggestions if a better way
	copy(base.ary, tmpa)
	return
}

//' InsertAll replaces offset case of Java's addAll; called by InsertAll so must handle offset to end.
//' Also must panic if ary nil, and return false if ary empty
func (base *ArrayList) InsertAll(pos int, ary []Obj) bool {
	if ary == nil {
		panic("ArrayList.InsertAll input array nil")
	}
	//' check capacity sufficient before shift, expand if needed
	addlen := len(ary)
	if addlen < 1 {
		return false
	} //' fail silently, empty insert could be a design choice
	oldlen := len(base.ary)
	newmin := oldlen + addlen
	if cap(base.ary) < newmin {
		base.EnsureCapacity(newmin)
	}
	//' alway need append loop
	for ix := 0; ix < addlen; ix++ {
		base.ary = append(base.ary, 0) //' expand first with dummy zero value
	}
	//' now can do shift
	if pos < oldlen { //' shift following entries before put new one
		copy(base.ary[pos+addlen:], base.ary[pos:])
	}
	//fmt.Println("<diag.ArrayList.InsertAll after shift> base:", base)
	//' and a loop to copy ary contents
	for ix, elem := range ary {
		//fmt.Println("<diag.ArrayList.InsertAll loop> ix:", ix, "; pos +ix:", pos+ix, "; elem:", elem)
		base.ary[pos+ix] = elem
	} //' for
	return true
} //' InsertAll

//' AppendAll replaces simple case of Java's addAll; calls InsertAll with pos beyond end of existing.
//' Also must panic if ary nil, and return false if ary empty
func (base *ArrayList) AppendAll(ary []Obj) bool {
	if ary == nil {
		panic("ArrayList.AppendAll input array nil")
	}
	return base.InsertAll(len(base.ary), ary)
}

//' ToArrayObj replaces non-param Java's toArray, returns array of type []Obj, roughly same as Java
//' version returning array of <Object> but with element type as found
func (base *ArrayList) ToArrayObj() (newary []Obj) {
	newary = make([]Obj, len(base.ary))
	copy(newary, base.ary)
	return
}

/*  ToArrayInt emulates param version of Java's ArrayList.toArray(arrayT) for type int only; returns
    updated input array if large enough, or a new array of same type if base.ary is larger than input 
    array.  User does NOT need to wrap array in struct Omni and extract array as with more general 
    ToArray of this package, and access to nth element is a simple 'elem := result[n]'.
    As a convenience, type-specific versions not needing wrapper are provided for 6 of the 17 basic
    Go types, eg ToArrayInt(ary); the other 11 can easily be added as needed for a specific project,
    or use the wrapper-based ToArray(ary).
    Ignored Java spec feature of nul (Go nil) after last element, Go slice makes that irrelevant.
*/
func (base *ArrayList) ToArrayInt(ary []int) []int {
	alen := len(base.ary)
	//fmt.Println("<diag.ArrayList.ToArrayInt> alen", alen, ", cap(ary)", cap(ary))
	if alen > cap(ary) {
		ary = make([]int, alen)
	}
	ary = ary[:alen] //' reslice in case len < cap or len > alen
	//fmt.Println("<diag.ArrayList.ToArrayInt> before loop, ary", ary)
	for ix, elem := range base.ary {
		ary[ix] = elem.(int)
	}
	//fmt.Println("<diag.ArrayList.ToArrayInt> after loop, ary", ary)
	return ary
}

/*  ToArrayBool emulates param version of Java's ArrayList.toArray(arrayT) for type bool only; see
    description of ToArrayInt for additional info, the same except type
*/
func (base *ArrayList) ToArrayBool(ary []bool) []bool {
	alen := len(base.ary)
	if alen > cap(ary) {
		ary = make([]bool, alen)
	}
	ary = ary[:alen] //' reslice in case len < cap or len > alen
	for ix, elem := range base.ary {
		ary[ix] = elem.(bool)
	}
	return ary
}

/*  ToArrayByte emulates param version of Java's ArrayList.toArray(arrayT) for type byte only; see
    description of ToArrayInt for additional info, the same except type
*/
func (base *ArrayList) ToArrayByte(ary []byte) []byte {
	alen := len(base.ary)
	if alen > cap(ary) {
		ary = make([]byte, alen)
	}
	ary = ary[:alen] //' reslice in case len < cap or len > alen
	for ix, elem := range base.ary {
		ary[ix] = elem.(byte)
	}
	return ary
}

/*  ToArrayFloat64 emulates param version of Java's ArrayList.toArray(arrayT) for type float64 only; see
    description of ToArrayInt for additional info, the same except type
*/
func (base *ArrayList) ToArrayFloat64(ary []float64) []float64 {
	alen := len(base.ary)
	if alen > cap(ary) {
		ary = make([]float64, alen)
	}
	ary = ary[:alen] //' reslice in case len < cap or len > alen
	for ix, elem := range base.ary {
		ary[ix] = elem.(float64)
	}
	return ary
}

/*  ToArrayRune emulates param version of Java's ArrayList.toArray(arrayT) for type rune only; see
    description of ToArrayInt for additional info, the same except type
*/
func (base *ArrayList) ToArrayRune(ary []rune) []rune {
	alen := len(base.ary)
	if alen > cap(ary) {
		ary = make([]rune, alen)
	}
	ary = ary[:alen] //' reslice in case len < cap or len > alen
	for ix, elem := range base.ary {
		ary[ix] = elem.(rune)
	}
	return ary
}

/*  ToArrayString emulates param version of Java's ArrayList.toArray(arrayT) for type string only; see
    description of ToArrayInt for additional info, the same except type
*/
func (base *ArrayList) ToArrayString(ary []string) []string {
	alen := len(base.ary)
	if alen > cap(ary) {
		ary = make([]string, alen)
	}
	ary = ary[:alen] //' reslice in case len < cap or len > alen
	for ix, elem := range base.ary {
		ary[ix] = elem.(string)
	}
	return ary
}

/*  ToArray emulates param version of Java's ArrayList.toArray(arrayT); returns updated input array
    if large enough, or a new array of same type if base.ary is larger than subarray of arrayT.
    User must wrap array in struct Omni and extract array, eg test example
        preAry1 := Omni{}; preAry1.aryInt = make([]int, 22); preAry1 = dmyALi.ToArray(preAry1)
        fmt.Println("<diag.ToArray>", preAry1.aryInt)
    Note that access to nth element would be 'elem := preAry1.aryInt[n]' not 'elem := preAry1[n]'
    As a convenience, type-specific versions not needing wrapper are also provided, eg ToArrayInt(ary)
    ((as of noon 2012.06.30, planned but not in place))
    Ignored Java spec feature of nul (Go nil) after last element, Go slice makes that irrelevant.
*/
func (base *ArrayList) ToArray(arrayT Omni) Omni {
	alen := len(base.ary)
	v := base.ary[0]
	switch v.(type) {

	case bool:
		if alen > cap(arrayT.aryBool) {
			arrayT.aryBool = make([]bool, alen)
		}
		arrayT.aryBool = arrayT.aryBool[:alen] //' reslice in case len < cap or len > alen
		for ix, elem := range base.ary {
			arrayT.aryBool[ix] = elem.(bool)
		}
		return arrayT

	case string:
		if alen > cap(arrayT.aryString) {
			arrayT.aryString = make([]string, alen)
		}
		arrayT.aryString = arrayT.aryString[:alen]
		for ix, elem := range base.ary {
			arrayT.aryString[ix] = elem.(string)
		}
		return arrayT

	case int:
		//fmt.Println("<diag.ArrayList.ToArray.case int> alen", alen, ", cap(arrayT.aryInt)", cap(arrayT.aryInt))
		if alen > cap(arrayT.aryInt) {
			arrayT.aryInt = make([]int, alen)
		}
		arrayT.aryInt = arrayT.aryInt[:alen] //' reslice in case len < cap or len > alen
		//fmt.Println("<diag.ArrayList.ToArray.case int> before loop, arrayT.aryInt", arrayT.aryInt)
		for ix, elem := range base.ary {
			arrayT.aryInt[ix] = elem.(int)
		}
		//fmt.Println("<diag.ArrayList.ToArray.case int> after loop, arrayT.aryInt", arrayT.aryInt)
		return arrayT

	case int8:
		if alen > cap(arrayT.aryInt8) {
			arrayT.aryInt8 = make([]int8, alen)
		}
		arrayT.aryInt8 = arrayT.aryInt8[:alen]
		for ix, elem := range base.ary {
			arrayT.aryInt8[ix] = elem.(int8)
		}
		return arrayT

	case int16:
		if alen > cap(arrayT.aryInt16) {
			arrayT.aryInt16 = make([]int16, alen)
		}
		arrayT.aryInt16 = arrayT.aryInt16[:alen]
		for ix, elem := range base.ary {
			arrayT.aryInt16[ix] = elem.(int16)
		}
		return arrayT

		/*
			case Int32:
				if alen > cap(arrayT.aryInt32) {
					arrayT.aryInt32 = make([]Int32, alen)
				}
				arrayT.aryInt32 = arrayT.aryInt32[:alen]
				for ix, elem := range base.ary {
					arrayT.aryInt32[ix] = elem.(Int32)
				}
				return arrayT
		            //  */

		//' byte and uint8 or rune and int32 not allowed in same switch-case or struct,
		//' one alias of other so would count as redundant
	case rune:
		if alen > cap(arrayT.aryRune) {
			arrayT.aryRune = make([]rune, alen)
		}
		arrayT.aryRune = arrayT.aryRune[:alen]
		for ix, elem := range base.ary {
			arrayT.aryRune[ix] = elem.(rune)
		}
		return arrayT

	case int64:
		if alen > cap(arrayT.aryInt64) {
			arrayT.aryInt64 = make([]int64, alen)
		}
		arrayT.aryInt64 = arrayT.aryInt64[:alen]
		for ix, elem := range base.ary {
			arrayT.aryInt64[ix] = elem.(int64)
		}
		return arrayT

	case uint:
		if alen > cap(arrayT.aryUint) {
			arrayT.aryUint = make([]uint, alen)
		}
		arrayT.aryUint = arrayT.aryUint[:alen]
		for ix, elem := range base.ary {
			arrayT.aryUint[ix] = elem.(uint)
		}
		return arrayT
		/*
		   case uint8:
		       if alen > cap(arrayT.aryUint8) {
		           arrayT.aryUint8 = make([]uint8, alen)
		       }
		       arrayT.aryUint8 = arrayT.aryUint8[:alen]
		       for ix, elem := range base.ary {
		           arrayT.aryUint8[ix] = elem.(uint8)
		       }
		       return arrayT
		*/

		//' byte and uint8 or rune and int32 not allowed in same switch-case or struct,
		//' one alias of other so would count as redundant
	case byte:
		if alen > cap(arrayT.aryByte) {
			arrayT.aryByte = make([]byte, alen)
		}
		arrayT.aryByte = arrayT.aryByte[:alen]
		for ix, elem := range base.ary {
			arrayT.aryByte[ix] = elem.(byte)
		}
		return arrayT

	case uint16:
		if alen > cap(arrayT.aryUint16) {
			arrayT.aryUint16 = make([]uint16, alen)
		}
		arrayT.aryUint16 = arrayT.aryUint16[:alen]
		for ix, elem := range base.ary {
			arrayT.aryUint16[ix] = elem.(uint16)
		}
		return arrayT

	case uint32:
		if alen > cap(arrayT.aryUint32) {
			arrayT.aryUint32 = make([]uint32, alen)
		}
		arrayT.aryUint32 = arrayT.aryUint32[:alen]
		for ix, elem := range base.ary {
			arrayT.aryUint32[ix] = elem.(uint32)
		}
		return arrayT

	case uint64:
		if alen > cap(arrayT.aryUint64) {
			arrayT.aryUint64 = make([]uint64, alen)
		}
		arrayT.aryUint64 = arrayT.aryUint64[:alen]
		for ix, elem := range base.ary {
			arrayT.aryUint64[ix] = elem.(uint64)
		}
		return arrayT

	case uintptr:
		if alen > cap(arrayT.aryUintptr) {
			arrayT.aryUintptr = make([]uintptr, alen)
		}
		arrayT.aryUintptr = arrayT.aryUintptr[:alen]
		for ix, elem := range base.ary {
			arrayT.aryUintptr[ix] = elem.(uintptr)
		}
		return arrayT

	case float32:
		if alen > cap(arrayT.aryFloat32) {
			arrayT.aryFloat32 = make([]float32, alen)
		}
		arrayT.aryFloat32 = arrayT.aryFloat32[:alen]
		for ix, elem := range base.ary {
			arrayT.aryFloat32[ix] = elem.(float32)
		}
		return arrayT

	case float64:
		if alen > cap(arrayT.aryFloat64) {
			arrayT.aryFloat64 = make([]float64, alen)
		}
		arrayT.aryFloat64 = arrayT.aryFloat64[:alen]
		for ix, elem := range base.ary {
			arrayT.aryFloat64[ix] = elem.(float64)
		}
		return arrayT

	case complex64:
		if alen > cap(arrayT.aryComplex64) {
			arrayT.aryComplex64 = make([]complex64, alen)
		}
		arrayT.aryComplex64 = arrayT.aryComplex64[:alen]
		for ix, elem := range base.ary {
			arrayT.aryComplex64[ix] = elem.(complex64)
		}
		return arrayT

	case complex128:
		if alen > cap(arrayT.aryComplex128) {
			arrayT.aryComplex128 = make([]complex128, alen)
		}
		arrayT.aryComplex128 = arrayT.aryComplex128[:alen]
		for ix, elem := range base.ary {
			arrayT.aryComplex128[ix] = elem.(complex128)
		}
		return arrayT

	} //' switch-case

	//' default, same as ToArrayAny
	//ary.aryObj = make([]Obj, alen)
	//copy(arrayT, base.ary)
	if alen > cap(arrayT.aryObj) {
		arrayT.aryObj = make([]Obj, alen)
	}
	arrayT.aryObj = arrayT.aryObj[:alen]
	copy(arrayT.aryObj, base.ary)
	return arrayT
} //' ToArray
//  */

//' findObj is a local utility, not exported, in support of Contains, IndexOf, LastIndexOf & RemoveObj
func (base *ArrayList) findObj(dir int, obj Obj) (pos int) { //' name pos not used, is documentation
	end := len(base.ary) - 1
	if dir > 0 { //' search fwd from 0
		for ix, elem := range base.ary {
			//fmt.Println("<diag.ArrayList.findObj> ix:", ix, "; elem:", elem, "; search obj:", obj)
			//fmt.Printf("<diag.ArrayList.findObj> types, of elem: %T, of search obj: %T\n", obj, elem)
			if elem == obj {
				return ix
			} //' found
		}
		return -1 //' not found
	} //' else search backward from end
	for ix, _ := range base.ary {
		if base.ary[end-ix] == obj {
			return end - ix
		} //' found
	}
	return -1 //' not found
} //' findObj

//' Contains emulates Java's ArrayList.contains; calls local utility method findObj; returns true if
//' item is found, else false
func (base *ArrayList) Contains(obj Obj) bool {
	//fmt.Println("<diag.ArrayList.Contains calls findObj> search obj:", obj, "; base.findObj(1, obj):", base.findObj(1, obj))
	if base.findObj(1, obj) >= 0 {
		return true
	}
	return false
}

//' IndexOf emulates Java's indexOf; returns position where item first found, or -1 if not found
func (base *ArrayList) IndexOf(obj Obj) int {
	return base.findObj(1, obj)
}

//' LastIndexOf emulates Java's lastIndexOf; returns position where item last found, or -1 if not found
func (base *ArrayList) LastIndexOf(obj Obj) int {
	return base.findObj(-1, obj)
}

//' RemoveObj emulates Java's remove with object param, does not return value (calls Remove w pos param);
//' returns false if object not found
func (base *ArrayList) RemoveObj(obj Obj) bool {
	pos := base.findObj(1, obj)
	if pos > 0 {
		base.Remove(pos) //' discards returned value
		return true
	} //' no else case, silent if not found
	return false
}

//' RemoveRange emulates Java's ArrayList.remove with object param, does not return value; silent
//' return if toIndex <= fromIndex; reslices w/o changing capacity
func (base *ArrayList) RemoveRange(fromIndex, toIndex int) {
	if toIndex <= fromIndex {
		return
	} //' silent if no action
	newlen := len(base.ary) + fromIndex - toIndex
	//fmt.Println("<diag.ArrayList.RemoveRange input> fromIndex:", fromIndex, "; toIndex:", toIndex, "; oldlen:", len(base.ary), "; newlen:", newlen)
	copy(base.ary[fromIndex:], base.ary[toIndex:])
	base.ary = base.ary[:newlen] //' capacity not affected
	return
}

/* obsolete
//' func to partially streamline tests of 17 array types
func TestToArrayForType(tmpObjAry []Obj) {
	//' TestToArrayForType 'blew up' due to chg in ToArray; cmt-out for now, fix later
	    tstAL := ArrayListPreset(tmpObjAry)
	    //toAryT := tstAL.ToArray().aryInt8 //' oops, not w/o a case stmt!  --tmp!!
	    //' attempt to 'scope' toAryT, do not expect to wk; didn't
	    //toAryT := tmpObjAry
	    //' so have to duplicate Printf in each case, even though identical; minor irritation
	    //' minor savings: put the fmt into a string (much more readable, and chg fmt in one place)
	    fmtString := "<diag.tstAL.toAryT> input %T; output %T, output element %T; content: %v\n"
	    v := tmpObjAry[0]
	    switch v.(type) {
	    case bool:          toAryT := tstAL.ToArray().aryBool
	        fmt.Printf(fmtString, tmpObjAry[0], toAryT, toAryT[0], toAryT)
	    case string:        toAryT := tstAL.ToArray().aryString
	        fmt.Printf(fmtString, tmpObjAry[0], toAryT, toAryT[0], toAryT)
	    case int:           toAryT := tstAL.ToArray().aryInt
	        fmt.Printf(fmtString, tmpObjAry[0], toAryT, toAryT[0], toAryT)
	    case int8:          toAryT := tstAL.ToArray().aryInt8
	        fmt.Printf(fmtString, tmpObjAry[0], toAryT, toAryT[0], toAryT)
	    case int16:         toAryT := tstAL.ToArray().aryInt16
	        fmt.Printf(fmtString, tmpObjAry[0], toAryT, toAryT[0], toAryT)
	    case int32:         toAryT := tstAL.ToArray().aryInt32
	        fmt.Printf(fmtString, tmpObjAry[0], toAryT, toAryT[0], toAryT)
	    case int64:         toAryT := tstAL.ToArray().aryInt64
	        fmt.Printf(fmtString, tmpObjAry[0], toAryT, toAryT[0], toAryT)
	    case uint:          toAryT := tstAL.ToArray().aryUint
	        fmt.Printf(fmtString, tmpObjAry[0], toAryT, toAryT[0], toAryT)
	    case uint8:         toAryT := tstAL.ToArray().aryUint8
	        fmt.Printf(fmtString, tmpObjAry[0], toAryT, toAryT[0], toAryT)
	    case uint16:        toAryT := tstAL.ToArray().aryUint16
	        fmt.Printf(fmtString, tmpObjAry[0], toAryT, toAryT[0], toAryT)
	    case uint32:        toAryT := tstAL.ToArray().aryUint32
	        fmt.Printf(fmtString, tmpObjAry[0], toAryT, toAryT[0], toAryT)
	    case uint64:        toAryT := tstAL.ToArray().aryUint64
	        fmt.Printf(fmtString, tmpObjAry[0], toAryT, toAryT[0], toAryT)
	    case uintptr:       toAryT := tstAL.ToArray().aryUintptr
	        fmt.Printf(fmtString, tmpObjAry[0], toAryT, toAryT[0], toAryT)
	    case float32:       toAryT := tstAL.ToArray().aryFloat32
	        fmt.Printf(fmtString, tmpObjAry[0], toAryT, toAryT[0], toAryT)
	    case float64:       toAryT := tstAL.ToArray().aryFloat64
	        fmt.Printf(fmtString, tmpObjAry[0], toAryT, toAryT[0], toAryT)
	    case complex64:     toAryT := tstAL.ToArray().aryComplex64
	        fmt.Printf(fmtString, tmpObjAry[0], toAryT, toAryT[0], toAryT)
	    case complex128:    toAryT := tstAL.ToArray().aryComplex128
	        fmt.Printf(fmtString, tmpObjAry[0], toAryT, toAryT[0], toAryT)
	    default:            toAryT := tstAL.ToArray().aryObj
	        fmt.Printf(fmtString, tmpObjAry[0], toAryT, toAryT[0], toAryT)
	    } //' switch-case
}
    //  */

//' testSet is local func to streamline tests of most methods with all 17 array types
func testSet(tstWrap Omni, typ Obj) {
	//' setup
	tmpObj := []Obj{} //' for scope

	switch typ.(type) {

	//' note that 5 cases have extra code to test ToArrayInt etc; would be 6 but bool tested in main

	case bool: //' this Println is 99% for documentation, 1% for idiot-proofing myself
		fmt.Println("Type bool should be tested separately w/o testSet & dblr")
		return

	case int:
		n := len(tstWrap.aryInt)
		tmpObj = make([]Obj, n)
		tstIntAx := make([]int, n) //' expected array
		tstIntAr := make([]int, n) //' result array
		for ix := 0; ix < n; ix++ {
			tmpObj[ix] = tstWrap.aryInt[ix]
			tstIntAx[ix] = tstWrap.aryInt[ix]
		}
		tstAL := ArrayListPreset(tmpObj)
		tstIntAr = tstAL.ToArrayInt(tstIntAr)
		FlagIfIntAryNE("<testSet, ToArrayInt>", tstIntAx, tstIntAr)

		//' byte, float64, rune & string cases also modified; rearranged to follow int instead of slide 23 order
		/*
		tmp: old versions of byte, float64, rune & string; rearranged to follow int instead of slide 23 order

		    case byte:
		        n := len(tstWrap.aryByte)
		        tmpObj = make([]Obj, n)
		        for ix := 0; ix < n; ix++ {
		            tmpObj[ix] = tstWrap.aryByte[ix]
		        }

		    case float64:
		        n := len(tstWrap.aryFloat64)
		        tmpObj = make([]Obj, n)
		        for ix := 0; ix < n; ix++ {
		            tmpObj[ix] = tstWrap.aryFloat64[ix]
		        }

		    case rune:
		        n := len(tstWrap.aryRune)
		        tmpObj = make([]Obj, n)
		        for ix := 0; ix < n; ix++ {
		            tmpObj[ix] = tstWrap.aryRune[ix]
		        }

		    case string:
		        n := len(tstWrap.aryString)
		        tmpObj = make([]Obj, n)
		        for ix := 0; ix < n; ix++ {
		            tmpObj[ix] = tstWrap.aryString[ix]
		        }
		//  */

		//' byte, float64, rune & string cases also modified; rearranged to follow int instead of slide 23 order

		//' byte Byte
	case byte:
		n := len(tstWrap.aryByte)
		tmpObj = make([]Obj, n)
		tstByteAx := make([]byte, n) //' expected array
		tstByteAr := make([]byte, n) //' result array
		for ix := 0; ix < n; ix++ {
			tmpObj[ix] = tstWrap.aryByte[ix]
			tstByteAx[ix] = tstWrap.aryByte[ix]
		}
		tstAL := ArrayListPreset(tmpObj)
		tstByteAr = tstAL.ToArrayByte(tstByteAr)
		FlagIfByteAryNE("<testSet, ToArrayByte>", tstByteAx, tstByteAr)

		//' float64 Float64
	case float64:
		n := len(tstWrap.aryFloat64)
		tmpObj = make([]Obj, n)
		tstFloat64Ax := make([]float64, n) //' expected array
		tstFloat64Ar := make([]float64, n) //' result array
		for ix := 0; ix < n; ix++ {
			tmpObj[ix] = tstWrap.aryFloat64[ix]
			tstFloat64Ax[ix] = tstWrap.aryFloat64[ix]
		}
		tstAL := ArrayListPreset(tmpObj)
		tstFloat64Ar = tstAL.ToArrayFloat64(tstFloat64Ar)
		FlagIfFloat64AryNE("<testSet, ToArrayFloat64>", tstFloat64Ax, tstFloat64Ar)

		//' rune Rune
	case rune:
		n := len(tstWrap.aryRune)
		tmpObj = make([]Obj, n)
		tstRuneAx := make([]rune, n) //' expected array
		tstRuneAr := make([]rune, n) //' result array
		for ix := 0; ix < n; ix++ {
			tmpObj[ix] = tstWrap.aryRune[ix]
			tstRuneAx[ix] = tstWrap.aryRune[ix]
		}
		tstAL := ArrayListPreset(tmpObj)
		tstRuneAr = tstAL.ToArrayRune(tstRuneAr)
		FlagIfRuneAryNE("<testSet, ToArrayRune>", tstRuneAx, tstRuneAr)

		//' string String
	case string:
		n := len(tstWrap.aryString)
		tmpObj = make([]Obj, n)
		tstStringAx := make([]string, n) //' expected array
		tstStringAr := make([]string, n) //' result array
		for ix := 0; ix < n; ix++ {
			tmpObj[ix] = tstWrap.aryString[ix]
			tstStringAx[ix] = tstWrap.aryString[ix]
		}
		tstAL := ArrayListPreset(tmpObj)
		tstStringAr = tstAL.ToArrayString(tstStringAr)
		FlagIfStringAryNE("<testSet, ToArrayString>", tstStringAx, tstStringAr)

	case int8:
		n := len(tstWrap.aryInt8)
		tmpObj = make([]Obj, n)
		for ix := 0; ix < n; ix++ {
			tmpObj[ix] = tstWrap.aryInt8[ix]
		}

	case int16:
		n := len(tstWrap.aryInt16)
		tmpObj = make([]Obj, n)
		for ix := 0; ix < n; ix++ {
			tmpObj[ix] = tstWrap.aryInt16[ix]
		}

		//' byte and uint8 or rune and int32 not allowed in same switch-case or struct,
		//' one alias of other so would count as redundant
	/*
		case Int32:
			n := len(tstWrap.aryInt32)
			tmpObj = make([]Obj, n)
			for ix := 0; ix < n; ix++ {
				tmpObj[ix] = tstWrap.aryInt32[ix]
			}
	*/

	case int64:
		n := len(tstWrap.aryInt64)
		tmpObj = make([]Obj, n)
		for ix := 0; ix < n; ix++ {
			tmpObj[ix] = tstWrap.aryInt64[ix]
		}

	case uint:
		n := len(tstWrap.aryUint)
		tmpObj = make([]Obj, n)
		for ix := 0; ix < n; ix++ {
			tmpObj[ix] = tstWrap.aryUint[ix]
		}

		//' byte and uint8 or rune and int32 not allowed in same switch-case or struct,
		//' one alias of other so would count as redundant
	/*
		case uint8:
			n := len(tstWrap.aryUint8)
			tmpObj = make([]Obj, n)
			for ix := 0; ix < n; ix++ {
				tmpObj[ix] = tstWrap.aryUint8[ix]
			}
	*/

	case uint16:
		n := len(tstWrap.aryUint16)
		tmpObj = make([]Obj, n)
		for ix := 0; ix < n; ix++ {
			tmpObj[ix] = tstWrap.aryUint16[ix]
		}

	case uint32:
		n := len(tstWrap.aryUint32)
		tmpObj = make([]Obj, n)
		for ix := 0; ix < n; ix++ {
			tmpObj[ix] = tstWrap.aryUint32[ix]
		}

	case uint64:
		n := len(tstWrap.aryUint64)
		tmpObj = make([]Obj, n)
		for ix := 0; ix < n; ix++ {
			tmpObj[ix] = tstWrap.aryUint64[ix]
		}

	case uintptr:
		n := len(tstWrap.aryUintptr)
		tmpObj = make([]Obj, n)
		for ix := 0; ix < n; ix++ {
			tmpObj[ix] = tstWrap.aryUintptr[ix]
		}

	case float32:
		n := len(tstWrap.aryFloat32)
		tmpObj = make([]Obj, n)
		for ix := 0; ix < n; ix++ {
			tmpObj[ix] = tstWrap.aryFloat32[ix]
		}

	case complex64:
		n := len(tstWrap.aryComplex64)
		tmpObj = make([]Obj, n)
		for ix := 0; ix < n; ix++ {
			tmpObj[ix] = tstWrap.aryComplex64[ix]
		}

	case complex128:
		n := len(tstWrap.aryComplex128)
		tmpObj = make([]Obj, n)
		for ix := 0; ix < n; ix++ {
			tmpObj[ix] = tstWrap.aryComplex128[ix]
		}

	default:
		fmt.Printf("testSet cannot handle type %T\n", typ)
	} //' switch-case

	//fmt.Println("Running testSet with type", typ.(type))
	fmt.Printf("\nRunning testSet with type %T\n", typ)

	tstAL := ArrayListPreset(tmpObj)
	tstObj := tstAL.ToArrayObj()
	//invObj := invertSense(tstObj) //' separate sw-case, but most cases can be collapsed into one case (I hope)
	invObj := dblr(tstObj) //' separate sw-case, but most cases can be collapsed into one case (I hope)
	//' BTW, type bool should be tested separately w/o testSet & dblr
	tstRef := tstAL.Copy() //' for occasionally restoring to original preset w/o creating a new ArrayList
	//' end of setup

	//' actual tests
	tgt1 := tstAL.Get(3)
	//fmt.Println("<testSet, Get & Contains> expect true:", tstAL.Contains(tgt1))
	FlagIfValNE("<testSet, Get & Contains t>", true, tstAL.Contains(tgt1))
	tgt2 := invObj
	//fmt.Println("<testSet, Get & Contains> expect false:", tstAL.Contains(tgt2))
	FlagIfValNE("<testSet, Get & Contains f>", false, tstAL.Contains(tgt2))
	//fmt.Println("<testSet, IndexOf> tgt:", tgt1, ", expect 3:", tstAL.IndexOf(tgt1))
	FlagIfValNE("<testSet, IndexOf "+fmt.Sprint(tgt1)+">", 3, tstAL.IndexOf(tgt1))
	//fmt.Println("<testSet, IndexOf> tgt:", tgt2, ", expect -1:", tstAL.IndexOf(tgt2))
	FlagIfValNE("<testSet, IndexOf "+fmt.Sprint(tgt2)+">", -1, tstAL.IndexOf(tgt2))
	//fmt.Println("<testSet, LastIndexOf> tgt:", tgt1, ", expect 8:", tstAL.LastIndexOf(tgt1))
	FlagIfValNE("<testSet, LastIndexOf "+fmt.Sprint(tgt1)+">", 8, tstAL.LastIndexOf(tgt1))
	//fmt.Println("<testSet, LastIndexOf> tgt:", tgt2, ", expect -1:", tstAL.LastIndexOf(tgt2))
	FlagIfValNE("<testSet, LastIndexOf "+fmt.Sprint(tgt2)+">", -1, tstAL.LastIndexOf(tgt2))

	//fmt.Println("<testSet, initial sequence> ", tstObj)
	v := tstAL.Remove(2)
	//' assemble expected from subsets of array tstObj; copy not reslice
	n := len(tstObj)
	x := make([]Obj, n-1)
	copy(x[:2], tstObj[:2])
	copy(x[2:], tstObj[3:])
	//fmt.Println("<testSet, tstAL.Remove(2)> value removed:", v, ", result: ", tstAL.ary)
	FlagIfObjAryNE("<testSet, tstAL.Remove(2) "+fmt.Sprint(v)+">", x, tstAL.ary)

	//' next 2 ck the bool (found or not), *and* the 'result' (same for both)
	b := tstAL.RemoveObj(tgt1)
	x = make([]Obj, n-2)
	copy(x[:2], tstObj[:2])
	copy(x[2:], tstObj[4:])
	//fmt.Println("<testSet, RemoveObj> tgt:", tgt1, ", true:", b, ", result:", tstAL.ary)
	FlagIfValNE("<testSet, RemoveObj "+fmt.Sprint(tgt1)+"; found?>", true, b)
	FlagIfObjAryNE("<testSet, RemoveObj "+fmt.Sprint(tgt1)+"; result>", x, tstAL.ary)

	b = tstAL.RemoveObj(tgt2)
	//fmt.Println("<testSet, RemoveObj> tgt:", tgt2, ", false:", b, ", result ", tstAL.ary)
	FlagIfValNE("<testSet, RemoveObj "+fmt.Sprint(tgt2)+"; found?>", false, b)
	FlagIfObjAryNE("<testSet, RemoveObj "+fmt.Sprint(tgt2)+"; result>", x, tstAL.ary)

	x = make([]Obj, n-4)
	copy(x[:2], tstObj[:2])
	copy(x[2:], tstObj[6:])
	tstAL.RemoveRange(2, 4)
	//fmt.Println("<testSet, RemoveRange(2,4), tgt 5,1> result", tstAL.ary)
	FlagIfObjAryNE("<testSet, RemoveRange(2,4), tgt 5,1>", x, tstAL.ary)

	tstAL = tstRef.Copy()
	//fmt.Println("<testSet, restored sequence> ", tstAL.ary)

	v = tstAL.Get(2)
	tstAL.Append(v)
	x = make([]Obj, n+1)
	copy(x[:n], tstObj)
	copy(x[n:], tstObj[2:3])
	//fmt.Println("<testSet, tstAL.Append(v)> value appended:", v, ", result: ", tstAL.ary)
	//FlagIfValNE("<testSet, tstAL.Append("+fmt.Sprint(v)+"> value appended", 3, v) //' type mismatch
	FlagIfValNE("<testSet, tstAL.Append("+fmt.Sprint(v)+"> value appended", tstObj[2], v)
	FlagIfObjAryNE("<testSet, RemoveRange(2,4), tgt 5,1>", x, tstAL.ary)

	x = make([]Obj, n+2)
	copy(x[:4], tstObj[:4])
	copy(x[4:5], tstObj[2:3])
	copy(x[5:n+1], tstObj[4:])
	copy(x[n+1:], tstObj[2:3])
	tstAL.Insert(4, v)
	//fmt.Println("<testSet, tstAL.Insert(4,v)> value inserted:", v, ", result: ", tstAL.ary)
	FlagIfObjAryNE("<testSet, tstAL.Insert(4,"+fmt.Sprint(v)+">", x, tstAL.ary)

	xx := make([]Obj, n+5)
	copy(xx[:n+2], x)
	copy(xx[n+2:], tstObj[:3])
	shortAry := tstObj[:3] //' reslice is safe here
	tstAL.AppendAll(shortAry)
	//fmt.Println("<testSet, tstAL.AppendAll(shortAry)> values appended:", shortAry, " result:", tstAL.ary)
	//fmt.Println("<testSet debug, tstAL.AppendAll(shortAry)> values appended:", shortAry, ", xx:", xx, ", actual:", tstAL.ary)
	FlagIfObjAryNE("<testSet, tstAL.AppendAll("+fmt.Sprint(shortAry)+">", xx, tstAL.ary)
	//' test of ToArray is indirect, implicit in earlier code

	tstAL.Set(6, v)
	//fmt.Println("<testSet, tstAL.tstAL.Set(6, v)> value set:", v, ", result:", tstAL.ary)
	FlagIfValNE("<testSet, tstAL.tstAL.Set(6, "+fmt.Sprint(v)+")> ", tstObj[2], tstAL.Get(6))
	/*
		//  */

	return
}

//' dblr is a local convenience utility called by testSet; adds 100 to most types (negate
//' won't work for unsigned), flips bool, does rot-13 on string; complains if unknown type
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
	//case Int32:
	//    v = 2 * v.(Int32)
	//' byte and uint8 or rune and int32 not allowed in same switch-case or struct,
	//' one alias of other so would count as redundant
	case rune:
		v = 2 * v.(rune)
	case int64:
		v = 2 * v.(int64)
	case uint:
		v = 2 * v.(uint)
	//case uint8:
	//    v = 2 * v.(uint8)
	//' byte and uint8 or rune and int32 not allowed in same switch-case or struct,
	//' one alias of other so would count as redundant
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
}

//' Copy replaces but does not emulate Java ArrayList.clone, returns a new ArrayList with elements copied from existing
//' ArrayList; unlike Java ArrayList.clone, changes to values in the copy will not change values in the original
//' (for shallow copy in Go, simply use ':=' or '=', so there's no need for a method Clone to do that)
func (base *ArrayList) Copy() (AL *ArrayList) {
	//AL = &ArrayList{base.ary}
	//' need to test that changes to copy do not change the original -- oops, bad
	alen := len(base.ary)
	tmp := make([]Obj, alen, alen) //' create with size and capacity both alen
	for ix, elem := range base.ary {
		tmp[ix] = elem
	}
	AL = &ArrayList{tmp}
	return
} //' Copy

//' strToBoolAsObjAry is a local convenience util to convert a string of t's & f's to true's & false's
//' in a []Obj with elements of type bool
func strToBoolAsObjAry(str string) []Obj {
	ba := []byte(str)
	bo := make([]Obj, len(ba))
	for ix, elem := range ba {
		bo[ix] = (elem == 't')
	}
	return bo
}

//' FlagIfValNE takes place of assertEquals(msg, xpctd, actl)
func FlagIfValNE(msg string, expected, actual interface{}) {
	if expected != actual {
		fmt.Printf("%s expected [%v], found [%v]\n", msg, expected, actual)
	}
}

//' FlagIfBoolAryNE takes place of assertArrayEquals(msg, xpctd, actl) for bool arrays
func FlagIfBoolAryNE(msg string, expected, actual []bool) {
    FlagIfValNE(msg+"Array length mismatch, checking to shorter anyway", len(expected), len(actual))
    for i := 0; i < MinInt(len(actual), len(expected)); i++ {
        FlagIfValNE(msg+" Array mismatch at index "+fmt.Sprint(i)+":", expected[i], actual[i])
    }
}

//' FlagIfByteAryNE takes place of assertArrayEquals(msg, xpctd, actl) for byte arrays
func FlagIfByteAryNE(msg string, expected, actual []byte) {
    FlagIfValNE(msg+"Array length mismatch, checking to shorter anyway", len(expected), len(actual))
    for i := 0; i < MinInt(len(actual), len(expected)); i++ {
        FlagIfValNE(msg+" Array mismatch at index "+fmt.Sprint(i)+":", expected[i], actual[i])
    }
}

//' FlagIfRuneAryNE takes place of assertArrayEquals(msg, xpctd, actl) for rune arrays
func FlagIfRuneAryNE(msg string, expected, actual []rune) {
	FlagIfValNE(msg+"Array length mismatch, checking to shorter anyway", len(expected), len(actual))
	for i := 0; i < MinInt(len(actual), len(expected)); i++ {
		FlagIfValNE(msg+" Array mismatch at index "+fmt.Sprint(i)+":", expected[i], actual[i])
	}
}

//' FlagIfFloat64AryNE takes place of assertArrayEquals(msg, ary1, ary2) for float64 arrays
func FlagIfFloat64AryNE(msg string, expected, actual []float64) {
	FlagIfValNE(msg+"Array length mismatch, checking to shorter anyway", len(expected), len(actual))
	for i := 0; i < MinInt(len(actual), len(expected)); i++ {
		FlagIfValNE(msg+" Array mismatch at index "+fmt.Sprint(i)+":", expected[i], actual[i])
	}
}

//' FlagIfIntAryNE takes place of assertArrayEquals(msg, ary1, ary2) for int arrays
func FlagIfIntAryNE(msg string, expected, actual []int) {
	FlagIfValNE(msg+"Array length mismatch, checking to shorter anyway", len(expected), len(actual))
	for i := 0; i < MinInt(len(actual), len(expected)); i++ {
		FlagIfValNE(msg+" Array mismatch at index "+fmt.Sprint(i)+":", expected[i], actual[i])
	}
}

//' FlagIfObjAryNE takes place of assertArrayEquals(msg, ary1, ary2) for int arrays
func FlagIfObjAryNE(msg string, expected, actual []Obj) {
	FlagIfValNE(msg+"Array length mismatch, checking to shorter anyway", len(expected), len(actual))
	for i := 0; i < MinInt(len(actual), len(expected)); i++ {
		FlagIfValNE(msg+" Array mismatch at index "+fmt.Sprint(i)+":", expected[i], actual[i])
	}
}

//' FlagIfStringAryNE takes place of assertArrayEquals(msg, ary1, ary2) for string arrays
func FlagIfStringAryNE(msg string, expected, actual []string) {
	FlagIfValNE(msg+"Array length mismatch, checking to shorter anyway", len(expected), len(actual))
	for i := 0; i < MinInt(len(actual), len(expected)); i++ {
		FlagIfValNE(msg+" Array mismatch at index "+fmt.Sprint(i)+":", expected[i], actual[i])
	}
}

//' func MinInt is trivial, just to avoid multiple casts required by use of math.Min on pair of int
//  func MinInt(a, b int) c int {  c := int(math.Min(float64(a), float64(b))) }
func MinInt(a, b int) (c int) {
	c = a
	if c > b {
		c = b
	}
	return
}

/*
//  */
