// Copyright 2023 FlexCDN tjbrains@foxmail.com. All rights reserved. Official site: https://tjbrains.com .

package goman

import (
	"github.com/tjbrains/CloudLogger/internal/utils/zero"
	"runtime"
	"sync"
)

type TaskGroup struct {
	semi chan zero.Zero
	wg   *sync.WaitGroup
	mu   *sync.RWMutex
}

func NewTaskGroup() *TaskGroup {
	var concurrent = runtime.NumCPU()
	if concurrent <= 1 {
		concurrent = 2
	}
	return &TaskGroup{
		semi: make(chan zero.Zero, concurrent),
		wg:   &sync.WaitGroup{},
		mu:   &sync.RWMutex{},
	}
}

func (this *TaskGroup) Run(f func()) {
	this.wg.Add(1)
	go func() {
		defer this.wg.Done()

		this.semi <- zero.Zero{}

		f()

		<-this.semi
	}()
}

func (this *TaskGroup) Wait() {
	this.wg.Wait()
}

func (this *TaskGroup) Lock() {
	this.mu.Lock()
}

func (this *TaskGroup) Unlock() {
	this.mu.Unlock()
}
