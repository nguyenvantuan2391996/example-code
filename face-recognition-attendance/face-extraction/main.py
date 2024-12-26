from io import BytesIO

from fastapi import FastAPI, UploadFile, Request
import cv2
import uvicorn
import numpy as np
from deepface import DeepFace
from PIL import Image
from starlette.middleware.base import BaseHTTPMiddleware
from starlette.middleware.cors import CORSMiddleware
from starlette.responses import JSONResponse

app = FastAPI()

API_KEY = "GIyBK7ge2fLWK8G6hXDh47xbm5sKVCZd"


class APIKeyMiddleware(BaseHTTPMiddleware):
    async def dispatch(self, request: Request, call_next):
        api_key = request.headers.get("X-API-Key")

        if api_key != API_KEY:
            return JSONResponse(
                status_code=401,
                content={"error": "Unauthorized"},
            )

        response = await call_next(request)
        return response


app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_methods=["*"],
    allow_headers=["*"],
)
app.add_middleware(APIKeyMiddleware)


def extract_face_vector(image_input, model_name) -> np.ndarray:
    img = cv2.cvtColor(image_input, cv2.COLOR_RGB2BGR)
    embedding = DeepFace.represent(img, model_name=model_name)
    if embedding is not None and len(embedding) > 0:
        return embedding[0]["embedding"]
    return np.empty(())


@app.post("/api/v1/extract/")
async def extract(file: UploadFile):
    try:
        file_content = await file.read()
        image_preprocess = Image.open(BytesIO(file_content))
        vector = extract_face_vector(np.array(image_preprocess), "ArcFace")

        return {
            "vector": vector,
        }
    except Exception as e:
        return {"error": f"System error. {str(e)}"}


if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=5000)
