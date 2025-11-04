package handlers

import (
	"fmt"
	"html"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func FirstHandler(w http.ResponseWriter, r *http.Request) {
	cwd, err := os.Getwd()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "некорректный метод", http.StatusMethodNotAllowed)
		return
	}
	//filePath := filepath.Join("..", "index.html")
	filePath := filepath.Join(cwd, "index.html")
	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "файл не найден", http.StatusNotFound)
		return
	}
	defer file.Close()
	fi, err := file.Stat()
	if err != nil {
		http.Error(w, "ошибка получения информации о файле", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	http.ServeContent(w, r, "index.html", fi.ModTime(), file)
}

func SecondHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "некорректный метод", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "некорректный парсинг формы", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "невозможно получить файл из формы", http.StatusBadRequest)
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "невозможно прочитать файл", http.StatusInternalServerError)
		return
	}

	result, err := service.Convert(string(content))
	if err != nil {
		http.Error(w, "некорректная конвертация: "+err.Error(), http.StatusInternalServerError)
		return
	}

	filename := fmt.Sprintf("%s%s", time.Now().UTC().Format("20060102T150405"), filepath.Ext("index.html"))

	outFile, err := os.Create(filename)
	if err != nil {
		http.Error(w, "невозможно создать файл", http.StatusInternalServerError)
		return
	}
	defer outFile.Close()

	if _, err := outFile.WriteString(result); err != nil {
		http.Error(w, "невозможно записать файл", http.StatusInternalServerError)
		return
	}

	filenameEsc := html.EscapeString(filename)
	resultEsc := html.EscapeString(result)
	contentEsc := html.EscapeString(string(content))

	response := fmt.Sprintf(`
<html>
<head><title>Результат загрузки</title></head>
<body>
<p>Конвертация завершена.</p>
<p>Результат записан в файл: <strong>%s</strong></p>
<p>Код Морзе: %s</p>
<p>Исходный текст: %s</p>
</body>
</html>`, filenameEsc, resultEsc, contentEsc)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, err = w.Write([]byte(response))
	if err != nil {
		fmt.Println("Ошибка при отправке ответа:", err)
	}
}
