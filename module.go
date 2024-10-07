package main

import (
	"context"

	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/module"
	"go.viam.com/rdk/services/motion"

	"go.viam.com/utils"

	explore "github.com/jeremyrhyde/viam-explore/src"
)

func main() {
	utils.ContextualMain(mainWithArgs, module.NewLoggerFromArgs("Motion Random Explore Go Module"))
}

func mainWithArgs(ctx context.Context, args []string, logger logging.Logger) (err error) {
	// instantiates the module itself
	myMod, err := module.NewModuleFromArgs(ctx, logger)
	if err != nil {
		return err
	}

	// Models and APIs add helpers to the registry during their init().
	// They can then be added to the module here.
	err = myMod.AddModelFromRegistry(ctx, motion.API, explore.Model)
	if err != nil {
		return err
	}

	// Each module runs as its own process
	err = myMod.Start(ctx)
	defer myMod.Close(ctx)
	if err != nil {
		return err
	}
	<-ctx.Done()
	return nil
}
