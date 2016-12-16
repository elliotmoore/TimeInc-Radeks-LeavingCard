resource "terraform_remote_state" "infra" {
  backend = "s3"

  config {
    bucket = "tf-terraform-state-{{.Environment}}"
    key = "{{.AwsAccountId}}/{{.AppName}}/terraform.tfstate"
    region = "us-west-1"
  }
}
