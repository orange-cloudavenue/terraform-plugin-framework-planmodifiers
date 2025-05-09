---
hide:
    - navigation
---
# String Plan Modifiers

String plan modifiers are used to modify the plan of a string attribute.
It will be used into the `PlanModifiers` field of the `schema.StringAttribute` struct.

## How to use it

```go
import (
    fstringplanmodifier "github.com/orange-cloudavenue/terraform-plugin-framework-planmodifiers/stringplanmodifier"
)
```

## List of Plan Modifiers

### SetDefault

- [`SetDefault`](setdefault.md) - Sets a default value for the attribute.
- [`SetDefaultEnvVar`](setdefaultenvvar.md) - Sets a default value for the attribute from an environment variable.
- [`SetDefaultFunc`](setdefaultfunc.md) - Sets a default value for the attribute from a function.
- [`SetDefaultEmptyString`](setdefaultemptystring.md) - Sets a empty string as default value for the attribute.

### RequireReplace

- [`RequireReplaceIfBool`](requirereplaceifbool.md) - Forces the resource to be replaced when the specified boolean attribute is changed.

### StringChange

- [`ToLower`](tolower.md) - Converts the string to lowercase.
- [`ToUpper`](toupper.md) - Converts the string to uppercase.
