package documentdb

import (
	"github.com/aquasecurity/defsec/rules/aws/documentdb"
	"github.com/aquasecurity/tfsec/internal/app/tfsec/scanner"
	"github.com/aquasecurity/tfsec/pkg/rule"
)

func init() {
	scanner.RegisterCheckRule(rule.Rule{
		BadExample: []string{`
 resource "aws_docdb_cluster" "bad_example" {
   cluster_identifier      = "my-docdb-cluster"
   engine                  = "docdb"
   master_username         = "foo"
   master_password         = "mustbeeightchars"
   backup_retention_period = 5
   preferred_backup_window = "07:00-09:00"
   skip_final_snapshot     = true
   enabled_cloudwatch_logs_exports = "something"
 }
 `},
		GoodExample: []string{`
 resource "aws_docdb_cluster" "good_example" {
   cluster_identifier      = "my-docdb-cluster"
   engine                  = "docdb"
   master_username         = "foo"
   master_password         = "mustbeeightchars"
   backup_retention_period = 5
   preferred_backup_window = "07:00-09:00"
   skip_final_snapshot     = true
   enabled_cloudwatch_logs_exports = "audit"
 }
 `},
		Links: []string{
			"https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/docdb_cluster#enabled_cloudwatch_logs_exports",
		},
		RequiredTypes: []string{
			"resource",
		},
		RequiredLabels: []string{
			"aws_docdb_cluster",
		},
		Base: documentdb.CheckEnableLogExport,
	})
}
