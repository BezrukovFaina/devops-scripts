import uuid
from datetime import datetime

def generate_uuid() -> str:
    return str(uuid.uuid4())

def get_timestamp() -> str:
    return datetime.now().strftime('%Y-%m-%d %H:%M:%S')

def log(message: str) -> None:
    timestamp = get_timestamp()
    print(f"[{timestamp}] {message}")

def main() -> None:
    uuid_value = generate_uuid()
    log(f"Generated UUID: {uuid_value}")

if __name__ == "__main__":
    main()