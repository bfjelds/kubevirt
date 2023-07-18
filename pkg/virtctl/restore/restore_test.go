package restore_test

import (
	"context"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"kubevirt.io/client-go/api"

	"github.com/spf13/cobra"
	k8smetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1 "kubevirt.io/api/core/v1"
	"kubevirt.io/client-go/kubecli"

	"kubevirt.io/kubevirt/pkg/virtctl/restore"
	"kubevirt.io/kubevirt/tests/clientcmd"
)

var _ = Describe("Saving", func() {

	const vmName = "testvm"
	var vmInterface *kubecli.MockVirtualMachineInterface
	var vmiInterface *kubecli.MockVirtualMachineInstanceInterface
	var ctrl *gomock.Controller

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		kubecli.GetKubevirtClientFromClientConfig = kubecli.GetMockKubevirtClientFromClientConfig
		kubecli.MockKubevirtClientInstance = kubecli.NewMockKubevirtClient(ctrl)
		vmInterface = kubecli.NewMockVirtualMachineInterface(ctrl)
		vmiInterface = kubecli.NewMockVirtualMachineInstanceInterface(ctrl)
	})

	Context("With missing input parameters", func() {
		It("should fail a restore", func() {
			cmd := clientcmd.NewRepeatableVirtctlCommand(pause.COMMAND_SAVE)
			err := cmd()
			Expect(err).To(HaveOccurred())
		})
	})

	DescribeTable("should restore VMI", func(persistOptions *v1.PersistOptions) {

		vmi := api.NewMinimalVMI(vmName)

		kubecli.MockKubevirtClientInstance.EXPECT().VirtualMachineInstance(k8smetav1.NamespaceDefault).Return(vmiInterface).Times(1)
		vmiInterface.EXPECT().Restore(context.Background(), vmi.Name, persistOptions).Return(nil).Times(1)

		var command *cobra.Command
		if len(persistOptions.DryRun) == 0 {
			command = clientcmd.NewVirtctlCommand(restore.COMMAND_RESTORE, "vmi", vmName)
		} else {
			command = clientcmd.NewVirtctlCommand(restore.COMMAND_RESTORE, "--dry-run", "vmi", vmName)
		}
		Expect(command.Execute()).To(Succeed())
	},
		Entry("", &v1.PersistOptions{}),
		Entry("with dry-run option", &v1.PersistOptions{PersistPath: "", DryRun: []string{k8smetav1.DryRunAll}}),
	)

	DescribeTable("should restore VM", func(persistOptions *v1.PersistOptions) {

		vmi := api.NewMinimalVMI(vmName)
		vm := kubecli.NewMinimalVM(vmName)
		vm.Spec.Template = &v1.VirtualMachineInstanceTemplateSpec{
			Spec: vmi.Spec,
		}

		kubecli.MockKubevirtClientInstance.EXPECT().VirtualMachine(k8smetav1.NamespaceDefault).Return(vmInterface).Times(1)
		kubecli.MockKubevirtClientInstance.EXPECT().VirtualMachineInstance(k8smetav1.NamespaceDefault).Return(vmiInterface).Times(1)

		vmInterface.EXPECT().Get(vm.Name, &k8smetav1.GetOptions{}).Return(vm, nil).Times(1)
		vmiInterface.EXPECT().Restore(context.Background(), vm.Name, persistOptions).Return(nil).Times(1)

		var command *cobra.Command
		if len(persistOptions.DryRun) == 0 {
			command = clientcmd.NewVirtctlCommand(restore.COMMAND_RESTORE, "vm", vmName)
		} else {
			command = clientcmd.NewVirtctlCommand(restore.COMMAND_RESTORE, "--dry-run", "vm", vmName)
		}
		Expect(command.Execute()).To(Succeed())
	},
		Entry("", &v1.PersistOptions{}),
		Entry("with dry-run option", &v1.PersistOptions{PersistPath: "", DryRun: []string{k8smetav1.DryRunAll}}),
	)

})
