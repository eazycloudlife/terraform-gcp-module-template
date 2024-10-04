terraform {
  required_version = ">= 1.9.6"
}

# The simplest possible Terraform module: it just outputs "Hello, World!"
output "hello_world" {
  value = "Hello, World!"
}
