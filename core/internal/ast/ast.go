package ast

type Resource struct {
	Name       string
	Type       string
	Properties map[string]interface{}
	DependsOn  []string
}

type Program struct {
	CloudVendors map[string]*CloudVendor
	Variables    map[string]*Variable
	Resources    []*Resource
}

type CloudVendor struct {
	Name       string
	Properties map[string]interface{}
}

type Variable struct {
	Name  string
	Value interface{}
}

type Expression interface {
	Evaluate() interface{}
}

type StringLiteral struct {
	Value string
}

func (s *StringLiteral) Evaluate() interface{} {
	return s.Value
}

type NumberLiteral struct {
	Value float64
}

func (n *NumberLiteral) Evaluate() interface{} {
	return n.Value
}

type BooleanLiteral struct {
	Value bool
}

func (b *BooleanLiteral) Evaluate() interface{} {
	return b.Value
}

type ObjectLiteral struct {
	Properties map[string]Expression
}

func (o *ObjectLiteral) Evaluate() interface{} {
	result := make(map[string]interface{})
	for key, expr := range o.Properties {
		result[key] = expr.Evaluate()
	}
	return result
}

type ArrayLiteral struct {
	Elements []Expression
}

func (a *ArrayLiteral) Evaluate() interface{} {
	var result []interface{}
	for _, expr := range a.Elements {
		result = append(result, expr.Evaluate())
	}
	return result
}

type ForLoop struct {
	Iterator   string
	Collection Expression
	Body       []interface{}
}

type IdentifierExpression struct {
	Name string
}

func (i *IdentifierExpression) Evaluate() interface{} {
	return i.Name
}

type DataSource struct {
	Name       string
	Type       string
	Properties map[string]interface{}
	Result     map[string]interface{}
}

type Output struct {
	Name        string
	Value       interface{}
	Description string
}
