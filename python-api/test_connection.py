import db

conn = db.getConnection()

cursor = conn.cursor()

data = db.getAll(cursor)

print(data[0])

db.closeConnection(conn)