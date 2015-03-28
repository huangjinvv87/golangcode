package main

import (
	"os"
	"os/signal"
	"reflect"
	"syscall"
)

type signalHandler struct {
	call []*callFunc
}

type callFunc struct {
	function reflect.Value
	params   []reflect.Value
}

func NewSignalHandler() *signalHandler {
	return &signalHandler{}
}

func NewCallFunc() *callFunc {
	return &callFunc{}
}

func (this *signalHandler) registerCallFunc(function interface{}, params ...interface{}) {
	call := NewCallFunc()
	call.function = reflect.ValueOf(function)
	if len(params) != call.function.Type().NumIn() {
		return
	}
	if len(params) > 0 {
		for _, value := range params {
			call.params = append(call.params, reflect.ValueOf(value))
		}
	}
	this.call = append(this.call, call)
}

func (this *signalHandler) callBack() {
	if len(this.call) > 0 {
		for _, f := range this.call {
			f.function.Call(f.params)
		}
	}
}

func exitSignalDeamon() {
	handler := NewSignalHandler()
	handler.registerCallFunc(clear, "clear work")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL, syscall.SIGTERM)
	select {
	case <-sigChan:
		handler.callBack()
		os.Exit(0)
	}

}
