package main

import (
	"testing"

	ui "github.com/gizak/termui"
	"github.com/stretchr/testify/assert"
)

// TestCreateHostListType test if `createHostList()` return a value with a *termui.List type
func TestCreateHostListType(t *testing.T) {
	expected := &ui.List{}
	actual := createHostList()

	assert.IsType(t, expected, actual, "`createHostList()` should return a *termui.List")
}

// TestCreateCPUListType test if `createCPUList()` return a value with a *termui.List type
func TestCreateCPUListType(t *testing.T) {
	expected := &ui.List{}
	actual := createCPUList()

	assert.IsType(t, expected, actual, "`createCPUList()` should return a *termui.List")
}

// TestCreateBIOSListType test if `createBIOSList()` return a value with a *termui.List type
func TestCreateBIOSListType(t *testing.T) {
	expected := &ui.List{}
	actual := createBIOSList()

	assert.IsType(t, expected, actual, "`createBIOSList()` should return a *termui.List")
}

// TestCreateNetListType test if `createNetList()` return a value with a *termui.List type
func TestCreateNetListType(t *testing.T) {
	expected := &ui.List{}
	actual := createNetList()

	assert.IsType(t, expected, actual, "`createNetList()` should return a *termui.List")
}

// TestCreateProcListType test if `createProcList()` return a value with a *termui.List type
func TestCreateProcListType(t *testing.T) {
	expected := &ui.List{}
	actual := createProcList()

	assert.IsType(t, expected, actual, "`createProcList()` should return a *termui.List")
}

// TestCreateRAMGauge test if `createRAMGauge()` return a value with a *termui.Gauge type
func TestCreateRAMGaugeType(t *testing.T) {
	expected := &ui.Gauge{}
	actual := createRAMGauge()

	assert.IsType(t, expected, actual, "`createRAMGauge()` should return a *termui.Gauge")
}

// TestCreateSwapGaugeType test if `createSwapGauge()` return a value with a *termui.Gauge type
func TestCreateSwapGaugeType(t *testing.T) {
	expected := &ui.Gauge{}
	actual := createSwapGauge()

	assert.IsType(t, expected, actual, "`createSwapGauge()` should return a *termui.Gauge")
}

// TestCreateSwapGaugeType test if `createCPUGauge()` return a value with a *termui.Gauge type
func TestCreateCPUGaugeType(t *testing.T) {
	expected := &ui.Gauge{}
	actual := createCPUGauge()

	assert.IsType(t, expected, actual, "`createCPUGauge()` should return a *termui.Gauge")
}

// TestCreateDiskGaugeType test if `createDiskGauge()` return a value with a []*termui.Gauge array
func TestCreateDiskGaugeType(t *testing.T) {
	expected := []*ui.Gauge{}
	actual := createDiskGauge()

	assert.IsType(t, expected, actual, "`createDiskGauge()` should return a []*termui.Gauge")
}

// TestCreateBIOSListType test if `createQuitPar()` return a value with a *termui.Par type
func TestCreateQuitParType(t *testing.T) {
	expected := &ui.Par{}
	actual := createQuitPar()

	assert.IsType(t, expected, actual, "`createQuitPar()` should return a *termui.Par")
}
