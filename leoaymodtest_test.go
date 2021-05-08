package leoaymodtest

import "testing"

// 参数传入 t *testing.T
func TestStart(t *testing.T) {
	result := Start()
	if result != 899 {
		t.Error("mytools starts failed.") // Error 表示测试失败
	}
	t.Log("mytools starts successfully.") // Log 打印信息
}