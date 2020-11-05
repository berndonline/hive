package deprovision

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/openshift/installer/pkg/destroy/vsphere"
	"github.com/openshift/installer/pkg/types"
	typesvsphere "github.com/openshift/installer/pkg/types/vsphere"

	"github.com/openshift/hive/pkg/constants"
)

// vSphereOptions is the set of options to deprovision an vSphere cluster
type vSphereOptions struct {
	logLevel string
	infraID  string
	vCenter  string
	username string
	password string
}

// NewDeprovisionvSphereCommand is the entrypoint to create the vSphere deprovision subcommand
func NewDeprovisionvSphereCommand() *cobra.Command {
	opt := &vSphereOptions{}
	cmd := &cobra.Command{
		Use:   "vsphere INFRAID",
		Short: "Deprovision vSphere assets (as created by openshift-installer)",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if err := opt.Complete(cmd, args); err != nil {
				log.WithError(err).Fatal("failed to complete options")
			}
			if err := opt.Validate(cmd); err != nil {
				log.WithError(err).Fatal("validation failed")
			}
			if err := opt.Run(); err != nil {
				log.WithError(err).Fatal("Runtime error")
			}
		},
	}
	flags := cmd.Flags()
	flags.StringVar(&opt.logLevel, "loglevel", "info", "log level, one of: debug, info, warn, error, fatal, panic")
	flags.StringVar(&opt.vCenter, "vsphere-vcenter", "", "Domain name or IP address of the vCenter")
	return cmd
}

// Complete finishes parsing arguments for the command
func (o *vSphereOptions) Complete(cmd *cobra.Command, args []string) error {
	o.infraID = args[0]
	return nil
}

// Validate ensures that option values make sense
func (o *vSphereOptions) Validate(cmd *cobra.Command) error {
	if o.vCenter == "" {
		o.vCenter = os.Getenv(constants.VSphereVCenterEnvVar)
		if o.vCenter == "" {
			return fmt.Errorf("must provide --vsphere-vcenter or set %s env var", constants.VSphereVCenterEnvVar)
		}
	}
	o.username = os.Getenv(constants.VSphereUsernameEnvVar)
	if o.username == "" {
		return fmt.Errorf("No %s env var set, cannot proceed", constants.VSphereUsernameEnvVar)
	}
	o.password = os.Getenv(constants.VSpherePasswordEnvVar)
	if o.password == "" {
		return fmt.Errorf("No %s env var set, cannot proceed", constants.VSpherePasswordEnvVar)
	}
	return nil
}

// Run executes the command
func (o *vSphereOptions) Run() error {
	// Set log level
	level, err := log.ParseLevel(o.logLevel)
	if err != nil {
		log.WithError(err).Error("cannot parse log level")
		return err
	}

	logger := log.NewEntry(&log.Logger{
		Out: os.Stdout,
		Formatter: &log.TextFormatter{
			FullTimestamp: true,
		},
		Hooks: make(log.LevelHooks),
		Level: level,
	})

	metadata := &types.ClusterMetadata{
		InfraID: o.infraID,
		ClusterPlatformMetadata: types.ClusterPlatformMetadata{
			VSphere: &typesvsphere.Metadata{
				VCenter:  o.vCenter,
				Username: o.username,
				Password: o.password,
			},
		},
	}

	destroyer, err := vsphere.New(logger, metadata)
	if err != nil {
		return err
	}

	return destroyer.Run()
}
