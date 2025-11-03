package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func FirstHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//http.ServeFile(w, r, filepath.Join("..", "index.html"))
	http.ServeFile(w, r, filepath.Join(`D:\Study\go_sprint6_final_project_v1\index.html`))
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
		http.Error(w, "невозможно полкчить файл из формы", http.StatusBadRequest)
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

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	response := fmt.Sprintf(`
<html>
<head><title>Результат загрузки</title></head>
<body>
<p>Конвертация завершена.</p>
<p>Результат записан в файл: <strong>%s</strong></p>
<p>Код Морзе: %s</p>
<p>Исходный текст: %s</p>
</body>
</html>`, filename, result, string(content))
	w.Write([]byte(response))
}
