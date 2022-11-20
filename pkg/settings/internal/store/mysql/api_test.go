package mysql

import (
	"context"
	"github.com/w3liu/bull-gateway/config"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/w3liu/bull-gateway/infra/mysql"
	"github.com/w3liu/bull-gateway/pkg/settings/options"
	"github.com/w3liu/bull-gateway/pkg/settings/types"
)

func newStore() *mysql.Store {
	conf := &config.Config{}
	_, err := toml.DecodeFile("local/config.toml", conf)
	if err != nil {
		panic(err)
	}
	return mysql.NewStore(&conf.Mysql)
}

func TestApiStore_Create(t *testing.T) {
	store := newStore()
	defer store.Close()
	as := newApiStore(store)
	api := &types.Api{
		GroupId:          1,
		Name:             "test01",
		Description:      "desc",
		Status:           0,
		ReqPath:          "/user/add",
		ReqHttpMethod:    "GET",
		InputRequestMode: "",
	}
	err := as.Create(context.TODO(), api, options.CreateOptions{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestApiStore_Update(t *testing.T) {

}

func TestApiStore_Delete(t *testing.T) {

}

func TestApiStore_Page(t *testing.T) {

}
