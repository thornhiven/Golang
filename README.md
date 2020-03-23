получить список запросов
/gettasklist

удалить запрос
/deltask?id=[ID]

создать запрос
http://localhost:8080/task
{
	"method": "post",
	"url": "http://google.com",
	"header": "application/json",
	"body": "hello, world!"
}
