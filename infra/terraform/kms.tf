# CIV-IAC-005: KMS key with full wildcard.
resource "aws_kms_key" "pii" {
  description             = "Civicore citizen-PII key"
  deletion_window_in_days = 7
  policy = jsonencode({
    Version = "2012-10-17",
    Statement = [{
      Effect    = "Allow",
      Principal = "*",
      Action    = "kms:*",
      Resource  = "*"
    }]
  })
}
