# Naive Bayes Classifier

This repository contains simple implementations of the Naive Bayes classifier in PHP, Go, Python, and JavaScript. The Naive Bayes algorithm is used for classification tasks, assuming that features are independent given the class label, which is why it's called "naive."

## Implementation Details

### PHP

The PHP implementation can be found in the `php/naive_bayes_classifier.php` file. It includes a `NaiveBayesClassifier` class with methods for training the classifier and classifying new instances. The classifier works with categorical data and applies Laplace smoothing to avoid zero probabilities.

### Go

The Go implementation can be found in the `go/naive_bayes_classifier.go` file. It defines a `NaiveBayesClassifier` struct with methods for training and classification. Similar to the PHP version, it handles categorical data and uses Laplace smoothing.

### Python

The Python implementation can be found in the `python/naive_bayes_classifier.py` file. It provides a `NaiveBayesClassifier` class with methods for training the classifier and classifying new instances. Like the others, it works with categorical data and uses Laplace smoothing to handle unseen features.

### JavaScript

The JavaScript implementation can be found in the `javascript/naive_bayes_classifier.js` file. It defines a `NaiveBayesClassifier` class with functions for training and classification. It also assumes categorical data and uses Laplace smoothing to avoid zero probabilities.

## Usage

Each implementation has been provided with example code in its respective file to demonstrate how to use the classifier. The training data and test instances are included in the examples.

To run each implementation, you can execute the corresponding file using the appropriate interpreter or environment for each language.

## Note

These implementations are kept simple for educational purposes and are not meant for production use. For real-world applications, consider using more robust and efficient libraries or frameworks that handle different data types, scaling, and performance optimizations.

Feel free to explore and experiment with the Naive Bayes classifier implementations provided in this repository. Enjoy classifying!