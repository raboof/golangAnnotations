package restAnnotation

import "github.com/MarcGrol/golangAnnotations/annotation"

const (
	typeRestOperation = "RestOperation"
	typeRestService   = "RestService"
	paramPath         = "path"
	paramMethod       = "method"
)

// Register makes the annotation-registry aware of these annotation
func Register() {
	annotation.RegisterAnnotation(typeRestOperation, []string{paramMethod, paramPath}, validateRestOperationAnnotation)
	annotation.RegisterAnnotation(typeRestService, []string{paramPath}, validateRestServiceAnnotation)
}

func validateRestOperationAnnotation(annot annotation.Annotation) bool {
	if annot.Name == typeRestOperation {
		path, hasPath := annot.Attributes[paramPath]
		method, hasMethod := annot.Attributes[paramMethod]
		return ((hasPath && path != "") && hasMethod && method != "")
	}
	return false
}

func validateRestServiceAnnotation(annot annotation.Annotation) bool {
	if annot.Name == typeRestService {
		_, ok := annot.Attributes[paramPath]
		return ok
	}
	return false
}
