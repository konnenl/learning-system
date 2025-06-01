from flask import Flask, request, jsonify
from ml.train import train
from api.config import Config
import numbers

app = Flask(__name__)

model = train()  
with open(Config.MODEL_PATH, "wb") as f:
    import dill
    dill.dump(model, f)

level_label = ['A1','A2','B1','B2','С1','C2',]

@app.route('/predict', methods=['POST'])
def predict():
    if not request.is_json:
        return jsonify({"error": "Request must be JSON"}), 415
    
    data = request.get_json()

    if not data or 'x' not in data:
        return jsonify({"error": "Пропущено поле 'x'"}), 400
    
    if not isinstance(data['x'], list):
        return jsonify({"error": "'x' должно быть массивом"}), 400
    
    for i, value in enumerate(data['x']):
        if not isinstance(value, numbers.Number):
            return jsonify({"error": f"Элемент {i} не число"}), 400
        if value > 1:
            return jsonify({"error": f"Элемент {i} должен быть меньше 1"}), 400

    if len(data['x']) != 6:
        return jsonify({"error": "В массиве 'x' должно быть 6 элементов"}), 400

    result = level_label[model.predict([data['x']]).tolist()[0] - 1]
    
    return jsonify({
        "status": "success",
        "result": result,
        "received_data": data['x']
    })

if __name__ == "__main__":
  app.run(port=Config.FLASK_PORT)