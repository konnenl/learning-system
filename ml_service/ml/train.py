from sklearn.ensemble import RandomForestClassifier
from .data_processing import get_data
import dill as pickle
import os
from pathlib import Path
import multiprocessing
max_cpu = min(4, multiprocessing.cpu_count())
os.environ['LOKY_MAX_CPU_COUNT'] = str(max_cpu)

def train_and_save(filename):
    data_path = str(Path(__file__).parent / "data.csv")
    x_train, x_test, y_train, y_test = get_data(data_path)
    
    model = RandomForestClassifier(
        n_estimators=50, 
        max_depth=5, 
        min_samples_split=2, 
        max_features='sqrt', 
        min_samples_leaf=1, 
        n_jobs=-1, 
        random_state=42)
    model.fit(x_train, y_train)
    
    print(f"Test accuracy: {model.score(x_test, y_test):.2f}")
    with open(filename, 'wb') as file:
        pickle.dump(model, file)
    return model

if __name__ == "__main__":
    model = train_and_save()
    filename = 'model_v1.pk'
    with open(filename, 'wb') as file:
        pickle.dump(model, file)