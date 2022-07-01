from pony.orm import Database, Required
from decouple import config

db = Database()


class Sites(db.Entity):
    initial_link = Required(str)
    link = Required(str)
    counter = Required(int)


db.bind(
    provider='postgres',
    user=config('database_user'),
    password=config('database_pass'),
    host='127.0.0.1',
    database=config('database_db'),
    port="32700"
)
db.generate_mapping(create_tables=True)