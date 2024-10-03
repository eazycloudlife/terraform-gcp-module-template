locals {
  name       = "ezcl-${basename(path.cwd)}"
  managed_by = "eazycloudlife@gmail.com"
  region     = "us-east1"
  project    = "project"

  tags = {
    Name       = local.name
    Example    = basename(path.cwd)
    Repository = "https://github.com/eazycloudlife/terraform-aws-labels"
  }
}
