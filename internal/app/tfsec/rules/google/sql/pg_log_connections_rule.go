package sql

import (
	"github.com/aquasecurity/defsec/rules"
	"github.com/aquasecurity/defsec/rules/google/sql"
	"github.com/aquasecurity/tfsec/internal/app/tfsec/block"
	"github.com/aquasecurity/tfsec/internal/app/tfsec/scanner"
	"github.com/aquasecurity/tfsec/pkg/rule"
)

func init() {
	scanner.RegisterCheckRule(rule.Rule{
		BadExample: []string{`
 resource "google_sql_database_instance" "db" {
 	name             = "db"
 	database_version = "POSTGRES_12"
 	region           = "us-central1"
 	settings {
 		database_flags {
 			name  = "log_connections"
 			value = "off"
 		}
 	}
 }
 			`},
		GoodExample: []string{`
 resource "google_sql_database_instance" "db" {
 	name             = "db"
 	database_version = "POSTGRES_12"
 	region           = "us-central1"
 	settings {
 		database_flags {
 			name  = "log_connections"
 			value = "on"
 		}
 	}
 }
 			`},
		Links: []string{
			"https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/sql_database_instance",
		},
		RequiredTypes:  []string{"resource"},
		RequiredLabels: []string{"google_sql_database_instance"},
		Base:           sql.CheckPgLogConnections,
		CheckTerraform: func(resourceBlock block.Block, _ block.Module) (results rules.Results) {
			if !resourceBlock.GetAttribute("database_version").StartsWith("POSTGRES") {
				return
			}

			settingsBlock := resourceBlock.GetBlock("settings")
			if settingsBlock.IsNil() {
				results.Add("Resource is not configured to log connections", resourceBlock)
				return
			}

			for _, dbFlagBlock := range settingsBlock.GetBlocks("database_flags") {
				if dbFlagBlock.GetAttribute("name").Equals("log_connections") {
					if valueAttr := dbFlagBlock.GetAttribute("value"); valueAttr.Equals("off") {
						results.Add("Resource is configured not to log connections", valueAttr)
					}
					return
				}
			}

			results.Add("Resource is not configured to log connections", resourceBlock)
			return results
		},
	})
}
