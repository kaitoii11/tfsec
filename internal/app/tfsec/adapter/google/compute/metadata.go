package compute

import (
	"github.com/aquasecurity/defsec/provider/google/compute"
	"github.com/aquasecurity/defsec/types"
	"github.com/aquasecurity/tfsec/internal/app/tfsec/block"
	"github.com/zclconf/go-cty/cty"
)

func adaptProjectMetadata(modules block.Modules) (metadata compute.ProjectMetadata) {
	metadata.EnableOSLogin = types.BoolUnresolvable(
		types.NewUnmanagedMetadata(
			types.NewRange("default", 0, 0),
			&types.FakeReference{},
		),
	)
	for _, metadataBlock := range modules.GetResourcesByType("google_compute_project_metadata") {
		if metadataAttr := metadataBlock.GetAttribute("metadata"); metadataAttr.IsNotNil() {
			if val := metadataAttr.MapValue("enable-oslogin"); val.Type() == cty.Bool {
				metadata.EnableOSLogin = types.BoolExplicit(val.True(), metadataAttr.Metadata())
			}
		}
	}
	return metadata
}
