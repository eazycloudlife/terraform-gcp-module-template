##-----------------------------------------------------------------------------
## Labels module call.
##-----------------------------------------------------------------------------
module "labels" {
  source = "../../"

  name          = local.name
  environment   = "dev"
  label_order   = ["name", "environment"]
  business_unit = "Corp"
  attributes    = ["private"]
  managed_by    = local.managed_by

  extra_tags = merge(
    local.tags,
    {
      Application = "EazyCloudLife"
    }
  )
}
