package swagger

func NewRessourceListing(apiVersion string, apis []Api) (rl ResourceListing) {

    rl.ApiVersion = apiVersion
    rl.SwaggerVersion = swaggerVersion
    rl.Apis = apis
    return
}

func (rl *ResourceListing) AddApi(api Api) {

    apis := make([]Api, len(rl.Apis)+1)

    for i := range rl.Apis {
        apis[i] = rl.Apis[i]
    }
    apis[len(rl.Apis)] = api

    rl.Apis = apis
}
