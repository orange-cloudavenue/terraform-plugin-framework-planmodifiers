---
hide:
    - navigation
---
# Int32 Plan Modifiers

Int32 plan modifiers are used to modify the plan of a int32 attribute.
I will be used into the `PlanModifiers` field of the `schema.Int32Attribute` struct.

## How to use it

```go
import (
    fint32planmodifier "github.com/orange-cloudavenue/terraform-plugin-framework-planmodifiers/int32planmodifier"
)
```

## List of Plan Modifiers

- [`SetDefault`](setdefault.md) - Sets a default value for the attribute.
- [`SetDefaultEnvVar`](setdefaultenvvar.md) - Sets a default value for the attribute from an environment variable.
- [`SetDefaultFunc`](setdefaultfunc.md) - Sets a default value for the attribute from a function.

### RequireReplace

- [`RequireReplaceIfBool`](requirereplaceifbool.md) - Forces the resource to be replaced when the specified boolean attribute is changed.
