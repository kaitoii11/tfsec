package ecr

import (
	"github.com/aquasecurity/defsec/rules/aws/ecr"
	"github.com/aquasecurity/tfsec/internal/app/tfsec/scanner"
	"github.com/aquasecurity/tfsec/pkg/rule"
)

func init() {
	scanner.RegisterCheckRule(rule.Rule{
		LegacyID: "AWS078",
		BadExample: []string{`
 resource "aws_ecr_repository" "bad_example" {
   name                 = "bar"
   image_tag_mutability = "MUTABLE"
 
   image_scanning_configuration {
     scan_on_push = true
   }
 }
 `},
		GoodExample: []string{`
 resource "aws_ecr_repository" "good_example" {
   name                 = "bar"
   image_tag_mutability = "IMMUTABLE"
 
   image_scanning_configuration {
     scan_on_push = true
   }
 }
 `},
		Links: []string{
			"https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ecr_repository",
		},
		RequiredTypes:  []string{"resource"},
		RequiredLabels: []string{"aws_ecr_repository"},
		Base:           ecr.CheckEnforceImmutableRepository,
	})
}
