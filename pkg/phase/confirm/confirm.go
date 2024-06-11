package confirm

import (
	"github.com/kubesphere/kubekey/pkg/common"
	"github.com/kubesphere/kubekey/pkg/core/task"
)

type UpgradeK8sConfirmModule struct {
	common.KubeModule
}

func (u *UpgradeK8sConfirmModule) GetName() string {
	return "UpgradeK8sConfirmModule"
}

func (u *UpgradeK8sConfirmModule) Init() {
	u.Name = "UpgradeKsConfirmModule"
	u.Desc = "Display upgrade kubesphere confirmation form"

	display := &task.LocalTask{
		Name:   "ConfirmForm",
		Desc:   "Display confirmation form",
		Action: new(UpgradeK8sConfirm),
	}

	u.Tasks = []task.Interface{
		display,
	}
}

type UpgradeKsConfirmModule struct {
	common.KubeModule
}

func (u *UpgradeKsConfirmModule) GetName() string {
	return "UpgradeKsConfirmModule"
}

func (u *UpgradeKsConfirmModule) Init() {
	u.Name = "UpgradeKsConfirmModule"
	u.Desc = "Display upgrade kubesphere confirmation form"

	display := &task.LocalTask{
		Name:   "ConfirmForm",
		Desc:   "Display confirmation form",
		Action: new(UpgradeKsConfirm),
	}

	u.Tasks = []task.Interface{
		display,
	}
}

type CreateK8sConfirmModule struct {
	common.KubeModule
}

func (u *CreateK8sConfirmModule) GetName() string {
	return "CreateK8sConfirmModule"
}

func (u *CreateK8sConfirmModule) Init() {
	u.Name = "CreateKsConfirmModule"
	u.Desc = "Display Create kubesphere confirmation form"

	display := &task.LocalTask{
		Name:   "ConfirmForm",
		Desc:   "Display confirmation form",
		Action: new(CreateK8sConfirm),
	}

	u.Tasks = []task.Interface{
		display,
	}
}

type CreateKsConfirmModule struct {
	common.KubeModule
}

func (u *CreateKsConfirmModule) GetName() string {
	return "CreateKsConfirmModule"
}

func (u *CreateKsConfirmModule) Init() {
	u.Name = "CreateKsConfirmModule"
	u.Desc = "Display Create kubesphere confirmation form"

	display := &task.LocalTask{
		Name:   "ConfirmForm",
		Desc:   "Display confirmation form",
		Action: new(CreateKsConfirm),
	}

	u.Tasks = []task.Interface{
		display,
	}
}
