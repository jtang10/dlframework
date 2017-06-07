// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Wed, 07 Jun 2017 15:59:19 BST.
// By https://git.io/c-for-go. DO NOT EDIT.

package gomxnet

/*
#cgo LDFLAGS: -lmxnet
#include "c_api.h"
#include "c_predict_api.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import "unsafe"

// NDArrayHandle type as declared in mxnet/c_api.h:49
type NDArrayHandle unsafe.Pointer

// MXGenericCallback type as declared in mxnet/c_api.h:108
type MXGenericCallback func() int32

// CustomOpFBFunc type as declared in mxnet/c_api.h:133
type CustomOpFBFunc func(Arg0 int32, Arg1 []unsafe.Pointer, Arg2 []int32, Arg3 []int32, Arg4 int32, Arg5 unsafe.Pointer) int32

// CustomOpDelFunc type as declared in mxnet/c_api.h:136
type CustomOpDelFunc func(Arg0 unsafe.Pointer) int32

// CustomOpListFunc type as declared in mxnet/c_api.h:137
type CustomOpListFunc func(Arg0 [][][]byte, Arg1 unsafe.Pointer) int32

// CustomOpInferShapeFunc type as declared in mxnet/c_api.h:138
type CustomOpInferShapeFunc func(Arg0 int32, Arg1 []int32, Arg2 [][]uint32, Arg3 unsafe.Pointer) int32

// CustomOpInferTypeFunc type as declared in mxnet/c_api.h:140
type CustomOpInferTypeFunc func(Arg0 int32, Arg1 []int32, Arg2 unsafe.Pointer) int32

// CustomOpBwdDepFunc type as declared in mxnet/c_api.h:141
type CustomOpBwdDepFunc func(Arg0 []int32, Arg1 []int32, Arg2 []int32, Arg3 []int32, Arg4 [][]int32, Arg5 unsafe.Pointer) int32

// CustomOpCreateFunc type as declared in mxnet/c_api.h:144
type CustomOpCreateFunc func(Arg0 string, Arg1 int32, Arg2 [][]uint32, Arg3 []int32, Arg4 []int32, Arg5 []MXCallbackList, Arg6 unsafe.Pointer) int32

// CustomOpPropCreator type as declared in mxnet/c_api.h:148
type CustomOpPropCreator func(Arg0 string, Arg1 int32, Arg2 []string, Arg3 []string, Arg4 []MXCallbackList) int32

// MXKVStoreUpdater type as declared in mxnet/c_api.h:1370
type MXKVStoreUpdater func(Key int32, Recv NDArrayHandle, Local NDArrayHandle, Handle unsafe.Pointer)

// MXKVStoreServerController type as declared in mxnet/c_api.h:1465
type MXKVStoreServerController func(Head int32, Body string, ControllerHandle unsafe.Pointer)

// NDListHandle type as declared in mxnet/c_predict_api.h:28
type NDListHandle unsafe.Pointer

// MXCallbackList as declared in mxnet/c_api.h:110
type MXCallbackList struct {
	NumCallbacks  int32
	Callbacks     **func() int32
	Contexts      []unsafe.Pointer
	ref92f982b    *C.struct_MXCallbackList
	allocs92f982b interface{}
}

// NDArrayOpInfo as declared in mxnet/c_api.h:91
type NDArrayOpInfo struct {
	Forward                    *func(Arg0 int32, Arg1 []unsafe.Pointer, Arg2 []int32, Arg3 unsafe.Pointer) int32
	Backward                   *func(Arg0 int32, Arg1 []unsafe.Pointer, Arg2 []int32, Arg3 unsafe.Pointer) int32
	InferShape                 *func(Arg0 int32, Arg1 []int32, Arg2 [][]uint32, Arg3 unsafe.Pointer) int32
	ListOutputs                *func(Arg0 [][][]byte, Arg1 unsafe.Pointer) int32
	ListArguments              *func(Arg0 [][][]byte, Arg1 unsafe.Pointer) int32
	DeclareBackwardDependency  *func(Arg0 []int32, Arg1 []int32, Arg2 []int32, Arg3 []int32, Arg4 [][]int32, Arg5 unsafe.Pointer) int32
	PForward                   unsafe.Pointer
	PBackward                  unsafe.Pointer
	PInferShape                unsafe.Pointer
	PListOutputs               unsafe.Pointer
	PListArguments             unsafe.Pointer
	PDeclareBackwardDependency unsafe.Pointer
	ref94c0476e                *C.struct_NDArrayOpInfo
	allocs94c0476e             interface{}
}

// NativeOpInfo as declared in mxnet/c_api.h:77
type NativeOpInfo struct {
	Forward        *func(Arg0 int32, Arg1 [][]float32, Arg2 []int32, Arg3 [][]uint32, Arg4 []int32, Arg5 unsafe.Pointer)
	Backward       *func(Arg0 int32, Arg1 [][]float32, Arg2 []int32, Arg3 [][]uint32, Arg4 []int32, Arg5 unsafe.Pointer)
	InferShape     *func(Arg0 int32, Arg1 []int32, Arg2 [][]uint32, Arg3 unsafe.Pointer)
	ListOutputs    *func(Arg0 [][][]byte, Arg1 unsafe.Pointer)
	ListArguments  *func(Arg0 [][][]byte, Arg1 unsafe.Pointer)
	PForward       unsafe.Pointer
	PBackward      unsafe.Pointer
	PInferShape    unsafe.Pointer
	PListOutputs   unsafe.Pointer
	PListArguments unsafe.Pointer
	ref7ce11869    *C.struct_NativeOpInfo
	allocs7ce11869 interface{}
}