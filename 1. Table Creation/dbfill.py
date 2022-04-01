import psycopg2
import random

conn = psycopg2.connect(dbname='gorodnichin', user='gorodnichin', 
                        password='mysupersecretpass', host='95.217.232.188', port='7777')

cursor = conn.cursor()

def fem(name,surname):
    if (name[-1] == "а") or (name[-1] == "я"):
        return surname+"а"
    else:
        return surname

names = ["Иван", "Павел", "Владимир", "Андрей", "Юлия", "Арина", "Петр", "Олег", "Даниил", "Мария", "Александр", "Илья", "Диана", "Ольга", "Полина", "Алексей", "Данил", "Арсений", "Николай", "Наталья", "Ульяна", "Дарья"]
surnames = ["Синицын", "Лебедев", "Коршунов", "Орлов", "Уткин", "Воробьев", "Голубев", "Филин", "Соколов", "Снегирев", "Дроздов", "Дятлов", "Скворцов", "Петухов", "Гусев", "Воронов", "Ласточкин", "Сорокин", "Соловьев", "Перепелкин", "Попугаев"]
groups = [1132, 2281, 4242, 1337, 3602, 2112]

for i in range(20):
    name = random.choice(names)
    surname = fem(name,random.choice(surnames))
    cursor.execute('INSERT INTO students (n_z, name, surname, group_num, score, date_of_birth) VALUES (%s,%s,%s,%s,%s,%s)',(i,name,surname,random.choice(groups),random.random()*3+2,str(random.randint(1998,2003))+"-"+str(random.randint(1,12))+"-"+str(random.randint(1,30))))
conn.commit()


cursor.execute('SELECT * FROM students')
studs = cursor.fetchall()
print(studs)

cursor.execute('SELECT * FROM hobby')
hobbys = cursor.fetchall()
print(hobbys)

for i in range(30):
    year_start = random.randint(2019,2022)
    year_end = random.randint(year_start,2022)
    month_start = random.randint(1,12)
    if (year_end > year_start):
        month_end = random.randint(1,12)
    else:
        month_end = random.randint(month_start,12)
    day_start = random.randint(1,30)
    if (year_end > year_start) or (month_end > month_start):
        day_end = random.randint(1,30)
    else:
        day_end = random.randint(day_start,30)
    stud = random.choice(studs)
    hobby = random.choice(hobbys)
    if random.random() > 0.5:
        cursor.execute('INSERT INTO student_hobby (n_z, id_hobby, date_start) VALUES (%s,%s,%s)',(stud[0],hobby[0],str(year_start)+"-"+str(month_start)+"-"+str(day_start)))
    else:
        cursor.execute('INSERT INTO student_hobby (n_z, id_hobby, date_start, date_end) VALUES (%s,%s,%s,%s)',(stud[0],hobby[0],str(year_start)+"-"+str(month_start)+"-"+str(day_start),str(year_end)+"-"+str(month_end)+"-"+str(day_end)))
conn.commit()