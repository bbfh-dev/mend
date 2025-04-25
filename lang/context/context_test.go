package context_test

import (
	"strings"
	"testing"

	"github.com/bbfh-dev/mend/lang/context"
	"github.com/stretchr/testify/assert"
)

func TestContextGetAndSet(test *testing.T) {
	assert := assert.New(test)

	ctx := context.New()
	assert.NotNil(ctx)

	ctx.Set([]string{"a", "b", "c"}, "Hello World!")
	ctx.Set([]string{"a", "b", "d"}, "Another one!")
	ctx.Set([]string{"e"}, "1")

	value, err := ctx.Query([]string{"a", "b", "c"})
	assert.NoError(err)
	assert.Equal("Hello World!", value)

	value, err = ctx.Query([]string{"a", "b", "d"})
	assert.NoError(err)
	assert.Equal("Another one!", value)

	value, err = ctx.Query([]string{"e"})
	assert.NoError(err)
	assert.Equal("1", value)

	value, err = ctx.Query([]string{"unknown"})
	assert.Error(err)
	assert.Equal("", value)
}

func TestContextExpressions(test *testing.T) {
	assert := assert.New(test)

	ctx := context.New()
	assert.NotNil(ctx)

	ctx.Set([]string{"a", "b", "c"}, "Hello World!")
	ctx.Set([]string{"a", "path"}, "/tmp/some/filename.html")

	var cases = []struct {
		Expression string
		ExpectErr  bool
		Expect     string
	}{
		{
			Expression: "this.a.b.c",
			ExpectErr:  false,
			Expect:     "Hello World!",
		},
		{
			Expression: "this.a.b.c == 'Hello World!'",
			ExpectErr:  false,
			Expect:     "true",
		},
		{
			Expression: "this.a.b.c != 'Hello World!'",
			ExpectErr:  false,
			Expect:     "false",
		},
		{
			Expression: "this.a.unknown == 'Hello World!'",
			ExpectErr:  true,
		},
		{
			Expression: "this.a.b.c has 'Hello'",
			ExpectErr:  false,
			Expect:     "true",
		},
		{
			Expression: "this.a.b.c lacks 'Hello'",
			ExpectErr:  false,
			Expect:     "false",
		},
		{
			Expression: "this.a.b.c || 'Fallback'",
			ExpectErr:  false,
			Expect:     "Hello World!",
		},
		{
			Expression: "this.a.unknown || 'Fallback'",
			ExpectErr:  false,
			Expect:     "Fallback",
		},
		{
			Expression: "this.a.b.c.to_lower()",
			ExpectErr:  false,
			Expect:     "hello world!",
		},
		{
			Expression: "this.a.b.c.to_upper()",
			ExpectErr:  false,
			Expect:     "HELLO WORLD!",
		},
		{
			Expression: "this.a.b.c.to_pascal_case()",
			ExpectErr:  false,
			Expect:     "HelloWorld",
		},
		{
			Expression: "this.a.b.c.to_camel_case()",
			ExpectErr:  false,
			Expect:     "helloWorld",
		},
		{
			Expression: "this.a.b.c.to_snake_case()",
			ExpectErr:  false,
			Expect:     "hello_world!",
		},
		{
			Expression: "this.a.b.c.to_kebab_case()",
			ExpectErr:  false,
			Expect:     "hello-world!",
		},
		{
			Expression: "this.a.b.c.to_lower().capitalize()",
			ExpectErr:  false,
			Expect:     "Hello world!",
		},
		{
			Expression: "this.a.b.c.length()",
			ExpectErr:  false,
			Expect:     "12",
		},
		{
			Expression: "this.a.b.c.length() == this.a.b.c.to_camel_case().length()",
			ExpectErr:  false,
			Expect:     "false",
		},
		{
			Expression: "this.a.path.extension()",
			ExpectErr:  false,
			Expect:     ".html",
		},
		{
			Expression: "this.a.path.filename().trim_extension()",
			ExpectErr:  false,
			Expect:     "filename",
		},
		{
			Expression: "this.a.b.unknown == 69",
			ExpectErr:  true,
		},
		{
			Expression: "this.a.b.unknown == 69 || true",
			ExpectErr:  false,
			Expect:     "true",
		},
		{
			Expression: "this.a.b.c has '!' || 'Fallback'",
			ExpectErr:  false,
			Expect:     "true",
		},
		{
			Expression: "",
			ExpectErr:  false,
			Expect:     "",
		},
	}

	for _, testCase := range cases {
		test.Run(
			strings.ReplaceAll(testCase.Expression, " ", "__"),
			func(test *testing.T) {
				result, err := ctx.Compute(testCase.Expression)
				if testCase.ExpectErr {
					assert.Error(err)
					return
				}
				assert.NoError(err)
				assert.Equal(testCase.Expect, result)
			},
		)
	}
}
