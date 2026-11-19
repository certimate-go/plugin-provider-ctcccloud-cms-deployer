package cms

import (
	common "github.com/certimate-go/plugin-provider-ctcccloud-cms-deployer/internal/cloudsdk/ctyun/zz-shared-common"
)

func WithAkSk(ak, sk string) common.OptionsFunc {
	return common.WithAkSk(ak, sk)
}
