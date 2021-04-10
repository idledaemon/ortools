/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 4.0.2
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: ./ortools/sat/go/sat.i

package sat

/*
#define intgo swig_intgo
typedef void *swig_voidp;

#include <stdint.h>


typedef long long intgo;
typedef unsigned long long uintgo;



typedef struct { char *p; intgo n; } _gostring_;
typedef struct { void* array; intgo len; intgo cap; } _goslice_;


typedef long long swig_type_1;
typedef long long swig_type_2;
typedef long long swig_type_3;
typedef long long swig_type_4;
typedef long long swig_type_5;
typedef long long swig_type_6;
typedef _goslice_ swig_type_7;
typedef _goslice_ swig_type_8;
typedef _goslice_ swig_type_9;
typedef _goslice_ swig_type_10;
typedef _goslice_ swig_type_11;
typedef _goslice_ swig_type_12;
typedef _goslice_ swig_type_13;
typedef _goslice_ swig_type_14;
typedef _goslice_ swig_type_15;
typedef _gostring_ swig_type_16;
typedef _goslice_ swig_type_17;
typedef _gostring_ swig_type_18;
typedef _goslice_ swig_type_19;
typedef _gostring_ swig_type_20;
typedef _goslice_ swig_type_21;
typedef _goslice_ swig_type_22;
extern void _wrap_Swig_free_sat_58fca0c20b2ea3ed(uintptr_t arg1);
extern uintptr_t _wrap_Swig_malloc_sat_58fca0c20b2ea3ed(swig_intgo arg1);
extern uintptr_t _wrap__swig_NewDirectorSolutionCallbackSolutionCallback_sat_58fca0c20b2ea3ed(int);
extern void _wrap_DeleteDirectorSolutionCallback_sat_58fca0c20b2ea3ed(uintptr_t arg1);
extern void _wrap_delete_SolutionCallback_sat_58fca0c20b2ea3ed(uintptr_t arg1);
extern void _wrap_SolutionCallback_onSolutionCallback_sat_58fca0c20b2ea3ed(uintptr_t arg1);
extern swig_type_1 _wrap_SolutionCallback_numBooleans_sat_58fca0c20b2ea3ed(uintptr_t arg1);
extern swig_type_2 _wrap_SolutionCallback_numBranches_sat_58fca0c20b2ea3ed(uintptr_t arg1);
extern swig_type_3 _wrap_SolutionCallback_numConflicts_sat_58fca0c20b2ea3ed(uintptr_t arg1);
extern swig_type_4 _wrap_SolutionCallback_numBinaryPropagations_sat_58fca0c20b2ea3ed(uintptr_t arg1);
extern swig_type_5 _wrap_SolutionCallback_numIntegerPropagations_sat_58fca0c20b2ea3ed(uintptr_t arg1);
extern double _wrap_SolutionCallback_wallTime_sat_58fca0c20b2ea3ed(uintptr_t arg1);
extern double _wrap_SolutionCallback_userTime_sat_58fca0c20b2ea3ed(uintptr_t arg1);
extern double _wrap_SolutionCallback_objectiveValue_sat_58fca0c20b2ea3ed(uintptr_t arg1);
extern double _wrap_SolutionCallback_bestObjectiveBound_sat_58fca0c20b2ea3ed(uintptr_t arg1);
extern swig_type_6 _wrap_SolutionCallback_solutionIntegerValue_sat_58fca0c20b2ea3ed(uintptr_t arg1, swig_intgo arg2);
extern _Bool _wrap_SolutionCallback_solutionBooleanValue_sat_58fca0c20b2ea3ed(uintptr_t arg1, swig_intgo arg2);
extern void _wrap_SolutionCallback_stopSearch_sat_58fca0c20b2ea3ed(uintptr_t arg1);
extern swig_type_7 _wrap_SolutionCallback_Response_sat_58fca0c20b2ea3ed(uintptr_t arg1);
extern swig_type_8 _wrap_SatHelper_Solve_sat_58fca0c20b2ea3ed(swig_type_9 arg1);
extern swig_type_10 _wrap_SatHelper_SolveWithParameters_sat_58fca0c20b2ea3ed(swig_type_11 arg1, swig_type_12 arg2);
extern swig_type_13 _wrap_SatHelper_SolveWithParametersAndSolutionCallback_sat_58fca0c20b2ea3ed(swig_type_14 arg1, swig_type_15 arg2, uintptr_t arg3);
extern swig_type_16 _wrap_SatHelper_ModelStats_sat_58fca0c20b2ea3ed(swig_type_17 arg1);
extern swig_type_18 _wrap_SatHelper_SolverResponseStats_sat_58fca0c20b2ea3ed(swig_type_19 arg1);
extern swig_type_20 _wrap_SatHelper_ValidateModel_sat_58fca0c20b2ea3ed(swig_type_21 arg1);
extern uintptr_t _wrap_SatHelper_VariableDomain_sat_58fca0c20b2ea3ed(swig_type_22 arg1);
extern uintptr_t _wrap_new_SatHelper_sat_58fca0c20b2ea3ed(void);
extern void _wrap_delete_SatHelper_sat_58fca0c20b2ea3ed(uintptr_t arg1);
#undef intgo
*/
import "C"

import "unsafe"
import _ "runtime/cgo"
import "sync"
import "github.com/golang/protobuf/proto"


type _ unsafe.Pointer



var Swig_escape_always_false bool
var Swig_escape_val interface{}


type _swig_fnptr *byte
type _swig_memberptr *byte


type _ sync.Mutex


type swig_gostring struct { p uintptr; n int }
func swigCopyString(s string) string {
  p := *(*swig_gostring)(unsafe.Pointer(&s))
  r := string((*[0x7fffffff]byte)(unsafe.Pointer(p.p))[:p.n])
  Swig_free(p.p)
  return r
}

func Swig_free(arg1 uintptr) {
	_swig_i_0 := arg1
	C._wrap_Swig_free_sat_58fca0c20b2ea3ed(C.uintptr_t(_swig_i_0))
}

func Swig_malloc(arg1 int) (_swig_ret uintptr) {
	var swig_r uintptr
	_swig_i_0 := arg1
	swig_r = (uintptr)(C._wrap_Swig_malloc_sat_58fca0c20b2ea3ed(C.swig_intgo(_swig_i_0)))
	return swig_r
}

type _swig_DirectorSolutionCallback struct {
	SwigcptrSolutionCallback
	v interface{}
}

func (p *_swig_DirectorSolutionCallback) Swigcptr() uintptr {
	return p.SwigcptrSolutionCallback.Swigcptr()
}

func (p *_swig_DirectorSolutionCallback) SwigIsSolutionCallback() {
}

func (p *_swig_DirectorSolutionCallback) DirectorInterface() interface{} {
	return p.v
}

func NewDirectorSolutionCallback(v interface{}) SolutionCallback {
	p := &_swig_DirectorSolutionCallback{0, v}
	p.SwigcptrSolutionCallback = SwigcptrSolutionCallback(C._wrap__swig_NewDirectorSolutionCallbackSolutionCallback_sat_58fca0c20b2ea3ed(C.int(swigDirectorAdd(p))))
	return p
}

func DeleteDirectorSolutionCallback(arg1 SolutionCallback) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_DeleteDirectorSolutionCallback_sat_58fca0c20b2ea3ed(C.uintptr_t(_swig_i_0))
}

//export Swiggo_DeleteDirector_SolutionCallback_sat_58fca0c20b2ea3ed
func Swiggo_DeleteDirector_SolutionCallback_sat_58fca0c20b2ea3ed(c int) {
	swigDirectorLookup(c).(*_swig_DirectorSolutionCallback).SwigcptrSolutionCallback = 0
	swigDirectorDelete(c)
}

type _swig_DirectorInterfaceSolutionCallbackOnSolutionCallback interface {
	OnSolutionCallback()
}

func (swig_p *_swig_DirectorSolutionCallback) OnSolutionCallback() {
	if swig_g, swig_ok := swig_p.v.(_swig_DirectorInterfaceSolutionCallbackOnSolutionCallback); swig_ok {
		swig_g.OnSolutionCallback()
		return
	}
	panic("call to pure virtual method")
}

//export Swig_DirectorSolutionCallback_callback_onSolutionCallback_sat_58fca0c20b2ea3ed
func Swig_DirectorSolutionCallback_callback_onSolutionCallback_sat_58fca0c20b2ea3ed(swig_c int) {
	swig_p := swigDirectorLookup(swig_c).(*_swig_DirectorSolutionCallback)
	swig_p.OnSolutionCallback()
}

type SwigcptrSolutionCallback uintptr

func (p SwigcptrSolutionCallback) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrSolutionCallback) SwigIsSolutionCallback() {
}

func (p SwigcptrSolutionCallback) DirectorInterface() interface{} {
	return nil
}

func DeleteSolutionCallback(arg1 SolutionCallback) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_SolutionCallback_sat_58fca0c20b2ea3ed(C.uintptr_t(_swig_i_0))
}

func (arg1 SwigcptrSolutionCallback) OnSolutionCallback() {
	_swig_i_0 := arg1
	C._wrap_SolutionCallback_onSolutionCallback_sat_58fca0c20b2ea3ed(C.uintptr_t(_swig_i_0))
}

func (arg1 SwigcptrSolutionCallback) NumBooleans() (_swig_ret int64) {
	var swig_r int64
	_swig_i_0 := arg1
	swig_r = (int64)(C._wrap_SolutionCallback_numBooleans_sat_58fca0c20b2ea3ed(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrSolutionCallback) NumBranches() (_swig_ret int64) {
	var swig_r int64
	_swig_i_0 := arg1
	swig_r = (int64)(C._wrap_SolutionCallback_numBranches_sat_58fca0c20b2ea3ed(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrSolutionCallback) NumConflicts() (_swig_ret int64) {
	var swig_r int64
	_swig_i_0 := arg1
	swig_r = (int64)(C._wrap_SolutionCallback_numConflicts_sat_58fca0c20b2ea3ed(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrSolutionCallback) NumBinaryPropagations() (_swig_ret int64) {
	var swig_r int64
	_swig_i_0 := arg1
	swig_r = (int64)(C._wrap_SolutionCallback_numBinaryPropagations_sat_58fca0c20b2ea3ed(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrSolutionCallback) NumIntegerPropagations() (_swig_ret int64) {
	var swig_r int64
	_swig_i_0 := arg1
	swig_r = (int64)(C._wrap_SolutionCallback_numIntegerPropagations_sat_58fca0c20b2ea3ed(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrSolutionCallback) WallTime() (_swig_ret float64) {
	var swig_r float64
	_swig_i_0 := arg1
	swig_r = (float64)(C._wrap_SolutionCallback_wallTime_sat_58fca0c20b2ea3ed(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrSolutionCallback) UserTime() (_swig_ret float64) {
	var swig_r float64
	_swig_i_0 := arg1
	swig_r = (float64)(C._wrap_SolutionCallback_userTime_sat_58fca0c20b2ea3ed(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrSolutionCallback) ObjectiveValue() (_swig_ret float64) {
	var swig_r float64
	_swig_i_0 := arg1
	swig_r = (float64)(C._wrap_SolutionCallback_objectiveValue_sat_58fca0c20b2ea3ed(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrSolutionCallback) BestObjectiveBound() (_swig_ret float64) {
	var swig_r float64
	_swig_i_0 := arg1
	swig_r = (float64)(C._wrap_SolutionCallback_bestObjectiveBound_sat_58fca0c20b2ea3ed(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrSolutionCallback) SolutionIntegerValue(arg2 int) (_swig_ret int64) {
	var swig_r int64
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (int64)(C._wrap_SolutionCallback_solutionIntegerValue_sat_58fca0c20b2ea3ed(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1)))
	return swig_r
}

func (arg1 SwigcptrSolutionCallback) SolutionBooleanValue(arg2 int) (_swig_ret bool) {
	var swig_r bool
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (bool)(C._wrap_SolutionCallback_solutionBooleanValue_sat_58fca0c20b2ea3ed(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1)))
	return swig_r
}

func (arg1 SwigcptrSolutionCallback) StopSearch() {
	_swig_i_0 := arg1
	C._wrap_SolutionCallback_stopSearch_sat_58fca0c20b2ea3ed(C.uintptr_t(_swig_i_0))
}

func (arg1 SwigcptrSolutionCallback) Response() (_swig_ret CpSolverResponse) {
	var swig_r []uint8
	_swig_i_0 := arg1
	swig_r_p := C._wrap_SolutionCallback_Response_sat_58fca0c20b2ea3ed(C.uintptr_t(_swig_i_0))
	swig_r = *(*[]uint8)(unsafe.Pointer(&swig_r_p))
	var swig_r_1 CpSolverResponse
{

  var pb CpSolverResponse
  goString := swig_r

  if len(goString) == 0 {
    panic("trouble with reading")
  } else {
    m := make([]byte, len(goString))
    copy(m, goString)

    err := proto.Unmarshal(m, &pb)

    if err != nil {
      panic(err)
    }

    swig_r_1 = pb
  }

  return pb
}
	return swig_r_1
}

type SolutionCallback interface {
	Swigcptr() uintptr
	SwigIsSolutionCallback()
	DirectorInterface() interface{}
	OnSolutionCallback()
	NumBooleans() (_swig_ret int64)
	NumBranches() (_swig_ret int64)
	NumConflicts() (_swig_ret int64)
	NumBinaryPropagations() (_swig_ret int64)
	NumIntegerPropagations() (_swig_ret int64)
	WallTime() (_swig_ret float64)
	UserTime() (_swig_ret float64)
	ObjectiveValue() (_swig_ret float64)
	BestObjectiveBound() (_swig_ret float64)
	SolutionIntegerValue(arg2 int) (_swig_ret int64)
	SolutionBooleanValue(arg2 int) (_swig_ret bool)
	StopSearch()
	Response() (_swig_ret CpSolverResponse)
}

type SwigcptrSatHelper uintptr

func (p SwigcptrSatHelper) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrSatHelper) SwigIsSatHelper() {
}

func SatHelperSolve(arg1 CpModelProto) (_swig_ret CpSolverResponse) {
	var swig_r []uint8
	var _swig_i_0 []byte
{
  // hello go
  bytes, err := proto.Marshal(&arg1)
  if err != nil {
    panic(err)
  }
  _swig_i_0 = bytes
}
	swig_r_p := C._wrap_SatHelper_Solve_sat_58fca0c20b2ea3ed(*(*C.swig_type_9)(unsafe.Pointer(&_swig_i_0)))
	swig_r = *(*[]uint8)(unsafe.Pointer(&swig_r_p))
	if Swig_escape_always_false {
		Swig_escape_val = _swig_i_0
	}
	var swig_r_1 CpSolverResponse
{

  var pb CpSolverResponse
  goString := swig_r

  if len(goString) == 0 {
    panic("trouble with reading")
  } else {
    m := make([]byte, len(goString))
    copy(m, goString)

    err := proto.Unmarshal(m, &pb)

    if err != nil {
      panic(err)
    }

    swig_r_1 = pb
  }

  return pb
}
	return swig_r_1
}

func SatHelperSolveWithParameters(arg1 CpModelProto, arg2 SatParameters) (_swig_ret CpSolverResponse) {
	var swig_r []uint8
	var _swig_i_0 []byte
{
  // hello go
  bytes, err := proto.Marshal(&arg1)
  if err != nil {
    panic(err)
  }
  _swig_i_0 = bytes
}
	var _swig_i_1 []byte
{
  // hello go
  bytes, err := proto.Marshal(&arg2)
  if err != nil {
    panic(err)
  }
  _swig_i_1 = bytes
}
	swig_r_p := C._wrap_SatHelper_SolveWithParameters_sat_58fca0c20b2ea3ed(*(*C.swig_type_11)(unsafe.Pointer(&_swig_i_0)), *(*C.swig_type_12)(unsafe.Pointer(&_swig_i_1)))
	swig_r = *(*[]uint8)(unsafe.Pointer(&swig_r_p))
	if Swig_escape_always_false {
		Swig_escape_val = _swig_i_0
	}
	if Swig_escape_always_false {
		Swig_escape_val = _swig_i_1
	}
	var swig_r_1 CpSolverResponse
{

  var pb CpSolverResponse
  goString := swig_r

  if len(goString) == 0 {
    panic("trouble with reading")
  } else {
    m := make([]byte, len(goString))
    copy(m, goString)

    err := proto.Unmarshal(m, &pb)

    if err != nil {
      panic(err)
    }

    swig_r_1 = pb
  }

  return pb
}
	return swig_r_1
}

func SatHelperSolveWithParametersAndSolutionCallback(arg1 CpModelProto, arg2 SatParameters, arg3 SolutionCallback) (_swig_ret CpSolverResponse) {
	var swig_r []uint8
	var _swig_i_0 []byte
{
  // hello go
  bytes, err := proto.Marshal(&arg1)
  if err != nil {
    panic(err)
  }
  _swig_i_0 = bytes
}
	var _swig_i_1 []byte
{
  // hello go
  bytes, err := proto.Marshal(&arg2)
  if err != nil {
    panic(err)
  }
  _swig_i_1 = bytes
}
	_swig_i_2 := arg3.Swigcptr()
	swig_r_p := C._wrap_SatHelper_SolveWithParametersAndSolutionCallback_sat_58fca0c20b2ea3ed(*(*C.swig_type_14)(unsafe.Pointer(&_swig_i_0)), *(*C.swig_type_15)(unsafe.Pointer(&_swig_i_1)), C.uintptr_t(_swig_i_2))
	swig_r = *(*[]uint8)(unsafe.Pointer(&swig_r_p))
	if Swig_escape_always_false {
		Swig_escape_val = _swig_i_0
	}
	if Swig_escape_always_false {
		Swig_escape_val = _swig_i_1
	}
	var swig_r_1 CpSolverResponse
{

  var pb CpSolverResponse
  goString := swig_r

  if len(goString) == 0 {
    panic("trouble with reading")
  } else {
    m := make([]byte, len(goString))
    copy(m, goString)

    err := proto.Unmarshal(m, &pb)

    if err != nil {
      panic(err)
    }

    swig_r_1 = pb
  }

  return pb
}
	return swig_r_1
}

func SatHelperModelStats(arg1 CpModelProto) (_swig_ret string) {
	var swig_r string
	var _swig_i_0 []byte
{
  // hello go
  bytes, err := proto.Marshal(&arg1)
  if err != nil {
    panic(err)
  }
  _swig_i_0 = bytes
}
	swig_r_p := C._wrap_SatHelper_ModelStats_sat_58fca0c20b2ea3ed(*(*C.swig_type_17)(unsafe.Pointer(&_swig_i_0)))
	swig_r = *(*string)(unsafe.Pointer(&swig_r_p))
	if Swig_escape_always_false {
		Swig_escape_val = _swig_i_0
	}
	var swig_r_1 string
 swig_r_1 = swigCopyString(swig_r) 
	return swig_r_1
}

func SatHelperSolverResponseStats(arg1 CpSolverResponse) (_swig_ret string) {
	var swig_r string
	var _swig_i_0 []byte
{
  // hello go
  bytes, err := proto.Marshal(&arg1)
  if err != nil {
    panic(err)
  }
  _swig_i_0 = bytes
}
	swig_r_p := C._wrap_SatHelper_SolverResponseStats_sat_58fca0c20b2ea3ed(*(*C.swig_type_19)(unsafe.Pointer(&_swig_i_0)))
	swig_r = *(*string)(unsafe.Pointer(&swig_r_p))
	if Swig_escape_always_false {
		Swig_escape_val = _swig_i_0
	}
	var swig_r_1 string
 swig_r_1 = swigCopyString(swig_r) 
	return swig_r_1
}

func SatHelperValidateModel(arg1 CpModelProto) (_swig_ret string) {
	var swig_r string
	var _swig_i_0 []byte
{
  // hello go
  bytes, err := proto.Marshal(&arg1)
  if err != nil {
    panic(err)
  }
  _swig_i_0 = bytes
}
	swig_r_p := C._wrap_SatHelper_ValidateModel_sat_58fca0c20b2ea3ed(*(*C.swig_type_21)(unsafe.Pointer(&_swig_i_0)))
	swig_r = *(*string)(unsafe.Pointer(&swig_r_p))
	if Swig_escape_always_false {
		Swig_escape_val = _swig_i_0
	}
	var swig_r_1 string
 swig_r_1 = swigCopyString(swig_r) 
	return swig_r_1
}

func SatHelperVariableDomain(arg1 IntegerVariableProto) (_swig_ret Operations_research_Domain) {
	var swig_r Operations_research_Domain
	var _swig_i_0 []byte
{
  // hello go
  bytes, err := proto.Marshal(&arg1)
  if err != nil {
    panic(err)
  }
  _swig_i_0 = bytes
}
	swig_r = (Operations_research_Domain)(SwigcptrOperations_research_Domain(C._wrap_SatHelper_VariableDomain_sat_58fca0c20b2ea3ed(*(*C.swig_type_22)(unsafe.Pointer(&_swig_i_0)))))
	if Swig_escape_always_false {
		Swig_escape_val = _swig_i_0
	}
	return swig_r
}

func NewSatHelper() (_swig_ret SatHelper) {
	var swig_r SatHelper
	swig_r = (SatHelper)(SwigcptrSatHelper(C._wrap_new_SatHelper_sat_58fca0c20b2ea3ed()))
	return swig_r
}

func DeleteSatHelper(arg1 SatHelper) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_SatHelper_sat_58fca0c20b2ea3ed(C.uintptr_t(_swig_i_0))
}

type SatHelper interface {
	Swigcptr() uintptr
	SwigIsSatHelper()
}


type SwigcptrOperations_research_Domain uintptr
type Operations_research_Domain interface {
	Swigcptr() uintptr;
}
func (p SwigcptrOperations_research_Domain) Swigcptr() uintptr {
	return uintptr(p)
}



var swigDirectorTrack struct {
	sync.Mutex
	m map[int]interface{}
	c int
}

func swigDirectorAdd(v interface{}) int {
	swigDirectorTrack.Lock()
	defer swigDirectorTrack.Unlock()
	if swigDirectorTrack.m == nil {
		swigDirectorTrack.m = make(map[int]interface{})
	}
	swigDirectorTrack.c++
	ret := swigDirectorTrack.c
	swigDirectorTrack.m[ret] = v
	return ret
}

func swigDirectorLookup(c int) interface{} {
	swigDirectorTrack.Lock()
	defer swigDirectorTrack.Unlock()
	ret := swigDirectorTrack.m[c]
	if ret == nil {
		panic("C++ director pointer not found (possible	use-after-free)")
	}
	return ret
}

func swigDirectorDelete(c int) {
	swigDirectorTrack.Lock()
	defer swigDirectorTrack.Unlock()
	if swigDirectorTrack.m[c] == nil {
		if c > swigDirectorTrack.c {
			panic("C++ director pointer invalid (possible memory corruption")
		} else {
			panic("C++ director pointer not found (possible use-after-free)")
		}
	}
	delete(swigDirectorTrack.m, c)
}


