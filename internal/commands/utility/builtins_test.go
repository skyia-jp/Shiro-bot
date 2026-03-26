package utility

import (
	"testing"

	"github.com/skyia-jp/shiro-go/internal/commands"
	"go.uber.org/zap"
)

func TestRegisterBuiltinsContainsSimpleCommands(t *testing.T) {
	log := zap.NewNop().Sugar()
	r := commands.NewRouter(log, nil)

	if err := RegisterBuiltins(r); err != nil {
		t.Fatalf("RegisterBuiltins failed: %v", err)
	}

	defs := r.Definitions()
	got := make(map[string]bool, len(defs))
	for _, d := range defs {
		got[d.Name] = true
	}

	for _, name := range []string{"ping", "help", "uptime"} {
		if !got[name] {
			t.Fatalf("expected command %q to be registered", name)
		}
	}
	if got["reminder"] {
		t.Fatal("reminder command must not be registered")
	}
}
