package connection

import (
	"context"
	"fmt"
	"strings"

	"github.com/DTunnel0/CheckUser-Go/src/domain/contract"
)

type sshConnection struct {
	executor contract.Executor
	next     contract.CountConnection
}

func NewSSHConnection(executor contract.Executor) contract.CountConnection {
	return &sshConnection{executor: executor}
}

func (ssh *sshConnection) SetNext(connection contract.CountConnection) {
	ssh.next = connection
}

func (s *sshConnection) ByUsername(ctx context.Context, username string) (int, error) {
	result, err := s.executor.Execute(ctx, "ps -eo args")
	if err != nil {
		return 0, err
	}

	counts := parseSSHConnections(result)
	totalConnections := counts[username]
	if s.next != nil {
		count, err := s.next.ByUsername(ctx, username)
		if err == nil {
			totalConnections += count
		}
	}

	return totalConnections, err
}

func (s *sshConnection) All(ctx context.Context) (int, error) {
	result, err := s.executor.Execute(ctx, "ps -eo args")
	if err != nil {
		return 0, fmt.Errorf("failed to execute command: %w", err)
	}

	counts := parseSSHConnections(result)
	totalConnections := 0
	for _, count := range counts {
		totalConnections += count
	}

	if s.next != nil {
		count, err := s.next.All(ctx)
		if err == nil {
			totalConnections += count
		}
	}

	return totalConnections, nil
}

func parseSSHConnections(raw string) map[string]int {
	counts := make(map[string]int)
	for _, line := range strings.Split(raw, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if !strings.HasPrefix(line, "sshd:") {
			continue
		}

		payload := strings.TrimSpace(strings.TrimPrefix(line, "sshd:"))
		atIndex := strings.Index(payload, "@")
		if atIndex == -1 {
			continue
		}

		username := strings.TrimSpace(payload[:atIndex])
		if username == "" {
			continue
		}

		counts[username]++
	}

	return counts
}
