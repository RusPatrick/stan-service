News-service
-
Данный серивис осуществляет публикацию сообщений и чтение сообщений из очереди
 
ЧАСТЬ 1: Работа с сервисом
-
**Старт сервиса**  
```docker-compose up --build -d```  

**Остановка сервиса**  
```docker-compose down```  

ЧАСТЬ 2: Информация о сервисе
-
#### **2.1 Публикация новости:**  
**URL:**  
&emsp;&emsp; http://localhost:8000/api/v1/news  
**METHOD:** POST  
**REQUEST BODY**  
```
{
	"title": "test",        //string, Заголовок новости
	"date": "2019-01-01"    //date, дата (формат: гггг-мм-дд)
}
```
#### **2.2 Чтение новости:**  
**URL:**  
&emsp;&emsp; http://localhost:8000/api/v1/news/{readerName}  
**URL-ПАРАМЕТРЫ**  
&emsp;&emsp; `readerName` - имя читателя, string    
**METHOD:** GET  
**RESPONSE BODY**  
```
{
	"title": "test",        //string, Заголовок новости
	"date": "2019-01-01"    //date, дата (формат: гггг-мм-дд)
}
```
