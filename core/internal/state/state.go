package state

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type State struct {
	Version   string                    `json:"version"`
	Resources map[string]*ResourceState `json:"resources"`
}

type ResourceState struct {
	Name       string                 `json:"name"`
	Type       string                 `json:"type"`
	Status     string                 `json:"status"`
	Attributes map[string]interface{} `json:"attributes"`
}

type Manager struct {
	stateDir  string
	stateFile string
}

func NewManager(stateDir string) *Manager {
	return &Manager{
		stateDir:  stateDir,
		stateFile: filepath.Join(stateDir, "tblang.tbstate"),
	}
}

func (m *Manager) LoadState() (*State, error) {
	if _, err := os.Stat(m.stateFile); os.IsNotExist(err) {
		return &State{
			Version:   "1.0",
			Resources: make(map[string]*ResourceState),
		}, nil
	}

	data, err := os.ReadFile(m.stateFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read state file: %w", err)
	}

	var state State
	if err := json.Unmarshal(data, &state); err != nil {
		return nil, fmt.Errorf("failed to parse state file: %w", err)
	}

	return &state, nil
}

func (m *Manager) SaveState(state *State) error {

	if err := os.MkdirAll(m.stateDir, 0755); err != nil {
		return fmt.Errorf("failed to create state directory: %w", err)
	}

	state.Version = "1.0"

	data, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal state: %w", err)
	}

	if err := os.WriteFile(m.stateFile, data, 0644); err != nil {
		return fmt.Errorf("failed to write state file: %w", err)
	}

	return nil
}

func (m *Manager) ClearState() error {
	if _, err := os.Stat(m.stateFile); os.IsNotExist(err) {
		return nil
	}

	return os.Remove(m.stateFile)
}

func (m *Manager) BackupState() error {
	if _, err := os.Stat(m.stateFile); os.IsNotExist(err) {
		return nil
	}

	backupFile := m.stateFile + ".backup"

	data, err := os.ReadFile(m.stateFile)
	if err != nil {
		return fmt.Errorf("failed to read state file for backup: %w", err)
	}

	if err := os.WriteFile(backupFile, data, 0644); err != nil {
		return fmt.Errorf("failed to write backup file: %w", err)
	}

	return nil
}
