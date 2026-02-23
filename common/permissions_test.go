package c_test

import (
	"fmt"
	"testing"

	c "github.com/mychiux413/goutils/common"
	"github.com/stretchr/testify/assert"
)

func TestPermissions(t *testing.T) {
	assert := assert.New(t)
	ph := c.PermissionInHuman{Get: true}
	p := ph.ToPermission()
	assert.True(p.Can(c.PERM_CAN_GET))
	assert.False(p.Can(c.PERM_CAN_VIEW))
	assert.False(p.Can(c.PERM_CAN_UPDATE))
	assert.False(p.Can(c.PERM_CAN_DELETE))
	assert.False(p.Can(c.PERM_CAN_OTHER1))
	assert.False(p.Can(c.PERM_CAN_OTHER2))
	assert.False(p.Can(c.PERM_CAN_ROOT))

	p = c.NewPermission(c.PERM_CAN_GET)
	assert.True(p.Can(c.PERM_CAN_GET))

	ph = c.PermissionInHuman{Other2: true}
	p = ph.ToPermission()
	assert.True(p.Can(c.PERM_CAN_OTHER2))
	assert.False(p.Can(c.PERM_CAN_OTHER1))
	assert.False(p.Can(c.PERM_CAN_ROOT))
}

func BenchmarkPermissions(b *testing.B) {
	/*
		goos: linux
		goarch: amd64
		cpu: AMD Ryzen 9 5950X 16-Core Processor
		BenchmarkPermissions-32    	  248307	      4318 ns/op
	*/
	ph := c.PermissionInHuman{Get: true}
	p := ph.ToPermission()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			p.Can(c.PERM_CAN_GET)
			p.Can(c.PERM_CAN_VIEW)
			p.Can(c.PERM_CAN_CREATE, c.PERM_CAN_GET)
			p.Can(c.PERM_CAN_UPDATE, c.PERM_CAN_GET)
			p.Can(c.PERM_CAN_DELETE, c.PERM_CAN_GET)
			p.Can(c.PERM_CAN_OTHER1)
			p.Can(c.PERM_CAN_OTHER2)
			p.Can(c.PERM_CAN_ROOT)
		}

	}
}

func TestUpdatePermissionConstraint(t *testing.T) {
	assert := assert.New(t)
	manager := c.NewPermission(c.PERM_CAN_CREATE, c.PERM_CAN_VIEW, c.PERM_CAN_UPDATE)
	source := c.NewPermission(c.PERM_CAN_CREATE, c.PERM_CAN_VIEW, c.PERM_CAN_UPDATE, c.PERM_CAN_DELETE)
	target := c.PERM_NONE
	expected := c.NewPermission(c.PERM_CAN_DELETE)

	result := manager.UpdateConstraint(source, target)
	assert.Equal(expected, *result, fmt.Sprintf("expected: %#v\nresult: %#v", expected.ToPermissionInHuman(), result.ToPermissionInHuman()))

	target = source + c.NewPermission(c.PERM_CAN_OTHER1)
	result = manager.UpdateConstraint(source, target)
	assert.Nil(result, "manager沒有權限不能更新")

	target = source - c.NewPermission(c.PERM_CAN_DELETE)
	result = manager.UpdateConstraint(source, target)
	assert.Nil(result, "manager沒有權限不能更新")

	manager = c.PERM_ROOT_ALL
	target = c.NewPermission(c.PERM_CAN_ROOT)
	result = manager.UpdateConstraint(source, target)
	assert.False(result.Can(c.PERM_CAN_ROOT), "沒有人能給source ROOT權限")
}
