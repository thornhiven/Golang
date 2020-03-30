получить список запросов
/gettasklist

удалить запрос
/deltask/id

создать запрос
/task
{
	"method": "get",
	"url": "http://google.com",
	"header": {
        	"Cache-Control": ["no-cache"],
        	"Content-Security-Policy": ["default-src * data: blob: 'unsafe-eval' 'unsafe-inline'"],
        	"Content-Type": ["text/html;charset=UTF-8"],
        	"Date": ["Mon, 23 Mar 2020 07:37:55 GMT"]
    	},
	"body": "hello, world!"
}
