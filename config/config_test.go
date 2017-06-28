package config

import (
	"testing"
)

func TestRead(t *testing.T) {
	conf := &Config{}
	conf.InitConfig("./config.conf")
	if "werqwecxv" != conf.Read("www", "aaa") {
		t.Fatal("aaa=werqwecxv")
	}

	if "twer etrew%%&%*" != conf.Read("zzz", "vv") {
		t.Fatal("vv=twer etrew%%&%*")
	}

	if "c:/go" != conf.Read("default", "path") {
		t.Fatal("path= c:/go")
	}

	if "1.44" != conf.Read("default", "version") {
		t.Fatal("version = 1.44")
	}

	if "wrong" != conf.Read("test", "something") {
		t.Fatal("something  = wrong  #注释1")
	}

	if "666" != conf.Read("test", "num") {
		t.Fatal("num =	666")
	}

	if "" != conf.Read("test", "fdfdfd") {
		t.Fatal("#fdfdfd = fdfdfd    注释整行")
	}

	if "refer" != conf.Read("test", "refer") {
		t.Fatal("refer= refer       //注释3")
	}

	t.Log("success!")
}
