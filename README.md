# OpenCL bindings for Go

Documentation at <http://godoc.org/github.com/jgillich/go-opencl/cl>.

See the [test](cl/cl_test.go) for usage examples.

By default, the OpenCL 1.2 API is exported. To get OpenCL 1.0, set the build tag `cl10`.

## Forked Additions
- Compatible with held ptrs to OpenCL pointers and contexts including `clKernel`,
`clQueue`, `clContext`, `clProgram`, `clDevice`
- Adds in the OpenCL GL Memory bindings for binding to allocated OpenGL buffers
and sharing memory contexts with OpenGL

## Usage

OpenCL bindings and wrapper functions map nearly one to one with the OpenCL core
library calls in version `1.2`. If your machine or OpenCl is not compatible with `1.2`
please use the build tag `cl10`. This forked library mostly tests holding OpenCL state
outside of scopes where the pointers are initially returned.

In order to retain your OpenCL states into higher scopes you must return each wrapper
object individually. The file `cl_pointer_test.go` creates its own wrapper functions
for creating the underlying objects in the test file and serves as an example wrapper
API for the `go-opencl` module.

```
init, _ := NewInit(data)
mContext := &Context{}
mQueue := &CommandQueue{}
mProgram := &Program{}
mKernel := &Kernel{}
mDevice := &Device{}
mDevices := make([]*Device, 10)
mInput := &MemObject{}
mOutput := &MemObject{}
mContext, mDevice, mDevices = init.CreateContext()
init.context = mContext
init.device = mDevice
init.devices = mDevices
mQueue = init.CreateQueue()
init.queue = mQueue
mProgram = init.CreateProgram()
init.program = mProgram
mKernel = init.CreateKernel()
init.kernel = mKernel
mInput, mOutput = init.CreateBuffers()
init.input = mInput
init.output = mOutput
init.Run()
```

Note that in order to reduce dependencies the OpenGL binding calls are not tested in
this library. For MacOSX however the `clCreateFromGLTexture2D` and `clCreateFromGLTexture3D`
calls have been deprecated.
