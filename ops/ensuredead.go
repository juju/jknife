package ops

import (
	"github.com/juju/jknife/util"
)

func UnitEnsureDead(agentConfigPath string, modelUUID string, unitName string) error {
	pool, err := util.Connect(agentConfigPath)
	if err != nil {
		return err
	}
	defer pool.Close()
	model, ps, err := pool.GetModel(modelUUID)
	if err != nil {
		return err
	}
	defer ps.Release()
	unit, err := model.State().Unit(unitName)
	if err != nil {
		return err
	}
	return unit.EnsureDead()
}
