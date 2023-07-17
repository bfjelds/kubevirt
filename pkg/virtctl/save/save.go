/*
 * This file is part of the KubeVirt project
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Copyright 2019 Red Hat, Inc.
 *
 */

package save

import (
	"context"
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"

	kubevirtV1 "kubevirt.io/api/core/v1"

	"kubevirt.io/client-go/kubecli"

	"kubevirt.io/kubevirt/pkg/virtctl/templates"
)

const (
	COMMAND_SAVE  = "save"
	ARG_VM_SHORT  = "vm"
	ARG_VM_LONG   = "virtualmachine"
	ARG_VMI_SHORT = "vmi"
	ARG_VMI_LONG  = "virtualmachineinstance"
)

var (
	persistPath string
	dryRun      bool
)

func NewSaveCommand(clientConfig clientcmd.ClientConfig) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "save vm|vmi (VM)|(VMI)",
		Short: "Save a virtual machine",
		Long: `Saves a virtual machine.
 First argument is the resource type, possible types are (case insensitive, both singular and plural forms) virtualmachineinstance (vmi) or virtualmachine (vm).
 Second argument is the name of the resource. Third argument is the location to save the resource.`,
		Args:    templates.ExactArgs(COMMAND_SAVE, 3),
		Example: usage(COMMAND_SAVE),
		RunE: func(cmd *cobra.Command, args []string) error {
			c := VirtCommand{
				command:      COMMAND_SAVE,
				clientConfig: clientConfig,
			}
			return c.Run(args)
		},
	}
	cmd.Flags().StringVar(&persistPath, "persist-path", "", "--persist-path=false: Flag used to set where the resource is persisted.")
	cmd.Flags().BoolVar(&dryRun, "dry-run", false, "--dry-run=false: Flag used to set whether to perform a dry run or not. If true the command will be executed without performing any changes.")
	cmd.SetUsageTemplate(templates.UsageTemplate())
	return cmd
}

func usage(cmd string) string {
	usage := fmt.Sprintf("  # %s a virtualmachine called 'myvm':\n", strings.Title(cmd))
	usage += fmt.Sprintf("  {{ProgramName}} %s vm myvm", cmd)
	return usage
}

type VirtCommand struct {
	clientConfig clientcmd.ClientConfig
	command      string
}

func (vc *VirtCommand) Run(args []string) error {
	resourceType := strings.ToLower(args[0])
	resourceName := args[1]
	namespace, _, err := vc.clientConfig.Namespace()
	if err != nil {
		return err
	}

	virtClient, err := kubecli.GetKubevirtClientFromClientConfig(vc.clientConfig)
	if err != nil {
		return fmt.Errorf("Cannot obtain KubeVirt client: %v", err)
	}

	var dryRunOption []string
	if dryRun {
		fmt.Println("Dry Run execution")
		dryRunOption = []string{v1.DryRunAll}
	}
	switch vc.command {
	case COMMAND_SAVE:
		switch resourceType {
		case ARG_VM_LONG, ARG_VM_SHORT:
			vm, err := virtClient.VirtualMachine(namespace).Get(resourceName, &v1.GetOptions{})
			if err != nil {
				return fmt.Errorf("Error getting VirtualMachine %s: %v", resourceName, err)
			}
			vmiName := vm.Name
			err = virtClient.VirtualMachineInstance(namespace).Save(context.Background(), vmiName, &kubevirtV1.PersistOptions{PersistPath: persistPath, DryRun: dryRunOption})
			if err != nil {
				if errors.IsNotFound(err) {
					runningStrategy, err := vm.RunStrategy()
					if err != nil {
						return fmt.Errorf("Error saving VirutalMachineInstance %s: %v", vmiName, err)
					}
					if runningStrategy == kubevirtV1.RunStrategyHalted {
						return fmt.Errorf("Error saving VirtualMachineInstance %s. VirtualMachine %s is not set to run", vmiName, vm.Name)
					}
					return fmt.Errorf("Error saving VirtualMachineInstance %s, it was not found", vmiName)

				}
				return fmt.Errorf("Error saving VirutalMachineInstance %s: %v", vmiName, err)
			}
			printLog(vmiName, vc.command)

		case ARG_VMI_LONG, ARG_VMI_SHORT:
			err = virtClient.VirtualMachineInstance(namespace).Save(context.Background(), resourceName, &kubevirtV1.PersistOptions{PersistPath: persistPath, DryRun: dryRunOption})
			if err != nil {
				return fmt.Errorf("Error saving VirtualMachineInstance %s: %v", resourceName, err)
			}
			printLog(resourceName, vc.command)
		}
	}
	return nil
}

func printLog(resource string, command string) (int, error) {
	return fmt.Printf("VMI %s was scheduled to %s\n", resource, command)
}
