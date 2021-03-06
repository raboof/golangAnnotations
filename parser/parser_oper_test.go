package parser

import (
	"testing"

	"github.com/MarcGrol/golangAnnotations/model"
	"github.com/stretchr/testify/assert"
)

func TestStructOperationsInDir(t *testing.T) {
	harvest, err := ParseSourceDir("./operations", ".*")
	assert.Equal(t, nil, err)
	assert.Equal(t, 2, len(harvest.Operations))

	{
		o := harvest.Operations[0]
		assert.Equal(t, "operations", o.PackageName)
		assert.Equal(t, []string{"// docline for getPersons"}, o.DocLines)
		assert.Equal(t, "getPersons", o.Name)
		assertField(t, model.Field{Name: "serv", TypeName: "Service", IsPointer: true}, *o.RelatedStruct)

		assert.Equal(t, 0, len(o.InputArgs))

		assert.Equal(t, 2, len(o.OutputArgs))
		assertField(t, model.Field{TypeName: "Person", IsSlice: true}, o.OutputArgs[0])
		assertField(t, model.Field{TypeName: "error"}, o.OutputArgs[1])
	}
	{
		o := harvest.Operations[1]
		assert.Equal(t, "operations", o.PackageName)
		assert.Equal(t, []string{`// docline for getPerson`}, o.DocLines)
		assert.Equal(t, "getPerson", o.Name)
		assertField(t, model.Field{Name: "s", TypeName: "Service"}, *o.RelatedStruct)

		assert.Equal(t, 1, len(o.InputArgs))
		assertField(t, model.Field{Name: "uid", TypeName: "string"}, o.InputArgs[0])

		assert.Equal(t, 3, len(o.OutputArgs))
		assertField(t, model.Field{TypeName: "Person"}, o.OutputArgs[0])
		assertField(t, model.Field{TypeName: "Person", IsPointer: true}, o.OutputArgs[1])
		assertField(t, model.Field{TypeName: "error"}, o.OutputArgs[2])
	}
}
