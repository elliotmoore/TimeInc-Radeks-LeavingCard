deployment_state "s3" {
  region = "us-west-1"
  bucket = "tf-deployment-state-{{.Environment}}"
  prefix = "{{.AwsAccountId}}"
}
