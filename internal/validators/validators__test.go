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
		name, version, err := ValidateModuleName("react")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if name != "react" {
			t.Errorf("expected module name 'react', but got '%s'", name)
		}
		if version != "latest" {
			t.Errorf("expected version 'latest', but got '%s'", version)
		}
	})

	t.Run("should handle module with version", func(t *testing.T) {
		name, version, err := ValidateModuleName("react@18.2.0")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if name != "react" {
			t.Errorf("expected module name 'react', but got '%s'", name)
		}
		if version != "18.2.0" {
			t.Errorf("expected version '18.2.0', but got '%s'", version)
		}
	})

	t.Run("should handle module with trailing @", func(t *testing.T) {
		name, version, err := ValidateModuleName("react@")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if name != "react" {
			t.Errorf("expected module name 'react', but got '%s'", name)
		}
		if version != "latest" {
			t.Errorf("expected version 'latest', but got '%s'", version)
		}
	})

	t.Run("should handle scoped module", func(t *testing.T) {
		name, version, err := ValidateModuleName("@angular/core")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if name != "@angular/core" {
			t.Errorf("expected module name '@angular/core', but got '%s'", name)
		}
		if version != "latest" {
			t.Errorf("expected version 'latest', but got '%s'", version)
		}
	})

	t.Run("should handle scoped module with version", func(t *testing.T) {
		name, version, err := ValidateModuleName("@angular/core@14")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if name != "@angular/core" {
			t.Errorf("expected module name '@angular/core', but got '%s'", name)
		}
		if version != "14" {
			t.Errorf("expected version '14', but got '%s'", version)
		}
	})

	t.Run("should handle scoped module with trailing @", func(t *testing.T) {
		name, version, err := ValidateModuleName("@angular/core@")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if name != "@angular/core" {
			t.Errorf("expected module name '@angular/core', but got '%s'", name)
		}
		if version != "latest" {
			t.Errorf("expected version 'latest', but got '%s'", version)
		}
	})

	t.Run("should handle problematic scoped module", func(t *testing.T) {
		name, version, err := ValidateModuleName("@anthropic-ai/claudde-code")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if name != "@anthropic-ai/claudde-code" {
			t.Errorf("expected module name '@anthropic-ai/claudde-code', but got '%s'", name)
		}
		if version != "latest" {
			t.Errorf("expected version 'latest', but got '%s'", version)
		}
	})

	t.Run("should return error for invalid module name '@'", func(t *testing.T) {
		_, _, err := ValidateModuleName("@")
		if err == nil {
			t.Error("expected error, but got nil")
		}
	})

	t.Run("should handle module with tilde in version", func(t *testing.T) {
		name, version, err := ValidateModuleName("string@~1.4.0")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if name != "string" {
			t.Errorf("expected module name 'string', but got '%s'", name)
		}
		if version != "1.4.0" {
			t.Errorf("expected version '1.4.0', but got '%s'", version)
		}
	})
}
