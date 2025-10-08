# Validation System

The `validation` package handles validating structs of data associated with http request or response bodies. It uses a small number of field tags to do this. It also handles decoding request bodies, and encoding response bodies from and to JSON respectively.

Validation rules are defined on structs using field tags.

### Supported tags:
#### `required`
Ensures a field's value is not its zero-value

- Use a pointer field to denote a field that must be present in the JSON, but the zero-value of the corresponding Go type could be a valid value.

Ex. given a JSON string `'{"active": false}'`, if the `active` boolean field must be present, then the struct should be this:
```go
type Obj struct {
	Active *bool `json:"active" validate:"required"`
}
```
In this case, if `active` is not present in the JSON, it will be unmarshalled into a `nil` pointer, which would then not pass validation since the field is required. If we had not used a pointer, this would be unmarshalled into a `false` value, and we would lose the ability to detect if the field was originally present in the JSON.

> Note: a pointer field with no `validate:"required"` tag denotes an optional field

#### `oneOf=%s|%s,...`
Currently only implemented for string fields, checks if the value is one of the `%s` values
