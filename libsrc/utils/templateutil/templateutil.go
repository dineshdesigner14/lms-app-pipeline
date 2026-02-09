package templateutil

import (
	"bytes"
	"html/template"
	"lmsapieng/include/common/globaldef"
)

func GetTemplateString(templateDataMap map[string]interface{}, TemplateStr string) string {
	//trace.Lg("GetTemplateStr() called with TemplateStr[%s]", TemplateStr)
	textTemplate := template.New("Template")
	textTemplate, _ = textTemplate.Parse(TemplateStr)
	doc := &bytes.Buffer{}
	err := textTemplate.Execute(doc, &templateDataMap)
	if err != nil {
		//trace.Lg("textTemplate.Execute() failed with err(%s)", err)
		return ""
	}
	return doc.String()
}

func GetTemplateStr(templateDataMap map[string]interface{}, TemplateStrArgs ...string) []byte {
	TemplateStr := ""
	if len(TemplateStrArgs) == 0 {
		return []byte(globaldef.NOT_INITIALIZED)
	}
	if len(TemplateStrArgs) > 2 {
		if len(TemplateStrArgs[2]) != 0 {
			TemplateStr = TemplateStrArgs[2]
		} else if len(TemplateStrArgs[1]) != 0 {
			TemplateStr = TemplateStrArgs[1]
		} else {
			TemplateStr = TemplateStrArgs[0]
		}
	} else if len(TemplateStrArgs) > 1 {
		if len(TemplateStrArgs[1]) != 0 {
			TemplateStr = TemplateStrArgs[1]
		} else {
			TemplateStr = TemplateStrArgs[0]
		}
	} else {
		if len(TemplateStrArgs[0]) != 0 {
			TemplateStr = TemplateStrArgs[0]
		}
	}
	if len(TemplateStr) == 0 {
		return []byte(globaldef.NOT_INITIALIZED)
	}
	textTemplate := template.New("Template")
	textTemplate, _ = textTemplate.Parse(TemplateStr)
	doc := &bytes.Buffer{}
	err := textTemplate.Execute(doc, &templateDataMap)
	if err != nil {
		//trace.LgReq(reqbrokerutil.GetReqID(templateDataMap), "textTemplate.Execute() failed with err(%s)", err)
		return []byte(globaldef.NOT_INITIALIZED)
	}
	return doc.Bytes()
}
