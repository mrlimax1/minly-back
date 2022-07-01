from typing import Union
from .tables import Sites
from pony.orm import db_session, select

from ..utils import generate_link


@db_session
def sites(initial_link: str) -> Sites:
    s = select(s for s in Sites if s.initial_link == initial_link).first()
    if s:
        s.counter += 1
        return s
    return Sites(
        initial_link=initial_link,
        link=generate_link(),
        counter=1
    )


@db_session
def get_by_link(link: str) -> Union[Sites, None]:
    return select(s for s in Sites if s.link == link).first()
