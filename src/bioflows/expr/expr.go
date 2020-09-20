package expr

import "github.com/hoisie/mustache"

type ExprManager struct{

}

func (expr *ExprManager) Render(data string, context ...interface{}) string {
	return mustache.Render(data,context...)
}


