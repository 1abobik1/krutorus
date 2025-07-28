package utils

import (
	"baulin_proj/internal/request"
	"fmt"
)

func CreateServiceTgMsg(req request.ServiceReq) string {
	if req.Name == "" || req.Name == " " {
		req.Name = "не указано"
	}
	return fmt.Sprintf(
		"📋 Новая заявка!\n\n"+
			"🔹 Услуга: %s\n"+
			"🔹 Тип: %s\n"+
			"👤 Имя: %s\n"+
			"📞 Телефон: %s",
		req.Service, req.Type, req.Name, req.Phone,
	)
}

func CreateLayoutTgMsg(req request.Layout) string {
	if req.Name == "" || req.Name == " " {
		req.Name = "не указано"
	}

	return fmt.Sprintf(
		"📋 Новая заявка!\n\n"+
			"🔹 Проект: %s\n"+
			"👤 Имя: %s\n"+
			"📞 Телефон: %s",
		req.ProjectType, req.Name, req.Phone,
	)
}

func CreateCalcTgMsg(req request.Calc) string {
	if req.Name == "" || req.Name == " " {
		req.Name = "не указано"
	}

	return fmt.Sprintf(
		"📋 Новая заявка на рассчет стоимости!\n\n"+
			"🔹 Место: %s\n"+
			"🔹 Площадь: %s\n"+
			"🔹 Тип ремонта: %s\n"+
			"👤 Имя: %s\n"+
			"📞 Телефон: %s",
		req.Place, req.Square, req.Type, req.Name, req.Phone,
	)
}
