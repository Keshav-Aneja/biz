package validators

import (
	"testing"
)

func TestValidateModuleName(t *testing.T) {
	t.Run("should return error for empty module name", func(t *testing.T) {
		_, _, err := ValidateModuleName("")
		if err == nil {
			t.Error("expected error, but got nil")
		}
	})

	t.Run("should handle simple module name", func(t *testing.T) {
		name, latest, err := ValidateModuleName("react")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if name != "react" {
			t.Errorf("expected module name 'react', but got '%s'", name)
		}
		if !latest {
			t.Error("expected to fetch latest version, but it was false")
		}
	})

	t.Run("should handle module with version", func(t *testing.T) {
		name, latest, err := ValidateModuleName("react@18.2.0")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if name != "react@18.2.0" {
			t.Errorf("expected module name 'react@18.2.0', but got '%s'", name)
		}
		if latest {
			t.Error("expected not to fetch latest version, but it was true")
		}
	})

	t.Run("should handle module with trailing @", func(t *testing.T) {
		name, latest, err := ValidateModuleName("react@")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if name != "react" {
			t.Errorf("expected module name 'react', but got '%s'", name)
		}
		if !latest {
			t.Error("expected to fetch latest version, but it was false")
		}
	})

	t.Run("should handle scoped module", func(t *testing.T) {
		name, latest, err := ValidateModuleName("@angular/core")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if name != "@angular/core" {
			t.Errorf("expected module name '@angular/core', but got '%s'", name)
		}
		if !latest {
			t.Error("expected to fetch latest version, but it was false")
		}
	})

	t.Run("should handle scoped module with version", func(t *testing.T) {
		name, latest, err := ValidateModuleName("@angular/core@14")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if name != "@angular/core@14" {
			t.Errorf("expected module name '@angular/core@14', but got '%s'", name)
		}
		if latest {
			t.Error("expected not to fetch latest version, but it was true")
		}
	})

	t.Run("should handle scoped module with trailing @", func(t *testing.T) {
		name, latest, err := ValidateModuleName("@angular/core@")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if name != "@angular/core" {
			t.Errorf("expected module name '@angular/core', but got '%s'", name)
		}
		if !latest {
			t.Error("expected to fetch latest version, but it was false")
		}
	})

	t.Run("should handle problematic scoped module", func(t *testing.T) {
		name, latest, err := ValidateModuleName("@anthropic-ai/claudde-code")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if name != "@anthropic-ai/claudde-code" {
			t.Errorf("expected module name '@anthropic-ai/claudde-code', but got '%s'", name)
		}
		if !latest {
			t.Error("expected to fetch latest version, but it was false")
		}
	})

	t.Run("should return error for invalid module name '@'", func(t *testing.T) {
		_, _, err := ValidateModuleName("@")
		if err == nil {
			t.Error("expected error, but got nil")
		}
	})
}
