package routes

import (
    "github.com/gin-gonic/gin"
    "os"
    "GoMVC/config"
    "GoMVC/core"
)

func RouterInstance(routes *gin.Engine) *gin.Engine {
    routes.GET("/", func(context *gin.Context) {
        var data = make(map[string]string)
        data["version"] = os.Getenv("app_version")
        core.Output(context, config.SUCCESS, data, make(map[string]string))
    })

    //----- Swagger -----
    //swaggerUrl := ginSwagger.URL("http://localhost/swagger/doc.json")
    //routes.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, swaggerUrl))
    return routes
}
