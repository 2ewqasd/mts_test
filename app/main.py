from fastapi import FastAPI


app = FastAPI()


@app.get("/")
def get():
    return {"status": 200}

@app.post("/")
def post():
    return {"status": 200}