package mysqlx

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"reflect"
)

var db *gorm.DB

type Options struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

func New(viper *viper.Viper) (*gorm.DB, error) {
	var (
		err error
		o   = new(Options)
	)
	if err = viper.UnmarshalKey("db", o); err != nil {
		return nil, err
	}
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?"+
		"charset=utf8mb4&parseTime=True&loc=Local",
		o.User, o.Password,
		o.Host, o.Port,
		o.Name)
	db, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db, nil
}

// 可以看 https://github.com/win5do/go-microservice-demo/blob/main/docs/sections/gorm.md
type ctxTransactionKey struct{}

func CtxWithTransaction(ctx context.Context, tx *gorm.DB) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, ctxTransactionKey{}, tx)
}

func Transaction(ctx context.Context, fc func(txctx context.Context) error) error {
	db := db.WithContext(ctx)

	return db.Transaction(func(tx *gorm.DB) error {
		txctx := CtxWithTransaction(ctx, tx)
		return fc(txctx)
	})
}

// 如果使用跨模型事务则传参
func GetDB(ctx context.Context) *gorm.DB {
	iface := ctx.Value(ctxTransactionKey{})

	if iface != nil {
		tx, ok := iface.(*gorm.DB)
		if !ok {
			log.Panicf("unexpect context value type: %s", reflect.TypeOf(tx))
			return nil
		}

		return tx
	}

	return db.WithContext(ctx)
}