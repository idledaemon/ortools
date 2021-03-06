package sat

import (
	"fmt"

	genSat "github.com/idledaemon/ortools/gen/sat"
)

type IntVar struct {
	modelProto *genSat.CpModelProto
	varIndex   int
	VarProto   *genSat.IntegerVariableProto
	negation   *notBoolVar
}

func NewIntVarLowerUpperBounds(modelProto *genSat.CpModelProto, lb int64, ub int64, name string) *IntVar {
	return NewIntVarBounds(modelProto, []int64{lb, ub}, name)
}

func NewIntVarBounds(modelProto *genSat.CpModelProto, bounds []int64, name string) *IntVar {
	varProto := genSat.IntegerVariableProto{
		Name:   name,
		Domain: bounds,
	}

	intVar := &IntVar{
		modelProto: modelProto,
		varIndex:   len(modelProto.GetVariables()),
		VarProto:   &varProto,
	}

	modelProto.Variables = append(modelProto.Variables, intVar.VarProto)

	return intVar
}

func (i *IntVar) Name() string {
	return i.VarProto.GetName()
}

func (i *IntVar) Index() int {
	return i.varIndex
}

func (i *IntVar) NumElements() int {
	return 1
}

func (i *IntVar) Variable(index int) *IntVar {
	return i
}

func (i *IntVar) Coefficient(index int) int {
	return 1
}

/** Returns the negation of a boolean variable. */
func (i *IntVar) Not() Literal {
	if i.negation == nil {
		i.negation = newNotBoolVar(i)
	}
	return i.negation
}

/** Returns a short string describing the variable. */
func (i *IntVar) ShortString() string {
	varProto := i.VarProto
	if len(varProto.GetName()) == 0 {
		domain := varProto.GetDomain()
		if len(domain) == 2 && domain[0] == domain[1] {
			return fmt.Sprintf("%d", domain[0])
		} else {
			return fmt.Sprintf("var_%d(%s)", i.Index(), i.DisplayBounds())
		}
	} else {
		return fmt.Sprintf("%s(%s)", i.Name(), i.DisplayBounds())
	}
}

/** Returns the domain as a string without the enclosing []. */
func (i *IntVar) DisplayBounds() string {
	out := ""
	varProto := i.VarProto
	domain := varProto.GetDomain()
	for n := 0; n < len(domain); n += 2 {
		if n != 0 {
			out += ", "
		}
		if domain[n] == domain[n+1] {
			out += fmt.Sprintf("%d", domain[n])
		} else {
			out += fmt.Sprintf("%d..%d", domain[n], domain[n+1])
		}
	}
	return out
}

func (i *IntVar) String() string {
	return i.modelProto.String()
}
