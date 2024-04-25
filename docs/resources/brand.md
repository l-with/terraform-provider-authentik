---
page_title: "authentik_brand Resource - terraform-provider-authentik"
subcategory: "System"
description: |-
  
---

# authentik_brand (Resource)



## Example Usage

```terraform
# Create/manage a default brand

resource "authentik_brand" "default" {
  domain         = "."
  default        = true
  branding_title = "test"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `domain` (String)

### Optional

- `attributes` (String) JSON format expected. Use jsonencode() to pass objects. Defaults to `{}`.
- `branding_favicon` (String)
- `branding_logo` (String)
- `branding_title` (String) Defaults to `authentik`.
- `default` (Boolean) Defaults to `false`.
- `flow_authentication` (String)
- `flow_device_code` (String)
- `flow_invalidation` (String)
- `flow_recovery` (String)
- `flow_unenrollment` (String)
- `flow_user_settings` (String)
- `web_certificate` (String)

### Read-Only

- `id` (String) The ID of this resource.