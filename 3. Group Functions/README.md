# Групповые функции
## Содержание
 - [Таблицы](#таблицы)
 - [Групповые функции](#запросы)

# Запросы
## Выведите на экран номера групп и количество студентов, обучающихся в них
Исполняем SQL-запрос, указанный ниже.  
`GROUP BY ...` определяет колонку, по которой группируются атрибуты.  
`COUNT(...)` производит заданную операцию над сгруппированными данными, в данном случае - подсчет количества строк.  
```SQL
SELECT st.group_num, COUNT(st.n_z) FROM students st GROUP BY st.group_num
```
![Результат](groups_1.png)  
*Результат*  

## Выведите на экран для каждой группы максимальный средний балл
Исполняем SQL-запрос, указанный ниже.  
Используем групповую функцию `MAX(...)` для выбора максимального значения атрибута в группе.  
```SQL
SELECT st.group_num, MAX(st.score) FROM students st GROUP BY st.group_num
```
![Результат](groups_2.png)  
*Результат*  

## Подсчитать количество студентов с каждой фамилией
Исполняем SQL-запрос, указанный ниже.  
Изменяем атрибут, по которому выполняется группировка.  
```SQL
SELECT st.surname, COUNT(st.surname) FROM students st GROUP BY st.surname
```
![Результат](groups_3.png)  
*Результат*  

## Подсчитать студентов, которые родились в каждом году
Исполняем SQL-запрос, указанный ниже.  
`EXTRACT(... FROM ...)` позволяет получить некое определенное поле из даты.  
```SQL
SELECT 
  EXTRACT(YEAR FROM st.date_of_birth), 
  COUNT(EXTRACT(YEAR FROM st.date_of_birth)) 
FROM students st 
GROUP BY 
  EXTRACT(YEAR FROM st.date_of_birth)
```
![Результат](groups_4.png)  
*Результат*  

## Для студентов каждого курса подсчитать средний балл
Исполняем SQL-запрос, указанный ниже.  
Воспользуемся прежде написанной функцией, возвращающей курс студента.  
```SQL
SELECT 
  LEFT(st.group_num::VARCHAR,1) grade, 
  ROUND(AVG(st.score),2) score 
FROM students st 
GROUP BY 
  LEFT(st.group_num::VARCHAR,1)
```
![Результат](groups_5.png)  
*Результат*  

## Для студентов заданного курса вывести один номер группы с максимальным средним баллом
Исполняем SQL-запрос, указанный ниже.  
Берем номера групп...  
...которые с первого курса...  
...группируем данные по группам, так как нас интересует номер группы...  
...сортируем по среднему баллу...  
...получаем только самый максимальный номер.  
```SQL
SELECT st.group_num 
FROM students st
WHERE
  LEFT(st.group_num::VARCHAR,1) = '1'
GROUP BY
  st.group_num
ORDER BY 
  AVG(st.score) DESC
LIMIT 1
```
![Результат](groups_6.png)  
*Результат*  

## Для каждой группы подсчитать средний балл, вывести на экран только те номера групп и их средний балл, в которых он менее или равен 3.5. Отсортировать по от меньшего среднего балла к большему.
Исполняем SQL-запрос, указанный ниже.  
`HAVING` похож на `WHERE`, но используется для групповых функций.  
```SQL
SELECT st.group_num, ROUND(AVG(st.score),2) score
FROM students st
GROUP BY
  st.group_num
HAVING
  AVG(st.score) > 3.5
ORDER BY 
  AVG(st.score)
```
![Результат](groups_7.png)  
*Результат*  

## Для каждой группы в одном запросе вывести количество студентов, максимальный балл в группе, средний балл в группе, минимальный балл в группе
Исполняем SQL-запрос, указанный ниже.  
```SQL
SELECT 
  st.group_num,
  COUNT(st.n_z) student_count,
  MAX(st.score),
  ROUND(AVG(st.score),2),
  MIN(st.score)
FROM students st
GROUP BY
  st.group_num
```
![Результат](groups_8.png)  
*Результат*  

## Вывести студента/ов, который/ые имеют наибольший балл в заданной группе
Исполняем SQL-запрос, указанный ниже.  
Используем обращение к нескольким таблицам, создавая таблицу максимального количества баллов в каждой группе.  
```SQL
SELECT st.*
FROM 
  (SELECT * FROM 
    (SELECT st.group_num, MAX(st.score) FROM students st GROUP BY st.group_num) gr_max
  WHERE gr_max.group_num = '2281') temp_res,
  students st
WHERE
  temp_res.group_num = st.group_num AND temp_res.max = st.score
```
![Результат](groups_9.png)  
*Результат*  

## Аналогично предыдущему заданию, но вывести в одном запросе для каждой группы студента с максимальным баллом.
Исполняем SQL-запрос, указанный ниже.  
```SQL
SELECT st.*
FROM 
  (SELECT st.group_num, MAX(st.score) FROM students st GROUP BY st.group_num) gr_max,
  students st
WHERE
  gr_max.group_num = st.group_num AND gr_max.max = st.score
```
![Результат](groups_10.png)  
*Результат*  
