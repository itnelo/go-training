package advanced

import (
	"fmt"
	"strconv"
	"unsafe"
)

// struct alignment: 8

// https://go101.org/article/memory-layout.html

var (
	User struct {
		//bt byte         // 1 byte
		//i32  int32  // 4 bytes + padding 3 bytes
		i64 int64 // 8 bytes
		//ui64 uint64 // 8 bytes
		b bool   // 1 byte + padding 7 bytes
		s string // 16 bytes
		//fl32    float32    // 4 bytes + padding 4 bytes
		//fl64    float64    // 8 bytes
		//cmx64   complex64  // 8 bytes
		//cmx128  complex128 // 16 bytes
		//uiptr uintptr // 8 bytes
	}
)

func structAlignments() {
	fmt.Printf("\nunsafe.Sizeof(User) == %v bytes", unsafe.Sizeof(User))

	fmt.Printf("\nunsafe.Alignof(User.i64) == %v", unsafe.Alignof(User.i64))
	fmt.Printf("\nunsafe.Alignof(User.b) == %v", unsafe.Alignof(User.b))
	fmt.Printf("\nunsafe.Alignof(User.s) == %v", unsafe.Alignof(User.s))

	var offset_i64 uintptr = unsafe.Offsetof(User.i64)

	// way no1 to converet a uintptr to string - unsafe *int
	var offset_i64_int *int = (*int)(unsafe.Pointer(&offset_i64))

	fmt.Println("\nunsafe.Offsetof(User.i64) == " + strconv.Itoa(*offset_i64_int))

	var offset_b uintptr = unsafe.Offsetof(User.b)

	// way no2 - fmt.Sprint
	fmt.Println("\nunsafe.Offsetof(User.b) == " + fmt.Sprint(offset_b))
	fmt.Println("\nunsafe.Offsetof(User.s) == " + fmt.Sprint(unsafe.Offsetof(User.s)))
}

func memoryTraverse() {
	vals := []int{10, 20, 30, 40}
	start := unsafe.Pointer(&vals[0])
	size := unsafe.Sizeof(int(0))

	fmt.Printf("unsafe.Sizeof(int(0)) == %v\n", size)

	for i := 0; i < len(vals); i++ {
		var ptr unsafe.Pointer = unsafe.Pointer(uintptr(start) + size*uintptr(i))
		var item int = *(*int)(ptr)
		fmt.Printf("%d: %v\n", ptr, item)
	}
}

func absManual() {
	var x float64 = -67
	var x_uint uint64 = *(*uint64)(unsafe.Pointer(&x))
	x_uint = x_uint &^ (1 << 63)
	var y float64 = *(*float64)(unsafe.Pointer(&x_uint))
	fmt.Printf("binary: %5.2f\n", y)
}

func lowlevel() {
	var x uint8 = 24

	//24 / 2 = 12 (0)
	//12 / 2 = 6  (0)
	//6 / 2 = 3   (0)
	//3 / 2 = 1   (1)
	//        |__ (1)

	// 011000 = 1 * 2^4 + 1 * 2^3 = 16 + 8 = 24
	fmt.Printf("x(24) in binary: %b", x)

	// obtain memory address of variable "x"
	x_memaddr := unsafe.Pointer(&x)

	fmt.Printf("\nmem addr for x: %#v", x_memaddr)

	// copying value by memory address
	var y *uint8

	fmt.Printf("\ny == nil: %v", y == nil)

	y = (*uint8)(x_memaddr)

	// 24 / 16 = 1  (8)
	//           |_ (1)
	// 0x18
	fmt.Printf("\ncopied y value in hex: %#x", *y)

	var xx uint16 = 1<<16 - 1

	var xx_memaddr (unsafe.Pointer) = unsafe.Pointer(&xx)
	var b *byte
	b = (*byte)(xx_memaddr)

	fmt.Printf("\nunsafe.Pointer carries no type information, so value has been truncated: %v", *b)

	structAlignments()

	memoryTraverse()

	absManual()
}
