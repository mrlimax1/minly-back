from pydantic import BaseModel


class JsonSites(BaseModel):
    initial_link: str
    link: str
    counter: int

    class Config:
        orm_mode = True
