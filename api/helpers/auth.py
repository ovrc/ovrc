from datetime import datetime, timedelta
from typing import Optional

import jwt
from fastapi import HTTPException
from fastapi.openapi.models import OAuthFlows as OAuthFlowsModel
from fastapi.security import OAuth2
from fastapi.security.utils import get_authorization_scheme_param
from starlette.requests import Request
from starlette.status import HTTP_401_UNAUTHORIZED

from api.config import JWS_ALGORITHM, JWS_SECRET_KEY


def create_access_token(*, data: dict, expires_delta: timedelta = None):
    """
    Creates and return an encoded jwt token.
    """
    to_encode = data.copy()
    if expires_delta:
        expire = datetime.utcnow() + expires_delta
    else:
        expire = datetime.utcnow() + timedelta(minutes=15)
    to_encode.update({"exp": expire})

    encoded_jwt = jwt.encode(to_encode, JWS_SECRET_KEY, algorithm=JWS_ALGORITHM)

    return encoded_jwt


class OAuth2PasswordBearerWithCookie(OAuth2):
    """
    This is a class that is very similar to OAuth2PasswordBearer, but instead of reading
    the jwt access token from an Auth header, it is read from a (hopefully httsonly and
    secure) cookie.
    """

    def __init__(
        self,
        tokenUrl: str,
        scheme_name: str = None,
        scopes: dict = None,
        auto_error: bool = True,
    ):
        if not scopes:
            scopes = {}
        flows = OAuthFlowsModel(password={"tokenUrl": tokenUrl, "scopes": scopes})
        super().__init__(flows=flows, scheme_name=scheme_name, auto_error=auto_error)

    async def __call__(self, request: Request) -> Optional[str]:
        authorization: str = request.cookies.get("access_token")

        scheme, param = get_authorization_scheme_param(authorization)
        if not authorization or scheme.lower() != "bearer":
            if self.auto_error:
                raise HTTPException(
                    status_code=HTTP_401_UNAUTHORIZED,
                    detail="Not authenticated",
                    headers={"WWW-Authenticate": "Bearer"},
                )
            else:
                return None

        return param
