/* MIT License
#
# Copyright (c) 2020 Ferhat Geçdoğan All Rights Reserved.
# Distributed under the terms of the MIT License.
#
# */

package main

import (
	"fmt"
	"strconv"
)

// Default Background Color Definitions
var ( 
  Esc = 033
  Default = 39
  CodeBlack = 30
  CodeRed = 31
  CodeGreen = 32
  CodeYellow = 33
  CodeBlue = 34
  CodeMagenta = 35
  CodeCyan = 36
  CodeLightGray = 37
  CodeDarkGray = 90
  CodeLightRed = 91
  CodeLightGreen = 92
  CodeLightYellow = 93
  CodeLightBlue = 94
  CodeLightMagenta = 95
  CodeLightCyan = 96
  CodeWhite = 97
)

// Default Foreground Color Definitions
var (
  FDefault = 49
  FBlack = 40
  FRed = 41
  FGreen = 42
  FYellow = 43
  FBlue = 44
  FMagenta = 45
  FCyan = 46
  FLightGray = 47
  FDarkGray = 100
  FLightRed = 101
  FLightGreen = 102
  FLightYellow = 103
  FLightBlue = 104
  FLightMagenta = 105
  FLightCyan = 106
  FWhite = 107	
)

// Default Set Type Definitions
var (
  SDefault = 0
  Bold = 1
  Dim = 2
  Underlined = 4
  Blink = 5 
  Reverse = 7
  Hidden = 8
)

// Default Unset Type Definitions
var (
  UAll = 0
  UBold = 21
  UDim = 22
  UUnderlined = 24
  UBlink = 25
  UReverse = 27
  UHidden = 28
)

var (
	Semicolon = ";"
	Mark = "m"
	Template = "\033["
)


// Converter
func IntToString(converter int) string {
	return strconv.Itoa(converter)
}

var (
  Black   = Color("\033[1;30m%s\033[0m")
  Red     = Color("\033[1;31m%s\033[0m")
  Green   = Color("\033[1;32m%s\033[0m")
  Yellow  = Color("\033[1;33m%s\033[0m")
  Purple  = Color("\033[1;34m%s\033[0m")
  Magenta = Color("\033[1;35m%s\033[0m")
  Teal    = Color("\033[1;36m%s\033[0m")
  White   = Color("\033[1;37m%s\033[0m")
  Other   = Color(Template + IntToString(Bold) + Semicolon + IntToString(CodeGreen) + Mark + "%s\033[0m")
)


func Color(colorString string) func(...interface{}) string {
  sprint := func(args ...interface{}) string {
    return fmt.Sprintf(colorString,
      fmt.Sprint(args...))
  }
  return sprint
}

// For Test
func main() {
	fmt.Print(Other("Hello!"))
}

