import libsvm
import libsvm.svm
import torch.nn as nn

class SVM(nn.Module):
    def __init__(self, kernel, C, gamma):
        super(SVM, self).__init__()
        self.kernel = kernel
        self.C = C
        self.gamma = gamma
        self.model = libsvm.svm.svm_train

    def forward(self, x):
        return self.model(x, self.kernel, self.C, self.gamma)
    
    def predict(self, x):
        return libsvm.svm.svm_predict(x, self.model)
    