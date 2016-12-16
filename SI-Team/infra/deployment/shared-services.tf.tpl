resource "terraform_remote_state" "shared-services" {
  backend = "s3"

  config {
    bucket = "tf-terraform-state-{{.Environment}}"
    key = "{{.AwsAccountId}}/shared-services/terraform.tfstate"
    region = "us-west-1"
  }
}
