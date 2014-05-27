package swagger

func NewOperation(operation string, nickname string, typ string, summary string, notes string) (op Operation) {

    op.HttpMethod = operation
    op.Nickname = nickname
    op.Summary = summary
    op.Notes = notes
    return
}

func (op *Operation) AddParameter(param Parameter) {

    params := make([]Parameter, len(op.Parameters)+1)

    for i := range op.Parameters {
        params[i] = op.Parameters[i]
    }
    params[len(op.Parameters)] = param

    op.Parameters = params
}
