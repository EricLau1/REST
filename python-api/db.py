import pymysql

def getConnection():
    return pymysql.connect(host='localhost', user='root', passwd='', db='test')
#end getConnection

def getAll(cursor):

    sql = 'select * from produto'

    cursor.execute(sql)

    data = cursor.fetchall()

    return data
#end getAll

def find(id, cursor):

    sql = 'select * from produto where id = %s'

    cursor.execute(sql, (id));

    data = cursor.fetchone()

    return data
#end find


def create(produto, cursor):

    sql = 'insert into produto (descricao, quantidade, preco) values ( %s, %s, %s )'
    
    result = cursor.execute(sql, (produto['descricao'], produto['quantidade'], produto['preco']))

    return result
#end create

def update(produto, cursor):

    sql = 'update produto set descricao = %s, quantidade = %s, preco = %s where id = %s'

    result = cursor.execute(sql, (produto['descricao'], produto['quantidade'], produto['preco'], produto['id']))

    return result
#end update

def delete(id, cursor):

    sql = 'delete from produto where id = %s'

    result = cursor.execute(sql, (id))

    return result
# end delete

def register(dados, cursor):

    sql = "insert into login (nome, senha) values (%s, %s)"

    result = cursor.execute(sql, (dados['nome'], dados['senha']))

    return result
#end register

def auth(login, cursor):

    sql = "select * from login where nome = %s && senha = %s"

    cursor.execute(sql, (login['nome'], login['senha']))

    result = cursor.fetchone()

    return result
#end auth

def applyCommit(connection):
    connection.commit()

def closeConnection(connection):
    connection.close()


#closeConnection(connection)
