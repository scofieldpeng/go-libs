package config

import (
	"os"
	"testing"
    "fmt"
)

// TestInit 测试初始化
func TestInit(t *testing.T) {
	configPath := os.Getenv("GOPATH") + string(os.PathSeparator) + "src/github.com/scofieldpeng/go-libs/config/test.ini"
	if err := Config.Init(configPath); err != nil {
		t.Fatal(err)
	}
}

// TestReload 测试重载
func TestReload(t *testing.T) {
	if err := Config.Reload(); err != nil {
		t.Fatal(err)
	}
}

// TestListen 测试监听文件夹变化
func TestListen(t *testing.T) {
	Config.AddNodeListener("app_test:path", "test",lala)
}

func lala(oldPath, newPath string) {
	fmt.Println("监听到配置app:path变化，新配置:", newPath, ",旧配置：", oldPath)
	if oldPath == "path1" && newPath == "path2" {
		fmt.Println("配置文件变化正确")
	} else {
		fmt.Println("配置文件变化错误")
	}
}

func TestRead(t *testing.T) {
	value, ok := Config.Read("app_test:testIntRight")
	if !ok {
		t.Fatal("读取app_test:testIntRight失败")
	} else if value != "1" {
		t.Fatal("读取app_test:testIntRight的配置错误,读取到的值为：", value)
	}
}

func TestReadInt(t *testing.T) {
	value, ok := Config.ReadInt("app_test:testIntRight")
	if !ok {
		t.Fatal("读取app_test:testIntright失败")
	} else if value != 1 {
		t.Fatal("读取app_test:testIntRight配置错误，读取到的值为：", value)
	}
}
