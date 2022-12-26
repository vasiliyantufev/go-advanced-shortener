package app

import "sync"

type DataMap struct {
	mx   sync.RWMutex
	Data map[string]string
}

//
//Создаём map
func NewDM() *DataMap {
	return &DataMap{
		Data: make(map[string]string),
	}
}

//пишем в map
func (data *DataMap) Put(id string, o string) {
	//мьютекс блокировка с общим доступом для записи
	data.mx.Lock()
	defer data.mx.Unlock()
	data.Data[id] = o
}

//читаем из map
func (data *DataMap) Get(id string) (o string, b bool) {
	//мьютекс блокировка с общим доступом для чтение
	data.mx.RLock()
	defer data.mx.RLock()
	o, b = data.Data[id]
	return
}