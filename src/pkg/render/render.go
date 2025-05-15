package render
	// При старте сервера (main.go):
    //     Вызывается CreateTemplateCache() — все шаблоны загружаются в память.
    //     Кэш сохраняется в app.TemplateCache.
    // При запросе (например, /home):
    //     handlers.Home вызывает RenderTemplate(w, "home.page.tmpl").
    //     RenderTemplate достаёт шаблон из кэша и отправляет его клиенту.
    // Если шаблонов нет в site_templates:
    //     CreateTemplateCache вернёт ошибку, и сервер не запустится (логируется в main.go).
import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/monstrong/proyektnaya-practica/src/pkg/config"
)

//Что это? Пустая карта (map) для кастомных функций, которые можно использовать в шаблонах.

//Пример использования: Если добавить сюда функцию "uppercase": strings.ToUpper,
//  то в шаблоне можно писать {{ "hello" | uppercase }} → выведет "HELLO".
var functions = template.FuncMap{

}
//Хранит указатель на конфиг приложения (чтобы render мог достучаться до TemplateCache).
var app *config.AppConfig

//sets the config for the template packege
//Сохраняет переданный конфиг в переменную app
func NewTemplates(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	var tc map[string]*template.Template
	if app.UseCache {
		// get the template cache from the app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// Достаём шаблон из кэша:
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache") //Шаблон не найден в кэше
	}

	//Рендерим в буфер:
	buf := new(bytes.Buffer)
	_ = t.Execute(buf, nil)
	//  t.Execute — применяет шаблон, подставляя данные (второй аргумент nil — значит, данных нет).
    //	Результат пишется в buf (временный буфер в памяти), а не сразу в ResponseWriter.
    //	Зачем буфер? Чтобы избежать частичной отправки HTML, если в шаблоне будет ошибка.

	//	Отправка клиенту:
	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser: ", err)
	}
	//	Копируем HTML из буфера в ResponseWriter.
}

// создаем template cache как map
//Цель: Предварительно загрузить все шаблоны в память, 
// чтобы не парсить их при каждом запросе.
func CreateTemplateCache() (map[string]*template.Template, error) {

	// Пустая map, где ключ — имя файла (например, "home.page.tmpl"), 
	// а значение — готовый шаблон.
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./site_templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		fmt.Println("Парсим шаблон:", page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		//  template.New(name) — создаёт новый шаблон с именем файла.
    	//  .Funcs(functions) — подключает кастомные функции (если бы они были).
    	//  .ParseFiles(page) — парсит HTML-шаблон из файла.

		matches, err := filepath.Glob("./site_templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		//    Ищет файлы с расширением .layout.tmpl (например, base.layout.tmpl).
		//    Если найдены — добавляет их к текущему шаблону (например, чтобы хедер/футер были общими для всех страниц).
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./site_templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts  // сохранение в кэш. Например: myCache["home.page.tmpl"] = готовый_шаблон
	}
	return myCache, nil
}
