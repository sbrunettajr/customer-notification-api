terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.27"
    }
  }
  required_version = ">= 0.14.9"
}

provider "aws" {
  profile = "default"
  region  = "us-east-1"
}

#resource "aws_sqs_queue" "sqs_created_customer" {
#  name = "sqs-created-customer"
#}

resource "aws_sns_topic" "sns_created_customer" {
  name = "sns_created_customer"
}

resource "aws_sns_topic_subscription" "sns_subscription_created_customer_email" {
  endpoint  = var.sns_subscription_created_customer_email_endpoint
  protocol  = "email"
  topic_arn = aws_sns_topic.sns_created_customer.arn
}

#resource "aws_sns_topic_subscription" "sns_subscription_created_customer_sqs" {
#  endpoint  = aws_sqs_queue.sqs_created_customer.arn
#  protocol  = "sqs"
#  topic_arn = aws_sns_topic.sns_created_customer.arn
#}