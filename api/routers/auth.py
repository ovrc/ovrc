import uuid

from fastapi import APIRouter, Form
from starlette.responses import Response
from starlette.requests import Request

from api.routers import APIResponseModel

router = APIRouter()


@router.post("/login", tags=["auth"])
async def login(
    *,
    response: Response,
    req: Request,
    username: str = Form(...),
    password: str = Form(...)
):

    print(req.headers)
    session_id = str(uuid.uuid4())
    response.set_cookie(key="session_id", value=session_id, httponly=True, secure=True)

    return APIResponseModel()
