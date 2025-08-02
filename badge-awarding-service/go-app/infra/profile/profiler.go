package profile

import (
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

func SaveMemoryProfile(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	runtime.GC() //GCをしてメモリの情報を取得
	if err := pprof.WriteHeapProfile(f); err != nil {
		return err
	}
	log.Printf("Memory profile written to %s", path)
	return nil
}
