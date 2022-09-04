package mysql

import (
	"context"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/w3liu/bull-gateway/infra/mysql"
	"github.com/w3liu/bull-gateway/pkg/settings/options"
	"github.com/w3liu/bull-gateway/pkg/settings/types"
)

func newStore() *mysql.Store {
	conf := &mysql.Conf{}
	_, err := toml.DecodeFile(".test/mysql.toml", conf)
	if err != nil {
		panic(err)
	}
	return mysql.NewStore(conf)
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
		ReqHTTPMethod:    "GET",
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
