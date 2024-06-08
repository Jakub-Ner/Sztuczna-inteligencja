import pandas as pd
from tabulate import tabulate
from sklearn.metrics import accuracy_score, confusion_matrix, classification_report
from IPython.display import display, HTML, Markdown


def present_confusion_matrix(matrix, labels=['Low', 'Medium', 'High']):
    """
    Present the confusion matrix in a readable table format.

    Parameters:
    - matrix: list of lists or 2D array, the confusion matrix
    - labels: list, the class labels
    """
    print("Confusion Matrix:")
    df = pd.DataFrame(matrix, index=[f'Actual: {label}' for label in labels], columns=[f'Predicted: {label}' for label in labels])
    table = tabulate(df, headers='keys', tablefmt='pretty', stralign='center')
    print(table)

def present_metrics(metrics_dict):
    """
    Present the metrics dictionary in a readable table format.

    Parameters:
    - metrics_dict: dict, the metrics dictionary
    """
    print("\nClassification Report:")
    df = pd.DataFrame(metrics_dict).T  # Transpose to have labels as rows
    df = df.round(3)
    table = tabulate(df, headers='keys', tablefmt='pretty', stralign='center')
    print(table)


def evaluate_the_model(y_test, y_pred):
  print("Accuracy:", accuracy_score(y_test, y_pred), '\n')
  present_confusion_matrix(confusion_matrix(y_test, y_pred))
  present_metrics(classification_report(y_test, y_pred, output_dict=True))

def print_md(string):
    display(Markdown(string))