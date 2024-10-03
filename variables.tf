#Module      : LABEL
#Description : Terraform label module variables.
variable "name" {
  description = "Name  (e.g. `app` or `cluster`)."
  type        = string
  default     = ""
}

variable "environment" {
  description = "Environment (e.g. `prod`, `dev`, `staging`)."
  type        = string
  default     = ""
}

variable "repository" {
  description = "Terraform current module repo"
  type        = string
  default     = "https://github.com/eazycloudlife/terraform-aws-labels"
}

variable "business_unit" {
  description = "Top-level division of your company that owns the subscription or workload that the resource belongs to. In smaller organizations, this tag might represent a single corporate or shared top-level organizational element."
  type        = string
  default     = "Corp"
}

variable "label_order" {
  description = "Label order, e.g. sequence of application name and environment `name`,`environment`,'attribute' [`webserver`,`qa`,`devops`,`public`,] ."
  type        = list(any)
  default     = []
}

variable "attributes" {
  description = "Additional attributes (e.g. `1`)."
  type        = list(string)
  default     = []
}

variable "extra_tags" {
  description = "Additional tags (e.g. map(`BusinessUnit`,`XYZ`)."
  type        = map(string)
  default     = {}
}

variable "managed_by" {
  description = "managed_by, eg 'EazyCloudLife'."
  type        = string
  default     = "eazycloudlife@gmail.com"
}

variable "enabled" {
  description = "Set to false to prevent the module from creating any resources."
  type        = bool
  default     = true
}

variable "delimiter" {
  description = "Delimiter to be used between `organization`, `name`, `environment` and `attributes`."
  type        = string
  default     = "-"
}
