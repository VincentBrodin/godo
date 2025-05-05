package engine

import (
	"github.com/VincentBrodin/godo/pkg/parser"
	"github.com/VincentBrodin/godo/pkg/utils"
)

type ResolvedCommand struct {
	Run   []string
	Where string
	Type  string
}

func resolve(cmd parser.Command) (ResolvedCommand, error) {
	where, err := utils.GetDir(cmd)
	if err != nil {
		return ResolvedCommand{}, err
	}

	run, _t, err := utils.GetRunAndType(cmd)
	if err != nil {
		return ResolvedCommand{}, err
	}

	var t string
	if _t == nil {
		t = utils.GetDefaultType()
	} else {
		t = *_t
	}

	return ResolvedCommand{
		Run:   run,
		Where: where,
		Type:  t,
	}, nil
}
