resource "aws_db_instance" "civicore" {
  identifier        = "civicore-prod"
  engine            = "postgres"
  engine_version    = "15.4"
  instance_class    = "db.t3.medium"
  allocated_storage = 100
  username          = "civicore"
  password          = "Civicore2024!"

  # CIV-IAC-001: public.
  publicly_accessible    = true
  # CIV-IAC-002: no encryption.
  storage_encrypted      = false
  # CIV-IAC-003: deletion protection off.
  deletion_protection    = false
  skip_final_snapshot    = true
}
