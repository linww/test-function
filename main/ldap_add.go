package main

import (
	"context"
	"fmt"
	"github.com/aobco/log"
	"github.com/go-ldap/ldap/v3"
	"golang.org/x/sync/semaphore"
	"sync"
	"xsky.com/ocpf/pkg/usercenter/model"
)

func main() {
	ldapServer, err := model.NewLdapServer(&model.LdapConfigurationParam{
		Id:               100001,
		Name:             "ldap_27_1",
		Address:          "10.16.111.27:389",
		BaseDn:           "dc=xsky-test,dc=local",
		BindUserDn:       "cn=ldapadm,dc=xsky-test,dc=local",
		BindUserPassword: "f4f31942b9a7232694fb5124174f49cb",
		PropertyMapping: &model.LdapPropertyMapping{
			LoginName:   "uid",
			ObjectClass: []string{"organizationalPerson"},
		},
	})
	if err != nil {
		panic(err)
	}
	defer ldapServer.Close()
	wg := &sync.WaitGroup{}
	sem := semaphore.NewWeighted(10)
	for i := 0; i < 20100; i++ {
		if err = sem.Acquire(context.TODO(), 1); err != nil {
			panic(err)
		}
		wg.Add(1)
		go func(num int, wg *sync.WaitGroup, sem *semaphore.Weighted) {
			defer func() {
				wg.Done()
				sem.Release(1)
			}()
			name := fmt.Sprintf("test-%d", num)
			addRequest := &ldap.AddRequest{
				DN:         fmt.Sprintf("cn=%s,ou=测试2组,dc=xsky-test,dc=local", name),
				Attributes: nil,
				Controls:   nil,
			}
			attributes := make([]ldap.Attribute, 0)
			attributes = append(attributes, ldap.Attribute{
				Type: "objectClass",
				Vals: []string{"top", "inetOrgPerson"},
			})
			attributes = append(attributes, ldap.Attribute{
				Type: "cn",
				Vals: []string{name},
			})
			attributes = append(attributes, ldap.Attribute{
				Type: "sn",
				Vals: []string{name},
			})
			attributes = append(attributes, ldap.Attribute{
				Type: "uid",
				Vals: []string{name},
			})
			attributes = append(attributes, ldap.Attribute{
				Type: "userPassword",
				Vals: []string{"123456"},
			})
			addRequest.Attributes = attributes
			if err = ldapServer.Conn.Add(addRequest); err != nil {
				panic(err)
			}
		}(i, wg, sem)
	}
	wg.Wait()
	log.Info("done")
}
