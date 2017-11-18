package main

import (
	"bytes"
	"fasthttp/model"
	"fasthttp/template"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/buaazp/fasthttprouter"
	"github.com/revenuemonster/rm-api/db"
	"github.com/valyala/fasthttp"
)

const (
	dbMerchant                    = "Merchant"
	dbMerchantProductSubscription = "MerchantProductSubscription"
)

// Index : index page
func Index(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Hello, world!\n\n")

	fmt.Fprintf(ctx, "Request method is %q\n", ctx.Method())
	fmt.Fprintf(ctx, "RequestURI is %q\n", ctx.RequestURI())
	fmt.Fprintf(ctx, "Requested path is %q\n", ctx.Path())
	fmt.Fprintf(ctx, "Host is %q\n", ctx.Host())
	fmt.Fprintf(ctx, "Query string is %q\n", ctx.QueryArgs())
	fmt.Fprintf(ctx, "User-Agent is %q\n", ctx.UserAgent())
	fmt.Fprintf(ctx, "Connection has been established at %s\n", ctx.ConnTime())
	fmt.Fprintf(ctx, "Request has been started at %s\n", ctx.Time())
	fmt.Fprintf(ctx, "Serial request number for the current connection is %d\n", ctx.ConnRequestNum())
	fmt.Fprintf(ctx, "Your ip is %q\n\n", ctx.RemoteIP())

	fmt.Fprintf(ctx, "Raw request is:\n---CUT---\n%s\n---CUT---", &ctx.Request)

	ctx.SetContentType("text/plain; charset=utf8")

	// Set arbitrary headers
	ctx.Response.Header.Set("X-My-Header", "my-header-value")

	// Set cookies
	var c fasthttp.Cookie
	c.SetKey("cookie-name")
	c.SetValue("cookie-value")
	ctx.Response.Header.SetCookie(&c)
}

// GetMerchant : retrieve merchant information
func GetMerchant(ctx *fasthttp.RequestCtx) {
	id := ctx.UserValue("id").(string)
	fmt.Println(id)
	if _, err := strconv.Atoi(id); err != nil {
		buffer := new(bytes.Buffer)
		template.Error(buffer)
		ctx.SetContentType("text/html; charset=utf8")
		ctx.SetStatusCode(500)
		// // ctx.Response.SetBody(buffer.Bytes())
		if _, err := ctx.Write(buffer.Bytes()); err != nil {
			fmt.Println(err)
			// return err
		}
		// return err
		return
	}

	// merchantKey, err := helper.DecodeIDToKey(dbMerchant, id, nil)
	// if err != nil {
	// 	fmt.Fprintln(ctx, "Bad Request")
	// 	return
	// }

	// loyaltyProductSubscription := new(model.MerchantProductSubscription)
	// if err := db.Kind(dbMerchantProductSubscription).Where("ProductSubscriptionID =", constant.LoyaltyVoucherFeature).Ancestor(merchantKey).First(loyaltyProductSubscription); err != nil {
	// 	fmt.Fprintln(ctx, "Not Found")
	// 	return
	// }

	// socialMediaProductSubscription := new(model.MerchantProductSubscription)
	// if err := db.Kind(dbMerchantProductSubscription).Where("ProductSubscriptionID =", constant.SocialNetworkFeature).Ancestor(merchantKey).First(socialMediaProductSubscription); err != nil {
	// 	fmt.Fprintln(ctx, "Internal server error")
	// 	return
	// }

	// merchant := new(model.Merchant)
	// if err := db.Kind(dbMerchant).Find(merchantKey, merchant); err != nil {
	// 	fmt.Fprintln(ctx, "Not Found")
	// 	return
	// }

	merchant := &model.Merchant{
		Name:    "RevenueMonster",
		LogoURL: "",
	}

	buffer := new(bytes.Buffer)
	template.Main(merchant, buffer)
	ctx.SetContentType("text/html; charset=utf8")
	ctx.SetStatusCode(200)
	// // ctx.Response.SetBody(buffer.Bytes())
	if _, err := ctx.Write(buffer.Bytes()); err != nil {
		fmt.Println(err)
		// return err
	}

	return // nil
}

func main() {
	// pending
	// override NotAllow handler
	// override NotFound handler
	// write response struct
	// change webpack entry point to template/master.html (prevent injection)
	projectID := os.Getenv("DATASTORE_PROJECT_ID")
	log.Println(projectID)
	if err := db.Config(projectID); err != nil {
		os.Stderr.WriteString(err.Error())
		panic(err.Error())
	}

	router := fasthttprouter.New()
	router.MethodNotAllowed = func(ctx *fasthttp.RequestCtx) {
		fmt.Println("Not allow")
	}
	router.NotFound = func(ctx *fasthttp.RequestCtx) {
		fmt.Println("Not found")
	}
	router.GET("/", Index)
	router.GET("/:id", GetMerchant)
	log.Fatal(fasthttp.ListenAndServe(":8080", router.Handler))
}
