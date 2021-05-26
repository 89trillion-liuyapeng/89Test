package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/ini.v1"
	"io/ioutil"
	"os"
)

type soldierInfo struct {
	//稀有度
	Rarity string
	//解锁阶段
	UnlockArena string
	//士兵id
	id string
	//战力
	CombatPoints string
}

var soldierMap = map[string]soldierInfo{}

func main() {
	var info soldierInfo
	//读取配置信息，获取端口号
	cfginfo, err := ini.Load("File/app.ini")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cfginfo.Section("server").Key("HttpPort"))

	//解析json文件
	jsonInfo, err2 := os.Open("File/config.army.model.json")
	if err2 != nil {
		fmt.Println(err2)
	}
	defer jsonInfo.Close()
	value, err3 := ioutil.ReadAll(jsonInfo)

	json.Unmarshal([]byte(value), &soldierMap)
	if err3 != nil {
		fmt.Println(err3)
	}
	//fmt.Println(soldierMap)

	//根据稀有度获取解锁阶段
	ra1 := info.getUnlockArenaByRarity("1")

	//根据id获取稀有度
	ra2 := info.getRarityById("10101")

	//根据id获取战力
	ra3 := info.getCombatPointsById("10101")
	fmt.Println(ra1, ra2, ra3)

	//获取每个阶段解锁的相应士兵的信息
	info.getAllInfoByUnlockArena()
}

func (info soldierInfo) getValueByRarity(Rarity string) string {
	for _, value := range soldierMap {
		if Rarity == value.Rarity {
			return value.UnlockArena
		}
	}
	return ""
}

func (info soldierInfo) getValueById(id string) soldierInfo {
	for key, value := range soldierMap {
		if id == key {
			return value
		}
	}
	return soldierInfo{}
}

func (info soldierInfo) getValueByUnlockArena(UnlockArena string) soldierInfo {
	for _, value := range soldierMap {
		if UnlockArena == value.UnlockArena {
			return value
		}
	}
	return soldierInfo{}
}

//根据稀有度获取解锁阶段	？
func (info soldierInfo) getUnlockArenaByRarity(Rarity string) string {
	UnlockArena := info.getValueByRarity(Rarity)
	if UnlockArena != "" {
		return UnlockArena
	}
	return ""
}

//根据士兵id获取稀有度
func (info soldierInfo) getRarityById(id string) string {
	mapValue := info.getValueById(id)
	if mapValue != (soldierInfo{}) {
		return mapValue.Rarity
	}
	return ""
}

//根据id获取战力
func (info soldierInfo) getCombatPointsById(id string) string {
	mapValue := info.getValueById(id)
	if mapValue != (soldierInfo{}) {
		return mapValue.CombatPoints
	}
	return ""
}

//获取每个阶段解锁的相应士兵的信息
func (info soldierInfo) getAllInfoByUnlockArena() {
	map1 := map[string]string{}
	for _, value := range soldierMap {
		unValue := value.UnlockArena
		map1[unValue] = ""
	}
	for key, _ := range map1 {
		soldierValue := info.getValueByUnlockArena(key)
		fmt.Println("解锁阶段:", key, "  士兵信息:", soldierValue)
	}
}
