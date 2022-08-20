package main

const (
	// Shell constants
	bash = "/bin/bash"
	sh   = "/bin/sh"
)

func GetSystemShell() string {
	if exists(bash) {
		return bash
	}
	return sh
}
