package main

import (
	"fmt"
	"github.com/go-ole/go-ole"
	"log"
	"os"
	"time"

	"github.com/go-ole/go-ole/oleutil"
)

func main() {
	// Инициализация COM
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	// Создание экземпляра объекта Scales
	unknown, err := oleutil.CreateObject("MassaKDriver100.Scales")
	if err != nil {
		log.Fatalf("Ошибка создания объекта Scales: %v", err)
	}
	scale, err := unknown.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		log.Fatalf("Ошибка получения интерфейса IDispatch: %v", err)
	}
	defer scale.Release()

	// Установка параметра подключения (COM3)
	_, err = oleutil.PutProperty(scale, "Connection", "COM3")
	if err != nil {
		log.Fatalf("Ошибка установки параметра Connection: %v", err)
	}

	// Открытие соединения с весами
	res, _ := oleutil.CallMethod(scale, "OpenConnection")
	if errorCode, ok := res.Value().(int32); !ok || errorCode != 0 {
		log.Fatalf("Не удалось подключиться к весам. Код ошибки: %v", errorCode)
	}
	defer func() {
		// Закрытие соединения при завершении
		_, _ = oleutil.CallMethod(scale, "CloseConnection")
	}()

	// Открываем файл для записи данных
	file, err := os.Create("weights.log")
	if err != nil {
		log.Fatalf("Ошибка создания файла: %v", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("Ошибка закрытия файла: %v", err)
		}
	}()

	// Бесконечный цикл для получения данных
	for {
		// Чтение веса
		res, _ := oleutil.CallMethod(scale, "ReadWeight")
		if errorCode, ok := res.Value().(int32); !ok || errorCode != 0 {
			log.Printf("Ошибка чтения веса. Код ошибки: %v", errorCode)
			continue
		}

		// Получение значения веса
		weightVariant, err := oleutil.GetProperty(scale, "Weight")
		if err != nil {
			log.Printf("Ошибка получения значения Weight: %v", err)
			continue
		}
		weight := weightVariant.Val

		// Получение цены деления
		divisionVariant, err := oleutil.GetProperty(scale, "Division")
		if err != nil {
			log.Printf("Ошибка получения значения Division: %v", err)
			continue
		}
		division := int32(divisionVariant.Val) // Приводим к int32

		// Получение стабильности веса
		stableVariant, err := oleutil.GetProperty(scale, "Stable")
		if err != nil {
			log.Printf("Ошибка получения значения Stable: %v", err)
			continue
		}
		stable := stableVariant.Val == int64(1) // Сравниваем с int32(1)

		// Если вес нестабилен, игнорируем данные
		if !stable {
			fmt.Println("Вес нестабилен. Данные игнорируются.")
			continue
		}

		// Преобразование веса в зависимости от Division
		switch division {
		case 0:
			weightGrams := float64(weight) / 100.0 // 100 мг = 0.1 г
			fmt.Printf("Вес: %.3f грамм\n", weightGrams)
		case 1:
			fmt.Printf("Вес: %d грамм\n", weight)
		case 2:
			weightGrams := float64(weight) * 10.0
			fmt.Printf("Вес: %.1f грамм\n", weightGrams)
		case 3:
			weightGrams := float64(weight) * 100.0
			fmt.Printf("Вес: %.1f грамм\n", weightGrams)
		case 4:
			weightKg := float64(weight) / 100.0
			fmt.Printf("Вес: %.2f кг\n", weightKg)
		default:
			log.Printf("Неизвестная цена деления: %d", division)
		}

		// Запись данных в файл
		_, err = file.WriteString(fmt.Sprintf("Вес: %v, Стабильность: %v, Цена деления: %v\n", weight, stable, division))
		if err != nil {
			log.Printf("Ошибка записи в файл: %v", err)
			continue
		}

		// Пауза перед следующим запросом
		time.Sleep(1 * time.Second)
	}
}
