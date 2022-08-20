package eval

// An Expr is an arithmetic expression.
// Expr 是一个算术表达式
type Expr interface {
	// Eval returns the value of this Expr in the environment env.
	// Eval 返回环境 env 中这个 Expr 的值。
	Eval(env Env) float64

	// Check reports errors in this Expr and adds its Vars to the set
	// 检查报告此 Expr 中的错误并将其 Vars 添加到集合中
	Check(vars map[Var]bool) error
}

// A Var identifies a variable, e.g., x.
// Var 标识一个变量，例如 x。
type Var string

// A literal is a numeric constant, e.g., 3.141.
// 文字是数字常量，例如 3.141。
type literal float64

// A unary represents a unary operator expression, e.g., -x.
// 一元表示一元运算符表达式，例如 -x。
type unary struct {
	op rune // one of '+','-'
	x  Expr
}

// A binary represents a binary operator expression, e.g., x+y.
// 二进制表示二元运算符表达式，例如 x+y。
type binary struct {
	op   rune // one of '+','-','*','/'
	x, y Expr
}

// A call represents a function call expression, e.g., sin(x).
// 调用表示函数调用表达式，例如 sin(x)。
type call struct {
	fn   string //one of "pow", "sin", "sqrt"
	args []Expr
}
