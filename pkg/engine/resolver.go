package engine

import (
	"github.com/VincentBrodin/godo/pkg/parser"
	"github.com/VincentBrodin/godo/pkg/utils"
)

type ResolvedCommand struct {
	Run   []string
	Times int32
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

	var times int32 = 1
	if cmd.Times != nil {
		times = *cmd.Times
	} 

	return ResolvedCommand{
		Run:   run,
		Times: times,
		Where: where,
		Type:  t,
	}, nil
}
