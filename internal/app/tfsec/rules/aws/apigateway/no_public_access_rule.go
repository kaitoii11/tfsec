package apigateway

import (
	"github.com/aquasecurity/defsec/rules/aws/apigateway"
	"github.com/aquasecurity/tfsec/internal/app/tfsec/scanner"
	"github.com/aquasecurity/tfsec/pkg/rule"
)

func init() {
	scanner.RegisterCheckRule(rule.Rule{
		BadExample: []string{`
 resource "aws_api_gateway_rest_api" "MyDemoAPI" {
	
 }

 resource "aws_api_gateway_method" "bad_example" {
   rest_api_id   = aws_api_gateway_rest_api.MyDemoAPI.id
   resource_id   = aws_api_gateway_resource.MyDemoResource.id
   http_method   = "GET"
   authorization = "NONE"
 }
 `},
		GoodExample: []string{`
 resource "aws_api_gateway_rest_api" "MyDemoAPI" {
	
 }

 resource "aws_api_gateway_method" "good_example" {
   rest_api_id   = aws_api_gateway_rest_api.MyDemoAPI.id
   resource_id   = aws_api_gateway_resource.MyDemoResource.id
   http_method   = "GET"
   authorization = "AWS_IAM"
 }
 `,
			`
 resource "aws_api_gateway_rest_api" "MyDemoAPI" {
	
 }

 resource "aws_api_gateway_method" "good_example" {
   rest_api_id      = aws_api_gateway_rest_api.MyDemoAPI.id
   resource_id      = aws_api_gateway_resource.MyDemoResource.id
   http_method      = "GET"
   authorization    = "NONE"
   api_key_required = true
 }
 `,
			`
 resource "aws_api_gateway_rest_api" "MyDemoAPI" {
	
 }

 resource "aws_api_gateway_method" "good_example" {
   rest_api_id   = aws_api_gateway_rest_api.MyDemoAPI.id
   resource_id   = aws_api_gateway_resource.MyDemoResource.id
   http_method   = "OPTION"
   authorization = "NONE"
 }
 `},
		Links: []string{
			"https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/api_gateway_method#authorization",
		},
		Base: apigateway.CheckNoPublicAccess,
	})
}
