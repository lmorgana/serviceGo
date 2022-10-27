# serviceGo
# Команды
>make - запуск

>make up - поднять контейнеры

>make down - остановить и удалить запущенные контейнеры

>make clean - удалить неиспользуемые контейнеры и образы

>make fclean - удалить все образы и контейнеры


# /getBalance.json
Метод получения баланса пользователя. Принимает id пользователя.

Params:

	Id_user
  
Response:

	Value
  
exit codes:

	200, 401, 412


# /topUp.json
Метод начисления средств на баланс. Принимает id пользователя и сколько средств зачислить.

Params:

	Id_user
  
	Value
  
Response:

Exit codes:

	200, 401, 412



# /reserving.json
Метод резервирования средств с основного баланса на отдельном счете. Принимает id пользователя, ИД услуги, ИД заказа, стоимость.

Params:

	Id_user
  
	Id_service
  
	Id_order
  
	Value
  
Response:

Exit codes:

	200, 401, 412


# /acceptFromReserve.json
Метод признания выручки – списывает из резерва деньги, добавляет данные в отчет для бухгалтерии. Принимает id пользователя, ИД услуги, ИД заказа, сумму.

Params:

	Id_user
  
	Id_service
  
	Id_order
  
	Value
  
Response:

Exit codes:

	200, 401 412




# /refuseFromReserve.json
Метод отклонения услуги – списывает из резерва деньги, возвращает деньги на счет пользователя. Принимает id пользователя, ИД услуги, ИД заказа, сумму.

Params:

	Id_user
  
	Id_service
  
	Id_order
  
	Value
  
Response:

Exit codes:

	200, 401 412





# Errors:

JSON example:
{
	“error”: “true”,
	“status”: 402,
	“message”: “invalid_id_user”,
	“description”: “invalid user id”
}
<img width="657" alt="Screen Shot 2022-10-27 at 2 05 18 PM" src="https://user-images.githubusercontent.com/89391330/198268490-3540fe88-03d7-4c0c-a3f4-9c7b8154767b.png">



