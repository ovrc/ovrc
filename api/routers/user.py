from fastapi import APIRouter
from starlette.requests import Request
from starlette.responses import Response
from starlette.status import HTTP_401_UNAUTHORIZED

from api.routers import APIResponseModel

router = APIRouter()


@router.get("/me", tags=["users"])
async def users_me(*, response: Response, request: Request):
    if "session_id" in request.cookies:
        # TODO: Validate user via session.
        return APIResponseModel(data={"user": {"email": "test@test.com"}})

    response.status_code = HTTP_401_UNAUTHORIZED
    return APIResponseModel(status="fail", message="could not validate user")
