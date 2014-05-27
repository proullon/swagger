package swagger

func NewApiDeclaration(version string, basePath string, ressourcePath string) (decl ApiDeclaration) {

    decl.ApiVersion = version
    decl.SwaggerVersion = swaggerVersion
    decl.BasePath = basePath
    decl.ResourcePath = ressourcePath

    decl.Produces = make([]string, 1)
    decl.Produces[0] = "application/json"
    return
}

func (decl *ApiDeclaration) AddApi(api Api) {

    apis := make([]Api, len(decl.Apis)+1)

    for i := range decl.Apis {
        apis[i] = decl.Apis[i]
    }
    apis[len(decl.Apis)] = api

    decl.Apis = apis
}
