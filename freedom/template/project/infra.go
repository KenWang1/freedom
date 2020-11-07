package project

import "fmt"

func init() {
	content["/infra/response.go"] = jsonResponseTemplate()
	content["/infra/request.go"] = jsonRequestTemplate()
}

func jsonRequestTemplate() string {
	return `
	//Package infra generated by 'freedom new-project {{.PackagePath}}'
	package infra

	import (
		"io/ioutil"
	
		"encoding/json"
		"github.com/8treenet/freedom"
		"gopkg.in/go-playground/validator.v9"
	)
	
	var validate *validator.Validate
	func init() {
		validate = validator.New()
		freedom.Prepare(func(initiator freedom.Initiator) {
			initiator.BindInfra(false, func() *Request {
				return &Request{}
			})
			initiator.InjectController(func(ctx freedom.Context) (com *Request) {
				initiator.GetInfra(ctx, &com)
				return
			})
		})
	}
	
	// Request .
	type Request struct {
		freedom.Infra
	}
	
	// BeginRequest .
	func (req *Request) BeginRequest(worker freedom.Worker) {
		req.Infra.BeginRequest(worker)
	}
	
	// ReadJSON .
	func (req *Request) ReadJSON(obj interface{}) error {
		rawData, err := ioutil.ReadAll(req.Worker.IrisContext().Request().Body)
		if err != nil {
			return err
		}
		err = json.Unmarshal(rawData, obj)
		if err != nil {
			return err
		}
	
		return validate.Struct(obj)
	}
	
	// ReadQuery .
	func (req *Request) ReadQuery(obj interface{}) error {
		if err := req.Worker.IrisContext().ReadQuery(obj); err != nil {
			return err
		}
		return validate.Struct(obj)
	}
	
	// ReadForm .
	func (req *Request) ReadForm(obj interface{}) error {
		if err := req.Worker.IrisContext().ReadForm(obj); err != nil {
			return err
		}
		return validate.Struct(obj)
	}	
	`
}
func jsonResponseTemplate() string {
	result := `
	//Package infra generated by 'freedom new-project {{.PackagePath}}'
	package infra

import (
	"strconv"

	"encoding/json"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/hero"
)

// JSONResponse .
type JSONResponse struct {
	Code        int
	Error       error
	contentType string
	content     []byte
	Object      interface{}
}

// Dispatch This is the middleware for HTTP output.
func (jrep JSONResponse) Dispatch(ctx context.Context) {
	jrep.contentType = "application/json"
	var repData struct {
		Code  int          %sjson:"code"%s
		Error string       %sjson:"error"%s
		Data  interface{}  %sjson:"data,omitempty"%s
	}

	repData.Data = jrep.Object
	repData.Code = jrep.Code
	if jrep.Error != nil {
		repData.Error = jrep.Error.Error()
	}
	if repData.Error != "" && repData.Code == 0 {
		repData.Code = 400
	}
	ctx.Values().Set("code", strconv.Itoa(repData.Code))

	jrep.content, _ = json.Marshal(repData)
	ctx.Values().Set("response", string(jrep.content))
	hero.DispatchCommon(ctx, 0, jrep.contentType, jrep.content, nil, nil, true)
}

`
	return fmt.Sprintf(result, "`", "`", "`", "`", "`", "`")
}
