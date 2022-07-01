import re
from typing import Optional
from fastapi.responses import JSONResponse

import uvicorn
from fastapi import FastAPI, Response
from starlette.middleware.cors import CORSMiddleware

from .database import sites, get_by_link
from .models import JsonSites
from .utils import SITE_URL_REG

app = FastAPI(docs_url=None, redoc_url=None)

app.add_middleware(
    CORSMiddleware,
    allow_origins=["http://localhost:3000"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)


@app.get("/get")
async def get(response: Response, initial_link: str = None):
    if not initial_link:
        return JSONResponse({"error": "link not found"})
    initial_link = 'https://' + initial_link if not re.search('(http://|https://)', initial_link) else initial_link
    initial_link = initial_link[0:-1] if initial_link[-1] == '/' else initial_link
    if initial_link.startswith('http://') or (re.search(SITE_URL_REG, initial_link)) is None:
        return JSONResponse({"error": "link is invalid"})

    return JsonSites.from_orm(sites(initial_link))


@app.get("/getByMinLy")
async def get_by_minly(link: str = None):
    if not link:
        return JSONResponse({"error": "link not found"})

    s = get_by_link(link)
    if s:
        return s.initial_link
    return JSONResponse({"error": "page not found"})


def start():
    uvicorn.run("src.main:app", host="0.0.0.0", port=8000, reload=True)
