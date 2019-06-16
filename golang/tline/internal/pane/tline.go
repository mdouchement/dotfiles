package pane

import (
	"fmt"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

const (
	colour16  = "colour16"
	colour154 = "colour154"
	colour239 = "colour239"
	colour250 = "colour250"
)

// StatusBar returns the pretty-printed status bar.
func StatusBar() string {
	return fmt.Sprint(
		withColors(colour154, "", ""),
		withColors(colour16, colour154, " ", topCPU(), " "),
		withColors(colour239, "", ""),
		withColors(colour16, colour154, " ", topRAM(), " "),
		withColors(colour239, colour154, ""),
		withColors(colour250, colour239, " ", now(), " "),
		"\n",
	)
}

func now() string {
	return time.Now().Format("2006-01-02 15:04")
}

func topCPU() string {
	// percs, err := cpu.Percent(0, false)
	// if err != nil {
	// 	return err.Error()
	// }
	// return fmt.Sprintf("%.2f%%", percs[0])

	avg, err := load.Avg()
	if err != nil {
		return err.Error()
	}
	return fmt.Sprintf("%.2f%%", avg.Load1)
}

func topRAM() string {
	vms, err := mem.VirtualMemory()
	if err != nil {
		return err.Error()
	}
	return humanize.Bytes(vms.Total - vms.Available)
}

func withColors(fg, bg string, payload ...string) string {
	// `#[default]` for reset color
	p := strings.Join(payload, "")
	if bg == "" {
		return fmt.Sprintf("#[fg=%s]%s", fg, p)
	}
	return fmt.Sprintf("#[fg=%s,bg=%s]%s", fg, bg, p)
}
