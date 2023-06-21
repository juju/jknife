// Copyright 2023 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package util

import (
	"strconv"

	"github.com/juju/clock"
	"github.com/juju/errors"
	"github.com/juju/juju/agent"
	"github.com/juju/juju/mongo"
	"github.com/juju/juju/state"
	"github.com/juju/juju/state/stateenvirons"
)

func Connect(agentConfigPath string) (*state.StatePool, error) {
	agentConfig, err := agent.ReadConfig(agentConfigPath)
	if err != nil {
		return nil, err
	}

	info, ok := agentConfig.MongoInfo()
	if !ok {
		return nil, errors.New("no state info available")
	}
	dialOpts, err := mongoDialOptions(mongo.DefaultDialOpts(), agentConfig)
	if err != nil {
		return nil, errors.Trace(err)
	}
	session, err := mongo.DialWithInfo(*info, dialOpts)
	if err != nil {
		return nil, errors.Trace(err)
	}
	defer session.Close()

	pool, err := state.OpenStatePool(state.OpenParams{
		Clock:              clock.WallClock,
		ControllerTag:      agentConfig.Controller(),
		ControllerModelTag: agentConfig.Model(),
		MongoSession:       session,
		NewPolicy:          stateenvirons.GetNewPolicyFunc(),
	})
	if err != nil {
		return nil, errors.Trace(err)
	}
	return pool, nil
}

func mongoDialOptions(
	baseOpts mongo.DialOpts,
	agentConfig agent.Config,
) (mongo.DialOpts, error) {
	dialOpts := baseOpts
	if limitStr := agentConfig.Value("MONGO_SOCKET_POOL_LIMIT"); limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			return mongo.DialOpts{}, errors.Errorf("invalid mongo socket pool limit %q", limitStr)
		}
		dialOpts.PoolLimit = limit
	}
	return dialOpts, nil
}
