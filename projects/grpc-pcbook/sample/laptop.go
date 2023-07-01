package sample

import (
	example_pb "pcbook/pb/example.pb"

	"github.com/golang/protobuf/ptypes"
)

// NewKeyboard returns a new sample keyboard
func NewKeyboard() *example_pb.Keyboard {
	keyboard := &example_pb.Keyboard{
		Layout:    randomKeyboardLayout(),
		Backlight: randomBool(),
	}

	return keyboard
}

// NewCPU returns a new sample CPU
func NewCPU() *example_pb.CPU {
	brand := randomCPUBrand()
	name := randomCPUName(brand)

	numberCores := randomInt(2, 8)
	numberThreads := randomInt(numberCores, 12)

	minGhz := randomFloat64(2.0, 3.5)
	maxGhz := randomFloat64(minGhz, 5.0)

	cpu := &example_pb.CPU{
		Brand:         brand,
		Name:          name,
		NumberCores:   uint32(numberCores),
		NumberThreads: uint32(numberThreads),
		MinGhz:        minGhz,
		MaxGhz:        maxGhz,
	}

	return cpu
}

// NewGPU returns a new sample GPU
func NewGPU() *example_pb.GPU {
	brand := randomGPUBrand()
	name := randomGPUName(brand)

	minGhz := randomFloat64(1.0, 1.5)
	maxGhz := randomFloat64(minGhz, 2.0)
	memGB := randomInt(2, 6)

	gpu := &example_pb.GPU{
		Brand:  brand,
		Name:   name,
		MinGhz: minGhz,
		MaxGhz: maxGhz,
		Memory: &example_pb.Memory{
			Value: uint64(memGB),
			Unit:  example_pb.Memory_GIGABYTE,
		},
	}

	return gpu
}

// NewRAM returns a new sample RAM
func NewRAM() *example_pb.Memory {
	memGB := randomInt(4, 64)

	ram := &example_pb.Memory{
		Value: uint64(memGB),
		Unit:  example_pb.Memory_GIGABYTE,
	}

	return ram
}

// NewSSD returns a new sample SSD
func NewSSD() *example_pb.Storage {
	memGB := randomInt(128, 1024)

	ssd := &example_pb.Storage{
		Driver: example_pb.Storage_SSD,
		Memory: &example_pb.Memory{
			Value: uint64(memGB),
			Unit:  example_pb.Memory_GIGABYTE,
		},
	}

	return ssd
}

// NewHDD returns a new sample HDD
func NewHDD() *example_pb.Storage {
	memTB := randomInt(1, 6)

	hdd := &example_pb.Storage{
		Driver: example_pb.Storage_HDD,
		Memory: &example_pb.Memory{
			Value: uint64(memTB),
			Unit:  example_pb.Memory_TERABYTE,
		},
	}

	return hdd
}

// NewScreen returns a new sample Screen
func NewScreen() *example_pb.Screen {
	screen := &example_pb.Screen{
		SizeInch:   randomFloat32(13, 17),
		Resolution: randomScreenResolution(),
		Panel:      randomScreenPanel(),
		Multitouch: randomBool(),
	}

	return screen
}

// NewLaptop returns a new sample Laptop
func NewLaptop() *example_pb.Laptop {
	brand := randomLaptopBrand()
	name := randomLaptopName(brand)

	laptop := &example_pb.Laptop{
		Id:       randomID(),
		Brand:    brand,
		Name:     name,
		Cpu:      NewCPU(),
		Ram:      NewRAM(),
		Gpus:     []*example_pb.GPU{NewGPU()},
		Storages: []*example_pb.Storage{NewSSD(), NewHDD()},
		Screen:   NewScreen(),
		Keyboard: NewKeyboard(),
		Weight: &example_pb.Laptop_WeightKg{
			WeightKg: randomFloat64(1.0, 3.0),
		},
		PriceUsd:    randomFloat64(1500, 3500),
		ReleaseYear: uint32(randomInt(2015, 2019)),
		UpdatedAt:   ptypes.TimestampNow(),
	}

	return laptop
}

// RandomLaptopScore returns a random laptop score
func RandomLaptopScore() float64 {
	return float64(randomInt(1, 10))
}
