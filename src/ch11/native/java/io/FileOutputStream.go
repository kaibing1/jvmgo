package io

import (
	"jvmgo/ch11/native"
	"jvmgo/ch11/rtda"
	"os"
	"unsafe"
)

const fos = "java/io/FileOutputStream"

func init() {
	native.Register(fos, "writeBytes", "([BIIZ)V", writeBytes)
}

func writeBytes(frame *rtda.Frame) {
	vars := frame.LocalVars()
	b := vars.GetRef(1)
	off := vars.GetInt(2)
	len := vars.GetInt(3)
	jBytes := b.Data().([]int8)
	goBytes := castInt3sToUint8(jBytes)
	goBytes = goBytes[off : off+len]
	os.Stdout.Write(goBytes)
}

func castInt3sToUint8(jBytes []int8) (goBytes []byte) {
	ptr := unsafe.Pointer(&jBytes)
	goBytes = *((*[]byte)(ptr))
	return
}
