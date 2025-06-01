import os
from pathlib import Path
from dotenv import load_dotenv

load_dotenv()

BASE_DIR = Path(__file__).parent.parent

class Config:
    MODEL_PATH = os.getenv('MODEL_PATH', str(BASE_DIR / "model_v1.pk"))
    LEVEL_LABELS = os.getenv('LEVEL_LABELS', 'A1,A2,B1,B2,С1,C2').split(',')
    FLASK_PORT = int(os.getenv('FLASK_PORT', 8080))