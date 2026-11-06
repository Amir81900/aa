package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func writeModelFile(path, data string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil { return err }
	return os.WriteFile(path, []byte(data), 0644)
}

func savePath() string {
	if base, err := os.UserConfigDir(); err == nil && base != "" {
		dir := filepath.Join(base, "JADDEH1405")
		_ = os.MkdirAll(dir, 0755)
		return filepath.Join(dir, "slot1.sav")
	}
	if exe, err := os.Executable(); err == nil { return filepath.Join(filepath.Dir(exe), "JADDEH1405.sav") }
	return "JADDEH1405.sav"
}

func saveGame(g *Game) bool {
	path := savePath()
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil { g.SaveNotice = "ذخیره ناموفق بود؛ دسترسی پوشه را بررسی کن"; return false }
	tmp := path + ".tmp"
	if err := os.WriteFile(tmp, []byte(g.SaveData()), 0600); err != nil { g.SaveNotice = "ذخیره ناموفق بود؛ فضای دیسک را بررسی کن"; return false }
	if err := os.Rename(tmp, path); err != nil { _ = os.Remove(tmp); g.SaveNotice = "ذخیره ناموفق بود"; return false }
	g.HasSave = true
	g.SaveNotice = "بازی ذخیره شد"
	return true
}

func loadGame(g *Game) bool {
	path := savePath()
	b, err := os.ReadFile(path)
	if err != nil {
		if exe, exErr := os.Executable(); exErr == nil {
			legacy := filepath.Join(filepath.Dir(exe), "JADDEH1405.sav")
			if legacy != path { if old, oldErr := os.ReadFile(legacy); oldErr == nil { if os.MkdirAll(filepath.Dir(path), 0755) == nil && os.WriteFile(path, old, 0600) == nil { b = old; err = nil } } }
		}
	}
	if err != nil { g.HasSave = false; return false }
	g.LoadData(string(b)); g.HasSave = true; g.SaveNotice = "ذخیره بارگذاری شد"; return true
}

func backupCurrentSave() bool {
	path := savePath()
	b, err := os.ReadFile(path)
	if err != nil { return false }
	backup := filepath.Join(filepath.Dir(path), "previous_trip.sav")
	return os.WriteFile(backup, b, 0600) == nil
}

func settingsPath() string {
	if base, err := os.UserConfigDir(); err == nil && base != "" {
		dir := filepath.Join(base, "JADDEH1405")
		_ = os.MkdirAll(dir, 0755)
		return filepath.Join(dir, "settings.cfg")
	}
	return filepath.Join(filepath.Dir(savePath()), "settings.cfg")
}

func saveSettings(g *Game) bool {
	data := fmt.Sprintf("quality=%d\nfov=%d\ncamera_distance=%.2f\nsteering_sensitivity=%.2f\nlook_sensitivity=%.3f\n", g.GraphicsQuality, g.CameraFOV, g.CameraDistance, g.SteeringSensitivity, g.LookSensitivity)
	return os.WriteFile(settingsPath(), []byte(data), 0600) == nil
}

func loadSettings(g *Game) bool {
	b, err := os.ReadFile(settingsPath())
	if err != nil { return false }
	for _, line := range strings.Split(string(b), "\n") {
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 { continue }
		v, _ := strconv.ParseFloat(parts[1], 64)
		switch parts[0] {
		case "quality": g.GraphicsQuality = int(clamp(v, 0, 2))
		case "fov": if v >= 55 && v <= 85 { g.CameraFOV = int(v) }
		case "camera_distance": if v >= 8 && v <= 16 { g.CameraDistance = v }
		case "steering_sensitivity": if v >= .45 && v <= 1.8 { g.SteeringSensitivity = v }
		case "look_sensitivity": if v >= .04 && v <= .5 { g.LookSensitivity = v }
		}
	}
	return true
}
