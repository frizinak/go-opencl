package cl

import (
	"fmt"
	"math/rand"
	"testing"
)

var kSource = `
__kernel void square(
   __global float* input,
   __global float* output,
   const unsigned int count)
{
   int i = get_global_id(0);
   if(i < count)
       output[i] = input[i] * input[i];
}
`

type Init struct {
	context *Context
	devices []*Device
	device  *Device
	kernel  *Kernel
	queue   *CommandQueue
	program *Program
	data    []float32
	input   *MemObject
	output  *MemObject
}

func TestInit(t *testing.T) {

	var data [1024]float32
	for i := 0; i < len(data); i++ {
		data[i] = rand.Float32()
	}

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

}

func NewInit(data [1024]float32) (Init, error) {
	init := Init{}
	init.data = data[:]

	return init, nil
}

func (p Init) CreateQueue() *CommandQueue {
	queue, err := p.context.CreateCommandQueue(p.device, 0)
	if err != nil {
		fmt.Errorf("CreateCommandQueue failed: %+v\n", err)
	}
	p.queue = queue
	return p.queue
}

func (p Init) CreateContext() (*Context, *Device, []*Device) {

	platforms, err := GetPlatforms()
	if err != nil {
		fmt.Errorf("Failed to get platforms: %+v\n", err)
	}
	for i, p := range platforms {
		fmt.Printf("Platform %d:\n", i)
		fmt.Printf("  Name: %s\n", p.Name())
		fmt.Printf("  Vendor: %s\n", p.Vendor())
		fmt.Printf("  Profile: %s\n", p.Profile())
		fmt.Printf("  Version: %s\n", p.Version())
		fmt.Printf("  Extensions: %s\n", p.Extensions())
	}
	platform := platforms[0]

	devices, err := platform.GetDevices(DeviceTypeAll)
	if err != nil {
		fmt.Errorf("Failed to get devices: %+v\n", err)
	}
	if len(devices) == 0 {
		fmt.Errorf("GetDevices return init ed no devices")
	}

	p.devices = devices

	deviceIndex := -1
	for i, d := range devices {
		if deviceIndex < 0 && d.Type() == DeviceTypeGPU {
			deviceIndex = i
		}
		fmt.Printf("Device %d (%s): %s\n", i, d.Type(), d.Name())
		fmt.Printf("  Address Bits: %d\n", d.AddressBits())
		fmt.Printf("  Available: %+v\n", d.Available())
		// fmt.Printf("  Built-In Kernels: %s\n", d.BuiltInKernels())
		fmt.Printf("  Compiler Available: %+v\n", d.CompilerAvailable())
		fmt.Printf("  Double FP Config: %s\n", d.DoubleFPConfig())
		fmt.Printf("  Driver Version: %s\n", d.DriverVersion())
		fmt.Printf("  Error Correction Supported: %+v\n", d.ErrorCorrectionSupport())
		fmt.Printf("  Execution Capabilities: %s\n", d.ExecutionCapabilities())
		fmt.Printf("  Extensions: %s\n", d.Extensions())
		fmt.Printf("  Global Memory Cache Type: %s\n", d.GlobalMemCacheType())
		fmt.Printf("  Global Memory Cacheline Size: %d KB\n", d.GlobalMemCachelineSize()/1024)
		fmt.Printf("  Global Memory Size: %d MB\n", d.GlobalMemSize()/(1024*1024))
		fmt.Printf("  Half FP Config: %s\n", d.HalfFPConfig())
		fmt.Printf("  Host Unified Memory: %+v\n", d.HostUnifiedMemory())
		fmt.Printf("  Image Support: %+v\n", d.ImageSupport())
		fmt.Printf("  Image2D Max Dimensions: %d x %d\n", d.Image2DMaxWidth(), d.Image2DMaxHeight())
		fmt.Printf("  Image3D Max Dimenionns: %d x %d x %d\n", d.Image3DMaxWidth(), d.Image3DMaxHeight(), d.Image3DMaxDepth())
		// fmt.Printf("  Image Max Buffer Size: %d\n", d.ImageMaxBufferSize())
		// fmt.Printf("  Image Max Array Size: %d\n", d.ImageMaxArraySize())
		// fmt.Printf("  Linker Available: %+v\n", d.LinkerAvailable())
		fmt.Printf("  Little Endian: %+v\n", d.EndianLittle())
		fmt.Printf("  Local Mem Size Size: %d KB\n", d.LocalMemSize()/1024)
		fmt.Printf("  Local Mem Type: %s\n", d.LocalMemType())
		fmt.Printf("  Max Clock Frequency: %d\n", d.MaxClockFrequency())
		fmt.Printf("  Max Compute Units: %d\n", d.MaxComputeUnits())
		fmt.Printf("  Max Constant Args: %d\n", d.MaxConstantArgs())
		fmt.Printf("  Max Constant Buffer Size: %d KB\n", d.MaxConstantBufferSize()/1024)
		fmt.Printf("  Max Mem Alloc Size: %d KB\n", d.MaxMemAllocSize()/1024)
		fmt.Printf("  Max Parameter Size: %d\n", d.MaxParameterSize())
		fmt.Printf("  Max Read-Image Args: %d\n", d.MaxReadImageArgs())
		fmt.Printf("  Max Samplers: %d\n", d.MaxSamplers())
		fmt.Printf("  Max Work Group Size: %d\n", d.MaxWorkGroupSize())
		fmt.Printf("  Max Work Item Dimensions: %d\n", d.MaxWorkItemDimensions())
		fmt.Printf("  Max Work Item Sizes: %d\n", d.MaxWorkItemSizes())
		fmt.Printf("  Max Write-Image Args: %d\n", d.MaxWriteImageArgs())
		fmt.Printf("  Memory Base Address Alignment: %d\n", d.MemBaseAddrAlign())
		fmt.Printf("  Native Vector Width Char: %d\n", d.NativeVectorWidthChar())
		fmt.Printf("  Native Vector Width Short: %d\n", d.NativeVectorWidthShort())
		fmt.Printf("  Native Vector Width Int: %d\n", d.NativeVectorWidthInt())
		fmt.Printf("  Native Vector Width Long: %d\n", d.NativeVectorWidthLong())
		fmt.Printf("  Native Vector Width Float: %d\n", d.NativeVectorWidthFloat())
		fmt.Printf("  Native Vector Width Double: %d\n", d.NativeVectorWidthDouble())
		fmt.Printf("  Native Vector Width Half: %d\n", d.NativeVectorWidthHalf())
		fmt.Printf("  OpenCL C Version: %s\n", d.OpenCLCVersion())
		// fmt.Printf("  Parent Device: %+v\n", d.ParentDevice())
		fmt.Printf("  Profile: %s\n", d.Profile())
		fmt.Printf("  Profiling Timer Resolution: %d\n", d.ProfilingTimerResolution())
		fmt.Printf("  Vendor: %s\n", d.Vendor())
		fmt.Printf("  Version: %s\n", d.Version())
	}
	if deviceIndex < 0 {
		deviceIndex = 0
	}
	device := devices[deviceIndex]
	fmt.Printf("Using device %d\n", deviceIndex)
	p.device = device
	context, err := CreateContext([]*Device{device})
	if err != nil {
		fmt.Errorf("CreateContext failed: %+v\n", err)
	}
	p.context = context

	return p.context, p.device, p.devices
}

func (p Init) CreateProgram() *Program {
	program, err := p.context.CreateProgramWithSource([]string{kSource})
	if err != nil {
		fmt.Errorf("CreateProgramWithSource failed: %+v\n", err)
	}

	p.program = program

	if err := p.program.BuildProgram(nil, ""); err != nil {
		fmt.Errorf("BuildProgram failed: %+v\n", err)
	}
	p.program = program
	return p.program
}

func (p Init) CreateKernel() *Kernel {
	kernel, err := p.program.CreateKernel("square")
	if err != nil {
		fmt.Errorf("CreateKernel failed: %+v\n", err)
	}
	p.kernel = kernel
	return p.kernel
}
func (p Init) CreateBuffers() (*MemObject, *MemObject) {
	for i := 0; i < 3; i++ {
		name, err := p.kernel.ArgName(i)
		if err == ErrUnsupported {
			break
		} else if err != nil {
			fmt.Errorf("GetKernelArgInfo for name failed: %+v\n", err)
			break
		} else {
			fmt.Printf("Kernel arg %d: %s\n", i, name)
		}
	}
	var err error
	p.input, err = p.context.CreateEmptyBuffer(MemReadOnly, 4*len(p.data))
	if err != nil {
		fmt.Errorf("CreateBuffer failed for input: %+v\n", err)
	}
	p.output, err = p.context.CreateEmptyBuffer(MemReadOnly, 4*len(p.data))
	if err != nil {
		fmt.Errorf("CreateBuffer failed for output: %+v\n", err)
	}
	if _, err = p.queue.EnqueueWriteBufferFloat32(p.input, true, 0, p.data, nil); err != nil {
		fmt.Errorf("EnqueueWriteBufferFloat32 failed: %+v\n", err)
	}
	if err = p.kernel.SetArgs(p.input, p.output, uint32(len(p.data))); err != nil {
		fmt.Errorf("SetKernelArgs failed: %+v\n", err)
	}
	return p.input, p.output
}

func (p Init) Run() error {
	local, err := p.kernel.WorkGroupSize(p.device)
	if err != nil {
		return fmt.Errorf("WorkGroupSize failed: %+v\n", err)
	}
	fmt.Printf("Work group size: %d\n", local)
	size, _ := p.kernel.PreferredWorkGroupSizeMultiple(nil)
	fmt.Printf("Preferred Work Group Size Multiple: %d\n", size)

	global := len(p.data)
	d := len(p.data) % local
	if d != 0 {
		global += local - d
	}
	if _, err := p.queue.EnqueueNDRangeKernel(p.kernel, nil, []int{global}, []int{local}, nil); err != nil {
		return fmt.Errorf("EnqueueNDRangeKernel failed: %+v\n", err)
	}

	if err := p.queue.Finish(); err != nil {
		return fmt.Errorf("Finish failed: %+v\n", err)
	}

	results := make([]float32, len(p.data))
	if _, err := p.queue.EnqueueReadBufferFloat32(p.output, true, 0, results, nil); err != nil {
		return fmt.Errorf("EnqueueReadBufferFloat32 failed: %+v\n", err)
	}

	correct := 0
	for i, v := range p.data {
		if results[i] == v*v {
			correct++
		}
	}

	if correct != len(p.data) {
		return fmt.Errorf("%d/%d correct values\n", correct, len(p.data))
	}
	return nil
}
