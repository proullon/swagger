package swagger

func NewParameter(name string, dataType string, description string, required bool, format string) (param Parameter) {

    param.ParamType = "query"
    param.DataType = dataType
    param.Description = description
    param.Name = name
    param.Required = required
    param.Format = format

    return
}
