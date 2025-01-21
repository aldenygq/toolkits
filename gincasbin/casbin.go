package gincasbin

import (
    gormadapter "github.com/casbin/gorm-adapter/v3"
    casbin "github.com/casbin/casbin/v3"
    "gorm.io/gorm"
)

type Casbin struct {
    Cas *casbin.Enforcer
}

func NewCasbin(tableName,rbacFile string,cas interface{},db *gorm.DB) (*Casbin,error)  {
    // 初始化CasBin适配器
	adapter, err := gormadapter.NewAdapterByDBWithCustomTable(db, &cas, tableName)
    if err != nil {
        return nil,err
    }
    // 初始化CasBin执行器
	c, err := casbin.NewEnforcer(rbacFile, adapter)
	if err != nil {
		return nil,err
	}
    err = c.LoadPolicy()
	if err != nil {
		return nil,err
	}
    return &Casbin{
        Cas:c,
    },nil
}
//校验角色继承规则是否存在
func (c *Casbin) CheckRoleRuleExist(params ...interface{}) bool {
    return c.Cas.HasGroupingPolicy(params...)
}
//添加角色继承规则,先校验，再添加
func (c *Casbin) AddRoleRule(params ...interface{}) error {
     if !c.CheckRoleRuleExist(params...) {
         _, err := c.Cas.AddGroupingPolicy(params...)
         if err != nil {
             return err
         }
     }
     err := c.Cas.SavePolicy()
     if err != nil {
        return err
     }
     return nil
}
//批量添加角色继承规则
func (c *Casbin) AddRoleRules(rules [][]string) error {
    // 批量添加策略规则
    _, err := c.Cas.AddGroupingPolicies(rules)
    if err != nil {
        return err
    }
    err = c.Cas.SavePolicy()
    if err != nil {
       return err
    }
    return nil
}
//校验策略规则是否存在
func (c *Casbin) CheckPolicyExist(params ...interface{}) bool {
    return c.Cas.HasPolicy(params...)
}
//添加规则，先校验，再添加
func (c *Casbin) AddPolicyRule(params ...interface{}) error {
    if !c.CheckPolicyExist(params...) {
        _, err := c.Cas.AddPolicy(params...)
        if err != nil {
            return err
        }
    }
    err := c.Cas.SavePolicy()
    if err != nil {
       return err
    }
    return nil
}
//批量添加策略规则
func (c *Casbin) AddPolicyRules(rules [][]string) error {
    // 批量添加策略规则
    _, err := c.Cas.AddPolicies(rules)
    if err != nil {
        return err
    }
    err = c.Cas.SavePolicy()
    if err != nil {
       return err
    }
    return nil
}

//删除角色继承策略
func (c *Casbin) DeleteRoleRule(params ...interface{}) error {
    // 批量添加策略规则
    _, err := c.Cas.RemoveGroupingPolicy(params...)
    if err != nil {
        return err
    }
    err = c.Cas.SavePolicy()
    if err != nil {
       return err
    }
    return nil
}
//批量删除继承策略
func (c *Casbin)  DeleteRoleRules(rules [][]string) error {
    // 批量添加策略规则
    _, err := c.Cas.RemoveGroupingPolicies(rules)
    if err != nil {
        return err
    }
    err = c.Cas.SavePolicy()
    if err != nil {
       return err
    }
    return nil
}

//删除权限策略
func (c *Casbin)  DeletePolicyRule(params ...interface{}) error {
    // 批量添加策略规则
    _, err := c.Cas.RemovePolicy(params...)
    if err != nil {
        return err
    }
    err = c.Cas.SavePolicy()
    if err != nil {
       return err
    }
    return nil
}

//批量删除权限策略
func (c *Casbin)  DeletePolicyRules(rules [][]string) error {
    // 批量添加策略规则
    _, err := c.Cas.RemovePolicies(rules)
    if err != nil {
        return err
    }
    err = c.Cas.SavePolicy()
    if err != nil {
       return err
    }
    return nil
}
