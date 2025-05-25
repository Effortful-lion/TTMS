package middleware


// 检查请求参数是否合法并打印日志
// TODO 只是先写出来，后续需要使用的话，需要对每一个handler进行修改，参数从context中获取

import (
	"log"
	"reflect"

	"github.com/gin-gonic/gin"
)

// 检查请求参数是否合法并打印日志
// TODO 只是先写出来，后续需要使用的话，需要对每一个handler进行修改，参数从context中获取
func CheckRequestParams(c *gin.Context, req interface{}) bool {
	// 根据请求方法和 Content-Type 获取参数
	switch c.Request.Method {
	case "GET":
		// 处理查询参数
		if err := c.ShouldBindQuery(req); err != nil {
			log.Printf("Failed to bind query parameters: %v", err)
			return false
		}
	case "POST", "PUT", "PATCH":
		contentType := c.GetHeader("Content-Type")
		switch {
		case contentType == "application/json":
			// 处理 JSON 数据
			if err := c.ShouldBindJSON(req); err != nil {
				log.Printf("Failed to bind JSON data: %v", err)
				return false
			}
		case contentType == "application/x-www-form-urlencoded":
			// 处理表单数据
			if err := c.ShouldBind(req); err != nil {
				log.Printf("Failed to bind form data: %v", err)
				return false
			}
		default:
			log.Println("Unsupported Content-Type")
			return false
		}
	default:
		log.Println("Unsupported HTTP method")
		return false
	}

	// 使用反射验证参数合法性
	if !validateRequest(req) {
		log.Println("Request parameters are invalid")
		return false
	}

	// 打印参数日志
	log.Printf("Request parameters: %+v", req)
	return true
}

// 使用反射验证请求参数合法性
func validateRequest(req interface{}) bool {
	value := reflect.ValueOf(req)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	if value.Kind() != reflect.Struct {
		return false
	}

	// 简单示例：检查结构体字段是否为空
	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		if isZeroValue(field) {
			return false
		}
	}
	return true
}

// 判断值是否为空
func isZeroValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Complex64, reflect.Complex128:
		return v.Complex() == 0
	case reflect.String:
		return v.String() == ""
	case reflect.Ptr, reflect.Interface, reflect.Slice, reflect.Map, reflect.Chan, reflect.Func:
		return v.IsNil()
	case reflect.Struct:
		// 递归检查结构体字段
		for i := 0; i < v.NumField(); i++ {
			if isZeroValue(v.Field(i)) {
				return true
			}
		}
		return false
	default:
		return false
	}
}