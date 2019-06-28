package xormt

import (
	"strings"
	"fmt"
)

type ErrParamEmpty struct {
	name string
}

func (epe *ErrParamEmpty)Error()string{
	if strings.TrimSpace(epe.name) == ""{
		return "the param is empty or nil"
	}else{
		return fmt.Sprintf("the param %s is empty or nil", epe.name)
	}
}

type ErrFieldEmpty struct {
	field string
}

func (efe* ErrFieldEmpty)Error()string{
	if strings.TrimSpace(efe.field) == ""{
		return "the field is empty or nil"
	}else{
		return fmt.Sprintf("the field %s is empty or nil", efe.field)
	}
}

type ErrDeaultTendarMissing struct{

}

func (efe* ErrDeaultTendarMissing)Error()string {
	return "the default tendar is missing"
}
