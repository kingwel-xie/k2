package utils

import (
	"encoding/json"
	"github.com/kingwel-xie/k2/core/storage"

	"github.com/chanxuehong/wechat/oauth2"
)

type WxCache struct {
	prefix          string
	wxTokenStoreKey string
	store           storage.AdapterCache
}

const (
	intervalTenant = "/"
)

// Token 获取微信oauth2 token
func (e WxCache) Token() (token *oauth2.Token, err error) {
	var str string
	str, err = e.store.Get(e.prefix + intervalTenant + e.wxTokenStoreKey)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(str), token)
	return
}

// PutToken 设置微信oauth2 token
func (e WxCache) PutToken(token *oauth2.Token) error {
	rb, err := json.Marshal(token)
	if err != nil {
		return err
	}
	return e.store.Set(e.prefix+intervalTenant+e.wxTokenStoreKey, string(rb), int(token.ExpiresIn)-200)
}
