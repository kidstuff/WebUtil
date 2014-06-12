package dblog

import (
	"errors"
	"io"
	"net/http"
)

var (
	ErrNoProvider = errors.New("kidstuff/WebUtil/log: no provider found")
)

type LoggerProvider interface {
	OpenLogger(*http.Request) (io.Writer, error)
}

var (
	mapProvider  = make(map[string]LoggerProvider)
	lastProvider LoggerProvider
)

func LoggerRegister(name string, c LoggerProvider) error {
	_, ok := mapProvider[name]
	if ok {
		return ErrNoProvider
	}

	mapProvider[name] = c
	lastProvider = c
	return nil
}

func GetProvider(name string) (LoggerProvider, error) {
	p, ok := mapProvider[name]
	if !ok {
		return nil, ErrNoProvider
	}

	return p, nil
}

// Provider returns the last provider added. It will panic if there's no one.
func Provider() LoggerProvider {
	if lastProvider == nil {
		panic(ErrNoProvider.Error())
	}
	return lastProvider
}
