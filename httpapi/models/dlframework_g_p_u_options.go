// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DlframeworkGPUOptions dlframework g p u options
// swagger:model dlframeworkGPUOptions
type DlframeworkGPUOptions struct {

	// The type of GPU allocation strategy to use.
	//
	// Allowed values:
	// "": The empty string (default) uses a system-chosen default
	//     which may change over time.
	//
	// "BFC": A "Best-fit with coalescing" algorithm, simplified from a
	//        version of dlmalloc.
	AllocatorType string `json:"allocator_type,omitempty"`

	// Force all tensors to be gpu_compatible. On a GPU-enabled TensorFlow,
	// enabling this option forces all CPU tensors to be allocated with Cuda
	// pinned memory. Normally, TensorFlow will infer which tensors should be
	// allocated as the pinned memory. But in case where the inference is
	// incomplete, this option can significantly speed up the cross-device memory
	// copy performance as long as it fits the memory.
	// Note that this option is not something that should be
	// enabled by default for unknown or very large models, since all Cuda pinned
	// memory is unpageable, having too much pinned memory might negatively impact
	// the overall host system performance.
	ForceGpuCompatible bool `json:"force_gpu_compatible,omitempty"`

	// A value between 0 and 1 that indicates what fraction of the
	// available GPU memory to pre-allocate for each process.  1 means
	// to pre-allocate all of the GPU memory, 0.5 means the process
	// allocates ~50% of the available GPU memory.
	PerProcessGpuMemoryFraction float64 `json:"per_process_gpu_memory_fraction,omitempty"`

	// A comma-separated list of GPU ids that determines the 'visible'
	// to 'virtual' mapping of GPU devices.  For example, if TensorFlow
	// can see 8 GPU devices in the process, and one wanted to map
	// visible GPU devices 5 and 3 as "/device:GPU:0", and "/device:GPU:1", then
	// one would specify this field as "5,3".  This field is similar in spirit to
	// the CUDA_VISIBLE_DEVICES environment variable, except it applies to the
	// visible GPU devices in the process.
	//
	// NOTE: The GPU driver provides the process with the visible GPUs
	// in an order which is not guaranteed to have any correlation to
	// the *physical* GPU id in the machine.  This field is used for
	// remapping "visible" to "virtual", which means this operates only
	// after the process starts.  Users are required to use vendor
	// specific mechanisms (e.g., CUDA_VISIBLE_DEVICES) to control the
	// physical to visible device mapping prior to invoking TensorFlow.
	VisibleDeviceList string `json:"visible_device_list,omitempty"`
}

// Validate validates this dlframework g p u options
func (m *DlframeworkGPUOptions) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DlframeworkGPUOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DlframeworkGPUOptions) UnmarshalBinary(b []byte) error {
	var res DlframeworkGPUOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
