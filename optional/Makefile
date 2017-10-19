.PHONY: generate
generate:
	go run cmd/optionalize/main.go -wrapper=String                  -wrapped=string
	go run cmd/optionalize/main.go -wrapper=Bool                    -wrapped=bool
	go run cmd/optionalize/main.go -wrapper=UInt8                   -wrapped=uint8
	go run cmd/optionalize/main.go -wrapper=UInt16                  -wrapped=uint16
	go run cmd/optionalize/main.go -wrapper=UInt32                  -wrapped=uint32
	go run cmd/optionalize/main.go -wrapper=UInt64                  -wrapped=uint64
	go run cmd/optionalize/main.go -wrapper=Int8                    -wrapped=int8
	go run cmd/optionalize/main.go -wrapper=Int16                   -wrapped=int16
	go run cmd/optionalize/main.go -wrapper=Int32                   -wrapped=int32
	go run cmd/optionalize/main.go -wrapper=Int64                   -wrapped=int64
	go run cmd/optionalize/main.go -wrapper=Float32                 -wrapped=float32
	go run cmd/optionalize/main.go -wrapper=Float64                 -wrapped=float64
	go run cmd/optionalize/main.go -wrapper=Complex64               -wrapped=complex64
	go run cmd/optionalize/main.go -wrapper=Complex128              -wrapped=complex128
	go run cmd/optionalize/main.go -wrapper=Int                     -wrapped=int
	go run cmd/optionalize/main.go -wrapper=UInt                    -wrapped=uint
	go run cmd/optionalize/main.go -wrapper=Bytes                   -wrapped=[]byte
	go run cmd/optionalize/main.go -wrapper=Rune                    -wrapped=rune
	go run cmd/optionalize/main.go -wrapper=Error                   -wrapped=error
	go run cmd/optionalize/main.go -wrapper=Time                    -wrapped=time.Time

	go run cmd/optionalize/main.go -wrapper=BoolSlice               -wrapped=[]bool
	go run cmd/optionalize/main.go -wrapper=UInt8Slice              -wrapped=[]uint8
	go run cmd/optionalize/main.go -wrapper=UInt16Slice             -wrapped=[]uint16
	go run cmd/optionalize/main.go -wrapper=UInt32Slice             -wrapped=[]uint32
	go run cmd/optionalize/main.go -wrapper=UInt64Slice             -wrapped=[]uint64
	go run cmd/optionalize/main.go -wrapper=Int8Slice               -wrapped=[]int8
	go run cmd/optionalize/main.go -wrapper=Int16Slice              -wrapped=[]int16
	go run cmd/optionalize/main.go -wrapper=Int32Slice              -wrapped=[]int32
	go run cmd/optionalize/main.go -wrapper=Int64Slice              -wrapped=[]int64
	go run cmd/optionalize/main.go -wrapper=Float32Slice            -wrapped=[]float32
	go run cmd/optionalize/main.go -wrapper=Float64Slice            -wrapped=[]float64
	go run cmd/optionalize/main.go -wrapper=Complex64Slice          -wrapped=[]complex64
	go run cmd/optionalize/main.go -wrapper=Complex128Slice         -wrapped=[]complex128
	go run cmd/optionalize/main.go -wrapper=StringSlice             -wrapped=[]string
	go run cmd/optionalize/main.go -wrapper=IntSlice                -wrapped=[]int
	go run cmd/optionalize/main.go -wrapper=UIntSlice               -wrapped=[]uint
	go run cmd/optionalize/main.go -wrapper=BytesSlice              -wrapped=[][]byte
	go run cmd/optionalize/main.go -wrapper=RuneSlice               -wrapped=[]rune
	go run cmd/optionalize/main.go -wrapper=ErrorSlice              -wrapped=[]error
	go run cmd/optionalize/main.go -wrapper=TimeSlice               -wrapped=[]time.Time

	go run cmd/optionalize/main.go -wrapper=OptionalBoolSlice       -wrapped=Bool
	go run cmd/optionalize/main.go -wrapper=OptionalUInt8Slice      -wrapped=UInt8
	go run cmd/optionalize/main.go -wrapper=OptionalUInt16Slice     -wrapped=UInt16
	go run cmd/optionalize/main.go -wrapper=OptionalUInt32Slice     -wrapped=UInt32
	go run cmd/optionalize/main.go -wrapper=OptionalUInt64Slice     -wrapped=UInt64
	go run cmd/optionalize/main.go -wrapper=OptionalInt8Slice       -wrapped=Int8
	go run cmd/optionalize/main.go -wrapper=OptionalInt16Slice      -wrapped=Int16
	go run cmd/optionalize/main.go -wrapper=OptionalInt32Slice      -wrapped=Int32
	go run cmd/optionalize/main.go -wrapper=OptionalInt64Slice      -wrapped=Int64
	go run cmd/optionalize/main.go -wrapper=OptionalFloat32Slice    -wrapped=Float32
	go run cmd/optionalize/main.go -wrapper=OptionalFloat64Slice    -wrapped=Float64
	go run cmd/optionalize/main.go -wrapper=OptionalComplex64Slice  -wrapped=Complex64
	go run cmd/optionalize/main.go -wrapper=OptionalComplex128Slice -wrapped=Complex128
	go run cmd/optionalize/main.go -wrapper=OptionalStringSlice     -wrapped=String
	go run cmd/optionalize/main.go -wrapper=OptionalIntSlice        -wrapped=Int
	go run cmd/optionalize/main.go -wrapper=OptionalUIntSlice       -wrapped=UInt
	go run cmd/optionalize/main.go -wrapper=OptionalBytesSlice      -wrapped=Bytes
	go run cmd/optionalize/main.go -wrapper=OptionalRuneSlice       -wrapped=Rune
	go run cmd/optionalize/main.go -wrapper=OptionalErrorSlice      -wrapped=Error
	go run cmd/optionalize/main.go -wrapper=OptionalTimeSlice       -wrapped=Time

	goimports -w *.go

.PHONY: clean
clean:
	rm -rf *bool*.go
	rm -rf *uint8*.go
	rm -rf *uint16*.go
	rm -rf *uint32*.go
	rm -rf *uint64*.go
	rm -rf *int8*.go
	rm -rf *int16*.go
	rm -rf *int32*.go
	rm -rf *int64*.go
	rm -rf *float32*.go
	rm -rf *float64*.go
	rm -rf *complex64*.go
	rm -rf *complex128*.go
	rm -rf *string*.go
	rm -rf *int*.go
	rm -rf *uint*.go
	rm -rf *bytes*.go
	rm -rf *rune*.go
	rm -rf *error*.go
	rm -rf *time*.go
