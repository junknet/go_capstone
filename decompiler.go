package main

// #cgo LDFLAGS: -L./c_lib/lib -lcapstone
// #cgo CFLAGS:  -I./c_lib/include/capstone
// #include "capstone.h"
import "C"
import (
	"reflect"
	"unsafe"
)

type Instruction struct {
	Address  uint
	Size     uint
	Mnemonic string
	OpStr    string
}

var handle C.csh

func init() {
	C.cs_open(C.cs_arch(C.CS_ARCH_ARM), C.cs_mode(C.CS_MODE_THUMB), &handle)

}

func decompiler(input []byte, address uint64) []Instruction {
	var insn *C.cs_insn
	bptr := (*C.uint8_t)(unsafe.Pointer(&input[0]))
	disassembled := C.cs_disasm(
		handle,
		bptr,
		C.size_t(len(input)),
		C.uint64_t(address),
		C.size_t(0), // count
		&insn,
	)
	var insns []C.cs_insn
	h := (*reflect.SliceHeader)(unsafe.Pointer(&insns))
	h.Data = uintptr(unsafe.Pointer(insn))
	h.Len = int(disassembled)
	h.Cap = int(disassembled)
	decomposed := []Instruction{}
	for _, insn := range insns {
		decomp := new(Instruction)
		decomp.Address = uint(insn.address)
		decomp.Size = uint(insn.size)
		decomp.Mnemonic = C.GoString(&insn.mnemonic[0])
		decomp.OpStr = C.GoString(&insn.op_str[0])
		decomposed = append(decomposed, *decomp)
	}
	return decomposed

}
