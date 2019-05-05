package testing_demo_3

import "testing"

func TestMonster_Store(t *testing.T) {
	//先创建一个Monster
	monster:=Monster{
		Name:"红孩儿",
		Age:10,
		Skill:"三味真火",
	}

	res:=monster.Store()
	if !res {
		t.Fatalf("monster.Store() 错误，希望为%v,实际为%v",true,res)
	}else {
		t.Logf("monster.Store() 正确，希望为%v,实际为%v",true,res)
	}
}

func TestMonster_Restore(t *testing.T) {
	//先创建一个Monster
	monster:=Monster{}

	res:=monster.Restore()
	if !res {
		t.Fatalf("monster.Store() 错误，希望为%v,实际为%v",true,res)
	}else {
		t.Logf("monster.Store() 正确，希望为%v,实际为%v,当前Monster为%v",true,res,monster)
	}
}

