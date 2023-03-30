package golang

import (
	"context"
	"github.com/intelops/compage/core/internal/languages"
)

const GoContextVars = "GoContextVars"

type GoValues struct {
	Values      *languages.Values
	LGoLangNode *LGolangNode
}

func AddValuesToContext(ctx context.Context) context.Context {
	values := ctx.Value(languages.LanguageContextVars).(languages.Values)
	v := GoValues{
		Values: &values,
		LGoLangNode: &LGolangNode{
			LanguageNode: values.LanguageNode,
		},
	}

	return context.WithValue(ctx, GoContextVars, v)
}