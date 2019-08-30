from datetime import timedelta

from fastapi import APIRouter, Depends
from fastapi.security import OAuth2PasswordRequestForm
from starlette.responses import Response

from api.config import JWS_ACCESS_TOKEN_EXPIRE_MINUTES
from api.helpers.auth import OAuth2PasswordBearerWithCookie, create_access_token

router = APIRouter()

oauth2_scheme = OAuth2PasswordBearerWithCookie(tokenUrl="/auth/token")


@router.post("/token", tags=["auth"])
async def auth_token(
    response: Response, form_data: OAuth2PasswordRequestForm = Depends()
):
    username = form_data.username
    password = form_data.password

    # Assume user is authenticated via username and password...
    # Add database checks later.

    access_token_expires = timedelta(minutes=JWS_ACCESS_TOKEN_EXPIRE_MINUTES)

    access_token = create_access_token(
        data={"sub": username}, expires_delta=access_token_expires
    ).decode("utf-8")

    response.set_cookie(
        key="access_token", value=f"Bearer {access_token}", httponly=True, secure=True
    )

    return
