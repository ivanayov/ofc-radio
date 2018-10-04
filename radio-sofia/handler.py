def handle(req):
    
    bytes = None
    with open("./function/index.html") as f:
        bytes = f.read()

    return str(bytes)