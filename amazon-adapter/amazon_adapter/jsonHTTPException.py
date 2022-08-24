class JSONHTTPException(Exception):

    def __init__(self, error: str, status: int, message: str):
        self.error = error
        self.status = status
        self.message = message
