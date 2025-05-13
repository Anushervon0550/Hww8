
package main

import (
	"fmt"
	"os"
)

type Storage interface {
	Save(data string)
	Load() string
}

type MemoryStorage struct {
	data string
}

func (m *MemoryStorage) Save(data string) {
	m.data = data
}

func (m *MemoryStorage) Load() string {
	return m.data
}

type FileStorage struct {
	filename string
}

func (f *FileStorage) Save(data string) {
	os.WriteFile(f.filename, []byte(data), 0644)
}

func (f *FileStorage) Load() string {
	content, _ := os.ReadFile(f.filename)
	return string(content)
}

type Action interface {
	Do()
}

type PrintAction struct {
	message string
}

func (p PrintAction) Do() {
	fmt.Println("Печать:", p.message)
}

type SaveAction struct {
	data string
}

func (s SaveAction) Do() {
	fmt.Println("Сохраняем:", s.data)
}

func LogAction(a Action) {
	fmt.Println("Выполняется действие:")
	a.Do()
}

func main() {
	mem := &MemoryStorage{}
	mem.Save("Привет из памяти")
	fmt.Println(mem.Load())

	file := &FileStorage{filename: "myfile.txt"}
	file.Save("Привет из файла")
	fmt.Println(file.Load())

	print := PrintAction{"Сообщение на экран"}
	save := SaveAction{"Данные для сохранения"}

	LogAction(print)
	LogAction(save)
}
