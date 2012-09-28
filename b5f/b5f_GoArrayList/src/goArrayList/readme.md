"ArrayList implemented in Go"

### GoArrayList is a Go language substitute for the Java class ArrayList, with very nearly all features. ###

No connection with Oracle nor with Google is claimed or implied; no one but Phil Stephens is 
responsible for any defects (unless they volunteer to help fix those defects).  Also, use at own risk.

Copyleft, use and modify as you see fit; giving credit optional but appreciated; eg 
    Project GoArrayList started by, and major initial Go coding done by, Phil Stephens.


### Contact info: ###
Please use PhilipRStephens@gmail.com for questions about this package, suggestions for improving it,
suggestions for other packages to consider coding, offers of employment or invitations to volunteer,
etc.  (Please see my github Jobs Profile if relevant).


### Motivation: ###
In converting a medium sized practice project from Java to Go, I found that neither array slice nor 
doubly linked list was a good fit for most of how I had used ArrayList in Java; after implementing the
small subset that I immediately needed, decided that implementing the whole feature set would be a good 
sharable work, as well as cause to learn a few more things about Go.  I have indeed learned from this; 
hope it also is useful to someone.

Ironically, I later abandoned using it in the project for which I developed it, but still a good 
learning experience, and subsets of it were directly adaptable for the int slices I did use.


### Known Exceptions to Completeness: ###
* Omitted field modCount as I could not find Java usage examples.  Can add it if I get requests 
  (please also give some sample Java that uses it).
* Replaced Java ArrayList.clone (some say to avoid using in Java; does essentially what '=' or ':='
  would do in Go: a new slice pointing to the same underlying array) with .Copy, returns a new 
  ArrayList with elements copied (changes to elements of the copy will NOT change elements of the 
  original).
* Did not find a way to have a single method return arrays with conveniently selected types, but with
  input and output arrays of interface Obj (effectively like a Java array of Objects), the elements can 
  be in any basic type such as int, and these elements (but not the array unless copy elements to a new 
  array) can be used directly as their own type, eg int or string.


### Also Note: ###
* Chose to implement the aliases Byte and Rune rather than their equivalent basic types uint8 and int32;
  had to choose one or the other for use in type-based switch-case, so went with what I as a user would 
  prefer in most projects.  Did some testing to confirm that they really are interchangable: yes.
* Any method of Java ArrayList that has 2 or more signatures with the same name has of course been split
  into 2 or more methods with different names here, as multiple signatures are not allowed in Go, and I
  did not see a way to use ellipses or interface to achieve the same effect, other than using Omni and Obj 
  as 2 kinds of wrapper of course (suggetions welcome if I missed a trick).

This release also includes a separate file main.go with unit tests; created my own funcs for this as I 
did not manage to get gotest working w/i Eclipse at this time.  (Nor godoc, which I have crudely simulated, 
further below, via manual edit).  The test is called 'main' in file 'main.go' because goclipse requires 
that.  Tested only in Windows XP via Eclipse with goclipse and only on one machine so far, and definitely 
not exhaustively complete tests.  Testing volunteers welcome, which is why I include main.go in this release.
Meanwhile I have a project to use this in, which may or may not reveal problems the unit tests missed.


### Addendum 2012.07.26: ###
A few days ago while doing web-search for another project, stumbled upon page "SliceTricks  How to do 
vector-esque things with slices" (<https://code.google.com/p/go-wiki/wiki/SliceTricks>, which in turn 
credits its source as <http://www.reddit.com/r/golang/comments/eh3gh/slice_support_and_containervector>).  

Today I took a break from the other project to see if I could use these 'tricks' in GoArrayList, and 
identified 7 funcs & methods to try them on.  All 7 worked, in several cases eliminating a loop for 
reduced lines of code and possibly faster execution (I haven't bothered to time it, as is not currently 
an issue).  The most dramatic lines-reduction is in InsertAll, from 27 to 8 lines total, and the changed 
portion replaces 18 lines with 1 line.  In most cases the changed portion replaces 3 lines with 1 line, 
except Insert replaces 6 lines with 1 line.  Not much chagrin that I didn't take advantage of slices this 
way in the first place, more delighted that I happened to find this set of hints.  Old code included as
comments for comparison, will probably delete in a future rev.  (Also fixed a deficiency in testSet, it 
now tests InsertAll).


### Addendum 2012.09.28: ###
Just updating the readme to attempt to use Markdown syntax <http://daringfireball.net/projects/markdown/>
rather than plain ascii text, with essentially no other changes except this Addendum and a small 2nd
paragraph added to the Motivation section.


Project Directory Structure
---------------------------

Here's a project directory structure (example is rev e) that works for me in Eclipse:

    b5e_GoArrayList/
        bin/
        pkg/
        src/
            goArrayList/
                doc.go
                GoArrayList.go
            main.go

...this was the result of some trial and error to get 'import  . "goArrayList" ' to work in main.go, so I 
know there are several other configurations that do NOT work, but there may be other legal ways to do it.


Some usage examples:
--------------------

### ArrayList of array: ###
    tmpAo = make([]Obj, 5)
    for ix, _ := range tmpAo {
        tmpAi = make([]int, 3)
        for ixx, _ := range tmpAi { tmpAi[ixx] = 3*ix + ixx }
        tmpAo[ix] = tmpAi
    }
    tstAL = ArrayListPreset(tmpAo)
    rtnAo := []Obj{};    rtnAo = tstAL.ToArray(rtnAo)
    fmt.Println("tstAL: ", tstAL, "\\n rtnAo: ", rtnAo)
    // tstAL:  &{[[0 1 2] [3 4 5] [6 7 8] [9 10 11] [12 13 14]]} 
    // rtnAo:  [[0 1 2] [3 4 5] [6 7 8] [9 10 11] [12 13 14]]

### ArrayList of struct: ###
    type City struct {
        Name    string
        Pop     int
        Lon, Lat    float64 //' already converted to miles
    }
    CaOrWa := []City{
        {"Longview", 0, 46.15, 122.95},
        {"Portland", 0, 45.54, 122.66},
        {"Cape Blanco State Park", 0, 42.84, 124.57},
        {"Oceanside", 0, 33.23, 117.31},
    }
    tmpAo = make([]Obj, 4)
    for ix, _ := range tmpAo {
        tmpAo[ix] = CaOrWa[ix]
    }
    tstAL = ArrayListPreset(tmpAo)
    rtnAo := []Obj{};    rtnAo = tstAL.ToArray(rtnAo)
    fmt.Println("tstAL: ", tstAL, ", rtnAo: ", rtnAo)
    //tstAL:  &{[{Longview 0 46.15 122.95} {Portland 0 45.54 122.66}
    // {Cape Blanco State Park 0 42.84 124.57} {Oceanside 0 33.23 117.31}]} 
    // rtnAo:  [{Longview 0 46.15 122.95} {Portland 0 45.54 122.66}
    // {Cape Blanco State Park 0 42.84 124.57} {Oceanside 0 33.23 117.31}]

### autodetect types in mixed ArrayList: ###
    tmpAi := []int{1, 2, 3, 4, 5};    tmpAo := []Obj{0, 0, 0, 0, 0}
    for ix, elem := range tmpAi { tmpAo[ix] = elem }
    tstAL = ArrayListPreset(tmpAo);    tmpAo = tstAL.ToArray(tmpAo)
    tstAL.Set(2, "three");    tstAL.Set(4, 5.0)
    for _, elem := range tmpAo {
        switch elem.(type) {
        case int: fmt.Print(elem, " is int.  ")
        case string: fmt.Print(elem, " is string.  ")
        default: fmt.Print(elem, " is neither int nor string.  ")
        }
    }
    fmt.Println()
    // 1 is int.  2 is int.  three is string.  4 is int.  5 is neither int nor string.  


The first line of each method, func or type, plus the preceding comment
-----------------------------------------------------------------------
(As a crude approximation of godoc, but generated manually)

    // GoArrayList project GoArrayList.go; see comments in doc.go
    package goArrayList
    
    //Obj interface matches any type; named to reflect similarity to Java 'object' in ArrayList etc.
    type Obj interface{}
    
    // ArrayList struct is intended as a fairly complete Go language substitute for the Java feature ArrayList.
    // Java documents describe ArrayList as: "Resizable-array implementation of the List interface".  ArrayList 
    // struct in package GoArrayList contains an array of type Obj, can contain any type of elements (this is 
    // comparable to Java spec, which states that ArrayList can only contain objects, not primatives); for 
    // conversion to an array of Obj (elements will usually be of some basic Go type, but can also be Obj, 
    // array slice, or struct), see ToArray method.
    // Note: nil is allowed in an element, much as null is allowed in Java ArrayList.
    // Skipped: field modCount not implemented; could be added if there is interest.
    type ArrayList struct {
    
    // ArrayListPreset func creates and returns a new ArrayList and preloads it with contents of preset, 
    // preserving their type (per entry, not per array/slice; see ToArray method for how to get an array of
    // Obj back out of an ArrayList).  Comparable to Java constructor "ArrayList(Collection<? extends E> c)".
    func ArrayListPreset(preset []Obj) (AL *ArrayList) { // func, not method, so no 'base'
    
    // ArrayListNew func creates and returns a new empty ArrayList with specified capacity.  Emulates Java
    // constructor "ArrayList(int initialCapacity)"; to emulate Java constructor "ArrayList()", call
    // ArrayListNew with param 10.
    func ArrayListNew(acap int) (AL *ArrayList) {
    
    // Clear method emulates Java ArrayList.clear; does a re-slice to change size to 0, leaves capacity 
    // unchanged.  Does not clear values, but this is usually harmless; if want to actually destroy old 
    // values (eg for security), use New instead (for security, overkill is reasonable).
    func (base *ArrayList) Clear() {
    
    // Append method emulates the append version of Java ArrayList.add; uses Go built-in 'append', which
    // handles capacity increase if needed; always returns 'true', since that is the behaviour in Java.
    func (base *ArrayList) Append(val Obj) bool {
    
    // Insert method emulates the insert version of Java ArrayList.add; uses Go built-in 'append', which
    // handles capacity increase if needed; inserts before entry already in position pos.
    func (base *ArrayList) Insert(pos int, val Obj) {
    
    // Get method emulates Java ArrayList.get, returns element at position pos (return signature is interface 
    // Obj, but can use directly as underlying type, this has been tested).
    func (base *ArrayList) Get(pos int) (val Obj) {
    
    // Set method emulates Java ArrayList.set, overwrites element at position pos (underlying type determined 
    // by param val; consistency not enforced); returns value removed, can ignore if not needed.
    func (base *ArrayList) Set(pos int, val Obj) Obj {
    
    // Remove method emulates indexed version of Java ArrayList.remove; returns value removed, can ignore if 
    // not needed.
    func (base *ArrayList) Remove(pos int) Obj {
    
    // EnsureCapacity method emulates Java ArrayList.ensureCapacity; if needed creates new array, copies 
    // existing elements to new array, makes new array the base.Ary, then reslices to old length.
    func (base *ArrayList) EnsureCapacity(minCapacity int) {
    
    // IsEmpty method emulates Java ArrayList.isEmpty, returns true if this list contains no elements.
    func (base *ArrayList) IsEmpty() bool {
    
    // Size method emulates Java ArrayList.size, returns size of slice.  Also implemented Cap to return slice 
    // capacity and SizeCap to return both at once; not in Java, but makes sense for Go.
    func (base *ArrayList) Size() int {
    
    // Cap method returns slice capacity; not in Java, but makes sense for Go.
    func (base *ArrayList) Cap() int {
    
    // SizeCap method returns both slice size and slice capacity; not in Java, but makes sense for Go.
    func (base *ArrayList) SizeCap() (int, int) { // convenience method
    
    // TrimToSize method emulates Java ArrayList.trimToSize, trims slice cap to slice len; actually saves old
    // values to a tmp array, makes a new one as base.Ary, and copies old values to new -- suggest to use 
    // sparingly, as TrimToSize copies array twice and will take time on a really large ArrayList, but has 
    // little value on a small one.
    func (base *ArrayList) TrimToSize() {
    
    // InsertAll method replaces offset version of Java's addAll; called by AppendAll so must handle offset to 
    // just beyond end.  Also, must panic if input ary nil, and return false if input ary empty.
    func (base *ArrayList) InsertAll(pos int, ary []Obj) bool {
    
    // AppendAll method replaces 1-param version of Java's addAll; calls InsertAll with pos just beyond end of 
    // existing ArrayList.  Also, must panic if input ary nil, and return false if input ary empty.
    func (base *ArrayList) AppendAll(ary []Obj) bool {
    
    // ToArray method emulates param version of Java's ArrayList.ToArray(arrayT) for interface Obj (both input 
    // and output are []Obj, but elements will usually be of some basic Go type); returns updated input array 
    // if large enough, or a new array of same type if base.Ary is larger than input array, with all elements 
    // copied (in order found) from base.Ary.  Access to nth element is a simple 'elem := result[n]', and each 
    // element can be used directly as its underlying type.  Type consistency of elements not enforced, but 
    // see doc.go for type-safe suggestions.
    // Note: ignored Java spec feature of nul (Go nil) after last element, as Go slice makes that unreachable.
    func (base *ArrayList) ToArray(ary []Obj) []Obj {
    
    // ToArrayNew method emulates non-param version of Java's ArrayList.ToArray(arrayT) for interface Obj; 
    // creates and returns a new array of interface Obj with all elements copied (in order found) from 
    // base.Ary.  Access to nth element is a simple 'elem := result[n]', and each element can be used directly 
    // as its underlying type.  Type consistency not enforced, but see doc.go for type-safe suggestions.
    // Note: ignored Java spec feature of nul (Go nil) after last element, as Go slice makes that unreachable.
    func (base *ArrayList) ToArrayNew() []Obj {
    
    // The findObj method is a local utility, not exported, in support of Contains, IndexOf, LastIndexOf & 
    // RemoveObj.  A refactoring of common code, direction allows it to search forward or backward.
    func (base *ArrayList) findObj(direction int, obj Obj) (pos int) { // name pos not used, is documentation
    
    // Contains method emulates Java's ArrayList.contains; calls local utility method findObj; returns true if
    // item is found, else false.
    func (base *ArrayList) Contains(obj Obj) bool {
    
    // IndexOf method emulates Java's ArrayList.indexOf; returns position where item first found, or -1 if 
    // not found.
    func (base *ArrayList) IndexOf(obj Obj) int {
    
    // LastIndexOf method emulates Java's ArrayList.lastIndexOf; returns position where item last found, or 
    // -1 if not found.
    func (base *ArrayList) LastIndexOf(obj Obj) int {
    
    // RemoveObj method emulates Java's ArrayList.remove with object param, does not return value (calls 
    // Remove with position param); returns false if object not found.
    func (base *ArrayList) RemoveObj(obj Obj) bool {
    
    // RemoveRange method emulates Java's ArrayList.remove with 2 int params, does not return value; silent
    // return if toIndex <= fromIndex; shifts elements and reslices w/o changing capacity.
    func (base *ArrayList) RemoveRange(fromIndex, toIndex int) {
    
    // Copy method replaces but does *not* emulate Java ArrayList.clone (which some advise not to use in Java), 
    // returns a new ArrayList with elements copied from existing ArrayList; unlike Java ArrayList.clone, 
    // changes to values in the copy will not change values in the original (for 'shallow copy' in Go, simply 
    // use ':=' or '=', so there's no need for a method Clone to do that).
    func (base *ArrayList) Copy() (AL *ArrayList) {
    
