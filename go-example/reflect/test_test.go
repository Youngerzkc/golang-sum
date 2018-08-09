package reflect

import (
	"testing"
	"fmt"
)
type UpdateRequest struct {
	Id           string   `protobuf:"bytes,1,opt,name=id" json:"id" xml:"id"`
	Name         string   `protobuf:"bytes,2,opt,name=name" json:"name" xml:"name"`
	SubjectId    string   `protobuf:"bytes,3,opt,name=subject_id,json=subjectId" json:"subject_id" xml:"subject_id"`
	UpdateFields []string `protobuf:"bytes,4,rep,name=update_fields,json=updateFields" json:"update_fields" xml:"update_fields"`
}
var req =&UpdateRequest{
	Id:"123456",
	Name:"younger",
	SubjectId:"12345678",
	UpdateFields:[]string{"id","subject_id"},
}


func TestExtractUpdates(t *testing.T){
	 	result:=ExtractUpdates(*req,req.UpdateFields)
		fmt.Println("result is ",result)	
	}
