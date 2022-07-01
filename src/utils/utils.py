import random

from decouple import config

from src.database.tables import Sites
from pony.orm import select
SITE_URL_REG = '(^https?)[a-zA-Z0-9\.\/\?\:@\-_=#]+\.([a-zA-Z]){2,6}([a-zA-Z0-9\.\&\/\?\:@\-_=#])*'
letters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789'


def generate_link():
    link = config('url') + ''.join(random.choice(letters) for i in range(5))
    while select(s for s in Sites if s.link == link):
        link = config('url') + ''.join(random.choice(letters) for i in range(5))
    return link

