package commands

import (
	"testing"

	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

func TestRegisterDuplicateCommand(t *testing.T) {
	log := zap.NewNop().Sugar()
	r := NewRouter(log, nil)

	cmd := Command{
		Definition: &discordgo.ApplicationCommand{Name: "ping", Description: "p"},
		Handler:    func(*Context) error { return nil },
	}

	if err := r.Register(cmd); err != nil {
		t.Fatalf("first register failed: %v", err)
	}
	if err := r.Register(cmd); err == nil {
		t.Fatal("expected duplicate register error")
	}
}

func TestDefinitionsSorted(t *testing.T) {
	log := zap.NewNop().Sugar()
	r := NewRouter(log, nil)

	_ = r.Register(Command{Definition: &discordgo.ApplicationCommand{Name: "zeta", Description: "z"}, Handler: func(*Context) error { return nil }})
	_ = r.Register(Command{Definition: &discordgo.ApplicationCommand{Name: "alpha", Description: "a"}, Handler: func(*Context) error { return nil }})

	defs := r.Definitions()
	if len(defs) != 2 {
		t.Fatalf("expected 2 definitions, got %d", len(defs))
	}
	if defs[0].Name != "alpha" || defs[1].Name != "zeta" {
		t.Fatalf("definitions are not sorted: %s, %s", defs[0].Name, defs[1].Name)
	}
}
