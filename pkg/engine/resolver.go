package engine

import "github.com/VincentBrodin/godo/pkg/parser"

type ResolvedCommand struct {
	Run   []string
	Where string
	Type  string
}

func resolve(cmd parser.Command) (ResolvedCommand, error) {
	where, err := getDir(cmd)
	if err != nil {
		return ResolvedCommand{}, err
	}

	run, _t, err := getRunAndType(cmd)
	if err != nil {
		return ResolvedCommand{}, err
	}

	var t string
	if _t == nil {
		t = "shell"
	} else {
		t = *_t
	}

	return ResolvedCommand{
		Run:   run,
		Where: where,
		Type:  t,
	}, nil
}
