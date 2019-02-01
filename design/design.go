package design

import (
        . "github.com/goadesign/goa/design"
        . "github.com/goadesign/goa/design/apidsl"
)

// API全体の定義
var _ = API("API Name", func() {
        Title("API Title")
        Version("0.1")
        Description("A simple goa service")
        Scheme("http")    // or https, both
        Host("localhost:8080")
})

// リソースの定義
// リソースのActionを定義していくのがメイン
var _ = Resource("bottle", func() {                // Resources group related API endpoints
        BasePath("/bottles")                       // together. They map to REST resources for REST
        DefaultMedia(BottleMedia)                  // services.

        Action("show", func() {                    // Actions define a single API endpoint together
                Description("Get bottle by id")    // with its path, parameters (both path
                Routing(GET("/:bottleID"))         // parameters and querystring values) and payload
                Params(func() {                    // (shape of the request body).
                        Param("bottleID", Integer, "Bottle ID") //リクエストパラメータ
                })
                Response(OK)                       // Responses define the shape and status code
                Response(NotFound)                 // of HTTP responses.
        })
        Action("create", func() {
                Description("Create new bottle")
                Routing(POST(""))
                Payload(CreateBottlePayload)
                Response(OK)
        })
})


// メディアタイプの定義
// メディアタイプはレスポンスの形式
var BottleMedia = MediaType("application/vnd.bottle+json", func() {  //慣習的に "application/vnd.<リソースの名前>+<形式>" が多いらしい
        Description("A bottle of wine")
        Attributes(func() {
                Attribute("id", Integer, "Unique bottle ID")  // id は整数型
                Attribute("href", String, "API href for making requests on the bottle") // href は文字列
                Attribute("name", String, "Name of wine") // name は文字列

                Required("id", "href", "name") // 上記のうちで必須なものをここに指定する
        })
        View("default", func() {                // default View は必須
                Attribute("id")
                Attribute("href")
                Attribute("name")
        })
})

// ペイロードの定義
// ペイロードはリクエストボディの形式（POST, PUT, PATCHを想定）
var CreateBottlePayload = Type("BottlePayload", func() {
        Member("bottleID", Integer, "Bottle ID", func(){
                Minimum(0)
                Maximum(127)
        })
        Member("category", String, "Category", func(){
                Enum("red", "whilte", "rose")
                Default("red")
        })
        Member("comment", String, "Comment", func(){
                MaxLength(256)
        })
        Required("bottleID", "category")
})
