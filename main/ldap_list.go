package main

import (
	"fmt"
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
	users, err := ldapServer.ListUsers()
	if err != nil {
		panic(err)
	}
	fmt.Println(len(users))
}
