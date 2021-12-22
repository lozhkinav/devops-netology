provider "aws" {
        region = "us-east-2"
}
resource "aws_s3_bucket" "bucket" {
  bucket = "netology-bucket-${terraform.workspace}"
  acl    = "private"
  tags = {
    Name        = "Bucket1"
    Environment = terraform.workspace
  }
}