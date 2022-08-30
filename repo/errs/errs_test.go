package errs

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_Err ...
func Test_Err(t *testing.T) {
	e := NewFrameworkErr(1, "framework error")
	t.Logf("err: %v", e)

	e = New(2, "business error")
	t.Logf("err: %v", e)

	code := ErrCode(e)
	assert.Equal(t, 2, code)
	msg := ErrMsg(e)
	assert.Equal(t, "business error", msg)

	e = Newf(3, "num %d is wrong", 5)
	code = ErrCode(e)
	assert.Equal(t, 3, code)
	msg = ErrMsg(e)
	assert.Equal(t, "num 5 is wrong", msg)

	var er *Err = nil
	code = ErrCode(er)
	assert.Equal(t, 0, code)
	msg = ErrMsg(er)
	assert.Equal(t, ErrMsgSuccess, msg)

	var err error
	code = ErrCode(err)
	assert.Equal(t, 0, code)
	msg = ErrMsg(err)
	assert.Equal(t, ErrMsgSuccess, msg)

	err = errors.New("test")
	code = ErrCode(err)
	assert.Equal(t, ErrCodeUnknown, code)
	msg = ErrMsg(err)
	assert.Equal(t, "test", msg)
}
