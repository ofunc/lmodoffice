/*
Copyright 2019 by ofunc

This software is provided 'as-is', without any express or implied warranty. In
no event will the authors be held liable for any damages arising from the use of
this software.

Permission is granted to anyone to use this software for any purpose, including
commercial applications, and to alter it and redistribute it freely, subject to
the following restrictions:

1. The origin of this software must not be misrepresented; you must not claim
that you wrote the original software. If you use this software in a product, an
acknowledgment in the product documentation would be appreciated but is not
required.

2. Altered source versions must be plainly marked as such, and must not be
misrepresented as being the original software.

3. This notice may not be removed or altered from any source distribution.
*/

// A simple Lua module for converting various office documents into OOXML format files.
package lmodoffice

import (
	"ofunc/lua"
)

// Open opens the module.
func Open(l *lua.State) int {
	l.NewTable(0, 4)

	l.Push("version")
	l.Push("1.0.0")
	l.SetTableRaw(-3)

	l.Push("toxlsx")
	l.Push(lToXLSX)
	l.SetTableRaw(-3)

	l.Push("todocx")
	l.Push(lToDOCX)
	l.SetTableRaw(-3)

	l.Push("topptx")
	l.Push(lToPPTX)
	l.SetTableRaw(-3)

	return 1
}
