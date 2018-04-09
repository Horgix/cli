package autoscalers

import (
	"github.com/Scalingo/cli/config"
	"github.com/Scalingo/cli/io"
	"gopkg.in/errgo.v1"
)

func Remove(app, containerType string) error {
	autoscaler, err := GetFromContainerType(app, containerType)
	if err != nil {
		if err == ErrNotFound {
			io.Error("Container type " + containerType + " has no autoscaler on the app " + app + ".")
			return nil
		}
		return errgo.Mask(err, errgo.Any)
	}

	c := config.ScalingoClient()
	err = c.AutoscalerRemove(app, autoscaler.ID)
	if err != nil {
		return errgo.Mask(err, errgo.Any)
	}

	io.Status("Autoscaler removed on", app, "for", containerType, "containers")
	return nil
}
