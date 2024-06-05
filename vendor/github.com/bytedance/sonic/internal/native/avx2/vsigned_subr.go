// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

package avx2

import (
	`github.com/bytedance/sonic/loader`
)

const (
    _entry__vsigned = 0
)

const (
    _stack__vsigned = 16
)

const (
    _size__vsigned = 320
)

var (
    _pcsp__vsigned = [][2]uint32{
        {1, 0},
        {4, 8},
        {112, 16},
        {113, 8},
        {114, 0},
        {125, 16},
        {126, 8},
        {127, 0},
        {260, 16},
        {261, 8},
        {262, 0},
        {266, 16},
        {267, 8},
        {268, 0},
        {306, 16},
        {307, 8},
        {308, 0},
        {316, 16},
        {317, 8},
        {320, 0},
    }
)

var _cfunc_vsigned = []loader.CFunc{
    {"_vsigned_entry", 0,  _entry__vsigned, 0, nil},
    {"_vsigned", _entry__vsigned, _size__vsigned, _stack__vsigned, _pcsp__vsigned},
}
