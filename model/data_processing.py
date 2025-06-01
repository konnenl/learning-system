import csv
from sklearn.model_selection import train_test_split
from imblearn.over_sampling import SMOTE
import os
import multiprocessing
max_cpu = min(4, multiprocessing.cpu_count())
os.environ['LOKY_MAX_CPU_COUNT'] = str(max_cpu)

def get_data(data_path):
    with open(data_path, encoding='utf-8') as file:
        raw_data = list(csv.reader(file))

    level_label = {
        'A1': 1,
        'A2': 2,
        'B1': 3,
        'B2': 4,
        'С1': 5,
        'C2': 6,
    }
    del raw_data[0]
    x_raw = []
    y = []

    for row in raw_data:
        x_raw.append(row[2:])
        y.append(level_label[row[1][:2]])

    x = [list(map(float, l)) for l in x_raw]
    smote = SMOTE(random_state=42, k_neighbors=3)
    x_balanced, y_balanced = smote.fit_resample(x, y)

    x_train, x_test, y_train, y_test = train_test_split(x_balanced, y_balanced, test_size=0.2, stratify=y_balanced, random_state=42)

    return x_train, x_test, y_train, y_test

if __name__ == "__main__":
    x_train, x_test, y_train, y_test = get_data("data.csv")