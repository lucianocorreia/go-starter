package flash

import (
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/lucianocorreia/go-starter/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFlash(t *testing.T) {
	e := echo.New()
	ctx, _ := tests.NewContext(e, "/")
	tests.InitSession(ctx)

	assertFlash := func(typ FlashType, message string) {
		ret := Get(ctx, typ)
		require.Len(t, ret, 1)
		assert.Equal(t, message, ret[0])
		ret = Get(ctx, typ)
		require.Len(t, ret, 0)
	}

	text := "aaa"
	Success(ctx, text)
	assertFlash(TypeSuccess, text)

	text = "bbb"
	Info(ctx, text)
	assertFlash(TypeInfo, text)

	text = "ccc"
	Danger(ctx, text)
	assertFlash(TypeDanger, text)

	text = "ddd"
	Warning(ctx, text)
	assertFlash(TypeWarning, text)

	text = "eee"
	Set(ctx, TypeSuccess, text)
	assertFlash(TypeSuccess, text)
}
