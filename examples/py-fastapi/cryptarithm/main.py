# vim: fileencoding=utf-8

import fire, uvicorn
from fastapi import FastAPI

from parse import Parser
from lexer import Token
from problem import Problem

from typing import Dict,List,Any

app = FastAPI()

@app.get("/")
async def read_root() -> Dict[str, List[str]]:
    return {"accessible_path": ["/docs", "/cryptarithm"]}

@app.get("/cryptarithm/")
async def solve_cryptarithm(expr: str = '') -> Dict[str, Any]:
    p = Parser(Token.tokenize(expr))
    pr = Problem(p.parse())
    status, solutions = pr.search_all_solution()
    return {"problem": expr, "answer": solutions, "status": status}

class Cli(object):
    def __init__(self):
        pass

    def server(self, port: int = 8000):
        """start `Hello World` server"""
        uvicorn.run(app, host='0.0.0.0', port=port)

    def cryptarithm(self, expr: str = ""):
        """solve cryptarithm problem"""
        print("Problem: {}".format(expr))

        p = Parser(Token.tokenize(expr))
        pr = Problem(p.parse())
        print(pr.search_all_solution())

if __name__ == '__main__':
    fire.Fire(Cli)
