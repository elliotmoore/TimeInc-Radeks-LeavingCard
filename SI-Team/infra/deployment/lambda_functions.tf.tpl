variable "app_name" {}
variable "environment" {}

module "thanks-radek" {
  source = "git@github.com:TimeIncOSS/tf-aws-lambda-function.git?ref=v1.0.0"

  runtime = "nodejs4.3"
  name = "${var.app_name}"
  handler = "thanks-radek.handler"
  description = "send message to radek"

  memory_size_in_mb = 128
  timeout_in_sec = 300

  # VPC settings
  subnet_ids = "${terraform_remote_state.shared-services.output.aws_private_subnet_ids}"
  security_group_ids = "${terraform_remote_state.infra.output.lambda_function_sg_id}"

  high_alarms_sns_topic_arn = "${terraform_remote_state.shared-services.output.high_alarm_sns_topic_arn}"
  low_alarms_sns_topic_arn = "${terraform_remote_state.shared-services.output.low_alarm_sns_topic_arn}"

  application = "${var.app_name}"
  environment = "${var.environment}"
}

resource "aws_sns_topic" "notifications" {
  name = "si-thanks-alerts"
}

resource "aws_sns_topic_subscription" "alert_sns_target" {
    topic_arn = "${aws_sns_topic.notifications.arn}"
    protocol  = "lambda"
    endpoint  = "${module.thanks-radek.function_arn}"
}

resource "aws_lambda_permission" "with_sns" {
    statement_id = "AllowExecutionFromSNS"
    action = "lambda:InvokeFunction"
    function_name = "${module.jobs-alerts.function_name}"
    principal = "sns.amazonaws.com"
    source_arn = "${aws_sns_topic.notifications.arn}"
}

resource "aws_iam_role_policy" "thanks-radek_role" {
    name = "lambda-si-thanks-alerts-iam"
    role = "${module.jobs-alerts.iam_role_name}"
    policy = <<POLICY
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": "sts:AssumeRole",
      "Resource": "*"
    }
  ]
}
POLICY
}
