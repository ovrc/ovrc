from typing import Dict

from pydantic import BaseModel, validator


class APIResponseModel(BaseModel):
    """
    Follows the JSend spec. Well, kind of. Some of the fields are optional (depending on
    the status) but all of them are returned as part of the response.
    """

    status: str = "success"
    data: Dict = {}
    message: str = None
    code: int = None

    @validator("status")
    def allowed_statuses(cls, v):
        allowed = ["success", "fail", "error"]

        if v not in allowed:
            raise ValueError("status is not allowed")

        return v
