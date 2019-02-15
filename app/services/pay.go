package services

import (
	"extended-design-golang/app/models"
)

//	获取支付方式 MAP
//	1.查询所有channel
//	2.依次在每个map中依次中查询出所有type

func GetPayMethodMap() []models.PayMethodMap {
	map_channels := []models.PayMethodMap{}

	//	查出所有channel
	db.Table("pay_channels").
		Select("" +
			"pay_channels.id, " +
			"pay_channels.name, " +
			"pay_channels.code " +
			"").
		Find(&map_channels)

	//	遍历channel, 依次在DB中找出每个channel所属type
	for i := 0; i < len(map_channels); i++ {
		db.Table("pay_map_channel_type").
			Select(""+
				"pay_map_channel_type.type_code as code, "+
				"pay_types.name as name, "+
				"pay_types.description as description "+
				"").
			Joins("LEFT JOIN pay_types ON pay_map_channel_type.type_code = pay_types.code").
			Where("pay_map_channel_type.channel_code = ?", map_channels[i].Code).
			Find(&map_channels[i].Types)

	}
	return map_channels
}

//	获取支付方式 MAP
//	1.一次查询出所有channel和map的映射
//	2.在代码中对map做整合

func GetAllPay() []models.PayMethodMap {
	// 查询所有channel types 的映射
	channels := []models.PayMethod{}

	db.Table("pay_map_channel_type").
		Select("" +
			"pay_channels.id 	AS 	channel_id, " +
			"pay_channels.name 	AS 	channel_name, " +
			"pay_channels.code 	AS 	channel_code, " +
			"pay_types.id   	AS	type_id, " +
			"pay_types.code   	AS	type_code, " +
			"pay_types.name   	AS	type_name, " +
			"pay_types.description   	AS	type_description" +
			"").
		Joins("LEFT JOIN pay_channels ON pay_map_channel_type.channel_code = pay_channels.code").
		Joins("LEFT JOIN pay_types ON pay_map_channel_type.type_code = pay_types.code").
		Find(&channels)

	//	去除重复channel

	map_channels := getRemoveDuplicatesChannel(channels)

	for i := 0; i < len(map_channels); i++ {
		for j := 0; j < len(channels); j++ {
			if map_channels[i].Code == channels[j].ChannelCode {
				map_channels[i].Types = append(map_channels[i].Types, models.PayMethodMapType{
					Code:        channels[i].TypeCode,
					Name:        channels[i].TypeName,
					Description: channels[i].TypeDescription,
				})
			}

		}
	}

	return map_channels
}

//	channel 去除重复
func getRemoveDuplicatesChannel(channels []models.PayMethod) []models.PayMethodMap {
	map_channels := []models.PayMethodMap{}

	for i := 0; i < len(channels); i++ {
		tmp_channel := models.PayMethodMap{}
		tmp_channel.Id = channels[i].ChannelId
		tmp_channel.Code = channels[i].ChannelCode
		tmp_channel.Name = channels[i].ChannelName

		if isExists(map_channels, tmp_channel) {
			continue
		}

		map_channels = append(map_channels, tmp_channel)
	}

	return map_channels
}

//	判断channel 是否存在
func isExists(arr []models.PayMethodMap, x models.PayMethodMap) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i].Code == x.Code {
			return true
		}
	}
	return false
}
