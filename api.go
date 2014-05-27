package swagger

func NewApi(path string, description string) (api Api) {
    api.Path = path
    api.Description = description

    return
}

func (api *Api) AddOperation(op Operation) {

    operations := make([]Operation, len(api.Operations)+1)

    for i := range api.Operations {
        operations[i] = api.Operations[i]
    }
    operations[len(api.Operations)] = op

    api.Operations = operations
}
