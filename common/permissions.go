package c

import (
	"net/url"
	"strconv"
)

// 用於高校能權限比對設定
type Permission uint8
type CanPermission uint8

const PERM_CAN_GET CanPermission = 0x01
const PERM_CAN_LIST CanPermission = 0x02
const PERM_CAN_CREATE CanPermission = 0x04
const PERM_CAN_UPDATE CanPermission = 0x08
const PERM_CAN_DELETE CanPermission = 0x10
const PERM_CAN_OTHER1 CanPermission = 0x20
const PERM_CAN_OTHER2 CanPermission = 0x40
const PERM_CAN_ROOT CanPermission = 0x80

// 有隨便一個權限都好
const PERM_ANY Permission = Permission(PERM_CAN_GET | PERM_CAN_LIST | PERM_CAN_CREATE | PERM_CAN_UPDATE | PERM_CAN_DELETE | PERM_CAN_OTHER1 | PERM_CAN_OTHER2)
const PERM_ROOT_ALL Permission = Permission(PERM_CAN_ROOT) | PERM_ANY
const PERM_NONE Permission = 0x00

type PermissionInHuman struct {
	Get    bool
	List   bool
	Create bool
	Update bool
	Delete bool
	Other1 bool
	Other2 bool
	Root   bool
}

func NewPermission(canPermissions ...CanPermission) Permission {
	var p CanPermission
	for _, pp := range canPermissions {
		p = p | pp
	}
	return Permission(p)
}

func (p *Permission) ToPermissionInHuman() *PermissionInHuman {
	return &PermissionInHuman{
		Get:    p.Can(PERM_CAN_GET),
		List:   p.Can(PERM_CAN_LIST),
		Create: p.Can(PERM_CAN_CREATE),
		Update: p.Can(PERM_CAN_UPDATE),
		Delete: p.Can(PERM_CAN_DELETE),
		Other1: p.Can(PERM_CAN_OTHER1),
		Other2: p.Can(PERM_CAN_OTHER2),
		Root:   p.Can(PERM_CAN_ROOT),
	}
}

// 給入多個表示AND，例如 p.Can(PERM_CAN_CREATE, PERM_CAN_GET) 代表要同時擁有Create跟Get權限才會回傳true
func (p *Permission) Can(canPermissions ...CanPermission) bool {
	cp := CanPermission(*p)
	if cp&PERM_CAN_ROOT > 0 {
		return true
	}
	if len(canPermissions) == 0 {
		return false
	}
	var ok bool = true
	for _, perm := range canPermissions {
		ok = ok && cp&perm > 0
	}
	return ok
}

/*
- managerPermission: 修改者的權限

- source: 被修改者的權限

- target: 希望修改後的權限

- 如果什麼都不會改變, 則回傳nil
*/
func (managerPermission *Permission) UpdateConstraint(source, target Permission) *Permission {
	// root can do anything
	if managerPermission.Can(PERM_CAN_ROOT) {
		return &target
	}
	// nobody can set ROOT
	if target.Can(PERM_CAN_ROOT) {
		target -= Permission(PERM_CAN_ROOT)
	}
	wannaChanges := source ^ target
	canChanges := wannaChanges & *managerPermission
	if canChanges == 0 {
		return nil
	}
	unChanges := PERM_ANY - canChanges
	result := wannaChanges&target + unChanges&source
	return &result
}

func (c *PermissionInHuman) ToPermission() Permission {
	var canPermissions []CanPermission
	if c.Get {
		canPermissions = append(canPermissions, PERM_CAN_GET)
	}
	if c.List {
		canPermissions = append(canPermissions, PERM_CAN_LIST)
	}
	if c.Create {
		canPermissions = append(canPermissions, PERM_CAN_CREATE)
	}
	if c.Update {
		canPermissions = append(canPermissions, PERM_CAN_UPDATE)
	}
	if c.Delete {
		canPermissions = append(canPermissions, PERM_CAN_DELETE)
	}
	if c.Other1 {
		canPermissions = append(canPermissions, PERM_CAN_OTHER1)
	}
	if c.Other2 {
		canPermissions = append(canPermissions, PERM_CAN_OTHER2)
	}
	if c.Root {
		canPermissions = append(canPermissions, PERM_CAN_ROOT)
	}
	return NewPermission(canPermissions...)
}

func (p *Permission) UnmarshalJSON(bytes []byte) error {
	str := string(bytes)
	i, err := strconv.Atoi(str)
	if err != nil {
		return err
	}
	*p = Permission(i)
	return nil
}

func (p *Permission) MarshalJSON() ([]byte, error) {
	str := strconv.Itoa(int(*p))
	return []byte(str), nil
}

func (p Permission) EncodeValues(key string, v *url.Values) error {
	str := strconv.Itoa(int(p))
	v.Add(key, str)
	return nil
}
