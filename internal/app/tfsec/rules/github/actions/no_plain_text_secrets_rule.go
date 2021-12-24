package actions

// ATTENTION!
// This rule was autogenerated!
// Before making changes, consider updating the generator.

import (
	"github.com/aquasecurity/defsec/rules"
	"github.com/aquasecurity/defsec/rules/github/actions"
	"github.com/aquasecurity/tfsec/internal/app/tfsec/block"
	"github.com/aquasecurity/tfsec/internal/app/tfsec/scanner"
	"github.com/aquasecurity/tfsec/pkg/rule"
)

func init() {
	scanner.RegisterCheckRule(rule.Rule{
		BadExample: []string{`
resource "github_actions_environment_secret" "bad_example" {	 
	repository       = "my repository name"
	environment       = "my environment"
	secret_name       = "my secret name"
	plaintext_value   = "sensitive secret string"
}
`},
		GoodExample: []string{`
resource "github_actions_environment_secret" "good_example" {
	repository       = "my repository name"
	environment       = "my environment"
	secret_name       = "my secret name"
	encrypted_value   = var.some_encrypted_secret_string
}
`},
		Links: []string{
			"https://registry.terraform.io/providers/integrations/github/latest/docs/resources/actions_environment_secret",
			"https://docs.github.com/en/actions/security-guides/security-hardening-for-github-actions",
		},
		RequiredTypes: []string{
			"resource",
		},
		RequiredLabels: []string{
			"github_actions_environment_secret",
		},
		CheckTerraform: func(resourceBlock block.Block, _ block.Module) (results rules.Results) {

			plaintextValue := resourceBlock.GetAttribute("plaintext_value")
			if plaintextValue.IsNotNil() {
				results.Add("Resource '%s' has plaintext value set. For security reasons encrypted value should be set instead.", resourceBlock)

			}
			return results
		},
		Base: actions.CheckNoPlainTextActionEnvironmentSecrets,
	})
}
