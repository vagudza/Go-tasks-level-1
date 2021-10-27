package tasks

import "fmt"

/*
21.	Реализовать паттерн «адаптер» на любом примере.

Паттерн Adapter принадлежит к структурным паттернам и используется для структурирования классов и объектов.
Необходимость использования паттерна Adapter возникает в случаях, когда нужно привести (адаптировать) одну систему к требованиям другой системы.
*/

// интерфейс USB в ПК
type USB interface {
	ConnectWithUsbCable()
	SendDataOnUSB(data []byte)
}

// структура "карта памяти" объемом 256 байт, необходимо скопировать данные c ПК на карту
type MemoryCard struct {
	data []byte
	size int
}

// функция, возвращающая "карту памяти". Определяем объем и форматируем
func NewMemoryCard(size int) *MemoryCard {
	card := MemoryCard{}
	card.size = size
	card.data = make([]byte, size)
	return &card
}

// запись данных на карту памяти
func (card *MemoryCard) RecieveDataToCard(data []byte) {
	copy(card.data, data)
	fmt.Printf("Data is writed on Memory Card: %v\n", card.data)
}

// Адаптер. Адаптируемый класс (карта памяти) становится одним из полей адаптера.
// Это логично, ведь в реальной жизни мы тоже вставляем карту внутрь кардридера, и она тоже становится его частью.
type CardReader struct {
	memCard *MemoryCard
}

// реализация интерфейса USB: подключение по кабелю USB
func (cr CardReader) ConnectWithUsbCable() {
	fmt.Println("Card Reader: connected with USB")
}

// реализация интерфейса USB: передача данных по кабелю USB
func (cr CardReader) SendDataOnUSB(data []byte) {
	fmt.Println("Card Reader: sending data from USB to Memory Card...")
	cr.memCard.RecieveDataToCard(data)
	fmt.Println("Card Reader: data sending is done!")
}

// подключение карты памяти к кардридеру:
func (cr *CardReader) InsertMemoryCard(memCard *MemoryCard) {
	cr.memCard = memCard
	fmt.Println("Memory Card: inserted!")
}

// компьютер, с которого нужно передать данные на карту памяти
type Computer struct {
	// у компьютера есть USB порт
	usbPort USB
}

// функция принимает интерфейс USB - любой девайс с USB интерфейсом.
// вызываем метод у device, реализующий подключение к USB
func (com *Computer) ConnectUSBDevice(device USB) {
	fmt.Println("Computer: new device found on USB")
	com.usbPort = device
	com.usbPort.ConnectWithUsbCable()
}

func (com *Computer) SendDataToUSBDevice() {
	data := []byte("some binary data")
	fmt.Println("Computer: send data on USB...")
	com.usbPort.SendDataOnUSB(data)
	fmt.Println("Computer: send data on USB done!")
}

func Task21() {
	// определяем сущности: кардридер, карта памяти и компьютер
	MyCardReader := CardReader{}
	// определяем объем карты в "конструкторе"
	MyMemoryCard := NewMemoryCard(64)
	MyComputer := Computer{}

	// подключаем к компьютеру кардридер через USB
	MyComputer.ConnectUSBDevice(&MyCardReader)
	// вставляем карту памяти в кардридер
	MyCardReader.InsertMemoryCard(MyMemoryCard)
	// передаем данные на карту памяти
	MyComputer.SendDataToUSBDevice()
}
